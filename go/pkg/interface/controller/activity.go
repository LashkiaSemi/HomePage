package controller

import (
	"homepage/pkg/domain/entity"
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
	GetByID(id int) (*ActivityResponse, error)
	GetUpcoming() (*ActivitiesResponse, error)
	GetForNotificaion() (*ActivitiesResponse, error)

	Create(activity, showDate, date, annotation string, isImportant, isNotify int) (int, error)
	UpdateByID(id int, activity, showDate, date, annotation string, isImportant, isNotify int) error

	DeleteByID(id int) error

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
		err = errors.Wrap(err, "failed to original data for response")
		return []*ActivitiesGroupByYearResponse{}, err
	}

	// responseづくり
	// [{ 年, [活動...] }, ... ]みたいな構造で、日付の若い順に並べてます
	var res = []*ActivitiesGroupByYearResponse{}
	var key = acts[0].Date[:4]
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
			tmp = append(tmp, convertToActivityResponse(act))
		}
	}
	if len(tmp) > 0 {
		res = append(res, &ActivitiesGroupByYearResponse{
			Year:       key,
			Activities: tmp,
		})
	}
	return res, nil
}

func (ac *activityController) GetUpcoming() (*ActivitiesResponse, error) {
	var res ActivitiesResponse
	datas, err := ac.ActivityInteractor.GetUpcoming()
	if err != nil {
		err = errors.Wrap(err, "failed to get original data")
		return &res, err
	}
	for _, data := range datas {
		res.Activities = append(res.Activities, convertToActivityResponse(data))
	}
	return &res, nil
}

func (ac *activityController) GetForNotificaion() (*ActivitiesResponse, error) {
	var res ActivitiesResponse
	datas, err := ac.ActivityInteractor.GetForNotification()
	if err != nil {
		err = errors.Wrap(err, "failed to get original data")
		return &res, err
	}
	for _, data := range datas {
		res.Activities = append(res.Activities, convertToActivityResponse(data))
	}
	return &res, nil
}

func (ac *activityController) GetByID(id int) (*ActivityResponse, error) {
	data, err := ac.ActivityInteractor.GetByID(id)
	if err != nil {
		err = errors.Wrap(err, "failed to get original data for response")
		return &ActivityResponse{}, err
	}
	return convertToActivityResponse(data), nil
}

func (ac *activityController) Create(activity, showDate, date, annotation string, isImportant, isNotify int) (int, error) {
	return ac.ActivityInteractor.Create(activity, showDate, date, annotation, isImportant, isNotify)
}

func (ac *activityController) UpdateByID(id int, activity, showDate, date, annotation string, isImportant, isNotify int) error {
	return ac.ActivityInteractor.UpdateByID(id, activity, showDate, date, annotation, isImportant, isNotify)
}

func (ac *activityController) DeleteByID(id int) error {
	return ac.ActivityInteractor.DeleteByID(id)
}

// admin
func (ac *activityController) AdminGetAll() ([]map[string]string, error) {
	var res []map[string]string
	acts, err := ac.ActivityInteractor.GetAll()
	if err != nil {
		err = errors.Wrap(err, "failed to get original data for response")
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
		err = errors.Wrap(err, "failed to get original data for response")
		return &res, err
	}
	res.Fields = append(res.Fields,
		&Field{Key: "ID", Value: data.ID},
		&Field{Key: "活動内容", Value: data.Activity},
		&Field{Key: "注釈", Value: data.Annotation},
		&Field{Key: "日付(表示用)", Value: data.ShowDate},
		&Field{Key: "日付(内部処理用)", Value: data.Date},
		&Field{Key: "重要", Value: data.IsImportant},
		&Field{Key: "通知する", Value: data.IsNotify},
	)
	res.ID = id
	return &res, nil
}

// ActivitiesGroupByYearResponse 年ごとに分けた活動内容
type ActivitiesGroupByYearResponse struct {
	Year       string
	Activities []*ActivityResponse
}

// ActivitiesResponse 活動内容を複数
type ActivitiesResponse struct {
	Activities []*ActivityResponse
}

// ActivityResponse 活動内容のレスポンス
type ActivityResponse struct {
	ID          int
	Activity    string
	Annotation  string
	ShowDate    string
	Date        string
	IsImportant bool
	IsNotify    bool
}

func convertToActivityResponse(data *entity.Activity) *ActivityResponse {
	return &ActivityResponse{
		ID:          data.ID,
		Activity:    data.Activity,
		Annotation:  data.Annotation,
		ShowDate:    data.ShowDate,
		Date:        data.Date,
		IsImportant: data.IsImportant == 1,
		IsNotify:    data.IsNotify == 1,
	}
}
