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

// ResearchHandler ハンドラ
type ResearchHandler interface {
	GetAll(w http.ResponseWriter, r *http.Request)
	GetByID(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type researchHandler struct {
	ResearchController controller.ResearchController
}

// NewResearchHandler ハンドラを作成する
func NewResearchHandler(sh repository.SQLHandler) ResearchHandler {
	return &researchHandler{
		ResearchController: controller.NewResearchController(
			interactor.NewResearchInteractor(
				repository.NewResearchRepository(sh),
			),
		),
	}
}

func (rh *researchHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	res, err := rh.ResearchController.ShowAll()
	if err != nil {
		response.HTTPError(w, err)
		return
	}
	response.Success(w, res)
}

func (rh *researchHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	resID, err := getIntParameter(r.URL.Path, "/researches/", "")
	if err != nil {
		logger.Warn("research get: params error. userID parameter is not int. researchID is ", resID)
		response.HTTPError(w, domain.BadRequest(errors.New("params error")))
		return
	}

	res, err := rh.ResearchController.ShowByID(resID)
	if err != nil {
		response.HTTPError(w, err)
		return
	}
	response.Success(w, res)
}

func (rh *researchHandler) Create(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Warn(err)
		response.HTTPError(w, domain.BadRequest(err))
		return
	}
	var req controller.UpdateResearchRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		logger.Error(err)
		response.HTTPError(w, domain.InternalServerError(err))
		return
	}

	res, err := rh.ResearchController.Create(&req)
	if err != nil {
		response.HTTPError(w, err)
		return
	}
	response.Success(w, res)
}

func (rh *researchHandler) Update(w http.ResponseWriter, r *http.Request) {
	resID, err := getIntParameter(r.URL.Path, "/researches/", "")
	if err != nil {
		logger.Warn("research get: params error. userID parameter is not int. researchID is ", resID)
		response.HTTPError(w, domain.BadRequest(errors.New("params error")))
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Warn(err)
		response.HTTPError(w, domain.BadRequest(err))
		return
	}
	var req controller.UpdateResearchRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		logger.Error(err)
		response.HTTPError(w, domain.InternalServerError(err))
		return
	}

	res, err := rh.ResearchController.Update(resID, &req)
	if err != nil {
		response.HTTPError(w, err)
		return
	}
	response.Success(w, res)
}

func (rh *researchHandler) Delete(w http.ResponseWriter, r *http.Request) {
	resID, err := getIntParameter(r.URL.Path, "/researches/", "")
	if err != nil {
		logger.Warn("research get: params error. userID parameter is not int. researchID is ", resID)
		response.HTTPError(w, domain.BadRequest(errors.New("params error")))
		return
	}

	err = rh.ResearchController.Delete(resID)
	if err != nil {
		response.HTTPError(w, err)
		return
	}
	response.NoContent(w)
}
