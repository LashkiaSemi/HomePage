package handler

import (
	"homepage/pkg/interface/repository"
)

// AppHandler アプリケーション全体のハンドラ
type AppHandler struct {
	UserHandler
	ActivityHandler
	SocietyHandler
	ResearchHandler
	JobHandler
	LectureHandler
}

// NewAppHandler ハンドラの作成
func NewAppHandler(sh repository.SQLHandler) *AppHandler {
	return &AppHandler{
		UserHandler:     NewUserHandler(sh),
		ActivityHandler: NewActivityHandler(),
		SocietyHandler:  NewSocietyHandler(sh),
		ResearchHandler: NewResearchHandler(sh),
		JobHandler:      NewJobHandler(sh),
		LectureHandler:  NewLectureHandler(sh),
	}
}
