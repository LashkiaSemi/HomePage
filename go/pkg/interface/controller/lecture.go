package controller

import (
	"homepage/pkg/domain/model"
	"homepage/pkg/usecase/interactor"
)

type lectureController struct {
	interactor.LectureInteractor
}

// LectureController レクチャーの入出力をコンバート
type LectureController interface {
	GetAll() (*LecturesResponse, error)
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

// LecturesResponse Lectures
type LecturesResponse struct {
	Lectures []*LectureResponse
}

// LectureResponse ID, Author, Title, File, Comment, Activation, CreatedAt
type LectureResponse struct {
	ID         int
	Author     *UserResponse
	Title      string
	File       string
	Comment    string
	Activation int
	CreatedAt  string
}

func convertToLectureResponse(lec *model.Lecture) *LectureResponse {
	return &LectureResponse{
		ID:         lec.ID,
		Author:     convertToUserResponse(lec.Author),
		Title:      lec.Title,
		File:       lec.File,
		Comment:    lec.Comment,
		Activation: lec.Activation,
		CreatedAt:  lec.CreatedAt,
	}
}
