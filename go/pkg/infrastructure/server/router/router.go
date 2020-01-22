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

	s.Handle("/societies", h.ManageSociety())
	s.Handle("/societies/", h.ManageOneSociety())

	s.Handle("/researches", h.ManageResearch())
	s.Handle("/researches/", h.ManageOneResearch())

	s.Handle("/jobs", h.ManageEmploy())
	s.Handle("/jobs/", h.ManageOneEmploy())
}
