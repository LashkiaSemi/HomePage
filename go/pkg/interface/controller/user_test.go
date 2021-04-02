package controller

import (
	mock_interactor "homepage/mock/interactor"
	"homepage/pkg/domain/entity"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestUser_GetAll(t *testing.T) {
	tests := []struct {
		name     string
		injector func(m *mock_interactor.MockUserInteractor)
		out      *UsersResponse
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to get",
			injector: func(m *mock_interactor.MockUserInteractor) {
				m.EXPECT().GetAll().Return(nil, errors.New("some error"))
			},
			out:      &UsersResponse{},
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_interactor.MockUserInteractor) {
				m.EXPECT().GetAll().Return([]*entity.User{
					{
						ID: 1,
					},
				}, nil)
			},
			out: &UsersResponse{
				Users: []*UserResponse{
					{
						ID: 1,
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
			m := mock_interactor.NewMockUserInteractor(ctrl)
			tt.injector(m)
			c := NewUserController(m)
			out, err := c.GetAll()
			assert.Equal(t, tt.out, out)
			tt.checkErr(t, err)
		})
	}
}

func TestUser_GetAllGroupByGrade(t *testing.T) {
	tests := []struct {
		name     string
		injector func(m *mock_interactor.MockUserInteractor)
		out      *UsersGroupByGradeResponse
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to get",
			injector: func(m *mock_interactor.MockUserInteractor) {
				m.EXPECT().GetAll().Return(nil, errors.New("some error"))
			},
			out:      &UsersGroupByGradeResponse{},
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_interactor.MockUserInteractor) {
				m.EXPECT().GetAll().Return([]*entity.User{
					{
						ID:    1,
						Grade: 1,
					},
					{
						ID:    2,
						Grade: 2,
					},
					{
						ID:    3,
						Grade: 2,
					},
				}, nil)
			},
			out: &UsersGroupByGradeResponse{
				GradeUsers: map[int][]*UserResponse{
					1: {
						{
							ID:    1,
							Grade: 1,
						},
					},
					2: {
						{
							ID:    2,
							Grade: 2,
						},
						{
							ID:    3,
							Grade: 2,
						},
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
			m := mock_interactor.NewMockUserInteractor(ctrl)
			tt.injector(m)
			c := NewUserController(m)
			out, err := c.GetAllGroupByGrade()
			assert.Equal(t, tt.out, out)
			tt.checkErr(t, err)
		})
	}
}

func TestUser_GetByID(t *testing.T) {
	tests := []struct {
		name     string
		injector func(m *mock_interactor.MockUserInteractor)
		in       int
		out      *UserResponse
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to get",
			injector: func(m *mock_interactor.MockUserInteractor) {
				m.EXPECT().GetByID(1).Return(nil, errors.New("some error"))
			},
			in:       1,
			out:      &UserResponse{},
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_interactor.MockUserInteractor) {
				m.EXPECT().GetByID(1).Return(&entity.User{
					ID: 1,
				}, nil)
			},
			in: 1,
			out: &UserResponse{
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
			m := mock_interactor.NewMockUserInteractor(ctrl)
			tt.injector(m)
			c := NewUserController(m)
			out, err := c.GetByID(tt.in)
			assert.Equal(t, tt.out, out)
			tt.checkErr(t, err)
		})
	}
}

