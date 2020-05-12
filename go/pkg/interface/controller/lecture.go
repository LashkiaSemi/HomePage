package controller

import (
	"fmt"
	"homepage/pkg/configs"
	"homepage/pkg/entity"
	"homepage/pkg/usecase/interactor"
)

type lectureController struct {
	interactor.LectureInteractor
}

// LectureController レクチャーの入出力をコンバート
type LectureController interface {
	GetAll() (*LecturesResponse, error)
	GetByID(id int) (*LectureResponse, error)
	Create(studentID, title, file, comment string, activation int) (*LectureResponse, error)
	UpdateByID(id int, title, comment string, activation int) (*LectureResponse, error)
	DeleteByID(id int) error
}

// NewLectureController コントローラを作成
func NewLectureController(li interactor.LectureInteractor) LectureController {
	return &lectureController{
		LectureInteractor: li,
	}
}

func (lc *lectureController) GetAll() (*LecturesResponse, error) {
	lecs, err := lc.LectureInteractor.GetAll()
	if err != nil {
		return &LecturesResponse{}, err
	}
	var res LecturesResponse
	for _, lec := range lecs {
		res.Lectures = append(res.Lectures, convertToLectureResponse(lec))
	}
	return &res, nil
}

func (lc *lectureController) GetByID(id int) (*LectureResponse, error) {
	lec, err := lc.LectureInteractor.GetByID(id)
	if err != nil {
		return &LectureResponse{}, err
	}
	return convertToLectureResponse(lec), err
}

func (lc *lectureController) Create(studentID, title, file, comment string, activation int) (*LectureResponse, error) {
	lecture, err := lc.LectureInteractor.Create(studentID, title, file, comment, activation)
	if err != nil {
		return &LectureResponse{}, err
	}
	return convertToLectureResponse(lecture), nil
}

func (lc *lectureController) UpdateByID(id int, title, comment string, activation int) (*LectureResponse, error) {
	lec, err := lc.LectureInteractor.UpdateByID(id, title, comment, activation)
	if err != nil {
		return &LectureResponse{}, err
	}
	return convertToLectureResponse(lec), nil
}

func (lc *lectureController) DeleteByID(id int) error {
	return lc.LectureInteractor.DeleteByID(id)
}

// LecturesResponse Lectures
type LecturesResponse struct {
	Lectures []*LectureResponse
}

// LectureResponse ID, Author, Title, File, Comment, Activation, CreatedAt
type LectureResponse struct {
	ID         int
	Author     *UserResponse
	Title      string
	FileName   string
	FilePath   string
	Comment    string
	Activation int
	CreatedAt  string
}

func convertToLectureResponse(lec *entity.Lecture) *LectureResponse {
	return &LectureResponse{
		ID:         lec.ID,
		Author:     convertToUserResponse(lec.Author),
		Title:      lec.Title,
		FileName:   lec.File,
		FilePath:   fmt.Sprintf("%s/%s", configs.SaveLectureFileDir, lec.File),
		Comment:    lec.Comment,
		Activation: lec.Activation,
		CreatedAt:  lec.CreatedAt,
	}
}
