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

func TestLecture_GetAll(t *testing.T) {
	// t.Parallel()

	tests := []struct {
		name     string
		injector func(m *mock_repository.MockLectureRepository)
		out      []*entity.Lecture
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to get",
			injector: func(m *mock_repository.MockLectureRepository) {
				m.EXPECT().FindAll().Return(nil, errors.New("some error"))
			},
			out:      nil,
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_repository.MockLectureRepository) {
				m.EXPECT().FindAll().Return([]*entity.Lecture{
					{
						ID: 1,
					},
				}, nil)
			},
			out: []*entity.Lecture{
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
			m := mock_repository.NewMockLectureRepository(ctrl)
			tt.injector(m)
			srv := NewLecture(m)
			out, err := srv.GetAll()
			assert.Equal(t, tt.out, out)
			tt.checkErr(t, err)
		})
	}
}

func TestLecture_GetByID(t *testing.T) {
	// t.Parallel()

	mockTime := "2020/02/02 00:00:00"
	tests := []struct {
		name     string
		injector func(m *mock_repository.MockLectureRepository)
		in       int
		out      *entity.Lecture
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to get",
			injector: func(m *mock_repository.MockLectureRepository) {
				m.EXPECT().FindByID(1).Return(nil, errors.New("some error"))
			},
			in:       1,
			out:      nil,
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_repository.MockLectureRepository) {
				m.EXPECT().FindByID(1).Return(&entity.Lecture{
					ID: 1,
				}, nil)
			},
			in: 1,
			out: &entity.Lecture{
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
			m := mock_repository.NewMockLectureRepository(ctrl)
			tt.injector(m)
			patch := monkey.Patch(helper.FormattedDateTimeNow, func() string { return mockTime })
			defer patch.Restore()
			srv := NewLecture(m)
			out, err := srv.GetByID(tt.in)
			assert.Equal(t, tt.out, out)
			tt.checkErr(t, err)
		})
	}
}

func TestLecture_Create(t *testing.T) {
	// t.Parallel()

	mockTime := "2020/02/02 00:00:00"
	user := entity.NewUser(
		"name001",
		"t317013",
		"password",
		"admin",
		"dep",
		"comment",
		4,
	)
	type in struct {
		title      string
		file       string
		comment    string
		activation int
		author     *entity.User
	}
	tests := []struct {
		name     string
		injector func(m *mock_repository.MockLectureRepository)
		in       in
		out      int
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to create",
			injector: func(m *mock_repository.MockLectureRepository) {
				m.EXPECT().Create(&entity.Lecture{
					Title:      "title001",
					File:       "file001",
					Comment:    "comment001",
					Activation: 1,
					Author:     user,
					CreatedAt:  mockTime,
					UpdatedAt:  mockTime,
				}).Return(0, errors.New("some error"))
			},
			in: in{
				title:      "title001",
				file:       "file001",
				comment:    "comment001",
				activation: 1,
				author:     user,
			},
			out:      0,
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_repository.MockLectureRepository) {
				m.EXPECT().Create(&entity.Lecture{
					Title:      "title001",
					File:       "file001",
					Comment:    "comment001",
					Activation: 1,
					Author:     user,
					CreatedAt:  mockTime,
					UpdatedAt:  mockTime,
				}).Return(1, nil)
			},
			in: in{
				title:      "title001",
				file:       "file001",
				comment:    "comment001",
				activation: 1,
				author:     user,
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
			m := mock_repository.NewMockLectureRepository(ctrl)
			tt.injector(m)
			patch := monkey.Patch(helper.FormattedDateTimeNow, func() string { return mockTime })
			defer patch.Restore()
			srv := NewLecture(m)
			out, err := srv.Create(tt.in.title, tt.in.file, tt.in.comment, tt.in.activation, tt.in.author)
			assert.Equal(t, tt.out, out)
			tt.checkErr(t, err)
		})
	}
}

func TestLecture_UpdateByID(t *testing.T) {
	// t.Parallel()

	mockTime := "2020/02/02 00:00:00"
	user := entity.NewUser(
		"name001",
		"t317013",
		"password",
		"admin",
		"dep",
		"comment",
		4,
	)
	type in struct {
		id         int
		title      string
		file       string
		comment    string
		activation int
		author     *entity.User
	}
	tests := []struct {
		name     string
		injector func(m *mock_repository.MockLectureRepository)
		in       in
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to get data",
			injector: func(m *mock_repository.MockLectureRepository) {
				m.EXPECT().FindByID(1).Return(nil, errors.New("some error"))
			},
			in: in{
				id: 1,
			},
			checkErr: assert.NotNil,
		},
		{
			name: "failed to update",
			injector: func(m *mock_repository.MockLectureRepository) {
				m.EXPECT().FindByID(1).Return(&entity.Lecture{
					ID:         1,
					Title:      "title001",
					File:       "file001",
					Comment:    "comment001",
					Activation: 1,
					Author:     user,
					CreatedAt:  mockTime,
					UpdatedAt:  mockTime,
				}, nil)
				m.EXPECT().UpdateByID(&entity.Lecture{
					ID:         1,
					Title:      "title001",
					File:       "file001",
					Comment:    "comment001",
					Activation: 1,
					Author:     user,
					CreatedAt:  mockTime,
					UpdatedAt:  mockTime,
				}).Return(errors.New("some error"))
			},
			in: in{
				id:         1,
				title:      "title001",
				file:       "file001",
				comment:    "comment001",
				activation: 1,
				author:     user,
			},
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_repository.MockLectureRepository) {
				m.EXPECT().FindByID(1).Return(&entity.Lecture{
					ID:         1,
					Title:      "title001",
					File:       "file001",
					Comment:    "comment001",
					Activation: 1,
					Author:     user,
					CreatedAt:  mockTime,
					UpdatedAt:  mockTime,
				}, nil)
				m.EXPECT().UpdateByID(&entity.Lecture{
					ID:         1,
					Title:      "title001",
					File:       "file001",
					Comment:    "comment001",
					Activation: 1,
					Author:     user,
					CreatedAt:  mockTime,
					UpdatedAt:  mockTime,
				}).Return(nil)
			},
			in: in{
				id:         1,
				title:      "title001",
				file:       "file001",
				comment:    "comment001",
				activation: 1,
				author:     user,
			},
			checkErr: assert.Nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			m := mock_repository.NewMockLectureRepository(ctrl)
			tt.injector(m)
			patch := monkey.Patch(helper.FormattedDateTimeNow, func() string { return mockTime })
			defer patch.Restore()
			srv := NewLecture(m)
			err := srv.UpdateByID(tt.in.id, tt.in.title, tt.in.file, tt.in.comment, tt.in.activation, tt.in.author)
			tt.checkErr(t, err)
		})
	}
}

func TestLecture_DeleteByID(t *testing.T) {
	// t.Parallel()

	tests := []struct {
		name     string
		injector func(m *mock_repository.MockLectureRepository)
		in       int
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to delete",
			injector: func(m *mock_repository.MockLectureRepository) {
				m.EXPECT().DeleteByID(1).Return(errors.New("some error"))
			},
			in:       1,
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_repository.MockLectureRepository) {
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
			m := mock_repository.NewMockLectureRepository(ctrl)
			tt.injector(m)
			srv := NewLecture(m)
			err := srv.DeleteByID(tt.in)
			tt.checkErr(t, err)
		})
	}
}