func TestUser_GetByStudentID(t *testing.T) {
	tests := []struct {
		name     string
		injector func(m *mock_interactor.MockUserInteractor)
		in       string
		out      *UserResponse
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to get",
			injector: func(m *mock_interactor.MockUserInteractor) {
				m.EXPECT().GetByStudentID("").Return(nil, errors.New("some error"))
			},
			in:       "",
			out:      &UserResponse{},
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_interactor.MockUserInteractor) {
				m.EXPECT().GetByStudentID("student001").Return(&entity.User{
					ID:        1,
					StudentID: "student001",
				}, nil)
			},
			in: "student001",
			out: &UserResponse{
				ID:        1,
				StudentID: "student001",
			},
			checkErr: assert.Nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			m := mock_interactor.NewMockUserInteractor(ctrl)
			tt.injector(m)
			c := NewUserController(m)
			out, err := c.GetByStudentID(tt.in)
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
		injector func(m *mock_interactor.MockUserInteractor)
		in       in
		out      *UserResponse
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to update",
			injector: func(m *mock_interactor.MockUserInteractor) {
				m.EXPECT().UpdateByID(1, "", "", "", "", 0).Return(nil, errors.New("some error"))
			},
			in: in{
				userID: 1,
			},
			out:      &UserResponse{},
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_interactor.MockUserInteractor) {
				m.EXPECT().UpdateByID(1, "name001", "student001", "dep001", "comment001", 1).Return(&entity.User{
					ID:         1,
					Name:       "name001",
					StudentID:  "student001",
					Department: "dep001",
					Comment:    "comment001",
					Grade:      1,
				}, nil)
			},
			in: in{
				userID:     1,
				name:       "name001",
				studentID:  "student001",
				department: "dep001",
				comment:    "comment001",
				grade:      1,
			},
			out: &UserResponse{
				ID:         1,
				Name:       "name001",
				StudentID:  "student001",
				Department: "dep001",
				Comment:    "comment001",
				Grade:      1,
			},
			checkErr: assert.Nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			m := mock_interactor.NewMockUserInteractor(ctrl)
			tt.injector(m)
			c := NewUserController(m)
			out, err := c.UpdateByID(tt.in.userID, tt.in.name, tt.in.studentID, tt.in.department, tt.in.comment, tt.in.grade)
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
		injector func(m *mock_interactor.MockUserInteractor)
		in       in
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to update",
			injector: func(m *mock_interactor.MockUserInteractor) {
				m.EXPECT().UpdatePasswordByStudentID("", "", "").Return(errors.New("some error"))
			},
			in:       in{},
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_interactor.MockUserInteractor) {
				m.EXPECT().UpdatePasswordByStudentID("student001", "oldPass", "newPass").Return(nil)
			},
			in: in{
				studentID:   "student001",
				oldPassword: "oldPass",
				newPassword: "newPass",
			},
			checkErr: assert.Nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			m := mock_interactor.NewMockUserInteractor(ctrl)
			tt.injector(m)
			c := NewUserController(m)
			err := c.UpdatePasswordByStudentID(tt.in.studentID, tt.in.oldPassword, tt.in.newPassword)
			tt.checkErr(t, err)
		})
	}
}

func TestUser_DeleteByID(t *testing.T) {
	tests := []struct {
		name     string
		injector func(m *mock_interactor.MockUserInteractor)
		in       int
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to delete",
			injector: func(m *mock_interactor.MockUserInteractor) {
				m.EXPECT().DeleteByID(1).Return(errors.New("some error"))
			},
			in:       1,
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_interactor.MockUserInteractor) {
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
			m := mock_interactor.NewMockUserInteractor(ctrl)
			tt.injector(m)
			c := NewUserController(m)
			err := c.DeleteByID(tt.in)
			tt.checkErr(t, err)
		})
	}
}

func TestUser_Login(t *testing.T) {
	type in struct {
		studentID string
		password  string
	}
	tests := []struct {
		name     string
		injector func(m *mock_interactor.MockUserInteractor)
		in       in
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to login",
			injector: func(m *mock_interactor.MockUserInteractor) {
				m.EXPECT().AuthenticationByStudentID("", "").Return(errors.New("some error"))
			},
			in:       in{},
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_interactor.MockUserInteractor) {
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
			m := mock_interactor.NewMockUserInteractor(ctrl)
			tt.injector(m)
			c := NewUserController(m)
			err := c.Login(tt.in.studentID, tt.in.password)
			tt.checkErr(t, err)
		})
	}
}

func TestUser_AdminLogin(t *testing.T) {
	type in struct {
		studentID string
		password  string
	}
	tests := []struct {
		name     string
		injector func(m *mock_interactor.MockUserInteractor)
		in       in
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to login",
			injector: func(m *mock_interactor.MockUserInteractor) {
				m.EXPECT().AdminAuthorizationByStudentID("", "").Return(errors.New("some error"))
			},
			in:       in{},
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_interactor.MockUserInteractor) {
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
			m := mock_interactor.NewMockUserInteractor(ctrl)
			tt.injector(m)
			c := NewUserController(m)
			err := c.AdminLogin(tt.in.studentID, tt.in.password)
			tt.checkErr(t, err)
		})
	}
}

