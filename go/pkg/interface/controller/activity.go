package controller

import (
	"homepage/pkg/entity"
	"homepage/pkg/usecase/interactor"
	"strconv"

	"github.com/pkg/errors"
)

type activityController struct {
	interactor.ActivityInteractor
}

// ActivityController 活動内容の入出力を変換
type ActivityController interface {
	GetAllGroupByYear() ([]*ActivitiesGroupByYearResponse, error)

	// admin
	AdminGetAll() ([]map[string]string, error)
	AdminGetByID(id int) (*FieldsResponse, error)
}

// NewActivityController コントローラの作成
func NewActivityController(ai interactor.ActivityInteractor) ActivityController {
	return &activityController{
		ActivityInteractor: ai,
	}
}

func (ac *activityController) GetAllGroupByYear() ([]*ActivitiesGroupByYearResponse, error) {
	acts, err := ac.ActivityInteractor.GetAll()
	if err != nil {
		err = errors.Wrap(err, "controller")
		return []*ActivitiesGroupByYearResponse{}, err
	}

	// responseづくり...
	var res = []*ActivitiesGroupByYearResponse{}
	var key = ""
	var tmp = []*ActivityResponse{}
	for _, act := range acts {
		if key == act.Date[:4] {
			tmp = append(tmp, convertToActivityResponse(act))
			continue
		} else {
			if len(tmp) > 0 {
				res = append(res, &ActivitiesGroupByYearResponse{
					Year:       key,
					Activities: tmp,
				})
			}
			key = act.Date[:4]
			tmp = []*ActivityResponse{}
		}
	}
	if len(tmp) > 0 {
		res = append(res, &ActivitiesGroupByYearResponse{
			Year:       key,
			Activities: tmp,
		})
	}
	return res, err
}

func (ac *activityController) AdminGetAll() ([]map[string]string, error) {
	var res []map[string]string
	acts, err := ac.ActivityInteractor.GetAll()
	if err != nil {
		err = errors.Wrap(err, "AdminGetAll")
		return res, err
	}
	for _, act := range acts {
		res = append(res, map[string]string{
			"id":    strconv.Itoa(act.ID),
			"title": act.Activity,
		})
	}
	return res, nil
}

func (ac *activityController) AdminGetByID(id int) (*FieldsResponse, error) {
	var res FieldsResponse
	data, err := ac.ActivityInteractor.GetByID(id)
	if err != nil {
		err = errors.Wrap(err, "AdminGetByID")
		return &res, err
	}
	res.Fields = append(res.Fields,
		&Field{Key: "ID", Value: data.ID},
		&Field{Key: "活動内容", Value: data.Activity},
		&Field{Key: "日付", Value: data.Date},
	)
	res.ID = id
	return &res, nil
}

// ActivitiesGroupByYearResponse 年ごとに分けた活動内容
type ActivitiesGroupByYearResponse struct {
	Year       string
	Activities []*ActivityResponse
}

// ActivityResponse 活動内容のレスポンス
type ActivityResponse struct {
	ID       int
	Activity string
	Date     string
}

func convertToActivityResponse(data *entity.Activity) *ActivityResponse {
	return &ActivityResponse{
		ID:       data.ID,
		Activity: data.Activity,
		Date:     data.Date,
	}
}
