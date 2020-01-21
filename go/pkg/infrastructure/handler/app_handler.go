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
	Logout() http.HandlerFunc

	// user
	ManageUser() http.HandlerFunc
	ManageOneUser() http.HandlerFunc

	// activity
	ManageActivity() http.HandlerFunc
	ManageOneActivity() http.HandlerFunc
}

type appHandler struct {
	AccountHandler
	UserHandler
	ActivityHandler
}

// NewAppHandler アプリケーションハンドラを作成
func NewAppHandler(sh repository.SQLHandler, ah interactor.AuthHandler) AppHandler {
	return &appHandler{
		AccountHandler:  NewAccountHandler(sh, ah),
		UserHandler:     NewUserHandler(sh, ah),
		ActivityHandler: NewActivityHandler(sh),
	}
}

func (ah *appHandler) ManageAccount() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			middleware.Authorized(ah.AccountHandler.GetAccount).ServeHTTP(w, r)
		case http.MethodPost:
			ah.AccountHandler.CreateAccount(w, r)
		case http.MethodPut:
			middleware.Authorized(ah.AccountHandler.UpdateAccount).ServeHTTP(w, r)
		case http.MethodDelete:
			middleware.Authorized(ah.AccountHandler.DeleteAccount).ServeHTTP(w, r)
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

func (ah *appHandler) Logout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodDelete:
			middleware.Authorized(ah.AccountHandler.Logout).ServeHTTP(w, r)
		default:
			logger.Warn("method not allowed")
			response.HTTPError(w, domain.MethodNotAllowed(errors.New("method not allowed")))
		}
	}
}

func (ah *appHandler) ManageUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			ah.UserHandler.GetUsers(w, r)
		case http.MethodPost:
			middleware.Authorized(middleware.Permission(ah.UserHandler.CreateUser)).ServeHTTP(w, r)
		default:
			logger.Warn("method not allowed")
			response.HTTPError(w, domain.MethodNotAllowed(errors.New("method not allowed")))
		}
	}
}

func (ah *appHandler) ManageOneUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			ah.UserHandler.GetUserByUserID(w, r)
		case http.MethodPut:
			middleware.Authorized(middleware.Permission(ah.UserHandler.UpdateUser)).ServeHTTP(w, r)
		case http.MethodDelete:
			middleware.Authorized(middleware.Permission(ah.UserHandler.DeleteUser)).ServeHTTP(w, r)
		default:
			logger.Warn("method not allowed")
			response.HTTPError(w, domain.MethodNotAllowed(errors.New("method not allowed")))
		}
	}
}

func (ah *appHandler) ManageActivity() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			ah.ActivityHandler.GetActivities(w, r)
		case http.MethodPost:
			middleware.Authorized(middleware.Permission(ah.ActivityHandler.CreateActivity)).ServeHTTP(w, r)
		default:
			logger.Warn("method not allowed")
			response.HTTPError(w, domain.MethodNotAllowed(errors.New("method not allowed")))
		}
	}
}

func (ah *appHandler) ManageOneActivity() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			ah.ActivityHandler.GetActivityByID(w, r)
		case http.MethodPut:
			middleware.Authorized(middleware.Permission(ah.ActivityHandler.UpdateActivity)).ServeHTTP(w, r)
		case http.MethodDelete:
			middleware.Authorized(middleware.Permission(ah.ActivityHandler.DeleteActivity)).ServeHTTP(w, r)
		default:
			logger.Warn("method not allowed")
			response.HTTPError(w, domain.MethodNotAllowed(errors.New("method not allowed")))
		}
	}
}
