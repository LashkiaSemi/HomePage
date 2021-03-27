package interactor

import (
	mock_service "homepage/mock/service"
	"homepage/pkg/domain/entity"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestResearch_GetAll(t *testing.T) {
	tests := []struct {
		name     string
		injector func(m *mock_service.MockResearch)
		out      []*entity.Research
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to get",
			injector: func(m *mock_service.MockResearch) {
				m.EXPECT().GetAll().Return(nil, errors.New("some error"))
			},
			out:      nil,
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_service.MockResearch) {
				m.EXPECT().GetAll().Return([]*entity.Research{
					{
						ID: 1,
					},
				}, nil)
			},
			out: []*entity.Research{
				{
					ID: 1,
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
			m := mock_service.NewMockResearch(ctrl)
			tt.injector(m)
			u := NewResearchInteractor(m)
			out, err := u.GetAll()
			assert.Equal(t, tt.out, out)
			tt.checkErr(t, err)
		})
	}
}

func TestResearch_GetByID(t *testing.T) {
	tests := []struct {
		name     string
		injector func(m *mock_service.MockResearch)
		in       int
		out      *entity.Research
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to get",
			injector: func(m *mock_service.MockResearch) {
				m.EXPECT().GetByID(1).Return(nil, errors.New("some error"))
			},
			in:       1,
			out:      nil,
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_service.MockResearch) {
				m.EXPECT().GetByID(1).Return(&entity.Research{
					ID: 1,
				}, nil)
			},
			in: 1,
			out: &entity.Research{
				ID: 1,
			},
			checkErr: assert.Nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			m := mock_service.NewMockResearch(ctrl)
			tt.injector(m)
			u := NewResearchInteractor(m)
			out, err := u.GetByID(tt.in)
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
		injector func(m *mock_service.MockResearch)
		in       in
		out      int
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to create",
			injector: func(m *mock_service.MockResearch) {
				m.EXPECT().Create("title001", "author001", "file001", "comment001", 1).Return(0, errors.New("some error"))
			},
			in: in{
				title:      "title001",
				author:     "author001",
				file:       "file001",
				comment:    "comment001",
				activation: 1,
			},
			out:      0,
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_service.MockResearch) {
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
			m := mock_service.NewMockResearch(ctrl)
			tt.injector(m)
			u := NewResearchInteractor(m)
			out, err := u.Create(tt.in.title, tt.in.author, tt.in.file, tt.in.comment, tt.in.activation)
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
		injector func(m *mock_service.MockResearch)
		in       in
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to update",
			injector: func(m *mock_service.MockResearch) {
				m.EXPECT().UpdateByID(1, "", "", "", "", 0).Return(errors.New("some error"))
			},
			in: in{
				id: 1,
			},
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_service.MockResearch) {
				m.EXPECT().UpdateByID(1, "title001", "author001", "file001", "comment001", 1).Return(nil)
			},
			in: in{
				id:         1,
				title:      "title001",
				author:     "author001",
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
			m := mock_service.NewMockResearch(ctrl)
			tt.injector(m)
			u := NewResearchInteractor(m)
			err := u.UpdateByID(tt.in.id, tt.in.title, tt.in.author, tt.in.file, tt.in.comment, tt.in.activation)
			tt.checkErr(t, err)
		})
	}
}

func TestResearch_DeleteByID(t *testing.T) {
	tests := []struct {
		name     string
		injector func(m *mock_service.MockResearch)
		in       int
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to delete",
			injector: func(m *mock_service.MockResearch) {
				m.EXPECT().DeleteByID(1).Return(errors.New("some error"))
			},
			in:       1,
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_service.MockResearch) {
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
			m := mock_service.NewMockResearch(ctrl)
			tt.injector(m)
			u := NewResearchInteractor(m)
			err := u.DeleteByID(tt.in)
			tt.checkErr(t, err)
		})
	}
}
