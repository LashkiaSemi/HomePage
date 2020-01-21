package controller

import (
	"errors"
	"homepage/conf"
	"homepage/pkg/domain"
	"homepage/pkg/domain/logger"
	"homepage/pkg/usecase/interactor"
	"time"
)

// SocietyController コントローラ
type SocietyController interface {
	ShowSocieties() (GetSocietiesResponse, error)
	ShowSocietyByID(socID int) (GetSocietyResponse, error)
	CreateSociety(req *UpdateSocietyRequest) (GetSocietyResponse, error)
	UpdateSociety(socID int, req *UpdateSocietyRequest) (GetSocietyResponse, error)
	DeleteSociety(socID int) error
}

type societyController struct {
	SocietyInteractor interactor.SocietyInteractor
}

// NewSocietyController コントローラの作成
func NewSocietyController(si interactor.SocietyInteractor) SocietyController {
	return &societyController{
		SocietyInteractor: si,
	}
}

func (sc *societyController) ShowSocieties() (res GetSocietiesResponse, err error) {
	socs, err := sc.SocietyInteractor.FetchSocieties()
	if err != nil {
		return
	}

	for _, soc := range socs {
		res.Societies = append(res.Societies, GetSocietyResponse{
			ID:      soc.ID,
			Title:   soc.Title,
			Author:  soc.Author,
			Society: soc.Society,
			Award:   soc.Award,
			Date:    soc.Date,
		})
	}
	return
}

// GetSocietiesResponse 学会発表の配列レスポンス
type GetSocietiesResponse struct {
	Societies []GetSocietyResponse `json:"societies"`
}

// GetSocietyResponse 学会発表のレスポンス
type GetSocietyResponse struct {
	ID      int    `json:"id"`
	Title   string `json:"string"`
	Author  string `json:"author"`
	Society string `json:"society"`
	Award   string `json:"award"`
	Date    string `json:"date"`
}

func (sc *societyController) ShowSocietyByID(socID int) (res GetSocietyResponse, err error) {
	soc, err := sc.SocietyInteractor.FetchSocietyByID(socID)
	if err != nil {
		return
	}
	res.ID = soc.ID
	res.Title = soc.Title
	res.Author = soc.Author
	res.Society = soc.Society
	res.Award = soc.Award
	res.Date = soc.Date
	return
}

func (sc *societyController) CreateSociety(req *UpdateSocietyRequest) (res GetSocietyResponse, err error) {
	if req.Society == "" {
		logger.Warn("createSociety: society is empty")
		return res, domain.BadRequest(errors.New("society is empty"))
	}
	if req.Date == "" {
		logger.Warn("createSociety: date is empty")
		return res, domain.BadRequest(errors.New("date is empty"))
	}

	date, err := time.Parse(conf.DateFormat, req.Date)
	if err != nil {
		logger.Warn("createSociety: fail to pars. string to date")
		return res, domain.BadRequest(errors.New("fail time parse"))
	}

	soc, err := sc.SocietyInteractor.AddSociety(req.Title, req.Author, req.Society, req.Award, date)
	res.ID = soc.ID
	res.Title = soc.Title
	res.Author = soc.Author
	res.Society = soc.Society
	res.Award = soc.Award
	res.Date = soc.Date
	return
}

// UpdateSocietyRequest 作成と更新の時のリクエスト
type UpdateSocietyRequest struct {
	Title   string `json:"title"`
	Author  string `json:"author"`
	Society string `json:"society"`
	Award   string `json:"award"`
	Date    string `json:"date"`
}

func (sc *societyController) UpdateSociety(socID int, req *UpdateSocietyRequest) (res GetSocietyResponse, err error) {
	var date time.Time
	if req.Date != "" {
		date, err = time.Parse(conf.DateFormat, req.Date)
		if err != nil {
			logger.Warn("createSociety: fail to pars. string to date")
			return res, domain.BadRequest(errors.New("fail time parse"))
		}
	}
	soc, err := sc.SocietyInteractor.UpdateSociety(socID, req.Title, req.Author, req.Society, req.Award, date)
	res.ID = soc.ID
	res.Title = soc.Title
	res.Author = soc.Author
	res.Society = soc.Society
	res.Award = soc.Award
	res.Date = soc.Date
	return
}

func (sc *societyController) DeleteSociety(socID int) error {
	return sc.SocietyInteractor.DeleteSociety(socID)
}
