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

type activityHandler struct {
	controller.ActivityController
}

// ActivityHandler 活動内容の入出力を受付
type ActivityHandler interface {
	GetActivities(w http.ResponseWriter, r *http.Request)

	// admin
	AdminGetAll(w http.ResponseWriter, r *http.Request)
	AdminGetByID(w http.ResponseWriter, r *http.Request)
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
	info := createInfo(r, "members", auth.GetStudentIDFromCookie(r))
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
	response.AdminRender(w, "detail.html", info, res)
}
