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
			log.Println("delete cookie")
			cookie.MaxAge = -1
			http.SetCookie(w, cookie)
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		// cxtにstudentIdを書き込み
		dcontext.SetStudentID(ctx, auth.GetStudentIDFromJWT(token))

		// nextfnc
		nextFunc(w, r.WithContext(ctx))
	}
}

// AdminAuthorized adminのログイン済みを検証
func AdminAuthorized(nextFunc http.HandlerFunc) http.HandlerFunc {
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
			http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
			return
		}
		tokenString := cookie.Value

		// jwtの検証
		token, err := auth.VerifyToken(tokenString)
		if err != nil {
			log.Println("failed to verify token: ", err)
			log.Println("delete cookie")
			cookie.MaxAge = -1
			http.SetCookie(w, cookie)
			http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
			return
		}

		err = auth.CheckIsAdminSession(auth.GetStudentIDFromJWT(token), tokenString)
		if err != nil {
			log.Println("failed to check permission: ", err)
			http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
			return
		}

		// cxtにstudentIdを書き込み
		dcontext.SetStudentID(ctx, auth.GetStudentIDFromJWT(token))

		// nextfnc
		nextFunc(w, r.WithContext(ctx))
	}
}
