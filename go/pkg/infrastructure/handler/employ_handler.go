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

// EmployHandler ハンドラ
type EmployHandler interface {
	GetAll(w http.ResponseWriter, r *http.Request)
	GetByID(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type employHandler struct {
	EmployController controller.EmployController
}

// NewEmployHandler ハンドラの作成
func NewEmployHandler(sh repository.SQLHandler) EmployHandler {
	return &employHandler{
		EmployController: controller.NewEmployController(
			interactor.NewEmployInteractor(
				repository.NewEmployRepository(sh),
			),
		),
	}
}

func (eh *employHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	res, err := eh.EmployController.ShowAll()
	if err != nil {
		response.HTTPError(w, err)
		return
	}
	response.Success(w, res)
}

func (eh *employHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	compID, err := getIntParameter(r.URL.Path, "/employments/", "")
	if err != nil {
		logger.Warn("employ getByID: params error. companyID parameter is not int. companyID is ", compID)
		response.HTTPError(w, domain.BadRequest(errors.New("params error")))
		return
	}

	res, err := eh.EmployController.ShowByID(compID)
	if err != nil {
		response.HTTPError(w, err)
		return
	}
	response.Success(w, res)
}

func (eh *employHandler) Create(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Warn(err)
		response.HTTPError(w, domain.BadRequest(err))
		return
	}
	var req controller.UpdateCompanyRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		logger.Error(err)
		response.HTTPError(w, domain.InternalServerError(err))
		return
	}

	res, err := eh.EmployController.Create(&req)
	if err != nil {
		response.HTTPError(w, err)
		return
	}
	response.Success(w, res)
}

func (eh *employHandler) Update(w http.ResponseWriter, r *http.Request) {
	compID, err := getIntParameter(r.URL.Path, "/employments/", "")
	if err != nil {
		logger.Warn("employ update: params error. companyID parameter is not int. companyID is ", compID)
		response.HTTPError(w, domain.BadRequest(errors.New("params error")))
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Warn(err)
		response.HTTPError(w, domain.BadRequest(err))
		return
	}
	var req controller.UpdateCompanyRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		logger.Error(err)
		response.HTTPError(w, domain.InternalServerError(err))
		return
	}

	res, err := eh.EmployController.Update(compID, &req)
	if err != nil {
		response.HTTPError(w, err)
		return
	}
	response.Success(w, res)
}

func (eh *employHandler) Delete(w http.ResponseWriter, r *http.Request) {
	compID, err := getIntParameter(r.URL.Path, "/employments/", "")
	if err != nil {
		logger.Warn("employ delete: params error. companyID parameter is not int. companyID is ", compID)
		response.HTTPError(w, domain.BadRequest(errors.New("params error")))
		return
	}
	err = eh.EmployController.Delete(compID)
	if err != nil {
		response.HTTPError(w, err)
		return
	}
	response.NoContent(w)
}
