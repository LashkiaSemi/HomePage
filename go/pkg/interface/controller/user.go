package controller

import (
	"homepage/pkg/domain/model"
	"homepage/pkg/usecase/interactor"
)

type userController struct {
	interactor.UserInteractor
}

// UserController ユーザの入出力を変換
type UserController interface {
	GetAllGroupByGrade() (*UsersGroupByGradeResponse, error)
	GetByID(userID string) (*UserResponse, error)
	Login(studentID, password string) error
}

// NewUserController コントローラの作成
func NewUserController(ui interactor.UserInteractor) UserController {
	return &userController{
		UserInteractor: ui,
	}
}

func (uc *userController) GetAllGroupByGrade() (*UsersGroupByGradeResponse, error) {
	users, err := uc.UserInteractor.GetAll()
	if err != nil {
		return &UsersGroupByGradeResponse{}, err
	}
	// var res UsersGroupByGradeResponse
	var res = make(map[int][]*UserResponse)
	for _, user := range users {
		res[user.Grade] = append(res[user.Grade], convertToUserResponse(user))
	}

	return &UsersGroupByGradeResponse{GradeUsers: res}, nil
}

func (uc *userController) GetByID(userID string) (*UserResponse, error) {
	user, err := uc.UserInteractor.GetByID(userID)
	if err != nil {
		return &UserResponse{}, err
	}
	return convertToUserResponse(user), err
}

func (uc *userController) Login(studentID, password string) error {
	return uc.AuthenticationByStudentID(studentID, password)
}

// UsersResponse 複数ユーザのレスポンス
type UsersResponse struct {
	Users []*UserResponse
}

// UsersGroupByGradeResponse 学年別mapを作成する
type UsersGroupByGradeResponse struct {
	GradeUsers map[int][]*UserResponse
}

// UserResponse ユーザ一件分
type UserResponse struct {
	ID         int
	StudentID  string
	Name       string
	Department string
	Grade      int
	Comment    string
}

func convertToUserResponse(user *model.User) *UserResponse {
	return &UserResponse{
		ID:         user.ID,
		StudentID:  user.StudentID,
		Name:       user.Name,
		Department: user.Department,
		Grade:      user.Grade,
		Comment:    user.Comment,
	}
}
