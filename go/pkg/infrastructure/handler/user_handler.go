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

// UserHandler ハンドラ
type UserHandler interface {
	GetAll(w http.ResponseWriter, r *http.Request)
	GetByID(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type userHandler struct {
	UserController controller.UserController
}

// NewUserHandler ハンドラの作成
func NewUserHandler(sh repository.SQLHandler, ah interactor.AuthHandler) UserHandler {
	return &userHandler{
		UserController: controller.NewUserController(
			interactor.NewUserInteractor(
				repository.NewUserRepository(sh),
				ah,
			),
		),
	}
}

func (uh *userHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	res, err := uh.UserController.ShowAll()
	if err != nil {
		response.HTTPError(w, err)
		return
	}

	response.Success(w, res)
}

func (uh *userHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	userID, err := getIntParameter(r.URL.Path, "/users/", "")
	if err != nil {
		logger.Warn("UpdateUser: params error. userID parameter is not int.userID is ", userID)
		response.HTTPError(w, domain.BadRequest(errors.New("params error")))
		return
	}

	res, err := uh.UserController.ShowByID(userID)
	if err != nil {
		response.HTTPError(w, err)
		return
	}
	response.Success(w, res)
}

func (uh *userHandler) Create(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Warn(err)
		response.HTTPError(w, domain.BadRequest(err))
		return
	}
	var req controller.UpdateUserRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		logger.Error(err)
		response.HTTPError(w, domain.InternalServerError(err))
		return
	}

	res, err := uh.UserController.Create(&req)
	if err != nil {
		response.HTTPError(w, err)
		return
	}

	response.Success(w, res)
}

func (uh *userHandler) Update(w http.ResponseWriter, r *http.Request) {
	userID, err := getIntParameter(r.URL.Path, "/users/", "")
	if err != nil {
		logger.Warn("UpdateUser: params error. userID parameter is not int.userID is ", userID)
		response.HTTPError(w, domain.BadRequest(errors.New("params error")))
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Warn(err)
		response.HTTPError(w, domain.BadRequest(err))
		return
	}
	var req controller.UpdateUserRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		logger.Error(err)
		response.HTTPError(w, domain.InternalServerError(err))
		return
	}

	res, err := uh.UserController.Update(userID, &req)
	if err != nil {
		response.HTTPError(w, err)
		return
	}

	response.Success(w, res)
}

func (uh *userHandler) Delete(w http.ResponseWriter, r *http.Request) {
	userID, err := getIntParameter(r.URL.Path, "/users/", "")
	if err != nil {
		logger.Warn("UpdateUser: params error. userID parameter is not int. userID is ", userID)
		response.HTTPError(w, domain.BadRequest(errors.New("params error")))
		return
	}

	err = uh.UserController.Delete(userID)
	if err != nil {
		response.HTTPError(w, err)
		return
	}

	response.NoContent(w)
}
