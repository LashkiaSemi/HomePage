package handler

import (
	"fmt"
	"homepage/pkg/configs"
	"homepage/pkg/infrastructure/auth"
	"homepage/pkg/infrastructure/server/response"
	"homepage/pkg/interface/controller"
	"homepage/pkg/interface/repository"
	"homepage/pkg/usecase/interactor"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

type researchHandler struct {
	controller.ResearchController
}

// ResearchHandler 卒業研究の入出力を受け付け
type ResearchHandler interface {
	GetAll(w http.ResponseWriter, r *http.Request)

	Create(w http.ResponseWriter, r *http.Request)
	UpdateByID(w http.ResponseWriter, r *http.Request)

	// admin
	AdminGetAll(w http.ResponseWriter, r *http.Request)
	AdminGetByID(w http.ResponseWriter, r *http.Request)
}

// NewResearchHandler ハンドラの作成
func NewResearchHandler(sh repository.SQLHandler) ResearchHandler {
	return &researchHandler{
		ResearchController: controller.NewResearchController(
			interactor.NewResearchInteractor(
				repository.NewResearchRepository(sh),
			),
		),
	}
}

func (rh *researchHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	info := createInfo(r, "research", auth.GetStudentIDFromCookie(r))

	res, err := rh.ResearchController.GetAll()
	if err != nil {
		// log.Println("researchHandler: GetAll: ", err)
		response.InternalServerError(w, info)
	}
	response.Success(w, "research/index.html", info, res)
}

func (rh *researchHandler) Create(w http.ResponseWriter, r *http.Request) {
	info := createInfo(r, "researches", auth.GetStudentIDFromCookie(r))

	body := []*FormField{
		createFormField("title", "", "タイトル", "text", nil),
		createFormField("author", "", "著者", "text", nil),
		createFormField("file", "", "ファイル", "file", nil),
		createFormField("comment", "", "コメント", "textarea", nil),
		createFormField("activation", "public", "公開する", "checkbox", nil),
	}

	if r.Method == "POST" {
		log.Println("research update: post request")
		title := r.PostFormValue("title")
		author := r.PostFormValue("author")
		comment := r.PostFormValue("comment")
		var activation int
		if r.PostFormValue("activation") == "public" {
			activation = 1
		} else {
			activation = 0
		}
		if title == "" || author == "" {
			info.Errors = append(info.Errors, "タイトル、著者は必須です")
			response.AdminRender(w, "edit.html", info, body)
			return
		}

		// file
		var fileName string
		file, fileHeader, err := r.FormFile("file")
		defer file.Close()
		if err != nil {
			log.Println("empty file", err)
			fileName = ""
		} else {
			// TODO: funcにしたい
			fileName = fileHeader.Filename
			var saveImage *os.File
			saveImage, err = os.Create(fmt.Sprintf("%s/%s", configs.SaveResearchFileDir, fileName))
			if err != nil {
				log.Println("failed to create file: ", err)
				// TODO: 驚き最小じゃない気がする
				response.InternalServerError(w, info)
				return
			}
			defer saveImage.Close()
			_, err = io.Copy(saveImage, file)
			if err != nil {
				log.Println("failed to save file: ", err)
				// 驚き最小じゃない気がする
				response.InternalServerError(w, info)
				return
			}
			log.Println("complete to save file")
		}
		id, err := rh.ResearchController.Create(title, author, fileName, comment, activation)
		if err != nil {
			log.Println(err)
			response.InternalServerError(w, info)
			return
		}
		log.Println("success create research")
		http.Redirect(w, r, fmt.Sprintf("/admin/researches/%d", id), http.StatusSeeOther)
	}

	response.AdminRender(w, "edit.html", info, body)
}

func (rh *researchHandler) UpdateByID(w http.ResponseWriter, r *http.Request) {
	info := createInfo(r, "researches", auth.GetStudentIDFromCookie(r))
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Println("failed to parse path parameter", err)
		response.InternalServerError(w, info)
		return
	}
	// 初期値の取得
	data, err := rh.ResearchController.GetByID(id)
	if err != nil {
		log.Println("failed to get target: ", err)
		response.InternalServerError(w, info)
		return
	}
	body := []*FormField{
		createFormField("title", data.Title, "タイトル", "text", nil),
		createFormField("author", data.Author, "著者", "text", nil),
		createFormField("file", data.FileName, "ファイル", "file", nil),
		createFormField("comment", data.Comment, "コメント", "textarea", nil),
		createFormField("activation", "public", "公開する", "checkbox", nil),
	}

	if r.Method == "POST" {
		log.Println("research update: post request")
		title := r.PostFormValue("title")
		author := r.PostFormValue("author")
		comment := r.PostFormValue("comment")
		var activation int
		if r.PostFormValue("activation") == "public" {
			activation = 1
		} else {
			activation = 0
		}
		if title == "" || author == "" {
			info.Errors = append(info.Errors, "タイトル、著者は必須です")
			response.AdminRender(w, "edit.html", info, body)
			return
		}

		// file
		var fileName string
		file, fileHeader, err := r.FormFile("file")
		if err != nil {
			log.Println("empty file", err)
			fileName = ""
		} else {
			// TODO: funcにしたい
			fileName = fileHeader.Filename
			var saveImage *os.File
			saveImage, err = os.Create(fmt.Sprintf("%s/%s", configs.SaveResearchFileDir, fileName))
			if err != nil {
				log.Println("failed to create file: ", err)
				// TODO: 驚き最小じゃない気がする
				response.InternalServerError(w, info)
				return
			}
			defer saveImage.Close()
			defer file.Close()
			_, err = io.Copy(saveImage, file)
			if err != nil {
				log.Println("failed to save file: ", err)
				// 驚き最小じゃない気がする
				response.InternalServerError(w, info)
				return
			}
		}
		err = rh.ResearchController.UpdateByID(id, title, author, fileName, comment, activation)
		if err != nil {
			log.Println(err)
			response.InternalServerError(w, info)
			return
		}
		log.Println("success update research")
		http.Redirect(w, r, fmt.Sprintf("/admin/researches/%d", id), http.StatusSeeOther)
	}

	response.AdminRender(w, "edit.html", info, body)

}

// admin
func (rh *researchHandler) AdminGetAll(w http.ResponseWriter, r *http.Request) {
	info := createInfo(r, "researches", auth.GetStudentIDFromCookie(r))
	res, err := rh.ResearchController.AdminGetAll()
	if err != nil {
		log.Println("jobHandler: AdminGetAll: ", err)
		response.InternalServerError(w, info)
		return
	}
	response.AdminRender(w, "list.html", info, res)
}

func (rh *researchHandler) AdminGetByID(w http.ResponseWriter, r *http.Request) {
	info := createInfo(r, "researches", auth.GetStudentIDFromCookie(r))
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Println("researchHandler: AdminGetByID: failed to parse path param: ", err)
		response.InternalServerError(w, info)
		return
	}
	res, err := rh.ResearchController.AdminGetByID(id)
	if err != nil {
		log.Println("researchHandler: AdminGetByID: ", err)
		response.InternalServerError(w, info)
		return
	}
	res.ID = id
	response.AdminRender(w, "detail.html", info, res)
}
