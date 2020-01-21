package router

import (
	"homepage/pkg/infrastructure/handler"
	"homepage/pkg/infrastructure/server"
)

// SettingRouter urlのハンドリング
func SettingRouter(s server.Server, h handler.AppHandler) {
	s.Handle("/account", h.ManageAccount())
	s.Handle("/login", h.Login())
	s.Handle("/logout", h.Logout())

	s.Handle("/users", h.ManageUser())
	s.Handle("/users/", h.ManageOneUser())

	s.Handle("/activities", h.ManageActivity())
	s.Handle("/activities/", h.ManageOneActivity())
}
