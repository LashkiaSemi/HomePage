package controller

import (
	"fmt"
	mock_interactor "homepage/mock/interactor"
	"homepage/pkg/configs"
	"homepage/pkg/domain/entity"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

var (
	filePrefix = fmt.Sprintf("%s/", configs.SaveLectureFileDir)
)

func TestLecture_GetAll(t *testing.T) {
	tests := []struct {
		name     string
		injector func(m *mock_interactor.MockLectureInteractor)
		out      *LecturesResponse
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to get",
			injector: func(m *mock_interactor.MockLectureInteractor) {
				m.EXPECT().GetAll().Return(nil, errors.New("some error"))
			},
			out:      &LecturesResponse{},
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_interactor.MockLectureInteractor) {
				m.EXPECT().GetAll().Return([]*entity.Lecture{
					{
						ID:     1,
						Author: &entity.User{},
					},
				}, nil)
			},
			out: &LecturesResponse{
				Lectures: []*LectureResponse{
					{
						ID:       1,
						FilePath: filePrefix,
						Author:   &UserResponse{},
					},
				},
			},
			checkErr: assert.Nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			m := mock_interactor.NewMockLectureInteractor(ctrl)
			tt.injector(m)
			c := NewLectureController(m)
			out, err := c.GetAll()
			assert.Equal(t, tt.out, out)
			tt.checkErr(t, err)
		})
	}
}

func TestLecture_GetByID(t *testing.T) {
	tests := []struct {
		name     string
		injector func(m *mock_interactor.MockLectureInteractor)
		in       int
		out      *LectureResponse
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to get",
			injector: func(m *mock_interactor.MockLectureInteractor) {
				m.EXPECT().GetByID(1).Return(nil, errors.New("some error"))
			},
			in:       1,
			out:      &LectureResponse{},
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_interactor.MockLectureInteractor) {
				m.EXPECT().GetByID(1).Return(&entity.Lecture{
					ID:     1,
					Author: &entity.User{},
				}, nil)
			},
			in: 1,
			out: &LectureResponse{
				ID:       1,
				FilePath: filePrefix,
				Author:   &UserResponse{},
			},
			checkErr: assert.Nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			m := mock_interactor.NewMockLectureInteractor(ctrl)
			tt.injector(m)
			c := NewLectureController(m)
			out, err := c.GetByID(tt.in)
			assert.Equal(t, tt.out, out)
			tt.checkErr(t, err)
		})
	}
}
func TestLecture_Create(t *testing.T) {
	type in struct {
		studentID  string
		title      string
		file       string
		comment    string
		activation int
	}
	tests := []struct {
		name     string
		injector func(m *mock_interactor.MockLectureInteractor)
		in       in
		out      int
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to create",
			injector: func(m *mock_interactor.MockLectureInteractor) {
				m.EXPECT().Create("", "", "", "", 0).Return(0, errors.New("some error"))
			},
			in:       in{},
			out:      0,
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_interactor.MockLectureInteractor) {
				m.EXPECT().Create("student001", "title001", "file001", "comment001", 0).Return(1, nil)
			},
			in: in{
				studentID: "student001",
				title:     "title001",
				file:      "file001",
				comment:   "comment001",
			},
			out:      1,
			checkErr: assert.Nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			m := mock_interactor.NewMockLectureInteractor(ctrl)
			tt.injector(m)
			c := NewLectureController(m)
			out, err := c.Create(tt.in.studentID, tt.in.title, tt.in.file, tt.in.comment, tt.in.activation)
			assert.Equal(t, tt.out, out)
			tt.checkErr(t, err)
		})
	}
}

