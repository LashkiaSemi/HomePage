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

type societyHandler struct {
	controller.SocietyController
}

// SocietyHandler 学会発表の入出力の受付
type SocietyHandler interface {
	GetAll(w http.ResponseWriter, r *http.Request)

	// admin
	AdminGetAll(w http.ResponseWriter, r *http.Request)
	AdminGeByID(w http.ResponseWriter, r *http.Request)
}

// NewSocietyHandler ハンドラの作成
func NewSocietyHandler(sh repository.SQLHandler) SocietyHandler {
	// TODO;!!!!
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
		response.InternalServerError(w, info)
		return
	}
	// response
	response.Success(w, "society/index.html", info, res)
}

// admin
func (sh *societyHandler) AdminGetAll(w http.ResponseWriter, r *http.Request) {
	info := createInfo(r, "societies", auth.GetStudentIDFromCookie(r))
	res, err := sh.SocietyController.AdminGetAll()
	if err != nil {
		log.Println("societyHandler: ", err)
		response.InternalServerError(w, info)
		return
	}
	response.AdminRender(w, "list.html", info, res)
}

func (sh *societyHandler) AdminGeByID(w http.ResponseWriter, r *http.Request) {
	info := createInfo(r, "societies", auth.GetStudentIDFromCookie(r))
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Println("societyHandler: AdminGetByID: failed to parse path param: ", err)
		response.InternalServerError(w, info)
		return
	}
	res, err := sh.SocietyController.AdminGetByID(id)
	if err != nil {
		log.Println("societyHandler: AdminGetByID: ", err)
		response.InternalServerError(w, info)
		return
	}
	response.AdminRender(w, "detail.html", info, res)

}
