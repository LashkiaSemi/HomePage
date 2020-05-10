package handler

import (
	"homepage/pkg/interface/repository"
)

type AppHandler struct {
	AuthHandler
	UserHandler
	ActivityHandler
	SocietyHandler
	JobHandler
}

func NewAppHandler(sh repository.SQLHandler) *AppHandler {
	return &AppHandler{
		AuthHandler:     NewAuthHandler(),
		UserHandler:     NewUserHandler(sh),
		ActivityHandler: NewActivityHandler(),
		SocietyHandler:  NewSocietyHandler(sh),
		JobHandler:      NewJobHandler(sh),
	}
}
