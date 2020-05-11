package handler

import (
	"homepage/pkg/domain/service"
	"homepage/pkg/infrastructure/auth"
	"homepage/pkg/infrastructure/server/response"
	"homepage/pkg/interface/controller"
	"homepage/pkg/interface/repository"
	"homepage/pkg/usecase/interactor"
	"log"
	"net/http"
)

// JobHandler 入出力の受付
type JobHandler interface {
	GetAll(w http.ResponseWriter, r *http.Request)
}

type jobHandler struct {
	controller.JobController
}

// NewJobHandler ハンドラの作成
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
	info := createInfo(r, "job", auth.GetStudentIDFromCookie(r))

	res, err := jh.JobController.GetAll()
	if err != nil {
		log.Println("jobHandler: GetAll: ", err)
		response.InternalServerError(w, info)
		return
	}
	response.Success(w, "job/index.html", info, res)
}
