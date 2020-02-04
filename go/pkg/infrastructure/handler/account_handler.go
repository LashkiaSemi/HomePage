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

// AccountHandler アカウント管理ハンドラ
type AccountHandler interface {
	Get(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)

	Login(w http.ResponseWriter, r *http.Request)
	Logout(w http.ResponseWriter, r *http.Request)
}

type accountHandler struct {
	AccountController controller.AccountController
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

func (ah *accountHandler) Get(w http.ResponseWriter, r *http.Request) {
	// コンテキストからuserIDの取得
	userID := dcontext.GetUserIDFromContext(r.Context())

	res, err := ah.AccountController.ShowByID(userID)
	if err != nil {
		response.HTTPError(w, err)
		return
	}

	// レスポンス
	response.Success(w, res)
}

func (ah *accountHandler) Create(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Warn(err)
		response.HTTPError(w, domain.BadRequest(err))
		return
	}
	var req controller.UpdateAccountRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		logger.Error(err)
		response.HTTPError(w, domain.InternalServerError(err))
		return
	}

	res, err := ah.AccountController.Create(&req)
	if err != nil {
		response.HTTPError(w, err)
		return
	}

	response.Success(w, res)
}

func (ah *accountHandler) Update(w http.ResponseWriter, r *http.Request) {
	userID := dcontext.GetUserIDFromContext(r.Context())

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Warn(err)
		response.HTTPError(w, domain.BadRequest(err))
		return
	}
	var req controller.UpdateAccountRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		logger.Error(err)
		response.HTTPError(w, domain.InternalServerError(err))
		return
	}

	res, err := ah.AccountController.Update(userID, &req)
	if err != nil {
		logger.Error(err)
		response.HTTPError(w, err)
		return
	}

	// session 情報の更新
	sess, err := domain.GetSessionByUserID(userID)
	if err != nil {
		logger.Error(err)
		response.HTTPError(w, err)
		return
	}
	sess.StudentID = res.StudentID
	sess.Role = res.Role

	response.Success(w, res)

}

func (ah *accountHandler) Delete(w http.ResponseWriter, r *http.Request) {
	userID := dcontext.GetUserIDFromContext(r.Context())

	// sessionの削除
	sess, err := session.Store.Get(r, conf.CookieName)
	if err != nil {
		response.HTTPError(w, domain.InternalServerError(err))
		return
	}
	sess.Options.MaxAge = -1
	err = sess.Save(r, w)
	if err != nil {
		response.HTTPError(w, domain.InternalServerError(err))
		return
	}

	err = ah.AccountController.Delete(userID)
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

	// TODO: session.SetSessionDataに移行したい
	sess, err := session.Store.Get(r, conf.CookieName)
	if err != nil {
		response.HTTPError(w, domain.InternalServerError(err))
		return
	}
	sess.Values["sessionID"] = sessData.SessionID
	// sess.Values["studentID"] = sessData.StudentID
	sess.Values["userID"] = sessData.UserID
	err = sess.Save(r, w)
	if err != nil {
		response.HTTPError(w, domain.InternalServerError(err))
		return
	}
	sessData.SetSessionList()

	response.Success(w, res)
}

func (ah *accountHandler) Logout(w http.ResponseWriter, r *http.Request) {
	// session
	sess, err := session.Store.Get(r, conf.CookieName)
	if err != nil {
		response.HTTPError(w, domain.InternalServerError(err))
		return
	}
	sess.Options.MaxAge = -1
	err = sess.Save(r, w)
	if err != nil {
		response.HTTPError(w, domain.InternalServerError(err))
		return
	}
	response.NoContent(w)
}
