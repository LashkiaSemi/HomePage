package controller

import (
	"homepage/pkg/domain"
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
		res.Researches = append(res.Researches, convertResearchToResponse(&data))
	}
	return
}

func (rc *researchController) ShowByID(resID int) (res GetResearchResponse, err error) {
	data, err := rc.ResearchInteractor.FetchByID(resID)
	if err != nil {
		return
	}
	return convertResearchToResponse(&data), nil
}

func (rc *researchController) Create(req *UpdateResearchRequest) (res GetResearchResponse, err error) {
	isPublic := 1
	if !req.IsPublic {
		isPublic = 0
	}
	data, err := rc.ResearchInteractor.Add(req.Title, req.Author, req.File, req.Comment, isPublic)
	if err != nil {
		return
	}
	return convertResearchToResponse(&data), nil

}

func (rc *researchController) Update(resID int, req *UpdateResearchRequest) (res GetResearchResponse, err error) {
	isPublic := 1
	if !req.IsPublic {
		isPublic = 0
	}
	data, err := rc.ResearchInteractor.Update(resID, req.Title, req.Author, req.File, req.Comment, isPublic)
	if err != nil {
		return
	}
	return convertResearchToResponse(&data), nil
}

func (rc *researchController) Delete(resID int) error {
	return rc.ResearchInteractor.Delete(resID)
}

func convertResearchToResponse(res *domain.Research) GetResearchResponse {
	return GetResearchResponse{
		ID:        res.ID,
		Title:     res.Title,
		Author:    res.Author,
		File:      res.File,
		Comment:   res.Comment,
		CreatedAt: res.CreatedAt,
	}
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

// GetResearchesResponse 複数のデータを返す
type GetResearchesResponse struct {
	Researches []GetResearchResponse `json:"researches"`
}

// UpdateResearchRequest 新規、更新時のリクエスト
type UpdateResearchRequest struct {
	Title    string `json:"title"`
	Author   string `json:"author"`
	File     string `json:"file"`
	Comment  string `json:"comment"`
	IsPublic bool   `json:"is_public"`
}
