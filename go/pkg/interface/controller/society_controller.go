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
	ShowAll() (GetSocietiesResponse, error)
	ShowByID(socID int) (GetSocietyResponse, error)
	Create(req *UpdateSocietyRequest) (GetSocietyResponse, error)
	Update(socID int, req *UpdateSocietyRequest) (GetSocietyResponse, error)
	Delete(socID int) error
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

func (sc *societyController) ShowAll() (res GetSocietiesResponse, err error) {
	socs, err := sc.SocietyInteractor.FetchAll()
	if err != nil {
		return
	}

	for _, soc := range socs {
		res.Societies = append(res.Societies, convertSocietyToResponse(&soc))
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
	Title   string `json:"title"`
	Author  string `json:"author"`
	Society string `json:"society"`
	Award   string `json:"award"`
	Date    string `json:"date"`
}

func (sc *societyController) ShowByID(socID int) (res GetSocietyResponse, err error) {
	soc, err := sc.SocietyInteractor.FetchByID(socID)
	if err != nil {
		return
	}
	return convertSocietyToResponse(&soc), nil

}

func (sc *societyController) Create(req *UpdateSocietyRequest) (res GetSocietyResponse, err error) {
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

	soc, err := sc.SocietyInteractor.Add(req.Title, req.Author, req.Society, req.Award, date)
	if err != nil {
		return
	}
	return convertSocietyToResponse(&soc), nil
}

// UpdateSocietyRequest 作成と更新の時のリクエスト
type UpdateSocietyRequest struct {
	Title   string `json:"title"`
	Author  string `json:"author"`
	Society string `json:"society"`
	Award   string `json:"award"`
	Date    string `json:"date"`
}

func (sc *societyController) Update(socID int, req *UpdateSocietyRequest) (res GetSocietyResponse, err error) {
	var date time.Time
	if req.Date != "" {
		date, err = time.Parse(conf.DateFormat, req.Date)
		if err != nil {
			logger.Warn("createSociety: fail to pars. string to date")
			return res, domain.BadRequest(errors.New("fail time parse"))
		}
	}
	soc, err := sc.SocietyInteractor.Update(socID, req.Title, req.Author, req.Society, req.Award, date)
	if err != nil {
		return
	}
	return convertSocietyToResponse(&soc), nil
}

func (sc *societyController) Delete(socID int) error {
	return sc.SocietyInteractor.Delete(socID)
}

func convertSocietyToResponse(soc *domain.Society) GetSocietyResponse {
	return GetSocietyResponse{
		ID:      soc.ID,
		Title:   soc.Title,
		Author:  soc.Author,
		Society: soc.Society,
		Award:   soc.Award,
		Date:    soc.Date,
	}
}
