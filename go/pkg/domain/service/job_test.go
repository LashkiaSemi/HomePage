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

func TestJob_GetAll(t *testing.T) {
	// t.Parallel()
	tests := []struct {
		name     string
		injector func(m *mock_repository.MockJobRepository)
		out      []*entity.Job
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to get",
			injector: func(m *mock_repository.MockJobRepository) {
				m.EXPECT().FindAll().Return(nil, errors.New("some error"))
			},
			out:      nil,
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_repository.MockJobRepository) {
				m.EXPECT().FindAll().Return([]*entity.Job{
					{
						ID: 1,
					},
				}, nil)
			},
			out: []*entity.Job{
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
			m := mock_repository.NewMockJobRepository(ctrl)
			tt.injector(m)
			srv := NewJob(m)
			out, err := srv.GetAll()
			assert.Equal(t, tt.out, out)
			tt.checkErr(t, err)
		})
	}
}

func TestJob_GetByID(t *testing.T) {
	// t.Parallel()
	tests := []struct {
		name     string
		injector func(m *mock_repository.MockJobRepository)
		in       int
		out      *entity.Job
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to get",
			injector: func(m *mock_repository.MockJobRepository) {
				m.EXPECT().FindByID(1).Return(nil, errors.New("some error"))
			},
			in:       1,
			out:      nil,
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_repository.MockJobRepository) {
				m.EXPECT().FindByID(1).Return(&entity.Job{
					ID: 1,
				}, nil)
			},
			in: 1,
			out: &entity.Job{
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
			m := mock_repository.NewMockJobRepository(ctrl)
			tt.injector(m)
			srv := NewJob(m)
			out, err := srv.GetByID(tt.in)
			assert.Equal(t, tt.out, out)
			tt.checkErr(t, err)
		})
	}
}

func TestJob_Create(t *testing.T) {
	// t.Parallel()
	mockTime := "2020/02/02 00:00:00"
	type in struct {
		company string
		job     string
	}
	tests := []struct {
		name     string
		injector func(m *mock_repository.MockJobRepository)
		in       in
		out      int
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to create",
			injector: func(m *mock_repository.MockJobRepository) {
				m.EXPECT().Create(&entity.Job{
					Job:       "job001",
					Company:   "com001",
					CreatedAt: mockTime,
					UpdatedAt: mockTime,
				}).Return(0, errors.New("some error"))
			},
			in: in{
				job:     "job001",
				company: "com001",
			},
			out:      0,
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_repository.MockJobRepository) {
				m.EXPECT().Create(&entity.Job{
					Job:       "job001",
					Company:   "com001",
					CreatedAt: mockTime,
					UpdatedAt: mockTime,
				}).Return(1, nil)
			},
			in: in{
				job:     "job001",
				company: "com001",
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
			m := mock_repository.NewMockJobRepository(ctrl)
			tt.injector(m)
			patch := monkey.Patch(helper.FormattedDateTimeNow, func() string { return mockTime })
			defer patch.Restore()
			srv := NewJob(m)
			out, err := srv.Create(tt.in.company, tt.in.job)
			assert.Equal(t, tt.out, out)
			tt.checkErr(t, err)
		})
	}
}

func TestJob_UpdateByID(t *testing.T) {
	// t.Parallel()
	mockTime := "2020/02/02 00:00:00"
	type in struct {
		id      int
		company string
		job     string
	}
	tests := []struct {
		name     string
		injector func(m *mock_repository.MockJobRepository)
		in       in
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to get data",
			injector: func(m *mock_repository.MockJobRepository) {
				m.EXPECT().FindByID(1).Return(nil, errors.New("some error"))
			},
			in: in{
				id:      1,
				company: "com001",
				job:     "job001",
			},
			checkErr: assert.NotNil,
		},
		{
			name: "failed to update",
			injector: func(m *mock_repository.MockJobRepository) {
				m.EXPECT().FindByID(1).Return(&entity.Job{
					ID:        1,
					Job:       "job001",
					Company:   "com001",
					CreatedAt: mockTime,
					UpdatedAt: mockTime,
				}, nil)
				m.EXPECT().UpdateByID(&entity.Job{
					ID:        1,
					Job:       "job001",
					Company:   "com001",
					CreatedAt: mockTime,
					UpdatedAt: mockTime,
				}).Return(errors.New("some error"))
			},
			in: in{
				id:      1,
				company: "com001",
				job:     "job001",
			},
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_repository.MockJobRepository) {
				m.EXPECT().FindByID(1).Return(&entity.Job{
					ID:        1,
					Job:       "job001",
					Company:   "com001",
					CreatedAt: mockTime,
					UpdatedAt: mockTime,
				}, nil)
				m.EXPECT().UpdateByID(&entity.Job{
					ID:        1,
					Job:       "job001",
					Company:   "com001",
					CreatedAt: mockTime,
					UpdatedAt: mockTime,
				}).Return(nil)
			},
			in: in{
				id:      1,
				company: "com001",
				job:     "job001",
			},
			checkErr: assert.Nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			m := mock_repository.NewMockJobRepository(ctrl)
			tt.injector(m)
			patch := monkey.Patch(helper.FormattedDateTimeNow, func() string { return mockTime })
			defer patch.Restore()
			srv := NewJob(m)
			err := srv.UpdateByID(tt.in.id, tt.in.company, tt.in.job)
			tt.checkErr(t, err)
		})
	}
}

func TestJob_DeleteByID(t *testing.T) {
	// t.Parallel()
	tests := []struct {
		name     string
		injector func(m *mock_repository.MockJobRepository)
		in       int
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to delete",
			injector: func(m *mock_repository.MockJobRepository) {
				m.EXPECT().DeleteByID(1).Return(errors.New("some error"))
			},
			in:       1,
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_repository.MockJobRepository) {
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
			m := mock_repository.NewMockJobRepository(ctrl)
			tt.injector(m)
			srv := NewJob(m)
			err := srv.DeleteByID(tt.in)
			tt.checkErr(t, err)
		})
	}
}
