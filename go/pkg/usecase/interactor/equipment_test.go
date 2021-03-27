package interactor

import (
	mock_service "homepage/mock/service"
	"homepage/pkg/domain/entity"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestEquipment_GetAll(t *testing.T) {
	tests := []struct {
		name     string
		injector func(m *mock_service.MockEquipment)
		out      []*entity.Equipment
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to get",
			injector: func(m *mock_service.MockEquipment) {
				m.EXPECT().GetAll().Return(nil, errors.New("some error"))
			},
			out:      nil,
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_service.MockEquipment) {
				m.EXPECT().GetAll().Return([]*entity.Equipment{
					{
						ID: 1,
					},
				}, nil)
			},
			out: []*entity.Equipment{
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
			m := mock_service.NewMockEquipment(ctrl)
			tt.injector(m)
			u := NewEquipmentInteractor(m)
			out, err := u.GetAll()
			assert.Equal(t, tt.out, out)
			tt.checkErr(t, err)
		})
	}
}

func TestEquipment_GetByID(t *testing.T) {
	tests := []struct {
		name     string
		injector func(m *mock_service.MockEquipment)
		in       int
		out      *entity.Equipment
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to get",
			injector: func(m *mock_service.MockEquipment) {
				m.EXPECT().GetByID(1).Return(nil, errors.New("some error"))
			},
			in:       1,
			out:      nil,
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_service.MockEquipment) {
				m.EXPECT().GetByID(1).Return(&entity.Equipment{
					ID: 1,
				}, nil)
			},
			in: 1,
			out: &entity.Equipment{
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
			m := mock_service.NewMockEquipment(ctrl)
			tt.injector(m)
			u := NewEquipmentInteractor(m)
			out, err := u.GetByID(tt.in)
			assert.Equal(t, tt.out, out)
			tt.checkErr(t, err)
		})
	}
}

func TestEquipment_Create(t *testing.T) {
	type in struct {
		name    string
		comment string
		stock   int
		tagID   int
	}
	tests := []struct {
		name     string
		injector func(m *mock_service.MockEquipment)
		in       in
		out      int
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to create",
			injector: func(m *mock_service.MockEquipment) {
				m.EXPECT().Create("name001", "comment001", 1, 1).Return(0, errors.New("some error"))
			},
			in: in{
				name:    "name001",
				comment: "comment001",
				stock:   1,
				tagID:   1,
			},
			out:      0,
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_service.MockEquipment) {
				m.EXPECT().Create("name001", "comment001", 1, 1).Return(1, nil)
			},
			in: in{
				name:    "name001",
				comment: "comment001",
				stock:   1,
				tagID:   1,
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
			m := mock_service.NewMockEquipment(ctrl)
			tt.injector(m)
			u := NewEquipmentInteractor(m)
			out, err := u.Create(tt.in.name, tt.in.comment, tt.in.stock, tt.in.tagID)
			assert.Equal(t, tt.out, out)
			tt.checkErr(t, err)
		})
	}
}

func TestEquipment_UpdateByID(t *testing.T) {
	type in struct {
		id      int
		name    string
		comment string
		stock   int
		tagID   int
	}
	tests := []struct {
		name     string
		injector func(m *mock_service.MockEquipment)
		in       in
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to update",
			injector: func(m *mock_service.MockEquipment) {
				m.EXPECT().UpdateByID(1, "name001", "comment001", 1, 1).Return(errors.New("some error"))
			},
			in: in{
				id:      1,
				name:    "name001",
				comment: "comment001",
				stock:   1,
				tagID:   1,
			},
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_service.MockEquipment) {
				m.EXPECT().UpdateByID(1, "name001", "comment001", 1, 1).Return(nil)
			},
			in: in{
				id:      1,
				name:    "name001",
				comment: "comment001",
				stock:   1,
				tagID:   1,
			},
			checkErr: assert.Nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			m := mock_service.NewMockEquipment(ctrl)
			tt.injector(m)
			u := NewEquipmentInteractor(m)
			err := u.UpdateByID(tt.in.id, tt.in.name, tt.in.comment, tt.in.stock, tt.in.tagID)
			tt.checkErr(t, err)
		})
	}
}

func TestEquipment_DeleteByID(t *testing.T) {
	tests := []struct {
		name     string
		injector func(m *mock_service.MockEquipment)
		in       int
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to delete",
			injector: func(m *mock_service.MockEquipment) {
				m.EXPECT().DeleteByID(1).Return(errors.New("some error"))
			},
			in:       1,
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_service.MockEquipment) {
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
			m := mock_service.NewMockEquipment(ctrl)
			tt.injector(m)
			u := NewEquipmentInteractor(m)
			err := u.DeleteByID(tt.in)
			tt.checkErr(t, err)
		})
	}
}
