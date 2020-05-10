package middleware

import (
	"context"
	"homepage/pkg/configs"
	"homepage/pkg/infrastructure/authentication"
	"homepage/pkg/infrastructure/dcontext"
	"log"
	"net/http"
)

// Authorized ログイン済みを検証する
func Authorized(nextFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context()
		if ctx == nil {
			ctx = context.Background()
		}

		// cookieからjwtを取得
		cookie, err := r.Cookie(configs.CookieName)
		if err != nil {
			log.Println("Cookie: ", err)
			return
		}
		tokenString := cookie.Value

		// jwtの検証
		token, err := authentication.VerifyToken(tokenString)
		if err != nil {
			log.Println("failed to verify token: ", err)
			return
		}

		// cxtにstudentIdを書き込み
		dcontext.SetStudentID(ctx, authentication.GetStudentIDFromJWT(token))

		// nextfnc
		nextFunc(w, r.WithContext(ctx))
	}
}
