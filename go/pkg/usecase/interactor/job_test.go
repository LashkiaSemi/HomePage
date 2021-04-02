package interactor

import (
	mock_service "homepage/mock/service"
	"homepage/pkg/domain/entity"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestJob_GetAll(t *testing.T) {
	tests := []struct {
		name     string
		injector func(m *mock_service.MockJob)
		out      []*entity.Job
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to get",
			injector: func(m *mock_service.MockJob) {
				m.EXPECT().GetAll().Return(nil, errors.New("some error"))
			},
			out:      nil,
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_service.MockJob) {
				m.EXPECT().GetAll().Return([]*entity.Job{
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
			m := mock_service.NewMockJob(ctrl)
			tt.injector(m)
			u := NewJobInteractor(m)
			out, err := u.GetAll()
			assert.Equal(t, tt.out, out)
			tt.checkErr(t, err)
		})
	}
}

func TestJob_GetByID(t *testing.T) {
	tests := []struct {
		name     string
		injector func(m *mock_service.MockJob)
		in       int
		out      *entity.Job
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to get",
			injector: func(m *mock_service.MockJob) {
				m.EXPECT().GetByID(1).Return(nil, errors.New("some error"))
			},
			in:       1,
			out:      nil,
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_service.MockJob) {
				m.EXPECT().GetByID(1).Return(&entity.Job{
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
			m := mock_service.NewMockJob(ctrl)
			tt.injector(m)
			u := NewJobInteractor(m)
			out, err := u.GetByID(tt.in)
			assert.Equal(t, tt.out, out)
			tt.checkErr(t, err)
		})
	}
}

func TestJob_Create(t *testing.T) {
	type in struct {
		company string
		job     string
	}
	tests := []struct {
		name     string
		injector func(m *mock_service.MockJob)
		in       in
		out      int
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to create",
			injector: func(m *mock_service.MockJob) {
				m.EXPECT().Create("comp001", "job001").Return(0, errors.New("some error"))
			},
			in: in{
				company: "comp001",
				job:     "job001",
			},
			out:      0,
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_service.MockJob) {
				m.EXPECT().Create("comp001", "job001").Return(1, nil)
			},
			in: in{
				company: "comp001",
				job:     "job001",
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
			m := mock_service.NewMockJob(ctrl)
			tt.injector(m)
			u := NewJobInteractor(m)
			out, err := u.Create(tt.in.company, tt.in.job)
			assert.Equal(t, tt.out, out)
			tt.checkErr(t, err)
		})
	}
}

func TestJob_UpdateByID(t *testing.T) {
	type in struct {
		id      int
		company string
		job     string
	}
	tests := []struct {
		name     string
		injector func(m *mock_service.MockJob)
		in       in
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to update",
			injector: func(m *mock_service.MockJob) {
				m.EXPECT().UpdateByID(1, "comp001", "job001").Return(errors.New("some error"))
			},
			in: in{
				id:      1,
				company: "comp001",
				job:     "job001",
			},
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_service.MockJob) {
				m.EXPECT().UpdateByID(1, "comp001", "job001").Return(nil)
			},
			in: in{
				id:      1,
				company: "comp001",
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
			m := mock_service.NewMockJob(ctrl)
			tt.injector(m)
			u := NewJobInteractor(m)
			err := u.UpdateByID(tt.in.id, tt.in.company, tt.in.job)
			tt.checkErr(t, err)
		})
	}
}

func TestJob_DeleteByID(t *testing.T) {
	tests := []struct {
		name     string
		injector func(m *mock_service.MockJob)
		in       int
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to delete",
			injector: func(m *mock_service.MockJob) {
				m.EXPECT().DeleteByID(1).Return(errors.New("some error"))
			},
			in:       1,
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_service.MockJob) {
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
			m := mock_service.NewMockJob(ctrl)
			tt.injector(m)
			u := NewJobInteractor(m)
			err := u.DeleteByID(tt.in)
			tt.checkErr(t, err)
		})
	}
}
