package handler

import (
	"homepage/pkg/infrastructure/auth"
	"homepage/pkg/infrastructure/server/response"
	"homepage/pkg/interface/controller"
	"homepage/pkg/interface/repository"
	"homepage/pkg/usecase/interactor"
	"net/http"
)

type equipmentHandler struct {
	controller.EquipmentController
}

// EquipmentHandler 備品関連の入出力を受け付け
type EquipmentHandler interface {
	GetAll(w http.ResponseWriter, r *http.Request)
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
	info := createInfo(r, "equipment", auth.GetStudentIDFromCookie(r))
	res, err := eh.EquipmentController.GetAll()
	if err != nil {
		response.InternalServerError(w, info)
		return
	}
	response.Success(w, "equipment/index.html", info, res)
}
