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

type JobHandler interface {
	GetAll(w http.ResponseWriter, r *http.Request)
}

type jobHandler struct {
	controller.JobController
}

func NewJobHandler(sh repository.SQLHandler) JobHandler {
	return &jobHandler{
		JobController: controller.NewJobController(
			interactor.NewJobInteractor(
				service.NewJobService(),
				repository.NewJobRepository(sh),
			),
		),
	}
}

func (jh *jobHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	info := createInfo(r, "job")

	res, err := jh.JobController.GetAll()
	if err != nil {
		log.Println("jobHandler: GetAll: ", err)
		response.InternalServerError(w, info)
		return
	}
	response.Success(w, "job/index.html", info, res)
}
