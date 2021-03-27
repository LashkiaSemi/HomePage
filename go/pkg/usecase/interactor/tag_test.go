package interactor

import (
	mock_service "homepage/mock/service"
	"homepage/pkg/domain/entity"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestTag_GetAll(t *testing.T) {
	tests := []struct {
		name     string
		injector func(m *mock_service.MockTag)
		out      []*entity.Tag
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to get",
			injector: func(m *mock_service.MockTag) {
				m.EXPECT().GetAll().Return(nil, errors.New("some error"))
			},
			out:      nil,
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_service.MockTag) {
				m.EXPECT().GetAll().Return([]*entity.Tag{
					{
						ID: 1,
					},
				}, nil)
			},
			out: []*entity.Tag{
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
			m := mock_service.NewMockTag(ctrl)
			tt.injector(m)
			u := NewTagInteractor(m)
			out, err := u.GetAll()
			assert.Equal(t, tt.out, out)
			tt.checkErr(t, err)
		})
	}
}

func TestTag_GetByID(t *testing.T) {
	tests := []struct {
		name     string
		injector func(m *mock_service.MockTag)
		in       int
		out      *entity.Tag
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to get",
			injector: func(m *mock_service.MockTag) {
				m.EXPECT().GetByID(1).Return(nil, errors.New("some error"))
			},
			in:       1,
			out:      nil,
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_service.MockTag) {
				m.EXPECT().GetByID(1).Return(&entity.Tag{
					ID: 1,
				}, nil)
			},
			in: 1,
			out: &entity.Tag{
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
			m := mock_service.NewMockTag(ctrl)
			tt.injector(m)
			u := NewTagInteractor(m)
			out, err := u.GetByID(tt.in)
			assert.Equal(t, tt.out, out)
			tt.checkErr(t, err)
		})
	}
}

func TestTag_Create(t *testing.T) {
	tests := []struct {
		name     string
		injector func(m *mock_service.MockTag)
		in       string
		out      int
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to create",
			injector: func(m *mock_service.MockTag) {
				m.EXPECT().Create("tag001").Return(0, errors.New("some error"))
			},
			in:       "tag001",
			out:      0,
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_service.MockTag) {
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
			m := mock_service.NewMockTag(ctrl)
			tt.injector(m)
			u := NewTagInteractor(m)
			out, err := u.Create(tt.in)
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
		injector func(m *mock_service.MockTag)
		in       in
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to update",
			injector: func(m *mock_service.MockTag) {
				m.EXPECT().UpdateByID(1, "tag001").Return(errors.New("some error"))
			},
			in: in{
				id:   1,
				name: "tag001",
			},
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_service.MockTag) {
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
			m := mock_service.NewMockTag(ctrl)
			tt.injector(m)
			u := NewTagInteractor(m)
			err := u.UpdateByID(tt.in.id, tt.in.name)
			tt.checkErr(t, err)
		})
	}
}

func TestTag_DeleteByID(t *testing.T) {
	tests := []struct {
		name     string
		injector func(m *mock_service.MockTag)
		in       int
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to delete",
			injector: func(m *mock_service.MockTag) {
				m.EXPECT().DeleteByID(1).Return(errors.New("some error"))
			},
			in:       1,
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_service.MockTag) {
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
			m := mock_service.NewMockTag(ctrl)
			tt.injector(m)
			u := NewTagInteractor(m)
			err := u.DeleteByID(tt.in)
			tt.checkErr(t, err)
		})
	}
}
