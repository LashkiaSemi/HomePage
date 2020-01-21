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

type ActivityHandler interface {
	GetActivities(w http.ResponseWriter, r *http.Request)
	GetActivityByID(w http.ResponseWriter, r *http.Request)
	CreateActivity(w http.ResponseWriter, r *http.Request)
	UpdateActivity(w http.ResponseWriter, r *http.Request)
	DeleteActivity(w http.ResponseWriter, r *http.Request)
}

type activityHandler struct {
	ActivityController controller.ActivityController
}

func NewActivityHandler(sh repository.SQLHandler) ActivityHandler {
	return &activityHandler{
		ActivityController: controller.NewActivityController(
			interactor.NewActivityInteractor(
				repository.NewActivityRepository(sh),
			),
		),
	}
}

func (ah *activityHandler) GetActivities(w http.ResponseWriter, r *http.Request) {
	res, err := ah.ActivityController.ShowActivities()
	if err != nil {
		response.HTTPError(w, err)
		return
	}

	response.Success(w, res)
}

func (ah *activityHandler) GetActivityByID(w http.ResponseWriter, r *http.Request) {
	actID, err := getIntParameter(r.URL.Path, "/activities/", "")
	if err != nil {
		logger.Warn("getActivity: can not get activityID from path")
		response.HTTPError(w, domain.BadRequest(errors.New("getActivity: cant get param from path")))
		return
	}

	res, err := ah.ActivityController.ShowActivityByID(actID)
	if err != nil {
		response.HTTPError(w, err)
		return
	}

	response.Success(w, res)
}

func (ah *activityHandler) CreateActivity(w http.ResponseWriter, r *http.Request) {
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

	res, err := ah.ActivityController.CreateActivity(&req)
	if err != nil {
		response.HTTPError(w, err)
		return
	}

	response.Success(w, res)
}

func (ah *activityHandler) UpdateActivity(w http.ResponseWriter, r *http.Request) {
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

	res, err := ah.ActivityController.UpdateActivity(actID, &req)
	if err != nil {
		response.HTTPError(w, err)
		return
	}
	response.Success(w, res)

}

func (ah *activityHandler) DeleteActivity(w http.ResponseWriter, r *http.Request) {
	actID, err := getIntParameter(r.URL.Path, "/activities/", "")
	if err != nil {
		logger.Warn("deleteActivity: can not get activityID from path")
		response.HTTPError(w, domain.BadRequest(errors.New("deleteActivity: cant get param from path")))
		return
	}

	err = ah.ActivityController.DeleteActivity(actID)
	if err != nil {
		response.HTTPError(w, err)
		return
	}
	response.NoContent(w)
}