func TestUser_AdminGetAll(t *testing.T) {
	tests := []struct {
		name     string
		injector func(m *mock_interactor.MockUserInteractor)
		out      []map[string]string
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to get",
			injector: func(m *mock_interactor.MockUserInteractor) {
				m.EXPECT().GetAll().Return(nil, errors.New("some error"))
			},
			out:      []map[string]string(nil),
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_interactor.MockUserInteractor) {
				m.EXPECT().GetAll().Return([]*entity.User{
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
			m := mock_interactor.NewMockUserInteractor(ctrl)
			tt.injector(m)
			c := NewUserController(m)
			out, err := c.AdminGetAll()
			assert.Equal(t, tt.out, out)
			tt.checkErr(t, err)
		})
	}
}

func TestUser_AdminGetByID(t *testing.T) {
	tests := []struct {
		name     string
		injector func(m *mock_interactor.MockUserInteractor)
		in       int
		out      *FieldsResponse
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to get",
			injector: func(m *mock_interactor.MockUserInteractor) {
				m.EXPECT().GetByID(1).Return(nil, errors.New("some error"))
			},
			in:       1,
			out:      &FieldsResponse{},
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_interactor.MockUserInteractor) {
				m.EXPECT().GetByID(1).Return(&entity.User{
					ID:         1,
					Name:       "name001",
					StudentID:  "student001",
					Department: "dep001",
					Grade:      2,
					Comment:    "comment001",
					Role:       "admin",
				}, nil)
			},
			in: 1,
			out: &FieldsResponse{
				ID: 1,
				Fields: []*Field{
					{Key: "ID", Value: 1},
					{Key: "名前", Value: "name001"},
					{Key: "学籍番号", Value: "student001"},
					{Key: "学部", Value: "dep001"},
					{Key: "学年", Value: "学部2年"},
					{Key: "コメント", Value: "comment001"},
					{Key: "権限", Value: "admin"},
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
			m := mock_interactor.NewMockUserInteractor(ctrl)
			tt.injector(m)
			c := NewUserController(m)
			out, err := c.AdminGetByID(tt.in)
			assert.Equal(t, tt.out, out)
			tt.checkErr(t, err)
		})
	}
}

func TestUser_AdminCreate(t *testing.T) {
	type in struct {
		name       string
		studentID  string
		department string
		comment    string
		password   string
		role       string
		grade      int
	}
	tests := []struct {
		name     string
		injector func(m *mock_interactor.MockUserInteractor)
		in       in
		out      int
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to create",
			injector: func(m *mock_interactor.MockUserInteractor) {
				m.EXPECT().AdminCreate("", "", "", "", "", "", 0).Return(0, errors.New("some error"))
			},
			in:       in{},
			out:      0,
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_interactor.MockUserInteractor) {
				m.EXPECT().AdminCreate("name001", "student001", "password", "admin", "dep001", "comment001", 2).Return(1, nil)
			},
			in: in{
				name:       "name001",
				studentID:  "student001",
				department: "dep001",
				comment:    "comment001",
				password:   "password",
				role:       "admin",
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
			m := mock_interactor.NewMockUserInteractor(ctrl)
			tt.injector(m)
			c := NewUserController(m)
			out, err := c.AdminCreate(tt.in.name, tt.in.studentID, tt.in.department, tt.in.comment, tt.in.password, tt.in.role, tt.in.grade)
			assert.Equal(t, tt.out, out)
			tt.checkErr(t, err)
		})
	}
}

func TestUser_AdminUpdateByID(t *testing.T) {
	type in struct {
		userID     int
		name       string
		studentID  string
		department string
		comment    string
		role       string
		grade      int
	}
	tests := []struct {
		name     string
		injector func(m *mock_interactor.MockUserInteractor)
		in       in
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to update",
			injector: func(m *mock_interactor.MockUserInteractor) {
				m.EXPECT().AdminUpdate(1, "", "", "", "", "", 0).Return(errors.New("some error"))
			},
			in: in{
				userID: 1,
			},
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_interactor.MockUserInteractor) {
				m.EXPECT().AdminUpdate(1, "name001", "student001", "admin", "dep001", "comment001", 2).Return(nil)
			},
			in: in{
				userID:     1,
				name:       "name001",
				studentID:  "student001",
				department: "dep001",
				comment:    "comment001",
				role:       "admin",
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
			m := mock_interactor.NewMockUserInteractor(ctrl)
			tt.injector(m)
			c := NewUserController(m)
			err := c.AdminUpdateByID(tt.in.userID, tt.in.name, tt.in.studentID, tt.in.department, tt.in.comment, tt.in.role, tt.in.grade)
			tt.checkErr(t, err)
		})
	}
}
