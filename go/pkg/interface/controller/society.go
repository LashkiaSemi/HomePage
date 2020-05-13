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
	GetByID(id int) (*SocietyResponse, error)

	Create(title, author, society, award, date string) (int, error)
	UpdateByID(id int, title, author, society, award, date string) error

	DeleteByID(id int) error

	// admin
	AdminGetAll() ([]map[string]string, error)
	AdminGetByID(id int) (*FieldsResponse, error)
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

func (sc *societyController) GetByID(id int) (*SocietyResponse, error) {
	data, err := sc.SocietyInteractor.GetByID(id)
	if err != nil {
		err = errors.Wrap(err, "failed to find data")
		return &SocietyResponse{}, err
	}
	return convertToSocietyResponse(data), nil
}

func (sc *societyController) Create(title, author, society, award, date string) (int, error) {
	id, err := sc.SocietyInteractor.Create(title, author, society, award, date)
	if err != nil {
		err = errors.Wrap(err, "controller")
	}
	return id, err
}

func (sc *societyController) UpdateByID(id int, title, author, society, award, date string) error {
	err := sc.SocietyInteractor.UpdateByID(id, title, author, society, award, date)
	if err != nil {
		err = errors.Wrap(err, "controller")
	}
	return err
}

func (sc *societyController) DeleteByID(id int) error {
	return sc.SocietyInteractor.DeleteByID(id)
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

// admin
func (sc *societyController) AdminGetByID(id int) (*FieldsResponse, error) {
	var res FieldsResponse
	data, err := sc.SocietyInteractor.GetByID(id)
	if err != nil {
		err = errors.Wrap(err, "societyController: AdminGetByID")
		return &res, err
	}
	res.Fields = append(res.Fields,
		&Field{Key: "ID", Value: data.ID},
		&Field{Key: "タイトル", Value: data.Title},
		&Field{Key: "氏名", Value: data.Author},
		&Field{Key: "学会", Value: data.Society},
		&Field{Key: "受賞", Value: data.Award},
		&Field{Key: "日付", Value: data.Date},
	)
	res.ID = data.ID
	return &res, nil
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
