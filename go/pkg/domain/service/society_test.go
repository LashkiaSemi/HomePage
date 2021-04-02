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

func TestSociety_GetAll(t *testing.T) {
	// t.Parallel()
	tests := []struct {
		name     string
		injector func(m *mock_repository.MockSocietyRepository)
		out      []*entity.Society
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "faile to get",
			injector: func(m *mock_repository.MockSocietyRepository) {
				m.EXPECT().FindAll().Return(nil, errors.New("some error"))
			},
			out:      nil,
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_repository.MockSocietyRepository) {
				m.EXPECT().FindAll().Return([]*entity.Society{
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
			m := mock_repository.NewMockSocietyRepository(ctrl)
			tt.injector(m)
			srv := NewSociety(m)
			out, e := srv.GetAll()
			assert.Equal(t, tt.out, out)
			tt.checkErr(t, e)
		})
	}
}

func TestSociety_GetByID(t *testing.T) {
	// t.Parallel()
	tests := []struct {
		name     string
		injector func(m *mock_repository.MockSocietyRepository)
		in       int
		out      *entity.Society
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to get",
			injector: func(m *mock_repository.MockSocietyRepository) {
				m.EXPECT().FindByID(1).Return(nil, errors.New("some error"))
			},
			in:       1,
			out:      nil,
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_repository.MockSocietyRepository) {
				m.EXPECT().FindByID(1).Return(&entity.Society{
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
			m := mock_repository.NewMockSocietyRepository(ctrl)
			tt.injector(m)
			srv := NewSociety(m)
			out, e := srv.GetByID(tt.in)
			assert.Equal(t, tt.out, out)
			tt.checkErr(t, e)
		})
	}
}

func TestSociety_Create(t *testing.T) {
	// t.Parallel()
	mockTime := "2020/02/02 00:00:00"
	type in struct {
		title   string
		author  string
		society string
		award   string
		date    string
	}
	tests := []struct {
		name     string
		injector func(m *mock_repository.MockSocietyRepository)
		in       in
		out      int
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to create",
			injector: func(m *mock_repository.MockSocietyRepository) {
				m.EXPECT().Create(&entity.Society{
					Title:     "title001",
					Author:    "author001",
					Society:   "soc001",
					Award:     "award001",
					Date:      mockTime,
					CreatedAt: mockTime,
					UpdatedAt: mockTime,
				}).Return(0, errors.New("some error"))
			},
			in: in{
				title:   "title001",
				author:  "author001",
				society: "soc001",
				award:   "award001",
				date:    mockTime,
			},
			out:      0,
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_repository.MockSocietyRepository) {
				m.EXPECT().Create(&entity.Society{
					Title:     "title001",
					Author:    "author001",
					Society:   "soc001",
					Award:     "award001",
					Date:      mockTime,
					CreatedAt: mockTime,
					UpdatedAt: mockTime,
				}).Return(1, nil)
			},
			in: in{
				title:   "title001",
				author:  "author001",
				society: "soc001",
				award:   "award001",
				date:    mockTime,
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
			m := mock_repository.NewMockSocietyRepository(ctrl)
			tt.injector(m)
			srv := NewSociety(m)
			out, e := srv.Create(tt.in.title, tt.in.author, tt.in.society, tt.in.award, tt.in.date)
			assert.Equal(t, tt.out, out)
			tt.checkErr(t, e)
		})
	}
}

func TestSociety_UpdateByID(t *testing.T) {
	// t.Parallel()
	mockTime := "2020/02/02 00:00:00"
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
		injector func(m *mock_repository.MockSocietyRepository)
		in       in
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to get",
			injector: func(m *mock_repository.MockSocietyRepository) {
				m.EXPECT().FindByID(1).Return(nil, errors.New("some error"))
			},
			in: in{
				id: 1,
			},
			checkErr: assert.NotNil,
		},
		{
			name: "failed to update",
			injector: func(m *mock_repository.MockSocietyRepository) {
				m.EXPECT().FindByID(1).Return(&entity.Society{
					ID:        1,
					Title:     "title001",
					Author:    "author001",
					Society:   "soc001",
					Award:     "award001",
					Date:      mockTime,
					CreatedAt: mockTime,
					UpdatedAt: mockTime,
				}, nil)
				m.EXPECT().UpdateByID(&entity.Society{
					ID:        1,
					Title:     "title001",
					Author:    "author001",
					Society:   "soc001",
					Award:     "award001",
					Date:      mockTime,
					CreatedAt: mockTime,
					UpdatedAt: mockTime,
				}).Return(errors.New("some error"))
			},
			in: in{
				id:      1,
				title:   "title001",
				author:  "author001",
				society: "soc001",
				award:   "award001",
				date:    mockTime,
			},
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_repository.MockSocietyRepository) {
				m.EXPECT().FindByID(1).Return(&entity.Society{
					ID:        1,
					Title:     "title001",
					Author:    "author001",
					Society:   "soc001",
					Award:     "award001",
					Date:      mockTime,
					CreatedAt: mockTime,
					UpdatedAt: mockTime,
				}, nil)
				m.EXPECT().UpdateByID(&entity.Society{
					ID:        1,
					Title:     "title001",
					Author:    "author001",
					Society:   "soc001",
					Award:     "award001",
					Date:      mockTime,
					CreatedAt: mockTime,
					UpdatedAt: mockTime,
				}).Return(nil)
			},
			in: in{
				id:      1,
				title:   "title001",
				author:  "author001",
				society: "soc001",
				award:   "award001",
				date:    mockTime,
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
			m := mock_repository.NewMockSocietyRepository(ctrl)
			tt.injector(m)
			srv := NewSociety(m)
			e := srv.UpdateByID(tt.in.id, tt.in.title, tt.in.author, tt.in.society, tt.in.award, tt.in.date)
			tt.checkErr(t, e)
		})
	}
}

func TestSociety_DeleteByID(t *testing.T) {
	// t.Parallel()
	tests := []struct {
		name     string
		injector func(m *mock_repository.MockSocietyRepository)
		in       int
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to delete",
			injector: func(m *mock_repository.MockSocietyRepository) {
				m.EXPECT().DeleteByID(1).Return(errors.New("some error"))
			},
			in:       1,
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_repository.MockSocietyRepository) {
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
			m := mock_repository.NewMockSocietyRepository(ctrl)
			tt.injector(m)
			srv := NewSociety(m)
			e := srv.DeleteByID(tt.in)
			tt.checkErr(t, e)
		})
	}
}
