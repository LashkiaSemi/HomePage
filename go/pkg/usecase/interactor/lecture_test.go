package interactor

import (
	mock_service "homepage/mock/service"
	"homepage/pkg/domain/entity"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestLecture_GetAll(t *testing.T) {
	tests := []struct {
		name     string
		injector func(ml *mock_service.MockLecture, mu *mock_service.MockUser)
		out      []*entity.Lecture
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to get",
			injector: func(ml *mock_service.MockLecture, mu *mock_service.MockUser) {
				ml.EXPECT().GetAll().Return(nil, errors.New("some error"))
			},
			out:      nil,
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(ml *mock_service.MockLecture, mu *mock_service.MockUser) {
				ml.EXPECT().GetAll().Return([]*entity.Lecture{
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
			ml := mock_service.NewMockLecture(ctrl)
			mu := mock_service.NewMockUser(ctrl)
			tt.injector(ml, mu)
			s := NewLectureInteractor(ml, mu)
			out, err := s.GetAll()
			assert.Equal(t, tt.out, out)
			tt.checkErr(t, err)
		})
	}
}

func TestLecture_GetByID(t *testing.T) {
	tests := []struct {
		name     string
		injector func(ml *mock_service.MockLecture, mu *mock_service.MockUser)
		in       int
		out      *entity.Lecture
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "fialed to get",
			injector: func(ml *mock_service.MockLecture, mu *mock_service.MockUser) {
				ml.EXPECT().GetByID(1).Return(nil, errors.New("some error"))
			},
			in:       1,
			out:      nil,
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(ml *mock_service.MockLecture, mu *mock_service.MockUser) {
				ml.EXPECT().GetByID(1).Return(&entity.Lecture{
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
			ml := mock_service.NewMockLecture(ctrl)
			mu := mock_service.NewMockUser(ctrl)
			tt.injector(ml, mu)
			s := NewLectureInteractor(ml, mu)
			out, err := s.GetByID(tt.in)
			assert.Equal(t, tt.out, out)
			tt.checkErr(t, err)
		})
	}
}

func TestLecture_Create(t *testing.T) {
	type in struct {
		studentID  string
		title      string
		file       string
		comment    string
		activation int
	}
	tests := []struct {
		name     string
		injector func(ml *mock_service.MockLecture, mu *mock_service.MockUser)
		in       in
		out      int
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to get author",
			injector: func(ml *mock_service.MockLecture, mu *mock_service.MockUser) {
				mu.EXPECT().GetByStudentID("student001").Return(nil, errors.New("some error"))
			},
			in: in{
				studentID: "student001",
			},
			out:      0,
			checkErr: assert.NotNil,
		},
		{
			name: "failed to create",
			injector: func(ml *mock_service.MockLecture, mu *mock_service.MockUser) {
				mu.EXPECT().GetByStudentID("student001").Return(&entity.User{
					ID: 1,
				}, nil)
				ml.EXPECT().Create("title001", "file001", "comment001", 1, &entity.User{
					ID: 1,
				}).Return(0, errors.New("some error"))
			},
			in: in{
				studentID:  "student001",
				title:      "title001",
				file:       "file001",
				comment:    "comment001",
				activation: 1,
			},
			out:      0,
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(ml *mock_service.MockLecture, mu *mock_service.MockUser) {
				mu.EXPECT().GetByStudentID("student001").Return(&entity.User{
					ID: 1,
				}, nil)
				ml.EXPECT().Create("title001", "file001", "comment001", 1, &entity.User{
					ID: 1,
				}).Return(1, nil)
			},
			in: in{
				studentID:  "student001",
				title:      "title001",
				file:       "file001",
				comment:    "comment001",
				activation: 1,
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
			ml := mock_service.NewMockLecture(ctrl)
			mu := mock_service.NewMockUser(ctrl)
			tt.injector(ml, mu)
			s := NewLectureInteractor(ml, mu)
			out, err := s.Create(tt.in.studentID, tt.in.title, tt.in.file, tt.in.comment, tt.in.activation)
			assert.Equal(t, tt.out, out)
			tt.checkErr(t, err)
		})
	}
}

func TestLecture_UpdateByID(t *testing.T) {
	type in struct {
		id         int
		studentID  string
		title      string
		file       string
		comment    string
		activation int
	}
	tests := []struct {
		name     string
		injector func(ml *mock_service.MockLecture, mu *mock_service.MockUser)
		in       in
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to get author",
			injector: func(ml *mock_service.MockLecture, mu *mock_service.MockUser) {
				mu.EXPECT().GetByStudentID("student001").Return(nil, errors.New("some error"))
			},
			in: in{
				studentID: "student001",
			},
			checkErr: assert.NotNil,
		},
		{
			name: "failed to update",
			injector: func(ml *mock_service.MockLecture, mu *mock_service.MockUser) {
				mu.EXPECT().GetByStudentID("student001").Return(&entity.User{
					ID: 1,
				}, nil)
				ml.EXPECT().UpdateByID(1, "title001", "file001", "comment001", 1, &entity.User{
					ID: 1,
				}).Return(errors.New("some error"))
			},
			in: in{
				id:         1,
				studentID:  "student001",
				title:      "title001",
				file:       "file001",
				comment:    "comment001",
				activation: 1,
			},
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(ml *mock_service.MockLecture, mu *mock_service.MockUser) {
				mu.EXPECT().GetByStudentID("student001").Return(&entity.User{
					ID: 1,
				}, nil)
				ml.EXPECT().UpdateByID(1, "title001", "file001", "comment001", 1, &entity.User{
					ID: 1,
				}).Return(nil)
			},
			in: in{
				id:         1,
				studentID:  "student001",
				title:      "title001",
				file:       "file001",
				comment:    "comment001",
				activation: 1,
			},
			checkErr: assert.Nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			ml := mock_service.NewMockLecture(ctrl)
			mu := mock_service.NewMockUser(ctrl)
			tt.injector(ml, mu)
			s := NewLectureInteractor(ml, mu)
			err := s.UpdateByID(tt.in.id, tt.in.studentID, tt.in.title, tt.in.file, tt.in.comment, tt.in.activation)
			tt.checkErr(t, err)
		})
	}
}

func TestLecture_DeleteByID(t *testing.T) {
	tests := []struct {
		name     string
		injector func(ml *mock_service.MockLecture, mu *mock_service.MockUser)
		in       int
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to delete",
			injector: func(ml *mock_service.MockLecture, mu *mock_service.MockUser) {
				ml.EXPECT().DeleteByID(1).Return(errors.New("some error"))
			},
			in:       1,
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(ml *mock_service.MockLecture, mu *mock_service.MockUser) {
				ml.EXPECT().DeleteByID(1).Return(nil)
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
			ml := mock_service.NewMockLecture(ctrl)
			mu := mock_service.NewMockUser(ctrl)
			tt.injector(ml, mu)
			s := NewLectureInteractor(ml, mu)
			err := s.DeleteByID(tt.in)
			tt.checkErr(t, err)
		})
	}
}
