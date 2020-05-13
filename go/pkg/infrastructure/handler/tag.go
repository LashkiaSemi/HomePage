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

type tagHandler struct {
	controller.TagController
}

// TagHandler タグ関連のリクエストを受けてレスポンスする
type TagHandler interface {
	// GetAll(w http.ResponseWriter, r *http.Request)
	// GetByID(w http.ResponseWriter, r *http.Request)

	Create(w http.ResponseWriter, r *http.Request)
	UpdateByID(w http.ResponseWriter, r *http.Request)

	AdminGetAll(w http.ResponseWriter, r *http.Request)
	AdminGetByID(w http.ResponseWriter, r *http.Request)
	AdminDeleteByID(w http.ResponseWriter, r *http.Request)
}

// NewTagHandler ハンドラの作成
func NewTagHandler(sh repository.SQLHandler) TagHandler {
	return &tagHandler{
		TagController: controller.NewTagController(
			interactor.NewTagInteractor(
				repository.NewTagRepository(sh),
			),
		),
	}
}

func (th *tagHandler) AdminGetAll(w http.ResponseWriter, r *http.Request) {
	info := createInfo(r, "tags", auth.GetStudentIDFromCookie(r))
	body, err := th.TagController.AdminGetAll()
	if err != nil {
		log.Println("tagHandler: AdminGetAll: ", err)
		response.InternalServerError(w, info)
		return
	}
	response.AdminRender(w, "list.html", info, body)
}

func (th *tagHandler) AdminGetByID(w http.ResponseWriter, r *http.Request) {
	info := createInfo(r, "tags", auth.GetStudentIDFromCookie(r))
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Println("tagHandler: AdminGetByID: failed to parse path param: ", err)
		response.InternalServerError(w, info)
		return
	}
	body, err := th.TagController.AdminGetByID(id)
	if err != nil {
		log.Println("tagHandler: AdminGetByID: ", err)
		response.InternalServerError(w, info)
		return
	}
	response.AdminRender(w, "detail.html", info, body)

}

func (th *tagHandler) Create(w http.ResponseWriter, r *http.Request) {
	info := createInfo(r, "tags", auth.GetStudentIDFromCookie(r))

	body := []*FormField{
		createFormField("name", "", "名前", "text", nil),
	}

	if r.Method == "POST" {
		log.Println("tag create: post request")
		name := r.PostFormValue("name")
		if name == "" {
			info.Errors = append(info.Errors, "名前は必須です")
			response.AdminRender(w, "edit.html", info, body)
			return
		}

		id, err := th.TagController.Create(name)
		if err != nil {
			log.Println(err)
			response.InternalServerError(w, info)
			return
		}
		log.Println("success to create tag")
		http.Redirect(w, r, fmt.Sprintf("/admin/tags/%d", id), http.StatusSeeOther)
	}

	response.AdminRender(w, "edit.html", info, body)
}

func (th *tagHandler) UpdateByID(w http.ResponseWriter, r *http.Request) {
	info := createInfo(r, "tags", auth.GetStudentIDFromCookie(r))
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Println("failed to parse path parameter", err)
		response.InternalServerError(w, info)
		return
	}
	data, err := th.TagController.GetByID(id)
	if err != nil {
		log.Println("failed to get target: ", err)
		response.InternalServerError(w, info)
		return
	}
	body := []*FormField{
		createFormField("name", data.Name, "名前", "text", nil),
	}

	if r.Method == "POST" {
		log.Println("tag update: post request")
		name := r.PostFormValue("name")
		if name == "" {
			info.Errors = append(info.Errors, "名前は必須です")
			response.AdminRender(w, "edit.html", info, body)
			return
		}

		err = th.TagController.UpdateByID(id, name)
		if err != nil {
			log.Println(err)
			response.InternalServerError(w, info)
			return
		}
		log.Println("success update tag")
		http.Redirect(w, r, fmt.Sprintf("/admin/tags/%d", id), http.StatusSeeOther)
	}

	response.AdminRender(w, "edit.html", info, body)
}

func (th *tagHandler) AdminDeleteByID(w http.ResponseWriter, r *http.Request) {
	info := createInfo(r, "tags", auth.GetStudentIDFromCookie(r))
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Println("failed to parse path parameter", err)
		response.InternalServerError(w, info)
		return
	}
	body, err := th.TagController.AdminGetByID(id)
	if err != nil {
		log.Println("AdminDeleteByID: ", err)
		response.InternalServerError(w, info)
		return
	}

	if r.Method == "POST" {
		log.Println("post request: delete tag")
		err = th.TagController.DeleteByID(id)
		if err != nil {
			log.Println("failed to delete")
			info.Errors = append(info.Errors, "削除に失敗しました")
			response.AdminRender(w, "delete.html", info, body)
			return
		}
		log.Println("success to delete tag")
		http.Redirect(w, r, "/admin/tags", http.StatusSeeOther)
	}
	response.AdminRender(w, "delete.html", info, body)
}
