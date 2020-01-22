package controller

import "homepage/pkg/usecase/interactor"

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
		res.Lectures = append(res.Lectures, GetLectureResponse{
			ID:      lec.ID,
			Title:   lec.Title,
			File:    lec.File,
			Comment: lec.Comment,
			User: GetUserResponse{
				ID:         lec.User.ID,
				Name:       lec.User.Name,
				StudentID:  lec.User.StudentID,
				Department: lec.User.Department,
				Grade:      lec.User.Grade,
				Comment:    lec.User.Comment,
				Role:       lec.User.Role,
			},
		})
	}
	return
}

func (lc *lectureController) ShowByID(lecID int) (res GetLectureResponse, err error) {
	lec, err := lc.LectureInteractor.FetchByID(lecID)
	if err != nil {
		return
	}
	res.ID = lec.ID
	res.Title = lec.Title
	res.File = lec.File
	res.Comment = lec.Comment
	res.User.ID = lec.User.ID
	res.User.Name = lec.User.Name
	res.User.StudentID = lec.User.StudentID
	res.User.Role = lec.User.Role
	res.User.Department = lec.User.Department
	res.User.Grade = lec.User.Grade
	res.User.Comment = lec.User.Comment
	return
}

func (lc *lectureController) Create(userID int, req *UpdateLectureRequest) (res GetLectureResponse, err error) {
	lec, err := lc.LectureInteractor.Add(req.Title, req.File, req.Comment, userID)
	if err != nil {
		return
	}
	res.ID = lec.ID
	res.Title = lec.Title
	res.File = lec.File
	res.Comment = lec.Comment
	res.User.ID = lec.User.ID
	res.User.Name = lec.User.Name
	res.User.StudentID = lec.User.StudentID
	res.User.Role = lec.User.Role
	res.User.Department = lec.User.Department
	res.User.Grade = lec.User.Grade
	res.User.Comment = lec.User.Comment
	return
}

func (lc *lectureController) Update(lecID, userID int, req *UpdateLectureRequest) (res GetLectureResponse, err error) {
	lec, err := lc.LectureInteractor.Update(lecID, req.Title, req.File, req.Comment, userID)
	if err != nil {
		return
	}
	res.ID = lec.ID
	res.Title = lec.Title
	res.File = lec.File
	res.Comment = lec.Comment
	res.User.ID = lec.User.ID
	res.User.Name = lec.User.Name
	res.User.StudentID = lec.User.StudentID
	res.User.Role = lec.User.Role
	res.User.Department = lec.User.Department
	res.User.Grade = lec.User.Grade
	res.User.Comment = lec.User.Comment
	return
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
