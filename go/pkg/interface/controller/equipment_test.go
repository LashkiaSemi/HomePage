package controller

import (
	mock_interactor "homepage/mock/interactor"
	"homepage/pkg/domain/entity"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestEquipment_GetAll(t *testing.T) {
	tests := []struct {
		name     string
		injector func(m *mock_interactor.MockEquipmentInteractor)
		out      *EquipmentsResponse
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to get",
			injector: func(m *mock_interactor.MockEquipmentInteractor) {
				m.EXPECT().GetAll().Return(nil, errors.New("some error"))
			},
			out:      &EquipmentsResponse{},
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_interactor.MockEquipmentInteractor) {
				m.EXPECT().GetAll().Return([]*entity.Equipment{
					{
						ID:  1,
						Tag: &entity.Tag{},
					},
				}, nil)
			},
			out: &EquipmentsResponse{
				Equipments: []*EquipmentResponse{
					{
						ID:  1,
						Tag: &TagResponse{},
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
			m := mock_interactor.NewMockEquipmentInteractor(ctrl)
			tt.injector(m)
			c := NewEquipmentController(m)
			out, err := c.GetAll()
			assert.Equal(t, tt.out, out)
			tt.checkErr(t, err)
		})
	}
}

func TestEquipment_GetByID(t *testing.T) {
	tests := []struct {
		name     string
		injector func(m *mock_interactor.MockEquipmentInteractor)
		in       int
		out      *EquipmentResponse
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to get",
			injector: func(m *mock_interactor.MockEquipmentInteractor) {
				m.EXPECT().GetByID(1).Return(nil, errors.New("some error"))
			},
			in:       1,
			out:      &EquipmentResponse{},
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_interactor.MockEquipmentInteractor) {
				m.EXPECT().GetByID(1).Return(&entity.Equipment{
					ID:  1,
					Tag: &entity.Tag{},
				}, nil)
			},
			in: 1,
			out: &EquipmentResponse{
				ID:  1,
				Tag: &TagResponse{},
			},
			checkErr: assert.Nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			m := mock_interactor.NewMockEquipmentInteractor(ctrl)
			tt.injector(m)
			c := NewEquipmentController(m)
			out, err := c.GetByID(tt.in)
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
		injector func(m *mock_interactor.MockEquipmentInteractor)
		in       in
		out      int
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to create",
			injector: func(m *mock_interactor.MockEquipmentInteractor) {
				m.EXPECT().Create("", "", 0, 0).Return(0, errors.New("some error"))
			},
			in:       in{},
			out:      0,
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_interactor.MockEquipmentInteractor) {
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
			m := mock_interactor.NewMockEquipmentInteractor(ctrl)
			tt.injector(m)
			c := NewEquipmentController(m)
			out, err := c.Create(tt.in.name, tt.in.comment, tt.in.stock, tt.in.tagID)
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
		injector func(m *mock_interactor.MockEquipmentInteractor)
		in       in
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to update",
			injector: func(m *mock_interactor.MockEquipmentInteractor) {
				m.EXPECT().UpdateByID(1, "", "", 0, 0).Return(errors.New("some error"))
			},
			in: in{
				id: 1,
			},
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_interactor.MockEquipmentInteractor) {
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
			m := mock_interactor.NewMockEquipmentInteractor(ctrl)
			tt.injector(m)
			c := NewEquipmentController(m)
			err := c.UpdateByID(tt.in.id, tt.in.name, tt.in.comment, tt.in.stock, tt.in.tagID)
			tt.checkErr(t, err)
		})
	}
}

func TestEquipment_DeleteByID(t *testing.T) {
	tests := []struct {
		name     string
		injector func(m *mock_interactor.MockEquipmentInteractor)
		in       int
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to delete",
			injector: func(m *mock_interactor.MockEquipmentInteractor) {
				m.EXPECT().DeleteByID(0).Return(errors.New("some error"))
			},
			in:       0,
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_interactor.MockEquipmentInteractor) {
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
			m := mock_interactor.NewMockEquipmentInteractor(ctrl)
			tt.injector(m)
			c := NewEquipmentController(m)
			err := c.DeleteByID(tt.in)
			tt.checkErr(t, err)
		})
	}
}

func TestEquipment_AdminGetAll(t *testing.T) {
	tests := []struct {
		name     string
		injector func(m *mock_interactor.MockEquipmentInteractor)
		out      []map[string]string
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to get",
			injector: func(m *mock_interactor.MockEquipmentInteractor) {
				m.EXPECT().GetAll().Return(nil, errors.New("some error"))
			},
			out:      []map[string]string(nil),
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_interactor.MockEquipmentInteractor) {
				m.EXPECT().GetAll().Return([]*entity.Equipment{
					{
						ID:   1,
						Name: "name001",
					},
				}, nil)
			},
			out: []map[string]string{
				{
					"id":    "1",
					"title": "name001",
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
			m := mock_interactor.NewMockEquipmentInteractor(ctrl)
			tt.injector(m)
			c := NewEquipmentController(m)
			out, err := c.AdminGetAll()
			assert.Equal(t, tt.out, out)
			tt.checkErr(t, err)
		})
	}
}

func TestEquipment_AdminGetByID(t *testing.T) {
	tests := []struct {
		name     string
		injector func(m *mock_interactor.MockEquipmentInteractor)
		in       int
		out      *FieldsResponse
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to get",
			injector: func(m *mock_interactor.MockEquipmentInteractor) {
				m.EXPECT().GetByID(0).Return(nil, errors.New("some error"))
			},
			in:       0,
			out:      &FieldsResponse{},
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_interactor.MockEquipmentInteractor) {
				m.EXPECT().GetByID(1).Return(&entity.Equipment{
					ID:      1,
					Name:    "name001",
					Comment: "comment001",
					Stock:   1,
					Tag: &entity.Tag{
						ID:   1,
						Name: "tag001",
					},
				}, nil)
			},
			in: 1,
			out: &FieldsResponse{
				Fields: []*Field{
					{Key: "ID", Value: 1},
					{Key: "備品名", Value: "name001"},
					{Key: "在庫", Value: 1},
					{Key: "コメント", Value: "comment001"},
					{Key: "タグ", Value: "tag001"},
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
			m := mock_interactor.NewMockEquipmentInteractor(ctrl)
			tt.injector(m)
			c := NewEquipmentController(m)
			out, err := c.AdminGetByID(tt.in)
			assert.Equal(t, tt.out, out)
			tt.checkErr(t, err)
		})
	}
}
