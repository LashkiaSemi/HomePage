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

type activityHandler struct {
	controller.ActivityController
}

// ActivityHandler 活動内容の入出力を受付
type ActivityHandler interface {
	GetActivities(w http.ResponseWriter, r *http.Request)

	Create(w http.ResponseWriter, r *http.Request)
	UpdateByID(w http.ResponseWriter, r *http.Request)

	// admin
	AdminGetAll(w http.ResponseWriter, r *http.Request)
	AdminGetByID(w http.ResponseWriter, r *http.Request)
	AdminDeleteByID(w http.ResponseWriter, r *http.Request)
}

// NewActivityHandler ハンドラの作成
func NewActivityHandler(sh repository.SQLHandler) ActivityHandler {
	return &activityHandler{
		ActivityController: controller.NewActivityController(
			interactor.NewActivityInteractor(
				repository.NewActivityRepository(sh),
			),
		),
	}
}

func (ah *activityHandler) GetActivities(w http.ResponseWriter, r *http.Request) {
	info := createInfo(r, "activity", auth.GetStudentIDFromCookie(r))

	res, err := ah.ActivityController.GetAllGroupByYear()
	if err != nil {
		log.Println(err)
		response.InternalServerError(w, info)
		return
	}
	response.Success(w, "activity/index.html", info, res)
}

func (ah *activityHandler) Create(w http.ResponseWriter, r *http.Request) {
	info := createInfo(r, "activities", auth.GetStudentIDFromCookie(r))

	body := []*FormField{
		createFormField("activity", "", "活動内容", "text", nil),
		createFormField("date", "", "日付", "date", nil),
	}

	if r.Method == "POST" {
		log.Println("activity create: post request")
		activity := r.PostFormValue("activity")
		date := r.PostFormValue("date")
		if activity == "" || date == "" {
			info.Errors = append(info.Errors, "活動内容、日付は必須です")
			response.AdminRender(w, "edit.html", info, body)
			return
		}

		id, err := ah.ActivityController.Create(activity, date)
		if err != nil {
			log.Println(err)
			response.InternalServerError(w, info)
			return
		}
		log.Println("success update activity")
		http.Redirect(w, r, fmt.Sprintf("/admin/activities/%d", id), http.StatusSeeOther)
	}

	response.AdminRender(w, "edit.html", info, body)
}

func (ah *activityHandler) UpdateByID(w http.ResponseWriter, r *http.Request) {
	info := createInfo(r, "activities", auth.GetStudentIDFromCookie(r))
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Println("failed to parse path parameter", err)
		response.InternalServerError(w, info)
		return
	}
	// 初期値の取得
	data, err := ah.ActivityController.GetByID(id)
	if err != nil {
		log.Println("failed to get target: ", err)
		response.InternalServerError(w, info)
		return
	}
	body := []*FormField{
		createFormField("activity", data.Activity, "活動内容", "text", nil),
		createFormField("date", data.Date, "日付", "date", nil),
	}

	if r.Method == "POST" {
		log.Println("activity update: post request")
		activity := r.PostFormValue("activity")
		date := r.PostFormValue("date")
		if activity == "" || date == "" {
			info.Errors = append(info.Errors, "活動内容、日付は必須です")
			response.AdminRender(w, "edit.html", info, body)
			return
		}

		err = ah.ActivityController.UpdateByID(id, activity, date)
		if err != nil {
			log.Println(err)
			response.InternalServerError(w, info)
			return
		}
		log.Println("success update activity")
		http.Redirect(w, r, fmt.Sprintf("/admin/activities/%d", id), http.StatusSeeOther)
	}

	response.AdminRender(w, "edit.html", info, body)
}

func (ah *activityHandler) AdminGetAll(w http.ResponseWriter, r *http.Request) {
	info := createInfo(r, "activities", auth.GetStudentIDFromCookie(r))
	res, err := ah.ActivityController.AdminGetAll()
	if err != nil {
		log.Println(err)
		response.InternalServerError(w, info)
		return
	}

	response.AdminRender(w, "list.html", info, res)

}

func (ah *activityHandler) AdminGetByID(w http.ResponseWriter, r *http.Request) {
	info := createInfo(r, "activities", auth.GetStudentIDFromCookie(r))
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Println("activityHandler: AdminGetByID: failed to parse path param: ", err)
		response.InternalServerError(w, info)
		return
	}
	res, err := ah.ActivityController.AdminGetByID(id)
	if err != nil {
		log.Println("activityHandler: AdminGetByID: ", err)
		response.InternalServerError(w, info)
		return
	}
	res.ID = id
	response.AdminRender(w, "detail.html", info, res)
}

func (ah *activityHandler) AdminDeleteByID(w http.ResponseWriter, r *http.Request) {
	info := createInfo(r, "activities", auth.GetStudentIDFromCookie(r))
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Println("failed to parse path parameter", err)
		response.InternalServerError(w, info)
		return
	}
	body, err := ah.ActivityController.AdminGetByID(id)
	if err != nil {
		log.Println("AdminDeleteByID: ", err)
		response.InternalServerError(w, info)
		return
	}

	if r.Method == "POST" {
		log.Println("post request: delete activity")
		err = ah.ActivityController.DeleteByID(id)
		if err != nil {
			log.Println("failed to delete")
			info.Errors = append(info.Errors, "削除に失敗しました")
			response.AdminRender(w, "delete.html", info, body)
			return
		}
		log.Println("success to delete activity")
		http.Redirect(w, r, "/admin/activities", http.StatusSeeOther)
	}
	response.AdminRender(w, "delete.html", info, body)
}
