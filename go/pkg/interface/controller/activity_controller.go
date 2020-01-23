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
	ShowAll() (GetActivitiesResponse, error)
	ShowByID(actID int) (GetActivityResponse, error)
	Create(req *UpdateActivityRequest) (GetActivityResponse, error)
	Update(actID int, req *UpdateActivityRequest) (GetActivityResponse, error)
	Delete(actID int) error
}

type activityController struct {
	ActivityInteractor interactor.ActivityInteractor
}

// NewActivityController コントローラの作成
func NewActivityController(ai interactor.ActivityInteractor) ActivityController {
	return &activityController{
		ActivityInteractor: ai,
	}
}

func (ac *activityController) ShowAll() (res GetActivitiesResponse, err error) {
	acts, err := ac.ActivityInteractor.FetchAll()
	if err != nil {
		return
	}

	for _, act := range acts {
		res.Activities = append(res.Activities, convertActivityToResponse(&act))
	}
	return
}

// GetActivitiesResponse 複数
type GetActivitiesResponse struct {
	Activities []GetActivityResponse `json:"activities"`
}

func (ac *activityController) ShowByID(actID int) (res GetActivityResponse, err error) {
	act, err := ac.ActivityInteractor.FetchByID(actID)
	if err != nil {
		return
	}
	return convertActivityToResponse(&act), nil
}

// GetActivityResponse 一件
type GetActivityResponse struct {
	ID       int    `json:"id"`
	Date     string `json:"date"`
	Activity string `json:"activity"`
}

func (ac *activityController) Create(req *UpdateActivityRequest) (res GetActivityResponse, err error) {
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

	act, err := ac.ActivityInteractor.Add(date, req.Activity)
	if err != nil {
		return res, err
	}

	return convertActivityToResponse(&act), nil
}

// UpdateActivityRequest 新規、更新
type UpdateActivityRequest struct {
	Date     string `json:"date"`
	Activity string `json:"activity"`
}

func (ac *activityController) Update(actID int, req *UpdateActivityRequest) (res GetActivityResponse, err error) {
	// 時刻をパース
	var date time.Time
	if req.Date != "" {
		date, err = time.Parse(conf.DateFormat, req.Date)
		if err != nil {
			logger.Warn("createActivity: fail to pars. string to date")
			return res, domain.BadRequest(errors.New("fail time parse"))
		}
	}

	act, err := ac.ActivityInteractor.Update(actID, date, req.Activity)
	if err != nil {
		return res, err
	}
	return convertActivityToResponse(&act), nil
}

func (ac *activityController) Delete(actID int) error {
	return ac.ActivityInteractor.Delete(actID)
}

func convertActivityToResponse(act *domain.Activity) GetActivityResponse {
	return GetActivityResponse{
		ID:       act.ID,
		Date:     act.Date,
		Activity: act.Activity,
	}
}
