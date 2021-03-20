//go:generate mockgen -source=$GOFILE -destination=../../../mock/$GOPACKAGE/$GOFILE -package=mock_$GOPACKAGE -build_flags=-mod=mod
package controller

import (
	"homepage/pkg/domain/entity"
	"homepage/pkg/usecase/interactor"
	"strconv"

	"github.com/pkg/errors"
)

type tagController struct {
	interactor.TagInteractor
}

// TagController タグの入出力を変換
type TagController interface {
	GetAll() (*TagsResponse, error)
	GetByID(id int) (*TagResponse, error)
	Create(name string) (int, error)
	UpdateByID(id int, name string) error
	DeleteByID(id int) error

	AdminGetAll() ([]map[string]string, error)
	AdminGetByID(id int) (*FieldsResponse, error)
}

// NewTagController コントローラの作成
func NewTagController(ti interactor.TagInteractor) TagController {
	return &tagController{
		TagInteractor: ti,
	}
}

func (tc *tagController) GetAll() (*TagsResponse, error) {
	tags, err := tc.TagInteractor.GetAll()
	if err != nil {
		err = errors.Wrap(err, "failed to original data for response")
		return &TagsResponse{}, err
	}
	var res TagsResponse
	for _, tag := range tags {
		res.Tags = append(res.Tags, convertToTagResponse(tag))
	}
	return &res, nil
}

func (tc *tagController) GetByID(id int) (*TagResponse, error) {
	tag, err := tc.TagInteractor.GetByID(id)
	if err != nil {
		err = errors.Wrap(err, "failed to original data for response")
		return &TagResponse{}, err
	}
	return convertToTagResponse(tag), nil
}

func (tc *tagController) Create(name string) (int, error) {
	return tc.TagInteractor.Create(name)
}

func (tc *tagController) UpdateByID(id int, name string) error {
	return tc.TagInteractor.UpdateByID(id, name)
}

func (tc *tagController) DeleteByID(id int) error {
	return tc.TagInteractor.DeleteByID(id)
}

// admin
func (tc *tagController) AdminGetAll() ([]map[string]string, error) {
	var res []map[string]string
	tags, err := tc.TagInteractor.GetAll()
	if err != nil {
		err = errors.Wrap(err, "failed to original data for response")
		return res, err
	}
	for _, tag := range tags {
		res = append(res, map[string]string{
			"id":    strconv.Itoa(tag.ID),
			"title": tag.Name,
		})
	}
	return res, nil
}

func (tc *tagController) AdminGetByID(id int) (*FieldsResponse, error) {
	var res FieldsResponse
	tag, err := tc.TagInteractor.GetByID(id)
	if err != nil {
		err = errors.Wrap(err, "failed to original data for response")
		return &res, err
	}
	res.Fields = append(res.Fields,
		&Field{Key: "ID", Value: tag.ID},
		&Field{Key: "名前", Value: tag.Name},
	)
	res.ID = tag.ID
	return &res, nil
}

// TagsResponse タグのレスポンス
type TagsResponse struct {
	Tags []*TagResponse
}

// TagResponse タグのレスポンス
type TagResponse struct {
	ID        int
	Name      string
	CreatedAt string
}

func convertToTagResponse(tag *entity.Tag) *TagResponse {
	return &TagResponse{
		ID:        tag.ID,
		Name:      tag.Name,
		CreatedAt: tag.CreatedAt,
	}
}
