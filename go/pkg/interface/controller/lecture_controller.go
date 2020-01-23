package controller

import "homepage/pkg/usecase/interactor"

import "homepage/pkg/domain"

// LectureController コントローラ
type LectureController interface {
	ShowAll() (GetLecturesResponse, error)
	ShowByID(lecID int) (GetLectureResponse, error)
	Create(userID int, req *UpdateLectureRequest) (GetLectureResponse, error)
	Update(lecID, userID int, req *UpdateLectureRequest) (GetLectureResponse, error)
	Delete(lecID int) error
}

type lectureController struct {
	LectureInteractor interactor.LectureInteractor
}

// NewLectureController コントローラの作成
func NewLectureController(li interactor.LectureInteractor) LectureController {
	return &lectureController{
		LectureInteractor: li,
	}
}

func (lc *lectureController) ShowAll() (res GetLecturesResponse, err error) {
	lecs, err := lc.LectureInteractor.FetchAll()
	if err != nil {
		return
	}
	for _, lec := range lecs {
		res.Lectures = append(res.Lectures, convertLectureToResponse(&lec))
	}
	return
}

func (lc *lectureController) ShowByID(lecID int) (res GetLectureResponse, err error) {
	lec, err := lc.LectureInteractor.FetchByID(lecID)
	if err != nil {
		return
	}
	return convertLectureToResponse(&lec), nil

}

func (lc *lectureController) Create(userID int, req *UpdateLectureRequest) (res GetLectureResponse, err error) {
	lec, err := lc.LectureInteractor.Add(req.Title, req.File, req.Comment, userID)
	if err != nil {
		return
	}
	return convertLectureToResponse(&lec), nil

}

func (lc *lectureController) Update(lecID, userID int, req *UpdateLectureRequest) (res GetLectureResponse, err error) {
	lec, err := lc.LectureInteractor.Update(lecID, req.Title, req.File, req.Comment, userID)
	if err != nil {
		return
	}
	return convertLectureToResponse(&lec), nil
}

func (lc *lectureController) Delete(lecID int) error {
	return lc.LectureInteractor.Delete(lecID)
}

// GetLecturesResponse 複数件レスポンス
type GetLecturesResponse struct {
	Lectures []GetLectureResponse `json:"lectures"`
}

// GetLectureResponse 一件
type GetLectureResponse struct {
	ID      int             `json:"id"`
	Title   string          `json:"title"`
	File    string          `json:"file"`
	Comment string          `json:"comment"`
	User    GetUserResponse `json:"user"`
}

// UpdateLectureRequest 新規、更新リクエスト
type UpdateLectureRequest struct {
	Title   string `json:"title"`
	File    string `json:"file"`
	Comment string `json:"comment"`
}

func convertLectureToResponse(lec *domain.Lecture) GetLectureResponse {
	return GetLectureResponse{
		ID:      lec.ID,
		Title:   lec.Title,
		File:    lec.File,
		Comment: lec.Comment,
		User:    convertUserToResponse(&lec.User),
	}
}
