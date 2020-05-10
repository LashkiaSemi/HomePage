package controller

import (
	// TODO: domainに...

	"homepage/pkg/domain/model"
	"homepage/pkg/usecase/interactor"
	"log"
)

type societyController struct {
	interactor.SocietyInteractor
}

type SocietyController interface {
	GetAll() (*SocietiesResponse, error)
}

func NewSocietyController(si interactor.SocietyInteractor) SocietyController {
	return &societyController{
		SocietyInteractor: si,
	}
}

func (sc *societyController) GetAll() (*SocietiesResponse, error) {

	datas, err := sc.SocietyInteractor.GetAll()
	if err != nil {
		log.Println("err", err)
	}

	var res SocietiesResponse
	for _, soc := range datas {
		res.Societies = append(res.Societies, convertToSocietyResponse(soc))
	}
	return &res, nil

}

type SocietiesResponse struct {
	Societies []*SocietyResponse
}

type SocietyResponse struct {
	Title   string
	Author  string
	Society string
	Award   string
	Date    string
}

func convertToSocietyResponse(soc *model.Society) *SocietyResponse {
	return &SocietyResponse{
		Title:   soc.Title,
		Author:  soc.Author,
		Society: soc.Society,
		Award:   soc.Award,
		Date:    soc.Date,
	}
}
