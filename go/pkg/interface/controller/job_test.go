package controller

import (
	mock_interactor "homepage/mock/interactor"
	"homepage/pkg/domain/entity"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestJob_GetAll(t *testing.T) {
	tests := []struct {
		name     string
		injector func(m *mock_interactor.MockJobInteractor)
		out      *JobsResponse
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to get",
			injector: func(m *mock_interactor.MockJobInteractor) {
				m.EXPECT().GetAll().Return(nil, errors.New("some error"))
			},
			out:      &JobsResponse{},
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_interactor.MockJobInteractor) {
				m.EXPECT().GetAll().Return([]*entity.Job{
					{
						ID: 1,
					},
				}, nil)
			},
			out: &JobsResponse{
				Jobs: []*JobResponse{
					{
						Company: "",
						Job:     "",
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
			m := mock_interactor.NewMockJobInteractor(ctrl)
			tt.injector(m)
			c := NewJobController(m)
			out, err := c.GetAll()
			assert.Equal(t, tt.out, out)
			tt.checkErr(t, err)
		})
	}
}

func TestJob_GetByID(t *testing.T) {
	tests := []struct {
		name     string
		injector func(m *mock_interactor.MockJobInteractor)
		in       int
		out      *JobResponse
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to get",
			injector: func(m *mock_interactor.MockJobInteractor) {
				m.EXPECT().GetByID(0).Return(nil, errors.New("some error"))
			},
			in:       0,
			out:      &JobResponse{},
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_interactor.MockJobInteractor) {
				m.EXPECT().GetByID(1).Return(&entity.Job{
					ID:      1,
					Company: "comp001",
					Job:     "job001",
				}, nil)
			},
			in: 1,
			out: &JobResponse{
				Company: "comp001",
				Job:     "job001",
			},
			checkErr: assert.Nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			m := mock_interactor.NewMockJobInteractor(ctrl)
			tt.injector(m)
			c := NewJobController(m)
			out, err := c.GetByID(tt.in)
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
		injector func(m *mock_interactor.MockJobInteractor)
		in       in
		out      int
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to create",
			injector: func(m *mock_interactor.MockJobInteractor) {
				m.EXPECT().Create("", "").Return(0, errors.New("some error"))
			},
			in:       in{},
			out:      0,
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_interactor.MockJobInteractor) {
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
			m := mock_interactor.NewMockJobInteractor(ctrl)
			tt.injector(m)
			c := NewJobController(m)
			out, err := c.Create(tt.in.company, tt.in.job)
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
		injector func(m *mock_interactor.MockJobInteractor)
		in       in
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to update",
			injector: func(m *mock_interactor.MockJobInteractor) {
				m.EXPECT().UpdateByID(0, "", "").Return(errors.New("some error"))
			},
			in:       in{},
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_interactor.MockJobInteractor) {
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
			m := mock_interactor.NewMockJobInteractor(ctrl)
			tt.injector(m)
			c := NewJobController(m)
			err := c.UpdateByID(tt.in.id, tt.in.company, tt.in.job)
			tt.checkErr(t, err)
		})
	}
}

func TestJob_DeleteByID(t *testing.T) {
	tests := []struct {
		name     string
		injector func(m *mock_interactor.MockJobInteractor)
		in       int
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to delete",
			injector: func(m *mock_interactor.MockJobInteractor) {
				m.EXPECT().DeleteByID(0).Return(errors.New("some error"))
			},
			in:       0,
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_interactor.MockJobInteractor) {
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
			m := mock_interactor.NewMockJobInteractor(ctrl)
			tt.injector(m)
			c := NewJobController(m)
			err := c.DeleteByID(tt.in)
			tt.checkErr(t, err)
		})
	}
}

func TestJob_AdminGetAll(t *testing.T) {
	tests := []struct {
		name     string
		injector func(m *mock_interactor.MockJobInteractor)
		out      []map[string]string
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to get",
			injector: func(m *mock_interactor.MockJobInteractor) {
				m.EXPECT().GetAll().Return(nil, errors.New("some error"))
			},
			out:      []map[string]string(nil),
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_interactor.MockJobInteractor) {
				m.EXPECT().GetAll().Return([]*entity.Job{
					{
						ID:      1,
						Company: "comp001",
						Job:     "job001",
					},
				}, nil)
			},
			out: []map[string]string{
				{
					"id":    "1",
					"title": "comp001 / job001",
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
			m := mock_interactor.NewMockJobInteractor(ctrl)
			tt.injector(m)
			c := NewJobController(m)
			out, err := c.AdminGetAll()
			assert.Equal(t, tt.out, out)
			tt.checkErr(t, err)
		})
	}
}

func TestJob_AdminGetByID(t *testing.T) {
	tests := []struct {
		name     string
		injector func(m *mock_interactor.MockJobInteractor)
		in       int
		out      *FieldsResponse
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to get",
			injector: func(m *mock_interactor.MockJobInteractor) {
				m.EXPECT().GetByID(0).Return(nil, errors.New("some error"))
			},
			in:       0,
			out:      &FieldsResponse{},
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_interactor.MockJobInteractor) {
				m.EXPECT().GetByID(1).Return(&entity.Job{
					ID:      1,
					Company: "comp001",
					Job:     "job001",
				}, nil)
			},
			in: 1,
			out: &FieldsResponse{
				Fields: []*Field{
					{Key: "ID", Value: 1},
					{Key: "企業名", Value: "comp001"},
					{Key: "仕事", Value: "job001"},
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
			m := mock_interactor.NewMockJobInteractor(ctrl)
			tt.injector(m)
			c := NewJobController(m)
			out, err := c.AdminGetByID(tt.in)
			assert.Equal(t, tt.out, out)
			tt.checkErr(t, err)
		})
	}
}
