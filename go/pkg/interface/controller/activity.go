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
	GetAllGroupByYear() (*ActivitiesGroupByYearResponse, error)
}

// NewActivityController コントローラの作成
func NewActivityController(ai interactor.ActivityInteractor) ActivityController {
	return &activityController{
		ActivityInteractor: ai,
	}
}

func (ac *activityController) GetAllGroupByYear() (*ActivitiesGroupByYearResponse, error) {
	acts, err := ac.ActivityInteractor.GetAll()
	if err != nil {
		err = errors.Wrap(err, "controller")
		return &ActivitiesGroupByYearResponse{}, err
	}
	var res = make(map[string][]*ActivityResponse)
	for _, act := range acts {
		res[act.Date[:4]] = append(res[act.Date[:4]], convertToActivityResponse(act))
	}
	// TODO: 降順にmap
	return &ActivitiesGroupByYearResponse{Activities: res}, err
}

// ActivitiesGroupByYearResponse 年ごとに分けた活動内容
type ActivitiesGroupByYearResponse struct {
	Activities map[string][]*ActivityResponse
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
