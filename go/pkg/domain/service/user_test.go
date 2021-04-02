package service

import (
	mock_repository "homepage/mock/repository"
	mock_service "homepage/mock/service"
	"homepage/pkg/domain/entity"
	"homepage/pkg/helper"
	"testing"

	"bou.ke/monkey"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestUser_GetAll(t *testing.T) {
	// t.Parallel()
	tests := []struct {
		name     string
		injector func(mr *mock_repository.MockUserRepository, mh *mock_service.MockVerifyHandler)
		out      []*entity.User
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "faile to get",
			injector: func(mr *mock_repository.MockUserRepository, mh *mock_service.MockVerifyHandler) {
				mr.EXPECT().FindAll().Return(nil, errors.New("some error"))
			},
			out:      nil,
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(mr *mock_repository.MockUserRepository, mh *mock_service.MockVerifyHandler) {
				mr.EXPECT().FindAll().Return([]*entity.User{
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
			mr := mock_repository.NewMockUserRepository(ctrl)
			mh := mock_service.NewMockVerifyHandler(ctrl)
			tt.injector(mr, mh)
			srv := NewUser(mr, mh)
			out, e := srv.GetAll()
			assert.Equal(t, tt.out, out)
			tt.checkErr(t, e)
		})
	}
}

func TestUser_GetByID(t *testing.T) {
	// t.Parallel()
	tests := []struct {
		name     string
		injector func(mr *mock_repository.MockUserRepository, mh *mock_service.MockVerifyHandler)
		in       int
		out      *entity.User
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to get",
			injector: func(mr *mock_repository.MockUserRepository, mh *mock_service.MockVerifyHandler) {
				mr.EXPECT().FindByID(1).Return(nil, errors.New("some error"))
			},
			in:       1,
			out:      nil,
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(mr *mock_repository.MockUserRepository, mh *mock_service.MockVerifyHandler) {
				mr.EXPECT().FindByID(1).Return(&entity.User{
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
			mr := mock_repository.NewMockUserRepository(ctrl)
			mh := mock_service.NewMockVerifyHandler(ctrl)
			tt.injector(mr, mh)
			srv := NewUser(mr, mh)
			out, e := srv.GetByID(tt.in)
			assert.Equal(t, tt.out, out)
			tt.checkErr(t, e)
		})
	}
}

func TestUser_GetByStudentID(t *testing.T) {
	// t.Parallel()
	tests := []struct {
		name     string
		injector func(mr *mock_repository.MockUserRepository, mh *mock_service.MockVerifyHandler)
		in       int
		out      *entity.User
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to get",
			injector: func(mr *mock_repository.MockUserRepository, mh *mock_service.MockVerifyHandler) {
				mr.EXPECT().FindByID(1).Return(nil, errors.New("some error"))
			},
			in:       1,
			out:      nil,
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(mr *mock_repository.MockUserRepository, mh *mock_service.MockVerifyHandler) {
				mr.EXPECT().FindByID(1).Return(&entity.User{
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
			mr := mock_repository.NewMockUserRepository(ctrl)
			mh := mock_service.NewMockVerifyHandler(ctrl)
			tt.injector(mr, mh)
			srv := NewUser(mr, mh)
			out, e := srv.GetByID(tt.in)
			assert.Equal(t, tt.out, out)
			tt.checkErr(t, e)
		})
	}
}

func TestUser_UpdateByID(t *testing.T) {
	// t.Parallel()
	mockTime := "2020/02/02 00:00:00"
	type in struct {
		id         int
		name       string
		studentID  string
		department string
		comment    string
		grade      int
	}
	tests := []struct {
		name     string
		injector func(mr *mock_repository.MockUserRepository, mh *mock_service.MockVerifyHandler)
		in       in
		out      *entity.User
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to get",
			injector: func(mr *mock_repository.MockUserRepository, mh *mock_service.MockVerifyHandler) {
				mr.EXPECT().FindByID(1).Return(nil, errors.New("some error"))
			},
			in: in{
				id: 1,
			},
			out:      nil,
			checkErr: assert.NotNil,
		},
		{
			name: "failed to update",
			injector: func(mr *mock_repository.MockUserRepository, mh *mock_service.MockVerifyHandler) {
				mr.EXPECT().FindByID(1).Return(&entity.User{
					ID:         1,
					StudentID:  "student001",
					Name:       "name001",
					Password:   "password",
					Department: "dep001",
					Grade:      2,
					Comment:    "comment001",
					CreatedAt:  mockTime,
					UpdatedAt:  mockTime,
				}, nil)
				mr.EXPECT().UpdateByID(&entity.User{
					ID:         1,
					StudentID:  "student001",
					Name:       "name001",
					Password:   "password",
					Department: "dep001",
					Grade:      2,
					Comment:    "comment001",
					CreatedAt:  mockTime,
					UpdatedAt:  mockTime,
				}).Return(errors.New("some error"))
			},
			in: in{
				id:         1,
				studentID:  "student001",
				name:       "name001",
				department: "dep001",
				comment:    "comment001",
				grade:      2,
			},
			out:      nil,
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(mr *mock_repository.MockUserRepository, mh *mock_service.MockVerifyHandler) {
				mr.EXPECT().FindByID(1).Return(&entity.User{
					ID:         1,
					StudentID:  "student001",
					Name:       "name001",
					Password:   "password",
					Department: "dep001",
					Grade:      2,
					Comment:    "comment001",
					CreatedAt:  mockTime,
					UpdatedAt:  mockTime,
				}, nil)
				mr.EXPECT().UpdateByID(&entity.User{
					ID:         1,
					StudentID:  "student001",
					Name:       "name001",
					Password:   "password",
					Department: "dep001",
					Grade:      2,
					Comment:    "comment001",
					CreatedAt:  mockTime,
					UpdatedAt:  mockTime,
				}).Return(nil)
			},
			in: in{
				id:         1,
				studentID:  "student001",
				name:       "name001",
				department: "dep001",
				comment:    "comment001",
				grade:      2,
			},
			out: &entity.User{
				ID:         1,
				StudentID:  "student001",
				Name:       "name001",
				Password:   "password",
				Department: "dep001",
				Grade:      2,
				Comment:    "comment001",
				CreatedAt:  mockTime,
				UpdatedAt:  mockTime,
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
			mr := mock_repository.NewMockUserRepository(ctrl)
			mh := mock_service.NewMockVerifyHandler(ctrl)
			tt.injector(mr, mh)
			srv := NewUser(mr, mh)
			out, e := srv.UpdateByID(tt.in.id, tt.in.name, tt.in.studentID, tt.in.department, tt.in.comment, tt.in.grade)
			assert.Equal(t, tt.out, out)
			tt.checkErr(t, e)
		})
	}
}

func TestUser_UpdatePasswordByStudentID(t *testing.T) {
	// t.Parallel()
	mockTime := "2020/02/02 00:00:00"
	type in struct {
		studentID   string
		oldPassword string
		newPassword string
	}
	tests := []struct {
		name     string
		injector func(mr *mock_repository.MockUserRepository, mh *mock_service.MockVerifyHandler)
		in       in
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to get auth info",
			injector: func(mr *mock_repository.MockUserRepository, mh *mock_service.MockVerifyHandler) {
				mr.EXPECT().FindAuthInfoByStudentID("student001").Return(nil, errors.New("some error"))
			},
			in: in{
				studentID: "student001",
			},
			checkErr: assert.NotNil,
		},
		{
			name: "failed to password verify",
			injector: func(mr *mock_repository.MockUserRepository, mh *mock_service.MockVerifyHandler) {
				mr.EXPECT().FindAuthInfoByStudentID("student001").Return(&entity.User{
					ID:         1,
					StudentID:  "student001",
					Name:       "name001",
					Password:   "password",
					Department: "dep001",
					Grade:      2,
					Comment:    "comment001",
					CreatedAt:  mockTime,
					UpdatedAt:  mockTime,
				}, nil)
				mh.EXPECT().PasswordVerify("password", "oldPassword").Return(errors.New("some error"))
			},
			in: in{
				studentID:   "student001",
				oldPassword: "oldPassword",
			},
			checkErr: assert.NotNil,
		},
		{
			name: "failed to password hash",
			injector: func(mr *mock_repository.MockUserRepository, mh *mock_service.MockVerifyHandler) {
				mr.EXPECT().FindAuthInfoByStudentID("student001").Return(&entity.User{
					ID:         1,
					StudentID:  "student001",
					Name:       "name001",
					Password:   "password",
					Department: "dep001",
					Grade:      2,
					Comment:    "comment001",
					CreatedAt:  mockTime,
					UpdatedAt:  mockTime,
				}, nil)
				mh.EXPECT().PasswordVerify("password", "oldPassword").Return(nil)
				mh.EXPECT().PasswordHash("newPassword").Return("", errors.New("some error"))
			},
			in: in{
				studentID:   "student001",
				oldPassword: "oldPassword",
				newPassword: "newPassword",
			},
			checkErr: assert.NotNil,
		},
		{
			name: "failed to update",
			injector: func(mr *mock_repository.MockUserRepository, mh *mock_service.MockVerifyHandler) {
				mr.EXPECT().FindAuthInfoByStudentID("student001").Return(&entity.User{
					ID:         1,
					StudentID:  "student001",
					Name:       "name001",
					Password:   "password",
					Department: "dep001",
					Grade:      2,
					Comment:    "comment001",
					CreatedAt:  mockTime,
					UpdatedAt:  mockTime,
				}, nil)
				mh.EXPECT().PasswordVerify("password", "oldPassword").Return(nil)
				mh.EXPECT().PasswordHash("newPassword").Return("hash", nil)
				mr.EXPECT().UpdatePasswordByStudentID("student001", "hash").Return(errors.New("some error"))
			},
			in: in{
				studentID:   "student001",
				oldPassword: "oldPassword",
				newPassword: "newPassword",
			},
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(mr *mock_repository.MockUserRepository, mh *mock_service.MockVerifyHandler) {
				mr.EXPECT().FindAuthInfoByStudentID("student001").Return(&entity.User{
					ID:         1,
					StudentID:  "student001",
					Name:       "name001",
					Password:   "password",
					Department: "dep001",
					Grade:      2,
					Comment:    "comment001",
					CreatedAt:  mockTime,
					UpdatedAt:  mockTime,
				}, nil)
				mh.EXPECT().PasswordVerify("password", "oldPassword").Return(nil)
				mh.EXPECT().PasswordHash("newPassword").Return("hash", nil)
				mr.EXPECT().UpdatePasswordByStudentID("student001", "hash").Return(nil)
			},
			in: in{
				studentID:   "student001",
				oldPassword: "oldPassword",
				newPassword: "newPassword",
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
			mr := mock_repository.NewMockUserRepository(ctrl)
			mh := mock_service.NewMockVerifyHandler(ctrl)
			tt.injector(mr, mh)
			srv := NewUser(mr, mh)
			e := srv.UpdatePasswordByStudentID(tt.in.studentID, tt.in.oldPassword, tt.in.newPassword)
			tt.checkErr(t, e)
		})
	}
}

func TestUser_DeleteByID(t *testing.T) {
	// t.Parallel()
	tests := []struct {
		name     string
		injector func(mr *mock_repository.MockUserRepository, mh *mock_service.MockVerifyHandler)
		in       int
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to delete",
			injector: func(mr *mock_repository.MockUserRepository, mh *mock_service.MockVerifyHandler) {
				mr.EXPECT().DeleteByID(1).Return(errors.New("some error"))
			},
			in:       1,
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(mr *mock_repository.MockUserRepository, mh *mock_service.MockVerifyHandler) {
				mr.EXPECT().DeleteByID(1).Return(nil)
			},
			in:       1,
			checkErr: assert.Nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mr := mock_repository.NewMockUserRepository(ctrl)
			mh := mock_service.NewMockVerifyHandler(ctrl)
			tt.injector(mr, mh)
			srv := NewUser(mr, mh)
			e := srv.DeleteByID(tt.in)
			tt.checkErr(t, e)
		})
	}
}