func TestLecture_UpdateByID(t *testing.T) {
	type in struct {
		id         int
		studentID  string
		title      string
		file       string
		comment    string
		activation int
	}
	tests := []struct {
		name     string
		injector func(m *mock_interactor.MockLectureInteractor)
		in       in
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to update",
			injector: func(m *mock_interactor.MockLectureInteractor) {
				m.EXPECT().UpdateByID(0, "", "", "", "", 0).Return(errors.New("some error"))
			},
			in:       in{},
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_interactor.MockLectureInteractor) {
				m.EXPECT().UpdateByID(1, "student001", "title001", "file001", "comment001", 1).Return(nil)
			},
			in: in{
				id:         1,
				studentID:  "student001",
				title:      "title001",
				file:       "file001",
				comment:    "comment001",
				activation: 1,
			},
			checkErr: assert.Nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			m := mock_interactor.NewMockLectureInteractor(ctrl)
			tt.injector(m)
			c := NewLectureController(m)
			err := c.UpdateByID(tt.in.id, tt.in.studentID, tt.in.title, tt.in.file, tt.in.comment, tt.in.activation)
			tt.checkErr(t, err)
		})
	}
}

func TestLecture_DeleteByID(t *testing.T) {
	tests := []struct {
		name     string
		injector func(m *mock_interactor.MockLectureInteractor)
		in       int
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to delete",
			injector: func(m *mock_interactor.MockLectureInteractor) {
				m.EXPECT().DeleteByID(1).Return(errors.New("some error"))
			},
			in:       1,
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_interactor.MockLectureInteractor) {
				m.EXPECT().DeleteByID(1).Return(nil)
			},
			in:       1,
			checkErr: assert.Nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			m := mock_interactor.NewMockLectureInteractor(ctrl)
			tt.injector(m)
			c := NewLectureController(m)
			err := c.DeleteByID(tt.in)
			tt.checkErr(t, err)
		})
	}
}

func TestLecture_AdminGetAll(t *testing.T) {
	tests := []struct {
		name     string
		injector func(m *mock_interactor.MockLectureInteractor)
		out      []map[string]string
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to get",
			injector: func(m *mock_interactor.MockLectureInteractor) {
				m.EXPECT().GetAll().Return(nil, errors.New("some error"))
			},
			out:      []map[string]string(nil),
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_interactor.MockLectureInteractor) {
				m.EXPECT().GetAll().Return([]*entity.Lecture{
					{
						ID:     1,
						Title:  "title001",
						Author: &entity.User{},
					},
				}, nil)
			},
			out: []map[string]string{
				{
					"id":    "1",
					"title": "title001",
				},
			},
			checkErr: assert.Nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			m := mock_interactor.NewMockLectureInteractor(ctrl)
			tt.injector(m)
			c := NewLectureController(m)
			out, err := c.AdminGetAll()
			assert.Equal(t, tt.out, out)
			tt.checkErr(t, err)
		})
	}
}

func TestLecture_AdminGetByID(t *testing.T) {
	tests := []struct {
		name     string
		injector func(m *mock_interactor.MockLectureInteractor)
		in       int
		out      *FieldsResponse
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to get",
			injector: func(m *mock_interactor.MockLectureInteractor) {
				m.EXPECT().GetByID(1).Return(nil, errors.New("some error"))
			},
			in:       1,
			out:      &FieldsResponse{},
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_interactor.MockLectureInteractor) {
				m.EXPECT().GetByID(1).Return(&entity.Lecture{
					ID:         1,
					Title:      "title001",
					File:       "file001",
					Comment:    "comment001",
					Activation: 1,
					Author: &entity.User{
						Name: "author001",
					},
				}, nil)
			},
			in: 1,
			out: &FieldsResponse{
				Fields: []*Field{
					{Key: "ID", Value: 1},
					{Key: "タイトル", Value: "title001"},
					{Key: "投稿者", Value: "author001"},
					{Key: "ファイル名", Value: "file001"},
					{Key: "コメント", Value: "comment001"},
					{Key: "公開", Value: 1},
				},
			},
			checkErr: assert.Nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			m := mock_interactor.NewMockLectureInteractor(ctrl)
			tt.injector(m)
			c := NewLectureController(m)
			out, err := c.AdminGetByID(tt.in)
			assert.Equal(t, tt.out, out)
			tt.checkErr(t, err)
		})
	}
}
