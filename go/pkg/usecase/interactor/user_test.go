package interactor

import (
	mock_service "homepage/mock/service"
	"homepage/pkg/domain/entity"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestUser_GetAll(t *testing.T) {
	tests := []struct {
		name     string
		injector func(m *mock_service.MockUser)
		out      []*entity.User
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to get",
			injector: func(m *mock_service.MockUser) {
				m.EXPECT().GetAll().Return(nil, errors.New("some error"))
			},
			out:      nil,
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_service.MockUser) {
				m.EXPECT().GetAll().Return([]*entity.User{
					{
						ID: 1,
					},
				}, nil)
			},
			out: []*entity.User{
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
			m := mock_service.NewMockUser(ctrl)
			tt.injector(m)
			u := NewUserInteractor(m)
			out, err := u.GetAll()
			assert.Equal(t, tt.out, out)
			tt.checkErr(t, err)
		})
	}
}

func TestUser_GetByID(t *testing.T) {
	tests := []struct {
		name     string
		injector func(m *mock_service.MockUser)
		in       int
		out      *entity.User
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to get",
			injector: func(m *mock_service.MockUser) {
				m.EXPECT().GetByID(1).Return(nil, errors.New("some error"))
			},
			in:       1,
			out:      nil,
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_service.MockUser) {
				m.EXPECT().GetByID(1).Return(&entity.User{
					ID: 1,
				}, nil)
			},
			in: 1,
			out: &entity.User{
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
			m := mock_service.NewMockUser(ctrl)
			tt.injector(m)
			u := NewUserInteractor(m)
			out, err := u.GetByID(tt.in)
			assert.Equal(t, tt.out, out)
			tt.checkErr(t, err)
		})
	}
}

func TestUser_GetByStudentID(t *testing.T) {
	tests := []struct {
		name     string
		injector func(m *mock_service.MockUser)
		in       string
		out      *entity.User
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to get",
			injector: func(m *mock_service.MockUser) {
				m.EXPECT().GetByStudentID("student001").Return(nil, errors.New("some error"))
			},
			in:       "student001",
			out:      nil,
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_service.MockUser) {
				m.EXPECT().GetByStudentID("student001").Return(&entity.User{
					ID: 1,
				}, nil)
			},
			in: "student001",
			out: &entity.User{
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
			m := mock_service.NewMockUser(ctrl)
			tt.injector(m)
			u := NewUserInteractor(m)
			out, err := u.GetByStudentID(tt.in)
			assert.Equal(t, tt.out, out)
			tt.checkErr(t, err)
		})
	}
}

func TestUser_UpdateByID(t *testing.T) {
	type in struct {
		userID     int
		name       string
		studentID  string
		department string
		comment    string
		grade      int
	}
	tests := []struct {
		name     string
		injector func(m *mock_service.MockUser)
		in       in
		out      *entity.User
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to update",
			injector: func(m *mock_service.MockUser) {
				m.EXPECT().UpdateByID(1, "", "", "", "", 0).Return(nil, errors.New("some error"))
			},
			in: in{
				userID: 1,
			},
			out:      nil,
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_service.MockUser) {
				m.EXPECT().UpdateByID(1, "name001", "student001", "department", "comment", 2).Return(&entity.User{
					ID: 1,
				}, nil)
			},
			in: in{
				userID:     1,
				name:       "name001",
				studentID:  "student001",
				department: "department",
				comment:    "comment",
				grade:      2,
			},
			out: &entity.User{
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
			m := mock_service.NewMockUser(ctrl)
			tt.injector(m)
			u := NewUserInteractor(m)
			out, err := u.UpdateByID(tt.in.userID, tt.in.name, tt.in.studentID, tt.in.department, tt.in.comment, tt.in.grade)
			assert.Equal(t, tt.out, out)
			tt.checkErr(t, err)
		})
	}
}

func TestUser_UpdatePasswordByStudentID(t *testing.T) {
	type in struct {
		studentID   string
		oldPassword string
		newPassword string
	}
	tests := []struct {
		name     string
		injector func(m *mock_service.MockUser)
		in       in
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to update",
			injector: func(m *mock_service.MockUser) {
				m.EXPECT().UpdatePasswordByStudentID("student001", "old001", "new001").Return(errors.New("some error"))
			},
			in: in{
				studentID:   "student001",
				oldPassword: "old001",
				newPassword: "new001",
			},
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_service.MockUser) {
				m.EXPECT().UpdatePasswordByStudentID("student001", "old001", "new001").Return(nil)
			},
			in: in{
				studentID:   "student001",
				oldPassword: "old001",
				newPassword: "new001",
			},
			checkErr: assert.Nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			m := mock_service.NewMockUser(ctrl)
			tt.injector(m)
			u := NewUserInteractor(m)
			err := u.UpdatePasswordByStudentID(tt.in.studentID, tt.in.oldPassword, tt.in.newPassword)
			tt.checkErr(t, err)
		})
	}
}

