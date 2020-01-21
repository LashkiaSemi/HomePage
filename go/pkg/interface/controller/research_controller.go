package controller

import (
	"homepage/pkg/usecase/interactor"
)

// ResearchController コントローラ
type ResearchController interface {
	ShowAll() (GetResearchesResponse, error)
	ShowByID(resID int) (GetResearchResponse, error)
	Create(req *UpdateResearchRequest) (GetResearchResponse, error)
	Update(resID int, req *UpdateResearchRequest) (GetResearchResponse, error)
	Delete(resID int) error
}

type researchController struct {
	ResearchInteractor interactor.ResearchInteractor
}

// NewResearchController コントローラの作成
func NewResearchController(ri interactor.ResearchInteractor) ResearchController {
	return &researchController{
		ResearchInteractor: ri,
	}
}

func (rc *researchController) ShowAll() (res GetResearchesResponse, err error) {
	datas, err := rc.ResearchInteractor.FetchAll()
	for _, data := range datas {
		res.Researches = append(res.Researches, GetResearchResponse{
			ID:        data.ID,
			Title:     data.Title,
			Author:    data.Author,
			File:      data.File,
			Comment:   data.Comment,
			CreatedAt: data.CreatedAt,
		})
	}
	return
}

// GetResearchesResponse 複数のデータを返す
type GetResearchesResponse struct {
	Researches []GetResearchResponse `json:"researches"`
}

// GetResearchResponse 卒研のデータ
type GetResearchResponse struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	File      string `json:"file"`
	Comment   string `json:"comment"`
	CreatedAt string `json:"created_at"`
}

func (rc *researchController) ShowByID(resID int) (res GetResearchResponse, err error) {
	data, err := rc.ResearchInteractor.FetchByID(resID)
	if err != nil {
		return
	}
	res.ID = data.ID
	res.Title = data.Title
	res.Author = data.Author
	res.File = data.File
	res.Comment = data.Comment
	res.CreatedAt = data.CreatedAt
	return
}

func (rc *researchController) Create(req *UpdateResearchRequest) (res GetResearchResponse, err error) {
	data, err := rc.ResearchInteractor.Add(req.Title, req.Author, req.File, req.Comment)
	if err != nil {
		return
	}
	res.ID = data.ID
	res.Title = data.Title
	res.Author = data.Author
	res.File = data.File
	res.Comment = data.Comment
	res.CreatedAt = data.CreatedAt
	return
}

// UpdateResearchRequest 新規、更新時のリクエスト
type UpdateResearchRequest struct {
	Title   string `json:"title"`
	Author  string `json:"author"`
	File    string `json:"file"`
	Comment string `json:"comment"`
}

func (rc *researchController) Update(resID int, req *UpdateResearchRequest) (res GetResearchResponse, err error) {

	data, err := rc.ResearchInteractor.Update(resID, req.Title, req.Author, req.File, req.Comment)
	if err != nil {
		return
	}
	res.ID = data.ID
	res.Title = data.Title
	res.Author = data.Author
	res.File = data.File
	res.Comment = data.Comment
	res.CreatedAt = data.CreatedAt
	return
}

func (rc *researchController) Delete(resID int) error {
	return rc.ResearchInteractor.Delete(resID)
}
