package controller

import (
	mock_interactor "homepage/mock/interactor"
	"homepage/pkg/domain/entity"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestTag_GetAll(t *testing.T) {
	tests := []struct {
		name     string
		injector func(m *mock_interactor.MockTagInteractor)
		out      *TagsResponse
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to get",
			injector: func(m *mock_interactor.MockTagInteractor) {
				m.EXPECT().GetAll().Return(nil, errors.New("some error"))
			},
			out:      &TagsResponse{},
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_interactor.MockTagInteractor) {
				m.EXPECT().GetAll().Return([]*entity.Tag{
					{
						ID: 1,
					},
				}, nil)
			},
			out: &TagsResponse{
				Tags: []*TagResponse{
					{
						ID: 1,
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
			m := mock_interactor.NewMockTagInteractor(ctrl)
			tt.injector(m)
			c := NewTagController(m)
			out, err := c.GetAll()
			assert.Equal(t, tt.out, out)
			tt.checkErr(t, err)
		})
	}
}

func TestTag_GetByID(t *testing.T) {
	tests := []struct {
		name     string
		injector func(m *mock_interactor.MockTagInteractor)
		in       int
		out      *TagResponse
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to get",
			injector: func(m *mock_interactor.MockTagInteractor) {
				m.EXPECT().GetByID(1).Return(nil, errors.New("some error"))
			},
			in:       1,
			out:      &TagResponse{},
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_interactor.MockTagInteractor) {
				m.EXPECT().GetByID(1).Return(&entity.Tag{
					ID: 1,
				}, nil)
			},
			in: 1,
			out: &TagResponse{
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
			m := mock_interactor.NewMockTagInteractor(ctrl)
			tt.injector(m)
			c := NewTagController(m)
			out, err := c.GetByID(tt.in)
			assert.Equal(t, tt.out, out)
			tt.checkErr(t, err)
		})
	}
}

func TestTag_Create(t *testing.T) {
	tests := []struct {
		name     string
		injector func(m *mock_interactor.MockTagInteractor)
		in       string
		out      int
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to create",
			injector: func(m *mock_interactor.MockTagInteractor) {
				m.EXPECT().Create("").Return(0, errors.New("some error"))
			},
			in:       "",
			out:      0,
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_interactor.MockTagInteractor) {
				m.EXPECT().Create("tag001").Return(1, nil)
			},
			in:       "tag001",
			out:      1,
			checkErr: assert.Nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			m := mock_interactor.NewMockTagInteractor(ctrl)
			tt.injector(m)
			c := NewTagController(m)
			out, err := c.Create(tt.in)
			assert.Equal(t, tt.out, out)
			tt.checkErr(t, err)
		})
	}
}

func TestTag_UpdateByID(t *testing.T) {
	type in struct {
		id   int
		name string
	}
	tests := []struct {
		name     string
		injector func(m *mock_interactor.MockTagInteractor)
		in       in
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to update",
			injector: func(m *mock_interactor.MockTagInteractor) {
				m.EXPECT().UpdateByID(1, "").Return(errors.New("some error"))
			},
			in: in{
				id: 1,
			},
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_interactor.MockTagInteractor) {
				m.EXPECT().UpdateByID(1, "tag001").Return(nil)
			},
			in: in{
				id:   1,
				name: "tag001",
			},
			checkErr: assert.Nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			m := mock_interactor.NewMockTagInteractor(ctrl)
			tt.injector(m)
			c := NewTagController(m)
			err := c.UpdateByID(tt.in.id, tt.in.name)
			tt.checkErr(t, err)
		})
	}
}

func TestTag_DeleteByID(t *testing.T) {
	tests := []struct {
		name     string
		injector func(m *mock_interactor.MockTagInteractor)
		in       int
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to delete",
			injector: func(m *mock_interactor.MockTagInteractor) {
				m.EXPECT().DeleteByID(1).Return(errors.New("some error"))
			},
			in:       1,
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_interactor.MockTagInteractor) {
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
			m := mock_interactor.NewMockTagInteractor(ctrl)
			tt.injector(m)
			c := NewTagController(m)
			err := c.DeleteByID(tt.in)
			tt.checkErr(t, err)
		})
	}
}

func TestTag_AdminGetAll(t *testing.T) {
	tests := []struct {
		name     string
		injector func(m *mock_interactor.MockTagInteractor)
		out      []map[string]string
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to get",
			injector: func(m *mock_interactor.MockTagInteractor) {
				m.EXPECT().GetAll().Return(nil, errors.New("some error"))
			},
			out:      []map[string]string(nil),
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_interactor.MockTagInteractor) {
				m.EXPECT().GetAll().Return([]*entity.Tag{
					{
						ID:   1,
						Name: "tag001",
					},
				}, nil)
			},
			out: []map[string]string{
				{
					"id":    "1",
					"title": "tag001",
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
			m := mock_interactor.NewMockTagInteractor(ctrl)
			tt.injector(m)
			c := NewTagController(m)
			out, err := c.AdminGetAll()
			assert.Equal(t, tt.out, out)
			tt.checkErr(t, err)
		})
	}
}

func TestTag_AdminGetByID(t *testing.T) {
	tests := []struct {
		name     string
		injector func(m *mock_interactor.MockTagInteractor)
		in       int
		out      *FieldsResponse
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to get",
			injector: func(m *mock_interactor.MockTagInteractor) {
				m.EXPECT().GetByID(1).Return(nil, errors.New("some error"))
			},
			in:       1,
			out:      &FieldsResponse{},
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_interactor.MockTagInteractor) {
				m.EXPECT().GetByID(1).Return(&entity.Tag{
					ID:   1,
					Name: "tag001",
				}, nil)
			},
			in: 1,
			out: &FieldsResponse{
				ID: 1,
				Fields: []*Field{
					{Key: "ID", Value: 1},
					{Key: "名前", Value: "tag001"},
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
			m := mock_interactor.NewMockTagInteractor(ctrl)
			tt.injector(m)
			c := NewTagController(m)
			out, err := c.AdminGetByID(tt.in)
			assert.Equal(t, tt.out, out)
			tt.checkErr(t, err)
		})
	}
}
