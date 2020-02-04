package handler

import (
	"encoding/json"
	"errors"
	"homepage/pkg/domain"
	"homepage/pkg/domain/logger"
	"homepage/pkg/infrastructure/server/response"
	"homepage/pkg/interface/controller"
	"homepage/pkg/interface/repository"
	"homepage/pkg/usecase/interactor"
	"io/ioutil"
	"net/http"
)

// ActivityHandler ハンドラ
type ActivityHandler interface {
	GetAll(w http.ResponseWriter, r *http.Request)
	GetByID(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type activityHandler struct {
	ActivityController controller.ActivityController
}

// NewActivityHandler ハンドラの作成
func NewActivityHandler(sh repository.SQLHandler) ActivityHandler {
	return &activityHandler{
		ActivityController: controller.NewActivityController(
			interactor.NewActivityInteractor(
				repository.NewActivityRepository(sh),
			),
		),
	}
}

func (ah *activityHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	res, err := ah.ActivityController.ShowAll()
	if err != nil {
		response.HTTPError(w, err)
		return
	}

	response.Success(w, res)
}

func (ah *activityHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	actID, err := getIntParameter(r.URL.Path, "/activities/", "")
	if err != nil {
		logger.Warn("getActivity: can not get activityID from path")
		response.HTTPError(w, domain.BadRequest(errors.New("getActivity: cant get param from path")))
		return
	}

	res, err := ah.ActivityController.ShowByID(actID)
	if err != nil {
		response.HTTPError(w, err)
		return
	}

	response.Success(w, res)
}

func (ah *activityHandler) Create(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Warn(err)
		response.HTTPError(w, domain.BadRequest(err))
		return
	}
	var req controller.UpdateActivityRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		logger.Error(err)
		response.HTTPError(w, domain.InternalServerError(err))
		return
	}

	res, err := ah.ActivityController.Create(&req)
	if err != nil {
		response.HTTPError(w, err)
		return
	}

	response.Success(w, res)
}

func (ah *activityHandler) Update(w http.ResponseWriter, r *http.Request) {
	actID, err := getIntParameter(r.URL.Path, "/activities/", "")
	if err != nil {
		logger.Warn("updateActivity: can not get activityID from path")
		response.HTTPError(w, domain.BadRequest(errors.New("updateActivity: cant get param from path")))
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Warn(err)
		response.HTTPError(w, domain.BadRequest(err))
		return
	}
	var req controller.UpdateActivityRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		logger.Error(err)
		response.HTTPError(w, domain.InternalServerError(err))
		return
	}

	res, err := ah.ActivityController.Update(actID, &req)
	if err != nil {
		response.HTTPError(w, err)
		return
	}
	response.Success(w, res)

}

func (ah *activityHandler) Delete(w http.ResponseWriter, r *http.Request) {
	actID, err := getIntParameter(r.URL.Path, "/activities/", "")
	if err != nil {
		logger.Warn("deleteActivity: can not get activityID from path")
		response.HTTPError(w, domain.BadRequest(errors.New("deleteActivity: cant get param from path")))
		return
	}

	err = ah.ActivityController.Delete(actID)
	if err != nil {
		response.HTTPError(w, err)
		return
	}
	response.NoContent(w)
}
