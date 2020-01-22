package controller

import (
	"homepage/pkg/usecase/interactor"
)

// TagController コントローラ
type TagController interface {
	ShowAll() (GetTagsResponse, error)
	ShowByID(tagID int) (GetTagResponse, error)
	Create(req *UpdateTagRequest) (GetTagResponse, error)
	Update(tagID int, req *UpdateTagRequest) (GetTagResponse, error)
	Delete(tagID int) error
}

type tagController struct {
	TagInteractor interactor.TagInteractor
}

// NewTagController コントローラの作成
func NewTagController(ti interactor.TagInteractor) TagController {
	return &tagController{
		TagInteractor: ti,
	}
}

func (tc *tagController) ShowAll() (res GetTagsResponse, err error) {
	tags, err := tc.TagInteractor.FetchAll()
	if err != nil {
		return
	}
	for _, tag := range tags {
		res.Tags = append(res.Tags, GetTagResponse{
			ID:   tag.ID,
			Name: tag.Name,
		})
	}
	return
}

func (tc *tagController) ShowByID(tagID int) (res GetTagResponse, err error) {
	tag, err := tc.TagInteractor.FetchByID(tagID)
	if err != nil {
		return
	}
	res.ID = tag.ID
	res.Name = tag.Name
	return
}

func (tc *tagController) Create(req *UpdateTagRequest) (res GetTagResponse, err error) {
	tag, err := tc.TagInteractor.Add(req.Name)
	if err != nil {
		return
	}
	res.ID = tag.ID
	res.Name = tag.Name
	return
}

func (tc *tagController) Update(tagID int, req *UpdateTagRequest) (res GetTagResponse, err error) {
	tag, err := tc.TagInteractor.Update(tagID, req.Name)
	if err != nil {
		return
	}
	res.ID = tag.ID
	res.Name = tag.Name
	return
}

func (tc *tagController) Delete(tagID int) error {
	return tc.TagInteractor.Delete(tagID)
}

// GetTagsResponse 複数のレスポンス
type GetTagsResponse struct {
	Tags []GetTagResponse `json:"tags"`
}

// GetTagResponse 一件
type GetTagResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// UpdateTagRequest 新規、更新のリクエスト
type UpdateTagRequest struct {
	Name string `json:"name"`
}
