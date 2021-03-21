package service

import (
	mock_repository "homepage/mock/repository"
	"homepage/pkg/domain/entity"
	"homepage/pkg/helper"
	"testing"

	"bou.ke/monkey"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestEquipment_GetAll(t *testing.T) {
	// t.Parallel()

	tests := []struct {
		name     string
		injector func(m *mock_repository.MockEquipmentRepository)
		out      []*entity.Equipment
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to get",
			injector: func(m *mock_repository.MockEquipmentRepository) {
				m.EXPECT().FindAll().Return(nil, errors.New("some error"))
			},
			out:      nil,
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_repository.MockEquipmentRepository) {
				m.EXPECT().FindAll().Return([]*entity.Equipment{
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
			m := mock_repository.NewMockEquipmentRepository(ctrl)
			tt.injector(m)
			srv := NewEquipment(m)
			out, e := srv.GetAll()
			assert.Equal(t, tt.out, out)
			tt.checkErr(t, e)
		})
	}
}

func TestEquipment_GetByID(t *testing.T) {
	// t.Parallel()
	tests := []struct {
		name     string
		injector func(m *mock_repository.MockEquipmentRepository)
		in       int
		out      *entity.Equipment
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to get",
			injector: func(m *mock_repository.MockEquipmentRepository) {
				m.EXPECT().FindByID(1).Return(nil, errors.New("some error"))
			},
			in:       1,
			out:      nil,
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_repository.MockEquipmentRepository) {
				m.EXPECT().FindByID(1).Return(&entity.Equipment{
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
			m := mock_repository.NewMockEquipmentRepository(ctrl)
			tt.injector(m)
			srv := NewEquipment(m)
			out, e := srv.GetByID(tt.in)
			assert.Equal(t, tt.out, out)
			tt.checkErr(t, e)
		})
	}
}

func TestEquipment_Create(t *testing.T) {
	// t.Parallel()
	mockTime := "2020/02/02 00:00:00"

	type in struct {
		name    string
		comment string
		stock   int
		tagID   int
	}
	tests := []struct {
		name     string
		injector func(m *mock_repository.MockEquipmentRepository)
		in       in
		out      int
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to create",
			injector: func(m *mock_repository.MockEquipmentRepository) {
				m.EXPECT().Create(gomock.Any()).Return(0, errors.New("some error"))
			},
			in:       in{},
			out:      0,
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_repository.MockEquipmentRepository) {
				m.EXPECT().Create(gomock.Any()).Return(1, nil)
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
			patch := monkey.Patch(helper.FormattedDateTimeNow, func() string { return mockTime })
			defer patch.Restore()
			m := mock_repository.NewMockEquipmentRepository(ctrl)
			tt.injector(m)
			srv := NewEquipment(m)
			out, e := srv.Create(tt.in.name, tt.in.comment, tt.in.stock, tt.in.tagID)
			assert.Equal(t, tt.out, out)
			tt.checkErr(t, e)
		})
	}
}

func TestEquipment_UpdateByID(t *testing.T) {
	// t.Parallel()
	mockTime := "2020/02/02 00:00:00"

	tag := &entity.Tag{
		ID: 1,
	}

	type in struct {
		id      int
		name    string
		comment string
		stock   int
		tagID   int
	}
	tests := []struct {
		name     string
		injector func(m *mock_repository.MockEquipmentRepository)
		in       in
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to get data",
			injector: func(m *mock_repository.MockEquipmentRepository) {
				m.EXPECT().FindByID(1).Return(nil, errors.New("some error"))
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
			name: "failed to update",
			injector: func(m *mock_repository.MockEquipmentRepository) {
				m.EXPECT().FindByID(1).Return(&entity.Equipment{
					ID:        1,
					Name:      "name001",
					Stock:     1,
					Comment:   "comment001",
					Tag:       tag,
					CreatedAt: mockTime,
					UpdatedAt: mockTime,
				}, nil)
				m.EXPECT().UpdateByID(&entity.Equipment{
					ID:        1,
					Name:      "name001",
					Stock:     1,
					Comment:   "comment001",
					Tag:       tag,
					CreatedAt: mockTime,
					UpdatedAt: mockTime,
				}).Return(errors.New("some error"))
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
			injector: func(m *mock_repository.MockEquipmentRepository) {
				m.EXPECT().FindByID(1).Return(&entity.Equipment{
					ID:      1,
					Name:    "name001",
					Stock:   1,
					Comment: "comment001",
					Tag: &entity.Tag{
						ID: 1,
					},
					CreatedAt: mockTime,
					UpdatedAt: mockTime,
				}, nil)
				m.EXPECT().UpdateByID(&entity.Equipment{
					ID:      1,
					Name:    "name001",
					Stock:   1,
					Comment: "comment001",
					Tag: &entity.Tag{
						ID: 1,
					},
					CreatedAt: mockTime,
					UpdatedAt: mockTime,
				}).Return(nil)
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
			patch := monkey.Patch(helper.FormattedDateTimeNow, func() string { return mockTime })
			defer patch.Restore()
			m := mock_repository.NewMockEquipmentRepository(ctrl)
			tt.injector(m)
			srv := NewEquipment(m)
			e := srv.UpdateByID(tt.in.id, tt.in.name, tt.in.comment, tt.in.stock, tt.in.tagID)
			tt.checkErr(t, e)
		})
	}
}

func TestEquipment_DeleteByID(t *testing.T) {
	// t.Parallel()
	tests := []struct {
		name     string
		injector func(m *mock_repository.MockEquipmentRepository)
		in       int
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to delete",
			injector: func(m *mock_repository.MockEquipmentRepository) {
				m.EXPECT().DeleteByID(1).Return(errors.New("some error"))
			},
			in:       1,
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_repository.MockEquipmentRepository) {
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
			m := mock_repository.NewMockEquipmentRepository(ctrl)
			tt.injector(m)
			srv := NewEquipment(m)
			e := srv.DeleteByID(tt.in)
			tt.checkErr(t, e)
		})
	}
}
