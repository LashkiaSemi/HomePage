package controller

import (
	"homepage/pkg/entity"
	"homepage/pkg/usecase/interactor"
	"strconv"

	"github.com/pkg/errors"
)

type researchController struct {
	interactor.ResearchInteractor
}

// ResearchController 卒業研究の入出力を変換
type ResearchController interface {
	GetAll() (*ResearchesResponse, error)

	// admin
	AdminGetAll() ([]map[string]string, error)
}

// NewResearchController コントローラの作成
func NewResearchController(ri interactor.ResearchInteractor) ResearchController {
	return &researchController{
		ResearchInteractor: ri,
	}
}

func (rc *researchController) GetAll() (*ResearchesResponse, error) {
	datas, err := rc.ResearchInteractor.GetAll()
	if err != nil {
		return &ResearchesResponse{}, err
	}
	var res ResearchesResponse
	for _, data := range datas {
		res.Researches = append(res.Researches, convertToResearchResponse(data))
	}
	return &res, nil
}

// admin
func (rc *researchController) AdminGetAll() ([]map[string]string, error) {
	var res []map[string]string
	datas, err := rc.ResearchInteractor.GetAll()
	if err != nil {
		err = errors.Wrap(err, "AdminGetAll")
		return res, err
	}
	for _, data := range datas {
		res = append(res, map[string]string{
			"id":    strconv.Itoa(data.ID),
			"title": data.Title,
		})
	}
	return res, nil
}

// ResearchesResponse 卒業研究のレスポンス
type ResearchesResponse struct {
	Researches []*ResearchResponse
}

// ResearchResponse 卒業研究のレスポンス
type ResearchResponse struct {
	ID         int
	Title      string
	Author     string
	File       string
	Comment    string
	Activation int
	CreatedAt  string
}

func convertToResearchResponse(data *entity.Research) *ResearchResponse {
	return &ResearchResponse{
		ID:         data.ID,
		Title:      data.Title,
		Author:     data.Author,
		File:       data.File,
		Comment:    data.Comment,
		Activation: data.Activation,
		CreatedAt:  data.CreatedAt,
	}
}
