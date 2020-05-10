package handler

import (
	"homepage/pkg/interface/repository"
)

type AppHandler struct {
	AuthHandler
	ActivityHandler
	SocietyHandler
}

func NewAppHandler(sh repository.SQLHandler) *AppHandler {
	return &AppHandler{
		AuthHandler:     NewAuthHandler(),
		ActivityHandler: NewActivityHandler(),
		SocietyHandler:  NewSocietyHandler(sh),
	}
}
