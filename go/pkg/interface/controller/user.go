package controller

import (
	"homepage/pkg/domain/model"
	"homepage/pkg/usecase/interactor"
)

type userController struct {
	interactor.UserInteractor
}

type UserController interface {
	GetAllGroupByGrade() (*UsersGroupByGradeResponse, error)
	Login(studentID, password string) error
}

func NewUserContoroller(ui interactor.UserInteractor) UserController {
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

func (uc *userController) Login(studentID, password string) error {
	return uc.AuthenticationByStudentID(studentID, password)
}

type UsersResponse struct {
	Users []*UserResponse
}

type UsersGroupByGradeResponse struct {
	GradeUsers map[int][]*UserResponse
}

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