func TestUser_AuthenticationByStudentID(t *testing.T) {
	type in struct {
		studentID string
		password  string
	}
	tests := []struct {
		name     string
		injector func(m *mock_service.MockUser)
		in       in
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to auth",
			injector: func(m *mock_service.MockUser) {
				m.EXPECT().AuthenticationByStudentID("student001", "password").Return(errors.New("some error"))
			},
			in: in{
				studentID: "student001",
				password:  "password",
			},
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_service.MockUser) {
				m.EXPECT().AuthenticationByStudentID("student001", "password").Return(nil)
			},
			in: in{
				studentID: "student001",
				password:  "password",
			},
			checkErr: assert.Nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			m := mock_service.NewMockUser(ctrl)
			tt.injector(m)
			u := NewUserInteractor(m)
			err := u.AuthenticationByStudentID(tt.in.studentID, tt.in.password)
			tt.checkErr(t, err)
		})
	}
}

func TestUser_AdminAuthorizationByStudentID(t *testing.T) {
	type in struct {
		studentID string
		password  string
	}
	tests := []struct {
		name     string
		injector func(m *mock_service.MockUser)
		in       in
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to auth",
			injector: func(m *mock_service.MockUser) {
				m.EXPECT().AdminAuthorizationByStudentID("student001", "password").Return(errors.New("some error"))
			},
			in: in{
				studentID: "student001",
				password:  "password",
			},
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_service.MockUser) {
				m.EXPECT().AdminAuthorizationByStudentID("student001", "password").Return(nil)
			},
			in: in{
				studentID: "student001",
				password:  "password",
			},
			checkErr: assert.Nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			m := mock_service.NewMockUser(ctrl)
			tt.injector(m)
			u := NewUserInteractor(m)
			err := u.AdminAuthorizationByStudentID(tt.in.studentID, tt.in.password)
			tt.checkErr(t, err)
		})
	}
}

func TestUser_AdminCreate(t *testing.T) {
	type in struct {
		name       string
		studentID  string
		password   string
		role       string
		department string
		comment    string
		grade      int
	}
	tests := []struct {
		name     string
		injector func(m *mock_service.MockUser)
		in       in
		out      int
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to create",
			injector: func(m *mock_service.MockUser) {
				m.EXPECT().AdminCreate("name001", "student001", "password", "admin", "dep001", "comment001", 2).Return(0, errors.New("some error"))
			},
			in: in{
				name:       "name001",
				studentID:  "student001",
				password:   "password",
				role:       "admin",
				department: "dep001",
				comment:    "comment001",
				grade:      2,
			},
			out:      0,
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_service.MockUser) {
				m.EXPECT().AdminCreate("name001", "student001", "password", "admin", "dep001", "comment001", 2).Return(1, nil)
			},
			in: in{
				name:       "name001",
				studentID:  "student001",
				password:   "password",
				role:       "admin",
				department: "dep001",
				comment:    "comment001",
				grade:      2,
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
			m := mock_service.NewMockUser(ctrl)
			tt.injector(m)
			u := NewUserInteractor(m)
			out, err := u.AdminCreate(tt.in.name, tt.in.studentID, tt.in.password, tt.in.role, tt.in.department, tt.in.comment, tt.in.grade)
			assert.Equal(t, tt.out, out)
			tt.checkErr(t, err)
		})
	}
}

func TestUser_AdminUpdate(t *testing.T) {
	type in struct {
		userID     int
		name       string
		studentID  string
		role       string
		department string
		comment    string
		grade      int
	}
	tests := []struct {
		name     string
		injector func(m *mock_service.MockUser)
		in       in
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to update",
			injector: func(m *mock_service.MockUser) {
				m.EXPECT().AdminUpdate(1, "name001", "student001", "admin", "dep001", "comment001", 2).Return(errors.New("some error"))
			},
			in: in{
				userID:     1,
				name:       "name001",
				studentID:  "student001",
				role:       "admin",
				department: "dep001",
				comment:    "comment001",
				grade:      2,
			},
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_service.MockUser) {
				m.EXPECT().AdminUpdate(1, "name001", "student001", "admin", "dep001", "comment001", 2).Return(nil)
			},
			in: in{
				userID:     1,
				name:       "name001",
				studentID:  "student001",
				role:       "admin",
				department: "dep001",
				comment:    "comment001",
				grade:      2,
			},
			checkErr: assert.Nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			m := mock_service.NewMockUser(ctrl)
			tt.injector(m)
			u := NewUserInteractor(m)
			err := u.AdminUpdate(tt.in.userID, tt.in.name, tt.in.studentID, tt.in.role, tt.in.department, tt.in.comment, tt.in.grade)
			tt.checkErr(t, err)
		})
	}
}

func TestUser_DeleteByID(t *testing.T) {
	tests := []struct {
		name     string
		injector func(m *mock_service.MockUser)
		in       int
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to delete",
			injector: func(m *mock_service.MockUser) {
				m.EXPECT().DeleteByID(1).Return(errors.New("some error"))
			},
			in:       1,
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_service.MockUser) {
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
			m := mock_service.NewMockUser(ctrl)
			tt.injector(m)
			u := NewUserInteractor(m)
			err := u.DeleteByID(tt.in)
			tt.checkErr(t, err)
		})
	}
}
