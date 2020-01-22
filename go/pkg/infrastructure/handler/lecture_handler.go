package handler

import (
	"encoding/json"
	"errors"
	"homepage/pkg/domain"
	"homepage/pkg/domain/logger"
	"homepage/pkg/infrastructure/dcontext"
	"homepage/pkg/infrastructure/server/response"
	"homepage/pkg/interface/controller"
	"homepage/pkg/interface/repository"
	"homepage/pkg/usecase/interactor"
	"io/ioutil"
	"net/http"
)

// LectureHandler ハンドラ
type LectureHandler interface {
	GetAll(w http.ResponseWriter, r *http.Request)
	GetByID(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type lectureHandler struct {
	LectureController controller.LectureController
}

// NewLectureHandler ハンドラの作成
func NewLectureHandler(sh repository.SQLHandler) LectureHandler {
	return &lectureHandler{
		LectureController: controller.NewLectureController(
			interactor.NewLectureInteractor(
				repository.NewLectureHandler(sh),
			),
		),
	}
}

func (lh *lectureHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	res, err := lh.LectureController.ShowAll()
	if err != nil {
		response.HTTPError(w, err)
		return
	}
	response.Success(w, res)
}

func (lh *lectureHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	lecID, err := getIntParameter(r.URL.Path, "/lectures/", "")
	if err != nil {
		logger.Warn("getByID lecture handler: can not get lectureID from path")
		response.HTTPError(w, domain.BadRequest(errors.New("cant get param from path")))
		return
	}

	res, err := lh.LectureController.ShowByID(lecID)
	if err != nil {
		response.HTTPError(w, err)
		return
	}
	response.Success(w, res)
}

func (lh *lectureHandler) Create(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Warn(err)
		response.HTTPError(w, domain.BadRequest(err))
		return
	}
	var req controller.UpdateLectureRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		logger.Error(err)
		response.HTTPError(w, domain.InternalServerError(err))
		return
	}

	userID := dcontext.GetUserIDFromContext(r.Context())

	res, err := lh.LectureController.Create(userID, &req)
	if err != nil {
		response.HTTPError(w, err)
		return
	}
	response.Success(w, res)
}

func (lh *lectureHandler) Update(w http.ResponseWriter, r *http.Request) {
	lecID, err := getIntParameter(r.URL.Path, "/lectures/", "")
	if err != nil {
		logger.Warn("update lecture handler: can not get lectureID from path")
		response.HTTPError(w, domain.BadRequest(errors.New("cant get param from path")))
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Warn(err)
		response.HTTPError(w, domain.BadRequest(err))
		return
	}
	var req controller.UpdateLectureRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		logger.Error(err)
		response.HTTPError(w, domain.InternalServerError(err))
		return
	}

	userID := dcontext.GetUserIDFromContext(r.Context())

	res, err := lh.LectureController.Update(lecID, userID, &req)
	if err != nil {
		response.HTTPError(w, err)
		return
	}
	response.Success(w, res)
}

func (lh *lectureHandler) Delete(w http.ResponseWriter, r *http.Request) {
	lecID, err := getIntParameter(r.URL.Path, "/lectures/", "")
	if err != nil {
		logger.Warn("update lecture handler: can not get lectureID from path")
		response.HTTPError(w, domain.BadRequest(errors.New("cant get param from path")))
		return
	}
	err = lh.LectureController.Delete(lecID)
	if err != nil {
		response.HTTPError(w, err)
		return
	}
	response.NoContent(w)
}
