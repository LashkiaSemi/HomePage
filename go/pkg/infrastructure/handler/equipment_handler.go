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

// EquipmentHandler ハンドラ
type EquipmentHandler interface {
	GetAll(w http.ResponseWriter, r *http.Request)
	GetByID(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type equipmentHandler struct {
	EquipmentController controller.EquipmentController
}

// NewEquipmentHandler ハンドラの作成
func NewEquipmentHandler(sh repository.SQLHandler) EquipmentHandler {
	return &equipmentHandler{
		EquipmentController: controller.NewEquipmentController(
			interactor.NewEquipmentInteractor(
				repository.NewEquipmentRepository(sh),
			),
		),
	}
}

func (eh *equipmentHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	res, err := eh.EquipmentController.ShowAll()
	if err != nil {
		response.HTTPError(w, err)
		return
	}
	response.Success(w, res)
}

func (eh *equipmentHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	equID, err := getIntParameter(r.URL.Path, "/equipments/", "")
	if err != nil {
		logger.Warn("Equipment getByID: can not get equipmentID from path")
		response.HTTPError(w, domain.BadRequest(errors.New("cant get param from path")))
		return
	}
	res, err := eh.EquipmentController.ShowByID(equID)
	if err != nil {
		response.HTTPError(w, err)
		return
	}
	response.Success(w, res)
}

func (eh *equipmentHandler) Create(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Warn(err)
		response.HTTPError(w, domain.BadRequest(err))
		return
	}
	var req controller.UpdateEquipmentRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		logger.Error(err)
		response.HTTPError(w, domain.InternalServerError(err))
		return
	}

	res, err := eh.EquipmentController.Create(&req)
	if err != nil {
		response.HTTPError(w, err)
		return
	}
	response.Success(w, res)
}

func (eh *equipmentHandler) Update(w http.ResponseWriter, r *http.Request) {
	equID, err := getIntParameter(r.URL.Path, "/equipments/", "")
	if err != nil {
		logger.Warn("updateEquipment: can not get equipmentID from path")
		response.HTTPError(w, domain.BadRequest(errors.New("cant get param from path")))
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Warn(err)
		response.HTTPError(w, domain.BadRequest(err))
		return
	}
	var req controller.UpdateEquipmentRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		logger.Error(err)
		response.HTTPError(w, domain.InternalServerError(err))
		return
	}

	res, err := eh.EquipmentController.Update(equID, &req)
	if err != nil {
		response.HTTPError(w, err)
		return
	}
	response.Success(w, res)
}

func (eh *equipmentHandler) Delete(w http.ResponseWriter, r *http.Request) {
	equID, err := getIntParameter(r.URL.Path, "/equipments/", "")
	if err != nil {
		logger.Warn("Equipment deleteHandler: can not get equipmentID from path")
		response.HTTPError(w, domain.BadRequest(errors.New("cant get param from path")))
		return
	}

	err = eh.EquipmentController.Delete(equID)
	if err != nil {
		response.HTTPError(w, err)
		return
	}
	response.NoContent(w)

}
