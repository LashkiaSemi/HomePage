package handler

import (
	"fmt"
	"homepage/pkg/configs"
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

type userHandler struct {
	controller.UserController
}

// UserHandler 入力と出力の受付
type UserHandler interface {
	GetAllGroupByGrade(w http.ResponseWriter, r *http.Request)
	GetByID(w http.ResponseWriter, r *http.Request)
	UpdateByID(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	Logout(w http.ResponseWriter, r *http.Request)
}

// NewUserHandler ハンドラの作成
func NewUserHandler(sh repository.SQLHandler) UserHandler {
	return &userHandler{
		UserController: controller.NewUserController(
			interactor.NewUserInteractor(
				repository.NewUserRepository(sh),
				auth.NewVerifyHandler(),
			),
		),
	}
}

func (uh *userHandler) GetAllGroupByGrade(w http.ResponseWriter, r *http.Request) {
	info := createInfo(r, "user", auth.GetStudentIDFromCookie(r))

	res, err := uh.UserController.GetAllGroupByGrade()
	if err != nil {
		response.InternalServerError(w, info)
		return
	}

	response.Success(w, "member/index.html", info, res)
}

func (uh *userHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	info := createInfo(r, "user", auth.GetStudentIDFromCookie(r))

	userID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Println("userHandler: GetByID: path param parse error: ", err)
		response.InternalServerError(w, info)
		return
	}

	res, err := uh.UserController.GetByID(userID)
	if err != nil {
		response.InternalServerError(w, info)
		return
	}
	response.Success(w, "member/detail.html", info, res)
}

func (uh *userHandler) UpdateByID(w http.ResponseWriter, r *http.Request) {
	info := createInfo(r, "user", auth.GetStudentIDFromCookie(r))
	// userID := mux.Vars(r)["id"]
	body, err := uh.UserController.GetByStudentID(info.StudentID)
	if err != nil {
		response.InternalServerError(w, info)
		return
	}
	if r.Method == "POST" {
		log.Println("user update: post")
		name := r.PostFormValue("name")
		studentID := r.PostFormValue("studentID")
		department := r.PostFormValue("department")
		comment := r.PostFormValue("comment")
		grade, err := strconv.Atoi(r.PostFormValue("grade"))
		if err != nil {
			// TODO: handling
			log.Println("int parse error")
		}
		// TODO: バリデーション!
		if name == "" || studentID == "" {
			log.Println("空地だめ!")
			response.Success(w, "member/edit.html", info, body)
			return
		}

		user, err := uh.UserController.UpdateByID(body.ID, name, studentID, department, comment, grade)
		if err != nil {
			response.InternalServerError(w, info)
			return
		}
		log.Println("user update: ", user)
		http.Redirect(w, r, fmt.Sprintf("/members/%d", body.ID), http.StatusSeeOther)
	}
	response.Success(w, "member/edit.html", info, body)

}

func (uh *userHandler) Login(w http.ResponseWriter, r *http.Request) {
	info := createInfo(r, "login", auth.GetStudentIDFromCookie(r))
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

		// jwtの作成
		token, err := auth.CreateToken(studentID)
		if err != nil {
			log.Println("failed to create token: ", err)
			response.InternalServerError(w, info)
			return
		}

		cookie := &http.Cookie{
			Name:  configs.CookieName,
			Value: token,
		}
		http.SetCookie(w, cookie)
		log.Println("redirect")
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	response.Success(w, "login.html", info, body)
}

func (uh *userHandler) Logout(w http.ResponseWriter, r *http.Request) {
	info := createInfo(r, "logout", auth.GetStudentIDFromCookie(r))

	// cookieの取得
	cookie, err := r.Cookie(configs.CookieName)
	if err != nil {
		log.Println("Cookie: ", err)
		// TODO: ?こここれでいいのか？
		response.InternalServerError(w, info)
		return
	}
	cookie.MaxAge = -1
	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/", http.StatusSeeOther)

}
