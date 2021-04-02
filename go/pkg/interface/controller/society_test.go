package controller

import (
	mock_interactor "homepage/mock/interactor"
	"homepage/pkg/domain/entity"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestSociety_GetAll(t *testing.T) {
	tests := []struct {
		name     string
		injector func(m *mock_interactor.MockSocietyInteractor)
		out      *SocietiesResponse
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to get",
			injector: func(m *mock_interactor.MockSocietyInteractor) {
				m.EXPECT().GetAll().Return(nil, errors.New("some error"))
			},
			out:      &SocietiesResponse{},
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_interactor.MockSocietyInteractor) {
				m.EXPECT().GetAll().Return([]*entity.Society{
					{
						ID:    1,
						Title: "title001",
					},
				}, nil)
			},
			out: &SocietiesResponse{
				Societies: []*SocietyResponse{
					{
						Title: "title001",
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
			m := mock_interactor.NewMockSocietyInteractor(ctrl)
			tt.injector(m)
			c := NewSocietyController(m)
			out, err := c.GetAll()
			assert.Equal(t, tt.out, out)
			tt.checkErr(t, err)
		})
	}
}

func TestSociety_GetByID(t *testing.T) {
	tests := []struct {
		name     string
		injector func(m *mock_interactor.MockSocietyInteractor)
		in       int
		out      *SocietyResponse
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to get",
			injector: func(m *mock_interactor.MockSocietyInteractor) {
				m.EXPECT().GetByID(1).Return(nil, errors.New("some error"))
			},
			in:       1,
			out:      &SocietyResponse{},
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_interactor.MockSocietyInteractor) {
				m.EXPECT().GetByID(1).Return(&entity.Society{
					ID:    1,
					Title: "title001",
				}, nil)
			},
			in: 1,
			out: &SocietyResponse{
				Title: "title001",
			},
			checkErr: assert.Nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			m := mock_interactor.NewMockSocietyInteractor(ctrl)
			tt.injector(m)
			c := NewSocietyController(m)
			out, err := c.GetByID(tt.in)
			assert.Equal(t, tt.out, out)
			tt.checkErr(t, err)
		})
	}
}

func TestSociety_Create(t *testing.T) {
	type in struct {
		title   string
		author  string
		society string
		award   string
		date    string
	}
	tests := []struct {
		name     string
		injector func(m *mock_interactor.MockSocietyInteractor)
		in       in
		out      int
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to create",
			injector: func(m *mock_interactor.MockSocietyInteractor) {
				m.EXPECT().Create("", "", "", "", "").Return(0, errors.New("some error"))
			},
			in:       in{},
			out:      0,
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_interactor.MockSocietyInteractor) {
				m.EXPECT().Create("title001", "author001", "soc001", "award001", "date001").Return(1, nil)
			},
			in: in{
				title:   "title001",
				author:  "author001",
				society: "soc001",
				award:   "award001",
				date:    "date001",
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
			m := mock_interactor.NewMockSocietyInteractor(ctrl)
			tt.injector(m)
			c := NewSocietyController(m)
			out, err := c.Create(tt.in.title, tt.in.author, tt.in.society, tt.in.award, tt.in.date)
			assert.Equal(t, tt.out, out)
			tt.checkErr(t, err)
		})
	}
}

func TestSociety_UpdateByID(t *testing.T) {
	type in struct {
		id      int
		title   string
		author  string
		society string
		award   string
		date    string
	}
	tests := []struct {
		name     string
		injector func(m *mock_interactor.MockSocietyInteractor)
		in       in
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to get",
			injector: func(m *mock_interactor.MockSocietyInteractor) {
				m.EXPECT().UpdateByID(1, "", "", "", "", "").Return(errors.New("some error"))
			},
			in: in{
				id: 1,
			},
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_interactor.MockSocietyInteractor) {
				m.EXPECT().UpdateByID(1, "title001", "author001", "soc001", "award001", "date001").Return(nil)
			},
			in: in{
				id:      1,
				title:   "title001",
				author:  "author001",
				society: "soc001",
				award:   "award001",
				date:    "date001",
			},
			checkErr: assert.Nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			m := mock_interactor.NewMockSocietyInteractor(ctrl)
			tt.injector(m)
			c := NewSocietyController(m)
			err := c.UpdateByID(tt.in.id, tt.in.title, tt.in.author, tt.in.society, tt.in.award, tt.in.date)
			tt.checkErr(t, err)
		})
	}
}

func TestSociety_DeleteByID(t *testing.T) {
	tests := []struct {
		name     string
		injector func(m *mock_interactor.MockSocietyInteractor)
		in       int
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to delete",
			injector: func(m *mock_interactor.MockSocietyInteractor) {
				m.EXPECT().DeleteByID(1).Return(errors.New("some error"))
			},
			in:       1,
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_interactor.MockSocietyInteractor) {
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
			m := mock_interactor.NewMockSocietyInteractor(ctrl)
			tt.injector(m)
			c := NewSocietyController(m)
			err := c.DeleteByID(tt.in)
			tt.checkErr(t, err)
		})
	}
}

func TestSociety_AdminGetAll(t *testing.T) {
	tests := []struct {
		name     string
		injector func(m *mock_interactor.MockSocietyInteractor)
		out      []map[string]string
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to get",
			injector: func(m *mock_interactor.MockSocietyInteractor) {
				m.EXPECT().GetAll().Return(nil, errors.New("some error"))
			},
			out:      []map[string]string(nil),
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_interactor.MockSocietyInteractor) {
				m.EXPECT().GetAll().Return([]*entity.Society{
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
			m := mock_interactor.NewMockSocietyInteractor(ctrl)
			tt.injector(m)
			c := NewSocietyController(m)
			out, err := c.AdminGetAll()
			assert.Equal(t, tt.out, out)
			tt.checkErr(t, err)
		})
	}
}

func TestSociety_AdminGetByID(t *testing.T) {
	tests := []struct {
		name     string
		injector func(m *mock_interactor.MockSocietyInteractor)
		in       int
		out      *FieldsResponse
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to get",
			injector: func(m *mock_interactor.MockSocietyInteractor) {
				m.EXPECT().GetByID(1).Return(nil, errors.New("some error"))
			},
			in:       1,
			out:      &FieldsResponse{},
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_interactor.MockSocietyInteractor) {
				m.EXPECT().GetByID(1).Return(&entity.Society{
					ID:      1,
					Title:   "title001",
					Author:  "author001",
					Society: "soc001",
					Award:   "award001",
					Date:    "date001",
				}, nil)
			},
			in: 1,
			out: &FieldsResponse{
				ID: 1,
				Fields: []*Field{
					{Key: "ID", Value: 1},
					{Key: "タイトル", Value: "title001"},
					{Key: "氏名", Value: "author001"},
					{Key: "学会", Value: "soc001"},
					{Key: "受賞", Value: "award001"},
					{Key: "日付", Value: "date001"},
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
			m := mock_interactor.NewMockSocietyInteractor(ctrl)
			tt.injector(m)
			c := NewSocietyController(m)
			out, err := c.AdminGetByID(tt.in)
			assert.Equal(t, tt.out, out)
			tt.checkErr(t, err)
		})
	}
}