func TestUser_AuthenticationByStudentID(t *testing.T) {
	// t.Parallel()
	mockTime := "2020/02/02 00:00:00"
	type in struct {
		studentID string
		password  string
	}
	tests := []struct {
		name     string
		injector func(mr *mock_repository.MockUserRepository, mh *mock_service.MockVerifyHandler)
		in       in
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to get auth info",
			injector: func(mr *mock_repository.MockUserRepository, mh *mock_service.MockVerifyHandler) {
				mr.EXPECT().FindAuthInfoByStudentID("student001").Return(nil, errors.New("some error"))
			},
			in: in{
				studentID: "student001",
			},
			checkErr: assert.NotNil,
		},
		{
			name: "failed to password verify",
			injector: func(mr *mock_repository.MockUserRepository, mh *mock_service.MockVerifyHandler) {
				mr.EXPECT().FindAuthInfoByStudentID("student001").Return(&entity.User{
					ID:         1,
					StudentID:  "student001",
					Name:       "name001",
					Role:       "member",
					Password:   "hash",
					Department: "dep001",
					Grade:      2,
					Comment:    "comment001",
					CreatedAt:  mockTime,
					UpdatedAt:  mockTime,
				}, nil)
				mh.EXPECT().PasswordVerify("hash", "password").Return(errors.New("some error"))
			},
			in: in{
				studentID: "student001",
				password:  "password",
			},
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(mr *mock_repository.MockUserRepository, mh *mock_service.MockVerifyHandler) {
				mr.EXPECT().FindAuthInfoByStudentID("student001").Return(&entity.User{
					ID:         1,
					StudentID:  "student001",
					Name:       "name001",
					Role:       "member",
					Password:   "hash",
					Department: "dep001",
					Grade:      2,
					Comment:    "comment001",
					CreatedAt:  mockTime,
					UpdatedAt:  mockTime,
				}, nil)
				mh.EXPECT().PasswordVerify("hash", "password").Return(nil)
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
			mr := mock_repository.NewMockUserRepository(ctrl)
			mh := mock_service.NewMockVerifyHandler(ctrl)
			tt.injector(mr, mh)
			srv := NewUser(mr, mh)
			e := srv.AuthenticationByStudentID(tt.in.studentID, tt.in.password)
			tt.checkErr(t, e)
		})
	}
}

