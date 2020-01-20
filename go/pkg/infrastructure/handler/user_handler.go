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
	"strconv"
)

type UserHandler interface {
	GetUsers(w http.ResponseWriter, r *http.Request)
	GetUserByUserID(w http.ResponseWriter, r *http.Request)
	CreateUser(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
}

type userHandler struct {
	UserController controller.UserController
}

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

func (uh *userHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	res, err := uh.UserController.ShowUsers()
	if err != nil {
		response.HTTPError(w, err)
		return
	}

	response.Success(w, res)
}

func (uh *userHandler) GetUserByUserID(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(getParameterInPath(r.URL.Path, "/users/", ""))
	if err != nil {
		logger.Warn("UpdateUser: params error. userID parameter is not int.userID is ", userID)
		response.HTTPError(w, domain.BadRequest(errors.New("params error")))
		return
	}

	res, err := uh.UserController.ShowUserByUserID(userID)
	if err != nil {
		response.HTTPError(w, err)
		return
	}
	response.Success(w, res)
}

func (uh *userHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
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

	res, err := uh.UserController.CreateUser(&req)
	if err != nil {
		response.HTTPError(w, err)
		return
	}

	response.Success(w, res)
}

func (uh *userHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(getParameterInPath(r.URL.Path, "/users/", ""))
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

	res, err := uh.UserController.UpdateUser(userID, &req)
	if err != nil {
		response.HTTPError(w, err)
		return
	}

	response.Success(w, res)
}

func (uh *userHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(getParameterInPath(r.URL.Path, "/users/", ""))
	if err != nil {
		logger.Warn("UpdateUser: params error. userID parameter is not int. userID is ", userID)
		response.HTTPError(w, domain.BadRequest(errors.New("params error")))
		return
	}

	err = uh.UserController.DeleteUser(userID)
	if err != nil {
		response.HTTPError(w, err)
		return
	}

	response.NoContent(w)

}
