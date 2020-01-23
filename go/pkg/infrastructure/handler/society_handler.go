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

// SocietyHandler ハンドラ
type SocietyHandler interface {
	GetAll(w http.ResponseWriter, r *http.Request)
	GetByID(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type societyHandler struct {
	SocietyController controller.SocietyController
}

// NewSocietyHandler ハンドラを作成する
func NewSocietyHandler(sh repository.SQLHandler) SocietyHandler {
	return &societyHandler{
		SocietyController: controller.NewSocietyController(
			interactor.NewSocietyInteractor(
				repository.NewSocietyRepository(sh),
			),
		),
	}
}

func (sh *societyHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	res, err := sh.SocietyController.ShowAll()
	if err != nil {
		response.HTTPError(w, err)
		return
	}
	response.Success(w, res)
}

func (sh *societyHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	socID, err := getIntParameter(r.URL.Path, "/societies/", "")
	if err != nil {
		logger.Warn("getSocietyByID: can not get societyID from path")
		response.HTTPError(w, domain.BadRequest(errors.New("can not get societyID from path")))
		return
	}
	res, err := sh.SocietyController.ShowByID(socID)
	if err != nil {
		response.HTTPError(w, err)
		return
	}
	response.Success(w, res)
}

func (sh *societyHandler) Create(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Warn(err)
		response.HTTPError(w, domain.BadRequest(err))
		return
	}
	var req controller.UpdateSocietyRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		logger.Error(err)
		response.HTTPError(w, domain.InternalServerError(err))
		return
	}

	res, err := sh.SocietyController.Create(&req)
	if err != nil {
		response.HTTPError(w, err)
		return
	}
	response.Success(w, res)
}

func (sh *societyHandler) Update(w http.ResponseWriter, r *http.Request) {
	socID, err := getIntParameter(r.URL.Path, "/societies/", "")
	if err != nil {
		logger.Warn("getSocietyByID: can not get societyID from path")
		response.HTTPError(w, domain.BadRequest(errors.New("can not get societyID from path")))
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Warn(err)
		response.HTTPError(w, domain.BadRequest(err))
		return
	}
	var req controller.UpdateSocietyRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		logger.Error(err)
		response.HTTPError(w, domain.InternalServerError(err))
		return
	}

	res, err := sh.SocietyController.Update(socID, &req)
	if err != nil {
		response.HTTPError(w, err)
		return
	}
	response.Success(w, res)

}

func (sh *societyHandler) Delete(w http.ResponseWriter, r *http.Request) {
	socID, err := getIntParameter(r.URL.Path, "/societies/", "")
	if err != nil {
		logger.Warn("getSocietyByID: can not get societyID from path")
		response.HTTPError(w, domain.BadRequest(errors.New("can not get societyID from path")))
		return
	}
	err = sh.SocietyController.Delete(socID)
	if err != nil {
		response.HTTPError(w, err)
		return
	}
	response.NoContent(w)
}
