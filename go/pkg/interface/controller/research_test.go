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
	researchFilePrefix = fmt.Sprintf("%s/", configs.SaveResearchFileDir)
)

func TestResearch_GetAll(t *testing.T) {
	tests := []struct {
		name     string
		injector func(m *mock_interactor.MockResearchInteractor)
		out      *ResearchesResponse
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to get",
			injector: func(m *mock_interactor.MockResearchInteractor) {
				m.EXPECT().GetAll().Return(nil, errors.New("some error"))
			},
			out:      &ResearchesResponse{},
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_interactor.MockResearchInteractor) {
				m.EXPECT().GetAll().Return([]*entity.Research{
					{
						ID: 1,
					},
				}, nil)
			},
			out: &ResearchesResponse{
				Researches: []*ResearchResponse{
					{
						ID:       1,
						FilePath: researchFilePrefix,
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
			m := mock_interactor.NewMockResearchInteractor(ctrl)
			tt.injector(m)
			c := NewResearchController(m)
			out, err := c.GetAll()
			assert.Equal(t, tt.out, out)
			tt.checkErr(t, err)
		})
	}
}

func TestResearch_GetByID(t *testing.T) {
	tests := []struct {
		name     string
		injector func(m *mock_interactor.MockResearchInteractor)
		in       int
		out      *ResearchResponse
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to get",
			injector: func(m *mock_interactor.MockResearchInteractor) {
				m.EXPECT().GetByID(1).Return(nil, errors.New("some error"))
			},
			in:       1,
			out:      &ResearchResponse{},
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_interactor.MockResearchInteractor) {
				m.EXPECT().GetByID(1).Return(&entity.Research{
					ID: 1,
				}, nil)
			},
			in: 1,
			out: &ResearchResponse{
				ID:       1,
				FilePath: researchFilePrefix,
			},
			checkErr: assert.Nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			m := mock_interactor.NewMockResearchInteractor(ctrl)
			tt.injector(m)
			c := NewResearchController(m)
			out, err := c.GetByID(tt.in)
			assert.Equal(t, tt.out, out)
			tt.checkErr(t, err)
		})
	}
}

func TestResearch_Create(t *testing.T) {
	type in struct {
		title      string
		author     string
		file       string
		comment    string
		activation int
	}
	tests := []struct {
		name     string
		injector func(m *mock_interactor.MockResearchInteractor)
		in       in
		out      int
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to create",
			injector: func(m *mock_interactor.MockResearchInteractor) {
				m.EXPECT().Create("", "", "", "", 0).Return(0, errors.New("some error"))
			},
			in:       in{},
			out:      0,
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_interactor.MockResearchInteractor) {
				m.EXPECT().Create("title001", "author001", "file001", "comment001", 1).Return(1, nil)
			},
			in: in{
				title:      "title001",
				author:     "author001",
				file:       "file001",
				comment:    "comment001",
				activation: 1,
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
			m := mock_interactor.NewMockResearchInteractor(ctrl)
			tt.injector(m)
			c := NewResearchController(m)
			out, err := c.Create(tt.in.title, tt.in.author, tt.in.file, tt.in.comment, tt.in.activation)
			assert.Equal(t, tt.out, out)
			tt.checkErr(t, err)
		})
	}
}

func TestResearch_UpdateByID(t *testing.T) {
	type in struct {
		id         int
		title      string
		author     string
		file       string
		comment    string
		activation int
	}
	tests := []struct {
		name     string
		injector func(m *mock_interactor.MockResearchInteractor)
		in       in
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to update",
			injector: func(m *mock_interactor.MockResearchInteractor) {
				m.EXPECT().UpdateByID(1, "", "", "", "", 0).Return(errors.New("some error"))
			},
			in: in{
				id: 1,
			},
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_interactor.MockResearchInteractor) {
				m.EXPECT().UpdateByID(1, "title001", "author001", "file001", "comment001", 0).Return(nil)
			},
			in: in{
				id:         1,
				title:      "title001",
				author:     "author001",
				file:       "file001",
				comment:    "comment001",
				activation: 0,
			},
			checkErr: assert.Nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			m := mock_interactor.NewMockResearchInteractor(ctrl)
			tt.injector(m)
			c := NewResearchController(m)
			err := c.UpdateByID(tt.in.id, tt.in.title, tt.in.author, tt.in.file, tt.in.comment, tt.in.activation)
			tt.checkErr(t, err)
		})
	}
}

func TestResearch_DeleteByID(t *testing.T) {
	tests := []struct {
		name     string
		injector func(m *mock_interactor.MockResearchInteractor)
		in       int
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to delete",
			injector: func(m *mock_interactor.MockResearchInteractor) {
				m.EXPECT().DeleteByID(1).Return(errors.New("some error"))
			},
			in:       1,
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_interactor.MockResearchInteractor) {
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
			m := mock_interactor.NewMockResearchInteractor(ctrl)
			tt.injector(m)
			c := NewResearchController(m)
			err := c.DeleteByID(tt.in)
			tt.checkErr(t, err)
		})
	}
}

func TestResearch_AdminGetAll(t *testing.T) {
	tests := []struct {
		name     string
		injector func(m *mock_interactor.MockResearchInteractor)
		out      []map[string]string
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to get",
			injector: func(m *mock_interactor.MockResearchInteractor) {
				m.EXPECT().GetAll().Return(nil, errors.New("some error"))
			},
			out:      []map[string]string(nil),
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_interactor.MockResearchInteractor) {
				m.EXPECT().GetAll().Return([]*entity.Research{
					{
						ID:    1,
						Title: "title001",
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
			m := mock_interactor.NewMockResearchInteractor(ctrl)
			tt.injector(m)
			c := NewResearchController(m)
			out, err := c.AdminGetAll()
			assert.Equal(t, tt.out, out)
			tt.checkErr(t, err)
		})
	}
}

func TestResearch_AdminGetByID(t *testing.T) {
	tests := []struct {
		name     string
		injector func(m *mock_interactor.MockResearchInteractor)
		in       int
		out      *FieldsResponse
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to get",
			injector: func(m *mock_interactor.MockResearchInteractor) {
				m.EXPECT().GetByID(1).Return(nil, errors.New("some error"))
			},
			in:       1,
			out:      &FieldsResponse{},
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_interactor.MockResearchInteractor) {
				m.EXPECT().GetByID(1).Return(&entity.Research{
					ID:         1,
					Title:      "title001",
					Author:     "author001",
					File:       "file001",
					Comment:    "comment001",
					Activation: 1,
				}, nil)
			},
			in: 1,
			out: &FieldsResponse{
				ID: 1,
				Fields: []*Field{
					{Key: "ID", Value: 1},
					{Key: "タイトル", Value: "title001"},
					{Key: "著者", Value: "author001"},
					{Key: "ファイル", Value: "file001"},
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
			m := mock_interactor.NewMockResearchInteractor(ctrl)
			tt.injector(m)
			c := NewResearchController(m)
			out, err := c.AdminGetByID(tt.in)
			assert.Equal(t, tt.out, out)
			tt.checkErr(t, err)
		})
	}
}
