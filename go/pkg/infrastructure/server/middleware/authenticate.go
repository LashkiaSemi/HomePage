package middleware

import (
	"context"
	"errors"
	"homepage/conf"
	"homepage/pkg/domain"
	"homepage/pkg/domain/logger"
	"homepage/pkg/infrastructure/dcontext"
	"homepage/pkg/infrastructure/server/response"
	"homepage/pkg/infrastructure/server/session"
	"net/http"
)

// Authorized sessionから認証を行う
func Authorized(nextFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sessList := domain.GetSessionList()
		// sessionListが空っぽの時。サーバを再起動した時とか
		if len(sessList) == 0 {
			logger.Warn("sessionList is empty.")
			response.HTTPError(w, domain.InternalServerError(errors.New("sessionList is empty. please re-login")))
			return
		}

		ctx := r.Context()
		if ctx == nil {
			ctx = context.Background()
		}

		// cookieからsessionを取得
		sess, err := session.Store.Get(r, conf.CookieName)
		if err != nil {
			logger.Warn(err)
			response.HTTPError(w, domain.Unauthorized(err))
			return
		}

		// studentIDを取得
		// studentID, ok := sess.Values["studentID"].(string)
		// if !ok || studentID == "" {
		// 	logger.Warn("studentID is empty")
		// 	response.HTTPError(w, domain.Unauthorized(errors.New("studentID is empty")))
		// 	return
		// }
		userID, ok := sess.Values["userID"].(int)
		if !ok || userID == 0 {
			logger.Warn("middleware authenticate: userID is empty")
			response.HTTPError(w, domain.Unauthorized(errors.New("userID is empty")))
			return
		}

		// sessionを確認
		if sessList[userID].SessionID != sess.Values["sessionID"].(string) {
			logger.Warn("wrong sessionID")
			response.HTTPError(w, domain.Unauthorized(errors.New("wrong sessionID")))
			return
		}

		// contextにuserIDを保存
		// ctx = dcontext.SetStudentID(ctx, studentID)
		ctx = dcontext.SetUserID(ctx, userID)

		// nextfnc
		nextFunc(w, r.WithContext(ctx))
	}
}

// Permission ownerの確認。Authorizedより後に
func Permission(nextFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sessList := domain.GetSessionList()
		// sessionListが空っぽの時。サーバを再起動した時とか
		// if len(sessList) == 0 {
		// 	logger.Warn("sessionList is empty.")
		// 	response.HTTPError(w, domain.InternalServerError(errors.New("sessionList is empty. please re-login")))
		// 	return
		// }

		userID := dcontext.GetUserIDFromContext(r.Context())

		// roleの取得
		if sessList[userID].Role != "owner" {
			logger.Warn("permission error.")
			response.HTTPError(w, domain.Unauthorized(errors.New("permission error")))
			return
		}

		nextFunc(w, r)
	}
}
