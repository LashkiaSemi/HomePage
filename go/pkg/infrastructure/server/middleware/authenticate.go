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

func Authorized(nextFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		if ctx == nil {
			ctx = context.Background()
		}

		// cookieからtokenを取得
		sess, err := session.Store.Get(r, conf.CookieName)
		if err != nil {
			logger.Warn(err)
			response.HTTPError(w, domain.Unauthorized(err))
			return
		}

		studentID := sess.Values["studentID"]
		if studentID == nil {
			logger.Warn("studentID is empty")
			response.HTTPError(w, domain.Unauthorized(errors.New("studentID is empty")))
			return
		}
		sid, _ := studentID.(string)

		// TODO: 一回サーバが止まってsessionListが吹っ飛ぶと、ログインできないね
		sessList := domain.GetSessionList()
		if sessList[sid] != sess.Values["sessionID"].(string) {
			logger.Warn("wrong sessionID")
			response.HTTPError(w, domain.Unauthorized(errors.New("wrong sessionID")))
			return
		}

		// contextにuserIDを保存
		ctx = dcontext.SetStudentID(ctx, sid)

		// nextfnc
		nextFunc(w, r.WithContext(ctx))
	}
}
