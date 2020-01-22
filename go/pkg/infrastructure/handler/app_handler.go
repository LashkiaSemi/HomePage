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

	// society
	ManageSociety() http.HandlerFunc
	ManageOneSociety() http.HandlerFunc

	// research
	ManageResearch() http.HandlerFunc
	ManageOneResearch() http.HandlerFunc

	// employment
	ManageEmploy() http.HandlerFunc
	ManageOneEmploy() http.HandlerFunc

	// equipment
	ManageEquipment() http.HandlerFunc
	ManageOneEquipment() http.HandlerFunc

	// lecture
	ManageLecture() http.HandlerFunc
	ManageOneLecture() http.HandlerFunc

	// tag
	ManageTag() http.HandlerFunc
	ManageOneTag() http.HandlerFunc
}

type appHandler struct {
	AccountHandler
	UserHandler
	ActivityHandler
	SocietyHandler
	ResearchHandler
	EmployHandler
	EquipmentHandler
	LectureHandler
	TagHandler
}

// NewAppHandler アプリケーションハンドラを作成
func NewAppHandler(sh repository.SQLHandler, ah interactor.AuthHandler) AppHandler {
	return &appHandler{
		AccountHandler:   NewAccountHandler(sh, ah),
		UserHandler:      NewUserHandler(sh, ah),
		ActivityHandler:  NewActivityHandler(sh),
		SocietyHandler:   NewSocietyHandler(sh),
		ResearchHandler:  NewResearchHandler(sh),
		EmployHandler:    NewEmployHandler(sh),
		EquipmentHandler: NewEquipmentHandler(sh),
		LectureHandler:   NewLectureHandler(sh),
		TagHandler:       NewTagHandler(sh),
	}
}

// Account
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

// Session
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

// User
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

// Activity
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

// Society
func (ah *appHandler) ManageSociety() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			ah.SocietyHandler.GetSocieties(w, r)
		case http.MethodPost:
			middleware.Authorized(middleware.Permission(ah.SocietyHandler.CreateSociety)).ServeHTTP(w, r)
		default:
			logger.Warn("method not allowed")
			response.HTTPError(w, domain.MethodNotAllowed(errors.New("method not allowed")))
		}
	}
}

func (ah *appHandler) ManageOneSociety() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			ah.SocietyHandler.GetSocietyByID(w, r)
		case http.MethodPut:
			middleware.Authorized(middleware.Permission(ah.SocietyHandler.UpdateSociety)).ServeHTTP(w, r)
		case http.MethodDelete:
			middleware.Authorized(middleware.Permission(ah.SocietyHandler.DeleteSociety)).ServeHTTP(w, r)
		default:
			logger.Warn("method not allowed")
			response.HTTPError(w, domain.MethodNotAllowed(errors.New("method not allowed")))
		}
	}
}

// research
func (ah *appHandler) ManageResearch() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			ah.ResearchHandler.GetAll(w, r)
		case http.MethodPost:
			middleware.Authorized(middleware.Permission(ah.ResearchHandler.Create)).ServeHTTP(w, r)
		default:
			logger.Warn("method not allowed")
			response.HTTPError(w, domain.MethodNotAllowed(errors.New("method not allowed")))
		}
	}
}

func (ah *appHandler) ManageOneResearch() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			ah.ResearchHandler.GetByID(w, r)
		case http.MethodPut:
			middleware.Authorized(middleware.Permission(ah.ResearchHandler.Update)).ServeHTTP(w, r)
		case http.MethodDelete:
			middleware.Authorized(middleware.Permission(ah.ResearchHandler.Delete)).ServeHTTP(w, r)
		default:
			logger.Warn("method not allowed")
			response.HTTPError(w, domain.MethodNotAllowed(errors.New("method not allowed")))
		}
	}
}

// employment
func (ah *appHandler) ManageEmploy() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			ah.EmployHandler.GetAll(w, r)
		case http.MethodPost:
			middleware.Authorized(middleware.Permission(ah.EmployHandler.Create)).ServeHTTP(w, r)
		default:
			logger.Warn("method not allowed")
			response.HTTPError(w, domain.MethodNotAllowed(errors.New("method not allowed")))
		}
	}
}

func (ah *appHandler) ManageOneEmploy() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			ah.EmployHandler.GetByID(w, r)
		case http.MethodPut:
			middleware.Authorized(middleware.Permission(ah.EmployHandler.Update)).ServeHTTP(w, r)
		case http.MethodDelete:
			middleware.Authorized(middleware.Permission(ah.EmployHandler.Delete)).ServeHTTP(w, r)
		default:
			logger.Warn("method not allowed")
			response.HTTPError(w, domain.MethodNotAllowed(errors.New("method not allowed")))
		}
	}
}

// equipment
func (ah *appHandler) ManageEquipment() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			ah.EquipmentHandler.GetAll(w, r)
		case http.MethodPost:
			middleware.Authorized(middleware.Permission(ah.EquipmentHandler.Create)).ServeHTTP(w, r)
		default:
			logger.Warn("method not allowed")
			response.HTTPError(w, domain.MethodNotAllowed(errors.New("method not allowed")))
		}
	}
}

func (ah *appHandler) ManageOneEquipment() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			ah.EquipmentHandler.GetByID(w, r)
		case http.MethodPut:
			middleware.Authorized(middleware.Permission(ah.EquipmentHandler.Update)).ServeHTTP(w, r)
		case http.MethodDelete:
			middleware.Authorized(middleware.Permission(ah.EquipmentHandler.Delete)).ServeHTTP(w, r)
		default:
			logger.Warn("method not allowed")
			response.HTTPError(w, domain.MethodNotAllowed(errors.New("method not allowed")))
		}
	}
}

// lecture
func (ah *appHandler) ManageLecture() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			ah.LectureHandler.GetAll(w, r)
		case http.MethodPost:
			middleware.Authorized(ah.LectureHandler.Create).ServeHTTP(w, r)
		default:
			logger.Warn("method not allowed")
			response.HTTPError(w, domain.MethodNotAllowed(errors.New("method not allowed")))
		}
	}
}

func (ah *appHandler) ManageOneLecture() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			ah.LectureHandler.GetByID(w, r)
		case http.MethodPut:
			middleware.Authorized(ah.LectureHandler.Update).ServeHTTP(w, r)
		case http.MethodDelete:
			middleware.Authorized(ah.LectureHandler.Delete).ServeHTTP(w, r)
		default:
			logger.Warn("method not allowed")
			response.HTTPError(w, domain.MethodNotAllowed(errors.New("method not allowed")))
		}
	}
}

// tag
func (ah *appHandler) ManageTag() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			ah.TagHandler.GetAll(w, r)
		case http.MethodPost:
			middleware.Authorized(middleware.Permission(ah.TagHandler.Create)).ServeHTTP(w, r)
		default:
			logger.Warn("method not allowed")
			response.HTTPError(w, domain.MethodNotAllowed(errors.New("method not allowed")))
		}
	}
}

func (ah *appHandler) ManageOneTag() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			ah.TagHandler.GetByID(w, r)
		case http.MethodPut:
			middleware.Authorized(middleware.Permission(ah.TagHandler.Update)).ServeHTTP(w, r)
		case http.MethodDelete:
			middleware.Authorized(middleware.Permission(ah.TagHandler.Delete)).ServeHTTP(w, r)
		default:
			logger.Warn("method not allowed")
			response.HTTPError(w, domain.MethodNotAllowed(errors.New("method not allowed")))
		}
	}
}
