package handler

import (
	"homepage/pkg/domain/service"
	"homepage/pkg/infrastructure/server/response"
	"homepage/pkg/interface/controller"
	"homepage/pkg/interface/repository"
	"homepage/pkg/usecase/interactor"
	"net/http"
)

type userHandler struct {
	controller.UserController
}

type UserHandler interface {
	GetAllGroupByGrade(w http.ResponseWriter, r *http.Request)
}

func NewUserHandler(sh repository.SQLHandler) UserHandler {
	return &userHandler{
		UserController: controller.NewUserContoroller(
			interactor.NewUserInteractor(
				service.NewUserService(),
				repository.NewUserRepository(sh),
			),
		),
	}
}

func (uh *userHandler) GetAllGroupByGrade(w http.ResponseWriter, r *http.Request) {
	info := createInfo(r, "user")

	res, err := uh.UserController.GetAllGroupByGrade()
	if err != nil {
		response.InternalServerError(w, info)
		return
	}

	response.Success(w, "member/index.html", info, res)
}
