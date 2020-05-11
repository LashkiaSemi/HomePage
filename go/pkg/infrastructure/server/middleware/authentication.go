package middleware

import (
	"context"
	"homepage/pkg/configs"
	"homepage/pkg/infrastructure/auth"
	"homepage/pkg/infrastructure/dcontext"
	"log"
	"net/http"
)

// Authorized ログイン済みを検証する
func Authorized(nextFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: nextでredirectできるといいよね

		ctx := r.Context()
		if ctx == nil {
			ctx = context.Background()
		}

		// cookieからjwtを取得
		cookie, err := r.Cookie(configs.CookieName)
		if err != nil {
			log.Println("Cookie: ", err)
			// cookieがない時
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		tokenString := cookie.Value

		// jwtの検証
		token, err := auth.VerifyToken(tokenString)
		if err != nil {
			log.Println("failed to verify token: ", err)
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		// cxtにstudentIdを書き込み
		dcontext.SetStudentID(ctx, auth.GetStudentIDFromJWT(token))

		// nextfnc
		nextFunc(w, r.WithContext(ctx))
	}
}
