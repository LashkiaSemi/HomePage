package handler

import (
	"homepage/pkg/domain/service"
	"homepage/pkg/infrastructure/server/response"
	"homepage/pkg/interface/controller"
	"homepage/pkg/interface/repository"
	"homepage/pkg/usecase/interactor"
	"log"
	"net/http"
)

type societyHandler struct {
	controller.SocietyController
}

type SocietyHandler interface {
	GetAll(w http.ResponseWriter, r *http.Request)
}

func NewSocietyHandler(sh repository.SQLHandler) SocietyHandler {
	// TODO;!!!!
	return &societyHandler{
		controller.NewSocietyController(
			interactor.NewSocietyInteractor(
				service.NewSocietyService(),
				repository.NewSocietyRepository(sh),
			),
		),
	}
}

func (sh *societyHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	info := createInfo(r, "society")

	// get data
	datas, err := sh.SocietyController.GetAll()
	if err != nil {
		log.Println(err)
	}
	res := datas
	// response
	response.Success(w, "society/index.html", info, res)
}
