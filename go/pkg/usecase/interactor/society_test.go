package interactor

import (
	mock_service "homepage/mock/service"
	"homepage/pkg/domain/entity"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestSociety_GetAll(t *testing.T) {
	tests := []struct {
		name     string
		injector func(m *mock_service.MockSociety)
		out      []*entity.Society
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to get",
			injector: func(m *mock_service.MockSociety) {
				m.EXPECT().GetAll().Return(nil, errors.New("some error"))
			},
			out:      nil,
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_service.MockSociety) {
				m.EXPECT().GetAll().Return([]*entity.Society{
					{
						ID: 1,
					},
				}, nil)
			},
			out: []*entity.Society{
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
			m := mock_service.NewMockSociety(ctrl)
			tt.injector(m)
			u := NewSocietyInteractor(m)
			out, err := u.GetAll()
			assert.Equal(t, tt.out, out)
			tt.checkErr(t, err)
		})
	}
}

func TestSociety_GetByID(t *testing.T) {
	tests := []struct {
		name     string
		injector func(m *mock_service.MockSociety)
		in       int
		out      *entity.Society
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to get",
			injector: func(m *mock_service.MockSociety) {
				m.EXPECT().GetByID(1).Return(nil, errors.New("some error"))
			},
			in:       1,
			out:      nil,
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_service.MockSociety) {
				m.EXPECT().GetByID(1).Return(&entity.Society{
					ID: 1,
				}, nil)
			},
			in: 1,
			out: &entity.Society{
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
			m := mock_service.NewMockSociety(ctrl)
			tt.injector(m)
			u := NewSocietyInteractor(m)
			out, err := u.GetByID(tt.in)
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
		injector func(m *mock_service.MockSociety)
		in       in
		out      int
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to create",
			injector: func(m *mock_service.MockSociety) {
				m.EXPECT().Create("title001", "author001", "soc001", "award001", "date001").Return(0, errors.New("some error"))
			},
			in: in{
				title:   "title001",
				author:  "author001",
				society: "soc001",
				award:   "award001",
				date:    "date001",
			},
			out:      0,
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_service.MockSociety) {
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
			m := mock_service.NewMockSociety(ctrl)
			tt.injector(m)
			u := NewSocietyInteractor(m)
			out, err := u.Create(tt.in.title, tt.in.author, tt.in.society, tt.in.award, tt.in.date)
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
		injector func(m *mock_service.MockSociety)
		in       in
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to update",
			injector: func(m *mock_service.MockSociety) {
				m.EXPECT().UpdateByID(1, "", "", "", "", "").Return(errors.New("some error"))
			},
			in: in{
				id: 1,
			},
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_service.MockSociety) {
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
			m := mock_service.NewMockSociety(ctrl)
			tt.injector(m)
			u := NewSocietyInteractor(m)
			err := u.UpdateByID(tt.in.id, tt.in.title, tt.in.author, tt.in.society, tt.in.award, tt.in.date)
			tt.checkErr(t, err)
		})
	}
}

func TestSociety_DeletByID(t *testing.T) {
	tests := []struct {
		name     string
		injector func(m *mock_service.MockSociety)
		in       int
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to delete",
			injector: func(m *mock_service.MockSociety) {
				m.EXPECT().DeleteByID(1).Return(errors.New("some error"))
			},
			in:       1,
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_service.MockSociety) {
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
			m := mock_service.NewMockSociety(ctrl)
			tt.injector(m)
			u := NewSocietyInteractor(m)
			err := u.DeleteByID(tt.in)
			tt.checkErr(t, err)
		})
	}
}
