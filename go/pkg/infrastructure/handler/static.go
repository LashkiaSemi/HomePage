package handler

import (
	"homepage/pkg/infrastructure/auth"
	"homepage/pkg/infrastructure/server/response"
	"homepage/pkg/interface/controller"
	"homepage/pkg/interface/repository"
	"homepage/pkg/usecase/interactor"
	"log"
	"net/http"
)

type staticPageHandler struct {
	controller.ActivityController
}

// StaticPageHandler 静的ページに一時的な情報を埋め込む。staticとは？
type StaticPageHandler interface {
	IndexHandler(w http.ResponseWriter, r *http.Request)
	AdminIndexHandler(w http.ResponseWriter, r *http.Request)
}

// NewStaticPageHandler ハンドラの作成
func NewStaticPageHandler(sh repository.SQLHandler) StaticPageHandler {
	return &staticPageHandler{
		ActivityController: controller.NewActivityController(
			interactor.NewActivityInteractor(
				repository.NewActivityRepository(sh),
			),
		),
	}
}

// IndexHandler home部分のhandler
func (sh *staticPageHandler) IndexHandler(w http.ResponseWriter, r *http.Request) {
	info := createInfo(r, "", auth.GetStudentIDFromCookie(r))
	body, err := sh.ActivityController.GetUpcoming()
	if err != nil {
		log.Printf("[error] failed to get data for response: %v", err)
		response.InternalServerError(w, info)
		return
	}
	response.Render(w, "index.html", info, body)
}

// AdminIndexHandler adminのトップページ
func (sh *staticPageHandler) AdminIndexHandler(w http.ResponseWriter, r *http.Request) {
	info := createInfo(r, "", auth.GetStudentIDFromCookie(r))
	var body interface{}
	response.AdminRender(w, "index.html", info, body)
}
