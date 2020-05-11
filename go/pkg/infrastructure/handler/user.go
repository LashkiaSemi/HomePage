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

// UserHandler 入力と出力の受付
type UserHandler interface {
	GetAllGroupByGrade(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	Logout(w http.ResponseWriter, r *http.Request)
}

// NewUserHandler ハンドラの作成
func NewUserHandler(sh repository.SQLHandler) UserHandler {
	return &userHandler{
		UserController: controller.NewUserController(
			interactor.NewUserInteractor(
				service.NewUserService(),
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