func TestUser_AdminAuthenticationByStudentID(t *testing.T) {
	// t.Parallel()
	mockTime := "2020/02/02 00:00:00"
	type in struct {
		studentID string
		password  string
	}
	tests := []struct {
		name     string
		injector func(mr *mock_repository.MockUserRepository, mh *mock_service.MockVerifyHandler)
		in       in
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to get auth info",
			injector: func(mr *mock_repository.MockUserRepository, mh *mock_service.MockVerifyHandler) {
				mr.EXPECT().FindAuthInfoByStudentID("student001").Return(nil, errors.New("some error"))
			},
			in: in{
				studentID: "student001",
			},
			checkErr: assert.NotNil,
		},
		{
			name: "failed to be not admin",
			injector: func(mr *mock_repository.MockUserRepository, mh *mock_service.MockVerifyHandler) {
				mr.EXPECT().FindAuthInfoByStudentID("student001").Return(&entity.User{
					ID:         1,
					StudentID:  "student001",
					Name:       "name001",
					Role:       "member",
					Password:   "hash",
					Department: "dep001",
					Grade:      2,
					Comment:    "comment001",
					CreatedAt:  mockTime,
					UpdatedAt:  mockTime,
				}, nil)
			},
			in: in{
				studentID: "student001",
			},
			checkErr: assert.NotNil,
		},
		{
			name: "failed to password verify",
			injector: func(mr *mock_repository.MockUserRepository, mh *mock_service.MockVerifyHandler) {
				mr.EXPECT().FindAuthInfoByStudentID("student001").Return(&entity.User{
					ID:         1,
					StudentID:  "student001",
					Name:       "name001",
					Role:       "admin",
					Password:   "hash",
					Department: "dep001",
					Grade:      2,
					Comment:    "comment001",
					CreatedAt:  mockTime,
					UpdatedAt:  mockTime,
				}, nil)
				mh.EXPECT().PasswordVerify("hash", "password").Return(errors.New("some error"))
			},
			in: in{
				studentID: "student001",
				password:  "password",
			},
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(mr *mock_repository.MockUserRepository, mh *mock_service.MockVerifyHandler) {
				mr.EXPECT().FindAuthInfoByStudentID("student001").Return(&entity.User{
					ID:         1,
					StudentID:  "student001",
					Name:       "name001",
					Role:       "admin",
					Password:   "hash",
					Department: "dep001",
					Grade:      2,
					Comment:    "comment001",
					CreatedAt:  mockTime,
					UpdatedAt:  mockTime,
				}, nil)
				mh.EXPECT().PasswordVerify("hash", "password").Return(nil)
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
			mr := mock_repository.NewMockUserRepository(ctrl)
			mh := mock_service.NewMockVerifyHandler(ctrl)
			tt.injector(mr, mh)
			srv := NewUser(mr, mh)
			e := srv.AdminAuthorizationByStudentID(tt.in.studentID, tt.in.password)
			tt.checkErr(t, e)
		})
	}
}

