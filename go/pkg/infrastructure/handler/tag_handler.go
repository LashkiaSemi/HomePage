package handler

import (
	"encoding/json"
	"errors"
	"homepage/pkg/domain"
	"homepage/pkg/domain/logger"
	"homepage/pkg/infrastructure/server/response"
	"homepage/pkg/interface/controller"
	"homepage/pkg/interface/repository"
	"homepage/pkg/usecase/interactor"
	"io/ioutil"
	"net/http"
)

// TagHandler ハンドラ
type TagHandler interface {
	GetAll(w http.ResponseWriter, r *http.Request)
	GetByID(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type tagHandler struct {
	TagController controller.TagController
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

func (th *tagHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	res, err := th.TagController.ShowAll()
	if err != nil {
		response.HTTPError(w, err)
		return
	}
	response.Success(w, res)
}

func (th *tagHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	tagID, err := getIntParameter(r.URL.Path, "/tags/", "")
	if err != nil {
		logger.Warn("tag getByID handler: can not get tagID from path")
		response.HTTPError(w, domain.BadRequest(errors.New("cant get param from path")))
		return
	}
	res, err := th.TagController.ShowByID(tagID)
	if err != nil {
		response.HTTPError(w, err)
		return
	}
	response.Success(w, res)
}

func (th *tagHandler) Create(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Warn(err)
		response.HTTPError(w, domain.BadRequest(err))
		return
	}
	var req controller.UpdateTagRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		logger.Error(err)
		response.HTTPError(w, domain.InternalServerError(err))
		return
	}

	res, err := th.TagController.Create(&req)
	if err != nil {
		response.HTTPError(w, err)
		return
	}
	response.Success(w, res)
}

func (th *tagHandler) Update(w http.ResponseWriter, r *http.Request) {
	tagID, err := getIntParameter(r.URL.Path, "/tags/", "")
	if err != nil {
		logger.Warn("update tag handler: can not get tagID from path")
		response.HTTPError(w, domain.BadRequest(errors.New("cant get param from path")))
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Warn(err)
		response.HTTPError(w, domain.BadRequest(err))
		return
	}
	var req controller.UpdateTagRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		logger.Error(err)
		response.HTTPError(w, domain.InternalServerError(err))
		return
	}

	res, err := th.TagController.Update(tagID, &req)
	if err != nil {
		response.HTTPError(w, err)
		return
	}
	response.Success(w, res)
}

func (th *tagHandler) Delete(w http.ResponseWriter, r *http.Request) {
	tagID, err := getIntParameter(r.URL.Path, "/tags/", "")
	if err != nil {
		logger.Warn("delete tag handler: can not get tagID from path")
		response.HTTPError(w, domain.BadRequest(errors.New("cant get param from path")))
		return
	}
	err = th.TagController.Delete(tagID)
	if err != nil {
		response.HTTPError(w, err)
		return
	}
	response.NoContent(w)
}
