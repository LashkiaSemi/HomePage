package handler

import (
	"encoding/json"
	"homepage/conf"
	"homepage/pkg/domain"
	"homepage/pkg/domain/logger"
	"homepage/pkg/infrastructure/dcontext"
	"homepage/pkg/infrastructure/server/response"
	"homepage/pkg/infrastructure/server/session"
	"homepage/pkg/interface/controller"
	"homepage/pkg/interface/repository"
	"homepage/pkg/usecase/interactor"
	"io/ioutil"
	"net/http"
)

type accountHandler struct {
	AccountController controller.AccountController
}

// AccountHandler アカウント管理ハンドラ
type AccountHandler interface {
	GetAccount(w http.ResponseWriter, r *http.Request)
	CreateAccount(w http.ResponseWriter, r *http.Request)
	UpdateAccount(w http.ResponseWriter, r *http.Request)
	DeleteAccount(w http.ResponseWriter, r *http.Request)

	Login(w http.ResponseWriter, r *http.Request)
}

// NewAccountHandler accountHandlerを作成
func NewAccountHandler(sh repository.SQLHandler, ah interactor.AuthHandler) AccountHandler {
	return &accountHandler{
		AccountController: controller.NewAccountController(
			interactor.NewAccountInteractor(
				repository.NewAccountRepository(sh),
				ah,
			),
		),
	}
}

func (ah *accountHandler) GetAccount(w http.ResponseWriter, r *http.Request) {
	// コンテキストからstudentIDの取得
	studentID := dcontext.GetStudentIDFromContext(r.Context())

	res, err := ah.AccountController.ShowAccountByStudentID(studentID)
	if err != nil {
		response.HTTPError(w, err)
		return
	}

	// レスポンス
	response.Success(w, res)
}

func (ah *accountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	logger.Debug("create account")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Warn(err)
		response.HTTPError(w, domain.BadRequest(err))
		return
	}
	var req controller.UpdateAccoutRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		logger.Error(err)
		response.HTTPError(w, domain.InternalServerError(err))
		return
	}

	res, err := ah.AccountController.CreateAccount(&req)
	if err != nil {
		response.HTTPError(w, err)
		return
	}

	response.Success(w, res)
}

func (ah *accountHandler) UpdateAccount(w http.ResponseWriter, r *http.Request) {
	logger.Debug("update account")
	// TODO: 実装して
	// studentID := dcontext.GetStudentIDFromContext(r.Context())
	userID := dcontext.GetUserIDFromContext(r.Context())

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Warn(err)
		response.HTTPError(w, domain.BadRequest(err))
		return
	}
	var req controller.UpdateAccoutRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		logger.Error(err)
		response.HTTPError(w, domain.InternalServerError(err))
		return
	}

	res, err := ah.AccountController.UpdateAccount(userID, &req)
	if err != nil {
		logger.Error(err)
		response.HTTPError(w, err)
		return
	}

	response.Success(w, res)

}

func (ah *accountHandler) DeleteAccount(w http.ResponseWriter, r *http.Request) {
	logger.Debug("delete account")
	// TODO: 実装して
	userID := dcontext.GetUserIDFromContext(r.Context())

	err := ah.AccountController.DeleteAccount(userID)
	if err != nil {
		response.HTTPError(w, err)
		return
	}

	response.NoContent(w)
}

func (ah *accountHandler) Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Warn(err)
		response.HTTPError(w, domain.BadRequest(err))
		return
	}
	var req controller.LoginRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		logger.Error(err)
		response.HTTPError(w, domain.InternalServerError(err))
		return
	}

	res, sessData, err := ah.AccountController.Login(&req)
	if err != nil {
		response.HTTPError(w, err)
		return
	}

	// session
	sess, err := session.Store.Get(r, conf.CookieName)
	if err != nil {
		response.HTTPError(w, domain.InternalServerError(err))
		return
	}
	sess.Values["sessionID"] = sessData.SessionID
	sess.Values["studentID"] = sessData.StudentID
	sess.Values["userID"] = sessData.UserID
	err = sess.Save(r, w)
	if err != nil {
		response.HTTPError(w, domain.InternalServerError(err))
		return
	}
	sessData.SetSessionList()

	response.Success(w, res)
}
