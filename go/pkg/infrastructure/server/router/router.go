package router

import (
	"homepage/pkg/infrastructure/handler"
	"homepage/pkg/infrastructure/server"
)

func SettingRouter(s server.Server, h handler.AppHandler) {
	s.Handle("/account", h.ManageAccount())
	s.Handle("/login", h.Login())
	s.Handle("/logout", h.Logout())
}
