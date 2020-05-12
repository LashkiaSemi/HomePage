package handler

import (
	"homepage/pkg/infrastructure/auth"
	"homepage/pkg/infrastructure/server/response"
	"homepage/pkg/interface/controller"
	"homepage/pkg/interface/repository"
	"homepage/pkg/usecase/interactor"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// JobHandler 入出力の受付
type JobHandler interface {
	GetAll(w http.ResponseWriter, r *http.Request)

	// admin
	AdminGetAll(w http.ResponseWriter, r *http.Request)
	AdminGetByID(w http.ResponseWriter, r *http.Request)
}

type jobHandler struct {
	controller.JobController
}

// NewJobHandler ハンドラの作成
func NewJobHandler(sh repository.SQLHandler) JobHandler {
	return &jobHandler{
		JobController: controller.NewJobController(
			interactor.NewJobInteractor(
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

// admin
func (jh *jobHandler) AdminGetAll(w http.ResponseWriter, r *http.Request) {
	info := createInfo(r, "jobs", auth.GetStudentIDFromCookie(r))
	res, err := jh.JobController.AdminGetAll()
	if err != nil {
		log.Println("jobHandler: AdminGetAll: ", err)
		response.InternalServerError(w, info)
		return
	}
	response.AdminRender(w, "list.html", info, res)
}

func (jh *jobHandler) AdminGetByID(w http.ResponseWriter, r *http.Request) {
	info := createInfo(r, "jobs", auth.GetStudentIDFromCookie(r))
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Println("jobHandler: AdminGetByID: failed to parse path param: ", err)
		response.InternalServerError(w, info)
		return
	}
	res, err := jh.JobController.AdminGetByID(id)
	if err != nil {
		log.Println("jobHandler: AdminGetByID: ", err)
		response.InternalServerError(w, info)
		return
	}
	response.AdminRender(w, "detail.html", info, res)
}
