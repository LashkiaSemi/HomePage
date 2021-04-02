//go:generate mockgen -source=$GOFILE -destination=../../../mock/$GOPACKAGE/$GOFILE -package=mock_$GOPACKAGE -build_flags=-mod=mod
package controller

import (
	"fmt"
	"homepage/pkg/configs"
	"homepage/pkg/domain/entity"
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
	GetByID(id int) (*ResearchResponse, error)

	Create(title, author, file, comment string, activation int) (int, error)
	UpdateByID(id int, title, author, file, comment string, activation int) error

	DeleteByID(id int) error

	// admin
	AdminGetAll() ([]map[string]string, error)
	AdminGetByID(id int) (*FieldsResponse, error)
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
		err = errors.Wrap(err, "failed to original data for response")
		return &ResearchesResponse{}, err
	}
	var res ResearchesResponse
	for _, data := range datas {
		res.Researches = append(res.Researches, convertToResearchResponse(data))
	}
	return &res, nil
}

func (rc *researchController) GetByID(id int) (*ResearchResponse, error) {
	data, err := rc.ResearchInteractor.GetByID(id)
	if err != nil {
		err = errors.Wrap(err, "failed to original data for response")
		return &ResearchResponse{}, err
	}
	return convertToResearchResponse(data), nil
}

func (rc *researchController) Create(title, author, file, comment string, activation int) (int, error) {
	return rc.ResearchInteractor.Create(title, author, file, comment, activation)
}

func (rc *researchController) UpdateByID(id int, title, author, file, comment string, activation int) error {
	return rc.ResearchInteractor.UpdateByID(id, title, author, file, comment, activation)
}

func (rc *researchController) DeleteByID(id int) error {
	return rc.ResearchInteractor.DeleteByID(id)
}

// admin
func (rc *researchController) AdminGetAll() ([]map[string]string, error) {
	var res []map[string]string
	datas, err := rc.ResearchInteractor.GetAll()
	if err != nil {
		err = errors.Wrap(err, "failed to original data for response")
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

func (rc *researchController) AdminGetByID(id int) (*FieldsResponse, error) {
	var res FieldsResponse
	data, err := rc.ResearchInteractor.GetByID(id)
	if err != nil {
		err = errors.Wrap(err, "failed to original data for response")
		return &res, err
	}
	res.Fields = append(res.Fields,
		&Field{Key: "ID", Value: data.ID},
		&Field{Key: "タイトル", Value: data.Title},
		&Field{Key: "著者", Value: data.Author},
		&Field{Key: "ファイル", Value: data.File},
		&Field{Key: "コメント", Value: data.Comment},
		&Field{Key: "公開", Value: data.Activation},
	)
	res.ID = data.ID
	return &res, nil
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
	FileName   string
	FilePath   string
	Comment    string
	Activation int
	CreatedAt  string
}

func convertToResearchResponse(data *entity.Research) *ResearchResponse {
	return &ResearchResponse{
		ID:         data.ID,
		Title:      data.Title,
		Author:     data.Author,
		FileName:   data.File,
		FilePath:   fmt.Sprintf("%s/%s", configs.SaveResearchFileDir, data.File),
		Comment:    data.Comment,
		Activation: data.Activation,
		CreatedAt:  data.CreatedAt,
	}
}
