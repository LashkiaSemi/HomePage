package handler

import (
	"errors"
	"homepage/pkg/domain"
	"homepage/pkg/domain/logger"
	"homepage/pkg/infrastructure/server/middleware"
	"homepage/pkg/infrastructure/server/response"
	"homepage/pkg/interface/repository"
	"homepage/pkg/usecase/interactor"
	"net/http"
)

// AppHandler アプリケーションハンドラ
type AppHandler interface {
	// account
	ManageAccount() http.HandlerFunc
	Login() http.HandlerFunc
}

type appHandler struct {
	AccountHandler
}

// NewAppHandler アプリケーションハンドラを作成
func NewAppHandler(sh repository.SQLHandler, ah interactor.AuthHandler) AppHandler {
	return &appHandler{
		AccountHandler: NewAccountHandler(sh, ah),
	}
}

func (ah *appHandler) ManageAccount() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			middleware.Authorized(ah.AccountHandler.GetAccount).ServeHTTP(w, r)
		case http.MethodPost:
			ah.AccountHandler.CreateAccount(w, r)
		default:
			logger.Warn("method not allowed")
			response.HTTPError(w, domain.MethodNotAllowed(errors.New("method not allowed")))
		}
	}
}

func (ah *appHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			ah.AccountHandler.Login(w, r)
		default:
			logger.Warn("method not allowed")
			response.HTTPError(w, domain.MethodNotAllowed(errors.New("method not allowed")))
		}
	}
}
