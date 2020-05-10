package service

import "homepage/pkg/domain/model"

type societyService struct{}

// SocietyService 学会発表のドメインロジックを実装
type SocietyService interface {
	Create() *model.Society
}

// NewSocietyService サービスを作成
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