func TestUser_AdminCreate(t *testing.T) {
	// t.Parallel()
	mockTime := "2020/02/02 00:00:00"
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
		injector func(mr *mock_repository.MockUserRepository, mh *mock_service.MockVerifyHandler)
		in       in
		out      int
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to password hash",
			injector: func(mr *mock_repository.MockUserRepository, mh *mock_service.MockVerifyHandler) {
				mh.EXPECT().PasswordHash("password").Return("", errors.New("some error"))
			},
			in: in{
				password: "password",
			},
			out:      0,
			checkErr: assert.NotNil,
		},
		{
			name: "failed to create",
			injector: func(mr *mock_repository.MockUserRepository, mh *mock_service.MockVerifyHandler) {
				mh.EXPECT().PasswordHash("password").Return("hash", nil)
				mr.EXPECT().AdminCreate(&entity.User{
					StudentID:  "student001",
					Name:       "name001",
					Role:       "member",
					Password:   "hash",
					Department: "dep001",
					Grade:      2,
					Comment:    "comment001",
					CreatedAt:  mockTime,
					UpdatedAt:  mockTime,
				}).Return(0, errors.New("some error"))
			},
			in: in{
				name:       "name001",
				studentID:  "student001",
				password:   "password",
				role:       "member",
				department: "dep001",
				grade:      2,
				comment:    "comment001",
			},
			out:      0,
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(mr *mock_repository.MockUserRepository, mh *mock_service.MockVerifyHandler) {
				mh.EXPECT().PasswordHash("password").Return("hash", nil)
				mr.EXPECT().AdminCreate(&entity.User{
					StudentID:  "student001",
					Name:       "name001",
					Role:       "member",
					Password:   "hash",
					Department: "dep001",
					Grade:      2,
					Comment:    "comment001",
					CreatedAt:  mockTime,
					UpdatedAt:  mockTime,
				}).Return(1, nil)
			},
			in: in{
				name:       "name001",
				studentID:  "student001",
				password:   "password",
				role:       "member",
				department: "dep001",
				grade:      2,
				comment:    "comment001",
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
			mr := mock_repository.NewMockUserRepository(ctrl)
			mh := mock_service.NewMockVerifyHandler(ctrl)
			tt.injector(mr, mh)
			srv := NewUser(mr, mh)
			out, e := srv.AdminCreate(tt.in.name, tt.in.studentID, tt.in.password, tt.in.role, tt.in.department, tt.in.comment, tt.in.grade)
			assert.Equal(t, tt.out, out)
			tt.checkErr(t, e)
		})
	}
}

