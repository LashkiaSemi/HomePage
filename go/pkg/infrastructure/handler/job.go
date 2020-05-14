package handler

import (
	"fmt"
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

	AdminCreate(w http.ResponseWriter, r *http.Request)
	AdminUpdateByID(w http.ResponseWriter, r *http.Request)

	// admin
	AdminGetAll(w http.ResponseWriter, r *http.Request)
	AdminGetByID(w http.ResponseWriter, r *http.Request)
	AdminDeleteByID(w http.ResponseWriter, r *http.Request)
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
		log.Printf("[error] failed to get data for response: %v", err)
		response.InternalServerError(w, info)
		return
	}
	response.Render(w, "job/index.html", info, res)
}

func (jh *jobHandler) AdminCreate(w http.ResponseWriter, r *http.Request) {
	info := createInfo(r, "jobs", auth.GetStudentIDFromCookie(r))

	body := []*FormField{
		createFormField("company", "", "企業名", "text", nil),
		createFormField("job", "", "職種", "text", nil),
	}

	if r.Method == "POST" {
		// log.Println("job create: post request")
		company := r.PostFormValue("company")
		job := r.PostFormValue("job")
		if company == "" {
			info.Errors = append(info.Errors, "企業名は必須です")
			response.AdminRender(w, "edit.html", info, body)
			return
		}

		id, err := jh.JobController.Create(company, job)
		if err != nil {
			log.Printf("[error] failed to create: %v", err)
			response.InternalServerError(w, info)
			return
		}
		// log.Println("success create job")
		http.Redirect(w, r, fmt.Sprintf("/admin/jobs/%d", id), http.StatusSeeOther)
	}

	response.AdminRender(w, "edit.html", info, body)
}

func (jh *jobHandler) AdminUpdateByID(w http.ResponseWriter, r *http.Request) {
	info := createInfo(r, "jobs", auth.GetStudentIDFromCookie(r))
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Printf("[error] failed to parse path parameter: %v", err)
		response.InternalServerError(w, info)
		return
	}
	data, err := jh.JobController.GetByID(id)
	if err != nil {
		log.Printf("[error] failed to get original data: %v", err)
		response.InternalServerError(w, info)
		return
	}
	body := []*FormField{
		createFormField("company", data.Company, "企業名", "text", nil),
		createFormField("job", data.Job, "職種", "text", nil),
	}

	if r.Method == "POST" {
		// log.Println("job update: post request")
		company := r.PostFormValue("company")
		job := r.PostFormValue("job")
		if company == "" {
			info.Errors = append(info.Errors, "企業名は必須です")
			response.AdminRender(w, "edit.html", info, body)
			return
		}

		err = jh.JobController.UpdateByID(id, company, job)
		if err != nil {
			log.Printf("[error] failed to update: %v", err)
			response.InternalServerError(w, info)
			return
		}
		// log.Println("success update job")
		http.Redirect(w, r, fmt.Sprintf("/admin/jobs/%d", id), http.StatusSeeOther)
	}
	response.AdminRender(w, "edit.html", info, body)
}

// admin
func (jh *jobHandler) AdminGetAll(w http.ResponseWriter, r *http.Request) {
	info := createInfo(r, "jobs", auth.GetStudentIDFromCookie(r))
	res, err := jh.JobController.AdminGetAll()
	if err != nil {
		log.Printf("[error] failed to get data for response: %v", err)
		response.InternalServerError(w, info)
		return
	}
	response.AdminRender(w, "list.html", info, res)
}

func (jh *jobHandler) AdminGetByID(w http.ResponseWriter, r *http.Request) {
	info := createInfo(r, "jobs", auth.GetStudentIDFromCookie(r))
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Printf("[error] failed to parse path param: %v", err)
		response.InternalServerError(w, info)
		return
	}
	res, err := jh.JobController.AdminGetByID(id)
	if err != nil {
		log.Printf("[error] failed to get data for response: %v", err)
		response.InternalServerError(w, info)
		return
	}
	res.ID = id
	response.AdminRender(w, "detail.html", info, res)
}

func (jh *jobHandler) AdminDeleteByID(w http.ResponseWriter, r *http.Request) {
	info := createInfo(r, "jobs", auth.GetStudentIDFromCookie(r))
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Printf("[error] failed to parse path parameter: %v", err)
		response.InternalServerError(w, info)
		return
	}
	body, err := jh.JobController.AdminGetByID(id)
	if err != nil {
		log.Printf("[error] failed to get original data: %v", err)
		response.InternalServerError(w, info)
		return
	}

	if r.Method == "POST" {
		// log.Println("post request: delete job")
		err = jh.JobController.DeleteByID(id)
		if err != nil {
			log.Printf("[error] failed to delete: %v", err)
			info.Errors = append(info.Errors, "削除に失敗しました")
			response.AdminRender(w, "delete.html", info, body)
			return
		}
		// log.Println("success to delete job")
		http.Redirect(w, r, "/admin/jobs", http.StatusSeeOther)
	}
	response.AdminRender(w, "delete.html", info, body)
}
