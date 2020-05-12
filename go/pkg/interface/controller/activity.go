package controller

import (
	"homepage/pkg/entity"
	"homepage/pkg/usecase/interactor"

	"github.com/pkg/errors"
)

type activityController struct {
	interactor.ActivityInteractor
}

// ActivityController 活動内容の入出力を変換
type ActivityController interface {
	GetAllGroupByYear() ([]*ActivitiesGroupByYearResponse, error)
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
