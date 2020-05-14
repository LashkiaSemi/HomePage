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

type societyHandler struct {
	controller.SocietyController
}

// SocietyHandler 学会発表の入出力の受付
type SocietyHandler interface {
	GetAll(w http.ResponseWriter, r *http.Request)

	AdminCreate(w http.ResponseWriter, r *http.Request)
	AdminUpdateByID(w http.ResponseWriter, r *http.Request)

	// admin
	AdminGetAll(w http.ResponseWriter, r *http.Request)
	AdminGeByID(w http.ResponseWriter, r *http.Request)
	AdminDeleteByID(w http.ResponseWriter, r *http.Request)
}

// NewSocietyHandler ハンドラの作成
func NewSocietyHandler(sh repository.SQLHandler) SocietyHandler {
	return &societyHandler{
		controller.NewSocietyController(
			interactor.NewSocietyInteractor(
				repository.NewSocietyRepository(sh),
			),
		),
	}
}

func (sh *societyHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	info := createInfo(r, "society", auth.GetStudentIDFromCookie(r))

	// get data
	res, err := sh.SocietyController.GetAll()
	if err != nil {
		log.Printf("[error] failed to get data for response: %v", err)
		response.InternalServerError(w, info)
		return
	}
	// response
	response.Render(w, "society/index.html", info, res)
}

func (sh *societyHandler) AdminCreate(w http.ResponseWriter, r *http.Request) {
	info := createInfo(r, "societies", auth.GetStudentIDFromCookie(r))

	body := []*FormField{
		createFormField("title", "", "タイトル", "text", nil),
		createFormField("author", "", "著者", "text", nil),
		createFormField("society", "", "学会", "text", nil),
		createFormField("award", "", "受賞", "text", nil),
		createFormField("date", "", "日付", "date", nil),
	}

	if r.Method == "POST" {
		// log.Println("society create: post request")
		title := r.PostFormValue("title")
		author := r.PostFormValue("author")
		society := r.PostFormValue("society")
		award := r.PostFormValue("award")
		date := r.PostFormValue("date")
		if title == "" || author == "" || society == "" || date == "" {
			info.Errors = append(info.Errors, "タイトル、著者、学会、日付は必須です")
		}
		if len(info.Errors) > 0 {
			response.AdminRender(w, "edit.html", info, body)
			return
		}

		id, err := sh.SocietyController.Create(title, author, society, award, date)
		if err != nil {
			log.Printf("[error] failed to create: %v", err)
			response.InternalServerError(w, info)
			return
		}
		// log.Println("success create society")
		http.Redirect(w, r, fmt.Sprintf("/admin/societies/%d", id), http.StatusSeeOther)
	}

	response.AdminRender(w, "edit.html", info, body)
}

func (sh *societyHandler) AdminUpdateByID(w http.ResponseWriter, r *http.Request) {
	info := createInfo(r, "societies", auth.GetStudentIDFromCookie(r))
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Printf("[error] failed to parse path parameter: %v", err)
		response.InternalServerError(w, info)
		return
	}
	// 初期値の取得
	data, err := sh.SocietyController.GetByID(id)
	if err != nil {
		log.Printf("[error] failed to get original data: %v", err)
		response.InternalServerError(w, info)
		return
	}
	body := []*FormField{
		createFormField("title", data.Title, "タイトル", "text", nil),
		createFormField("author", data.Author, "著者", "text", nil),
		createFormField("society", data.Society, "学会", "text", nil),
		createFormField("award", data.Author, "受賞", "text", nil),
		createFormField("date", data.Date, "日付", "date", nil),
	}

	if r.Method == "POST" {
		// log.Println("society update: post request")
		title := r.PostFormValue("title")
		author := r.PostFormValue("author")
		society := r.PostFormValue("society")
		award := r.PostFormValue("award")
		date := r.PostFormValue("date")
		if title == "" || author == "" || society == "" || date == "" {
			info.Errors = append(info.Errors, "タイトル、著者、学会、日付は必須です")
		}
		if len(info.Errors) > 0 {
			response.AdminRender(w, "edit.html", info, body)
			return
		}

		err = sh.SocietyController.UpdateByID(id, title, author, society, award, date)
		if err != nil {
			log.Printf("[error] failed to update: %v", err)
			response.InternalServerError(w, info)
			return
		}
		// log.Println("success update society")
		http.Redirect(w, r, fmt.Sprintf("/admin/societies/%d", id), http.StatusSeeOther)
	}

	response.AdminRender(w, "edit.html", info, body)
}

// admin
func (sh *societyHandler) AdminGetAll(w http.ResponseWriter, r *http.Request) {
	info := createInfo(r, "societies", auth.GetStudentIDFromCookie(r))
	res, err := sh.SocietyController.AdminGetAll()
	if err != nil {
		log.Printf("[error] failed to get data for response: %v", err)
		response.InternalServerError(w, info)
		return
	}
	response.AdminRender(w, "list.html", info, res)
}

func (sh *societyHandler) AdminGeByID(w http.ResponseWriter, r *http.Request) {
	info := createInfo(r, "societies", auth.GetStudentIDFromCookie(r))
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Printf("[error] failed to parse path param: %v", err)
		response.InternalServerError(w, info)
		return
	}
	res, err := sh.SocietyController.AdminGetByID(id)
	if err != nil {
		log.Printf("[error] failed to get data for response: %v", err)
		response.InternalServerError(w, info)
		return
	}
	res.ID = id
	response.AdminRender(w, "detail.html", info, res)

}

func (sh *societyHandler) AdminDeleteByID(w http.ResponseWriter, r *http.Request) {
	info := createInfo(r, "societies", auth.GetStudentIDFromCookie(r))
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Printf("[error] failed to parse path parameter: %v", err)
		response.InternalServerError(w, info)
		return
	}
	body, err := sh.SocietyController.AdminGetByID(id)
	if err != nil {
		log.Printf("[error] failed to get original data: %v", err)
		response.InternalServerError(w, info)
		return
	}

	if r.Method == "POST" {
		// log.Println("post request: delete society")
		err = sh.SocietyController.DeleteByID(id)
		if err != nil {
			log.Printf("[error] failed to delete: %v", err)
			info.Errors = append(info.Errors, "削除に失敗しました")
			response.AdminRender(w, "delete.html", info, body)
			return
		}
		// log.Println("success to delete society")
		http.Redirect(w, r, "/admin/societies", http.StatusSeeOther)
	}
	response.AdminRender(w, "delete.html", info, body)
}
