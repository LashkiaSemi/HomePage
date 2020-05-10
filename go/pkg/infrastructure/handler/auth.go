package handler

import (
	"homepage/pkg/configs"
	"net/http"
)

type AuthHandler interface {
	Login(w http.ResponseWriter, r *http.Request)
	// Logout() http.HandlerFunc
}

type authHandler struct {
}

func NewAuthHandler() AuthHandler {
	return &authHandler{}
}

func (ah *authHandler) Login(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:  configs.CookieName,
		Value: "studentID",
	}
	http.SetCookie(w, cookie)
}

// func (ah *authHandler) Logout() http.HandlerFunc {

// }
