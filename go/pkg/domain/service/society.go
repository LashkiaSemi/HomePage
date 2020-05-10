package service

import "homepage/pkg/domain/model"

type societyService struct{}

type SocietyService interface {
	Create() *model.Society
}

func NewSocietyService() SocietyService {
	return &societyService{}
}

func (ss *societyService) Create() *model.Society {
	return &model.Society{
		Title:   "new",
		Author:  "new",
		Society: "new",
		Award:   "new",
		Date:    "2020/02/02",
	}
}
