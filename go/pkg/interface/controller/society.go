package controller

import (
	// TODO: domainに...

	"homepage/pkg/entity"
	"homepage/pkg/usecase/interactor"
	"log"
	"strconv"

	"github.com/pkg/errors"
)

type societyController struct {
	interactor.SocietyInteractor
}

// SocietyController 学会発表の入出力を変換
type SocietyController interface {
	GetAll() (*SocietiesResponse, error)

	AdminGetAll() ([]map[string]string, error)
}

// NewSocietyController コントローラの作成
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

// admin
func (sc *societyController) AdminGetAll() ([]map[string]string, error) {
	var res []map[string]string
	socs, err := sc.SocietyInteractor.GetAll()
	if err != nil {
		err = errors.Wrap(err, "AdminGetAll")
		return res, err
	}
	for _, soc := range socs {
		res = append(res, map[string]string{
			"id":    strconv.Itoa(soc.ID),
			"title": soc.Title,
		})
	}
	return res, nil
}

// SocietiesResponse 学会発表の複数件分
type SocietiesResponse struct {
	Societies []*SocietyResponse
}

// SocietyResponse 学会発表一件分
type SocietyResponse struct {
	Title   string
	Author  string
	Society string
	Award   string
	Date    string
}

func convertToSocietyResponse(soc *entity.Society) *SocietyResponse {
	return &SocietyResponse{
		Title:   soc.Title,
		Author:  soc.Author,
		Society: soc.Society,
		Award:   soc.Award,
		Date:    soc.Date,
	}
}
