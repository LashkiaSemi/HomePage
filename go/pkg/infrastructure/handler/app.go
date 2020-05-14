package handler

import (
	"homepage/pkg/interface/repository"
)

// AppHandler アプリケーション全体のハンドラ
type AppHandler struct {
	StaticPageHandler
	UserHandler
	ActivityHandler
	SocietyHandler
	ResearchHandler
	JobHandler
	LectureHandler
	EquipmentHandler
	TagHandler
}

// NewAppHandler ハンドラの作成
func NewAppHandler(sh repository.SQLHandler) *AppHandler {
	return &AppHandler{
		StaticPageHandler: NewStaticPageHandler(sh),
		UserHandler:       NewUserHandler(sh),
		ActivityHandler:   NewActivityHandler(sh),
		SocietyHandler:    NewSocietyHandler(sh),
		ResearchHandler:   NewResearchHandler(sh),
		JobHandler:        NewJobHandler(sh),
		LectureHandler:    NewLectureHandler(sh),
		EquipmentHandler:  NewEquipmentHandler(sh),
		TagHandler:        NewTagHandler(sh),
	}
}
