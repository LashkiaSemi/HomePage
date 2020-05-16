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
	"strings"

	"github.com/gorilla/mux"
)

type activityHandler struct {
	controller.ActivityController
}

// ActivityHandler 活動内容の入出力を受付
type ActivityHandler interface {
	GetActivities(w http.ResponseWriter, r *http.Request)

	// admin
	AdminGetAll(w http.ResponseWriter, r *http.Request)
	AdminGetByID(w http.ResponseWriter, r *http.Request)
	AdminCreate(w http.ResponseWriter, r *http.Request)
	AdminUpdateByID(w http.ResponseWriter, r *http.Request)
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
	info := createInfo(r, "activities", auth.GetStudentIDFromCookie(r))

	res, err := ah.ActivityController.GetAllGroupByYear()
	if err != nil {
		log.Printf("[error] failed to get data for response: %v", err)
		response.InternalServerError(w, info)
		return
	}
	response.Render(w, "activity/index.html", info, res)
}

func (ah *activityHandler) AdminCreate(w http.ResponseWriter, r *http.Request) {
	info := createInfo(r, "activities", auth.GetStudentIDFromCookie(r))

	body := []*FormField{
		createFormField("activity", "", "活動内容", "text", nil),
		createFormField("annotation", "", "注釈", "text", nil),
		createFormField("showDate", "", "日付(表示用)", "text", nil),
		createFormField("lastDate", "", "日付(ソート、プレフィックス用。数日間の場合は、最終日の日付が良い)", "date", nil),
		createFormField("isImportant", "1", "重要", "checkbox", nil),
	}

	if r.Method == "POST" {
		// log.Println("activity create: post request")
		activity := r.PostFormValue("activity")
		showDate := r.PostFormValue("showDate")
		lastDate := r.PostFormValue("lastDate")
		annotation := r.PostFormValue("annotation")
		var isImportant int
		if activity == "" || lastDate == "" {
			info.Errors = append(info.Errors, "活動内容、ソート用日付は必須です")
		}
		if r.PostFormValue("isImportant") == "1" {
			isImportant = 1
		} else {
			isImportant = 0
		}
		if len(info.Errors) > 0 {
			response.AdminRender(w, "edit.html", info, body)
			return
		}

		id, err := ah.ActivityController.Create(activity, showDate, lastDate, annotation, isImportant)
		if err != nil {
			log.Printf("[error] failed to create: %v", err)
			response.InternalServerError(w, info)
			return
		}
		// log.Println("success update activity")
		http.Redirect(w, r, fmt.Sprintf("/admin/activities/%d", id), http.StatusSeeOther)
	}
	response.AdminRender(w, "edit.html", info, body)
}

func (ah *activityHandler) AdminUpdateByID(w http.ResponseWriter, r *http.Request) {
	info := createInfo(r, "activities", auth.GetStudentIDFromCookie(r))
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Printf("[error] failed to parse path parameter: %v", err)
		response.InternalServerError(w, info)
		return
	}
	// 初期値の取得
	data, err := ah.ActivityController.GetByID(id)
	if err != nil {
		log.Printf("[error] failed to get original data: %v", err)
		response.InternalServerError(w, info)
		return
	}
	data.LastDate = strings.Replace(data.LastDate, " ", "T", 1)
	body := []*FormField{
		createFormField("activity", data.Activity, "活動内容", "text", nil),
		createFormField("annotation", data.Annotation, "注釈", "text", nil),
		createFormField("showDate", data.ShowDate, "日付(表示用)", "text", nil),
		createFormField("lastDate", data.LastDate, "日付(ソート、プレフィックス用。数日間の場合は、最終日の日付が良い)", "date", nil),
		createFormField("isImportant", "1", "重要", "checkbox", nil),
	}

	if r.Method == "POST" {
		// log.Println("activity update: post request")
		activity := r.PostFormValue("activity")
		showDate := r.PostFormValue("showDate")
		lastDate := r.PostFormValue("lastDate")
		if activity == "" {
			info.Errors = append(info.Errors, "活動内容は必須です")
		}
		if lastDate == "" {
			lastDate = data.LastDate
		}
		annotation := r.PostFormValue("annotation")
		var isImportant int
		if r.PostFormValue("isImportant") == "1" {
			isImportant = 1
		} else {
			isImportant = 0
		}
		if len(info.Errors) > 0 {
			response.AdminRender(w, "edit.html", info, body)
			return
		}

		err = ah.ActivityController.UpdateByID(id, activity, showDate, lastDate, annotation, isImportant)
		if err != nil {
			log.Printf("[error] failed to update: %v", err)
			response.InternalServerError(w, info)
			return
		}
		// log.Println("success update activity")
		http.Redirect(w, r, fmt.Sprintf("/admin/activities/%d", id), http.StatusSeeOther)
	}

	response.AdminRender(w, "edit.html", info, body)
}

func (ah *activityHandler) AdminGetAll(w http.ResponseWriter, r *http.Request) {
	info := createInfo(r, "activities", auth.GetStudentIDFromCookie(r))
	res, err := ah.ActivityController.AdminGetAll()
	if err != nil {
		log.Printf("[error] failed to get data for response: %v", err)
		response.InternalServerError(w, info)
		return
	}

	response.AdminRender(w, "list.html", info, res)
}

func (ah *activityHandler) AdminGetByID(w http.ResponseWriter, r *http.Request) {
	info := createInfo(r, "activities", auth.GetStudentIDFromCookie(r))
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Printf("[error] failed to parse path parameter: %v", err)
		response.InternalServerError(w, info)
		return
	}
	res, err := ah.ActivityController.AdminGetByID(id)
	if err != nil {
		log.Printf("[error] failed to get data for response: %v", err)
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
		log.Printf("[error] failed to parse path parameter: %v", err)
		response.InternalServerError(w, info)
		return
	}
	body, err := ah.ActivityController.AdminGetByID(id)
	if err != nil {
		log.Printf("[error] failed to get original data: %v", err)
		response.InternalServerError(w, info)
		return
	}

	if r.Method == "POST" {
		// log.Println("post request: delete activity")
		err = ah.ActivityController.DeleteByID(id)
		if err != nil {
			log.Printf("[error] failed to delete: %v", err)
			info.Errors = append(info.Errors, "削除に失敗しました")
			response.AdminRender(w, "delete.html", info, body)
			return
		}
		// log.Println("success to delete activity")
		http.Redirect(w, r, "/admin/activities", http.StatusSeeOther)
	}
	response.AdminRender(w, "delete.html", info, body)
}
