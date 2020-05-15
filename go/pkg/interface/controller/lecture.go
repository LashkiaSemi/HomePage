package controller

import (
	"fmt"
	"homepage/pkg/configs"
	"homepage/pkg/entity"
	"homepage/pkg/usecase/interactor"
	"strconv"

	"github.com/pkg/errors"
)

type lectureController struct {
	interactor.LectureInteractor
}

// LectureController レクチャーの入出力をコンバート
type LectureController interface {
	GetAll() (*LecturesResponse, error)
	GetByID(id int) (*LectureResponse, error)

	Create(studentID, title, file, comment string, activation int) (int, error)
	UpdateByID(id int, studentID, title, file, comment string, activation int) error
	DeleteByID(id int) error

	// admin
	AdminGetAll() ([]map[string]string, error)
	AdminGetByID(id int) (*FieldsResponse, error)
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
		err = errors.Wrap(err, "failed to original data for response")
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
		err = errors.Wrap(err, "failed to original data for response")
		return &LectureResponse{}, err
	}
	return convertToLectureResponse(lec), err
}

func (lc *lectureController) Create(studentID, title, file, comment string, activation int) (int, error) {
	return lc.LectureInteractor.Create(studentID, title, file, comment, activation)
}

func (lc *lectureController) UpdateByID(id int, studentID, title, file, comment string, activation int) error {
	return lc.LectureInteractor.UpdateByID(id, studentID, title, file, comment, activation)
}

func (lc *lectureController) DeleteByID(id int) error {
	return lc.LectureInteractor.DeleteByID(id)
}

// admin
func (lc *lectureController) AdminGetAll() ([]map[string]string, error) {
	var res []map[string]string
	datas, err := lc.LectureInteractor.GetAll()
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

func (lc *lectureController) AdminGetByID(id int) (*FieldsResponse, error) {
	var res FieldsResponse
	data, err := lc.LectureInteractor.GetByID(id)
	if err != nil {
		err = errors.Wrap(err, "failed to original data for response")
		return &res, err
	}
	res.Fields = append(res.Fields,
		&Field{Key: "ID", Value: data.ID},
		&Field{Key: "タイトル", Value: data.Title},
		&Field{Key: "投稿者", Value: data.Author.Name},
		&Field{Key: "ファイル名", Value: data.File},
		&Field{Key: "コメント", Value: data.Comment},
		&Field{Key: "公開", Value: data.Activation},
	)
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