func TestUser_AdminUpdate(t *testing.T) {
	// t.Parallel()
	mockTime := "2020/02/02 00:00:00"
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
		injector func(mr *mock_repository.MockUserRepository, mh *mock_service.MockVerifyHandler)
		in       in
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to get",
			injector: func(mr *mock_repository.MockUserRepository, mh *mock_service.MockVerifyHandler) {
				mr.EXPECT().FindByID(1).Return(nil, errors.New("some error"))
			},
			in: in{
				userID: 1,
			},
			checkErr: assert.NotNil,
		},
		{
			name: "failed to update",
			injector: func(mr *mock_repository.MockUserRepository, mh *mock_service.MockVerifyHandler) {
				mr.EXPECT().FindByID(1).Return(&entity.User{
					ID:         1,
					StudentID:  "student001",
					Name:       "name001",
					Password:   "password",
					Role:       "member",
					Department: "dep001",
					Grade:      2,
					Comment:    "comment001",
					CreatedAt:  mockTime,
					UpdatedAt:  mockTime,
				}, nil)
				mr.EXPECT().AdminUpdateByID(&entity.User{
					ID:         1,
					StudentID:  "student001",
					Name:       "name001",
					Password:   "password",
					Role:       "member",
					Department: "dep001",
					Grade:      2,
					Comment:    "comment001",
					CreatedAt:  mockTime,
					UpdatedAt:  mockTime,
				}).Return(errors.New("some error"))
			},
			in: in{
				userID:     1,
				name:       "name001",
				studentID:  "student001",
				department: "dep001",
				role:       "member",
				comment:    "comment001",
				grade:      2,
			},
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(mr *mock_repository.MockUserRepository, mh *mock_service.MockVerifyHandler) {
				mr.EXPECT().FindByID(1).Return(&entity.User{
					ID:         1,
					StudentID:  "student001",
					Name:       "name001",
					Password:   "password",
					Role:       "member",
					Department: "dep001",
					Grade:      2,
					Comment:    "comment001",
					CreatedAt:  mockTime,
					UpdatedAt:  mockTime,
				}, nil)
				mr.EXPECT().AdminUpdateByID(&entity.User{
					ID:         1,
					StudentID:  "student001",
					Name:       "name001",
					Password:   "password",
					Role:       "member",
					Department: "dep001",
					Grade:      2,
					Comment:    "comment001",
					CreatedAt:  mockTime,
					UpdatedAt:  mockTime,
				}).Return(nil)
			},
			in: in{
				userID:     1,
				name:       "name001",
				studentID:  "student001",
				department: "dep001",
				role:       "member",
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
			patch := monkey.Patch(helper.FormattedDateTimeNow, func() string { return mockTime })
			defer patch.Restore()
			mr := mock_repository.NewMockUserRepository(ctrl)
			mh := mock_service.NewMockVerifyHandler(ctrl)
			tt.injector(mr, mh)
			srv := NewUser(mr, mh)
			e := srv.AdminUpdate(tt.in.userID, tt.in.name, tt.in.studentID, tt.in.role, tt.in.department, tt.in.comment, tt.in.grade)
			tt.checkErr(t, e)
		})
	}
}
