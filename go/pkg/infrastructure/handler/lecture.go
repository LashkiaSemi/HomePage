package handler

import (
	"homepage/pkg/domain/service"
	"homepage/pkg/infrastructure/auth"
	"homepage/pkg/infrastructure/server/response"
	"homepage/pkg/interface/controller"
	"homepage/pkg/interface/repository"
	"homepage/pkg/usecase/interactor"
	"net/http"
)

type lectureHandler struct {
	controller.LectureController
}

// LectureHandler レクチャーの入出力の受付
type LectureHandler interface {
	GetAll(w http.ResponseWriter, r *http.Request)
}

// NewLectureHandler ハンドラの作成
func NewLectureHandler(sh repository.SQLHandler) LectureHandler {
	return &lectureHandler{
		LectureController: controller.NewLectureController(
			interactor.NewLectureInteractor(
				service.NewLectureService(),
				repository.NewLectureRepository(sh),
			),
		),
	}
}

func (lh *lectureHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	info := createInfo(r, "lecture", auth.GetStudentIDFromCookie(r))
	res, err := lh.LectureController.GetAll()
	if err != nil {
		response.InternalServerError(w, info)
		return
	}
	response.Success(w, "lecture/index.html", info, res)
}
