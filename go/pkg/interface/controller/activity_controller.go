package controller

import (
	"errors"
	"homepage/conf"
	"homepage/pkg/domain"
	"homepage/pkg/domain/logger"
	"homepage/pkg/usecase/interactor"
	"time"
)

// ActivityController コントローラ
type ActivityController interface {
	ShowActivities() (GetActivitiesResponse, error)
	ShowActivityByID(actID int) (GetActivityResponse, error)
	CreateActivity(req *UpdateActivityRequest) (GetActivityResponse, error)
	UpdateActivity(actID int, req *UpdateActivityRequest) (GetActivityResponse, error)
	DeleteActivity(actID int) error
}

type activityController struct {
	ActivityInteractor interactor.ActivityInteractor
}

// NewActivityController コントローラを作成する
func NewActivityController(ai interactor.ActivityInteractor) ActivityController {
	return &activityController{
		ActivityInteractor: ai,
	}
}

func (ac *activityController) ShowActivities() (res GetActivitiesResponse, err error) {
	acts, err := ac.ActivityInteractor.FetchActivities()
	if err != nil {
		return
	}

	for _, act := range acts {
		res.Activities = append(res.Activities, GetActivityResponse{
			ID:       act.ID,
			Date:     act.Date,
			Activity: act.Activity,
		})
	}
	return
}

// GetActivitiesResponse レスポンス
type GetActivitiesResponse struct {
	Activities []GetActivityResponse `json:"activities"`
}

func (ac *activityController) ShowActivityByID(actID int) (res GetActivityResponse, err error) {
	act, err := ac.ActivityInteractor.FetchActivityByID(actID)
	if err != nil {
		return
	}
	res.ID = act.ID
	res.Date = act.Date
	res.Activity = act.Activity
	return
}

// GetActivityResponse レスポンス
type GetActivityResponse struct {
	ID       int    `json:"id"`
	Date     string `json:"date"`
	Activity string `json:"activity"`
}

func (ah *activityController) CreateActivity(req *UpdateActivityRequest) (res GetActivityResponse, err error) {
	// 入力へのバリデーション
	if req.Activity == "" {
		logger.Warn("createActivity: activity is empty")
		return res, domain.BadRequest(errors.New("activity is empty"))
	}
	if req.Date == "" {
		logger.Warn("createActivity: date is empty")
		return res, domain.BadRequest(errors.New("date is empty"))
	}

	// 時刻をパース
	date, err := time.Parse(conf.DateFormat, req.Date)
	if err != nil {
		logger.Warn("createActivity: fail to pars. string to date")
		return res, domain.BadRequest(errors.New("fail time parse"))
	}

	act, err := ah.ActivityInteractor.AddActiviry(date, req.Activity)
	if err != nil {
		return res, err
	}

	res.ID = act.ID
	res.Date = act.Date
	res.Activity = act.Activity
	return
}

// UpdateActivityRequest 新規、更新時のbody
type UpdateActivityRequest struct {
	Date     string `json:"date"`
	Activity string `json:"activity"`
}

func (ah *activityController) UpdateActivity(actID int, req *UpdateActivityRequest) (res GetActivityResponse, err error) {
	// 時刻をパース
	var date time.Time
	if req.Date != "" {
		date, err = time.Parse(conf.DateFormat, req.Date)
		if err != nil {
			logger.Warn("createActivity: fail to pars. string to date")
			return res, domain.BadRequest(errors.New("fail time parse"))
		}
	}

	act, err := ah.ActivityInteractor.UpdateActiviry(actID, date, req.Activity)
	if err != nil {
		return res, err
	}

	res.ID = act.ID
	res.Date = act.Date
	res.Activity = act.Activity
	return
}

func (ah *activityController) DeleteActivity(actID int) error {
	return ah.ActivityInteractor.DeleteActiviry(actID)
}
