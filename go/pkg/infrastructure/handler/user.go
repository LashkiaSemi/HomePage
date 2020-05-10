package handler

import (
	"homepage/pkg/configs"
	"homepage/pkg/domain/service"
	"homepage/pkg/infrastructure/auth"
	"homepage/pkg/infrastructure/server/response"
	"homepage/pkg/interface/controller"
	"homepage/pkg/interface/repository"
	"homepage/pkg/usecase/interactor"
	"log"
	"net/http"
)

type userHandler struct {
	controller.UserController
}

type UserHandler interface {
	GetAllGroupByGrade(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
}

func NewUserHandler(sh repository.SQLHandler) UserHandler {
	return &userHandler{
		UserController: controller.NewUserContoroller(
			interactor.NewUserInteractor(
				service.NewUserService(),
				repository.NewUserRepository(sh),
				auth.NewVerifyHandler(),
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

func (uh *userHandler) Login(w http.ResponseWriter, r *http.Request) {
	info := createInfo(r, "login")
	var body interface{}

	if r.Method == "POST" {
		studentID := r.PostFormValue("studentID")
		password := r.PostFormValue("password")

		if studentID == "" || password == "" {
			response.Success(w, "login.html", info, body)
			return
		}

		err := uh.UserController.Login(studentID, password)
		if err != nil {
			// TODO: ろぐいんしっぱいじの
			log.Println("failed to login: ", err)
			response.Success(w, "login.html", info, body)
			return
		}

		cookie := &http.Cookie{
			Name:  configs.CookieName,
			Value: studentID,
		}
		http.SetCookie(w, cookie)
		log.Println("redirect")
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	response.Success(w, "login.html", info, body)
}
