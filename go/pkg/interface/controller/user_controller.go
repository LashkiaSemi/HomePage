package controller

import (
	"errors"
	"homepage/pkg/domain"
	"homepage/pkg/domain/logger"
	"homepage/pkg/usecase/interactor"
)

type UserController interface {
	ShowUsers() (GetUsersResponse, error)
	ShowUserByUserID(userID int) (GetUserResponse, error)
	CreateUser(req *UpdateUserRequest) (GetUserResponse, error)
	UpdateUser(userID int, req *UpdateUserRequest) (GetUserResponse, error)
	DeleteUser(userID int) error
}

type userController struct {
	UserInteractor interactor.UserInteractor
}

func NewUserController(ui interactor.UserInteractor) UserController {
	return &userController{
		UserInteractor: ui,
	}
}

func (uc *userController) ShowUsers() (res GetUsersResponse, err error) {
	users, err := uc.UserInteractor.FetchUsers()
	if err != nil {
		return
	}
	for _, user := range users {
		res.Users = append(res.Users, GetUserResponse{
			ID:         user.ID,
			Name:       user.Name,
			StudentID:  user.StudentID,
			Role:       user.Role,
			Department: user.Department,
			Grade:      user.Grade,
			Comment:    user.Comment,
		})
	}
	return
}

type GetUsersResponse struct {
	Users []GetUserResponse `json:"users"`
}

func (uc *userController) ShowUserByUserID(userID int) (res GetUserResponse, err error) {
	if userID == 0 {
		logger.Warn("ShowUserByUserID: userID is empty")
		return res, domain.BadRequest(errors.New("userID is empty"))
	}

	user, err := uc.UserInteractor.FetchUserByUserID(userID)
	if err != nil {
		return
	}
	res.ID = user.ID
	res.Name = user.Name
	res.StudentID = user.StudentID
	res.Role = user.Role
	res.Department = user.Department
	res.Grade = user.Grade
	res.Comment = user.Comment
	return
}

type GetUserResponse struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	StudentID  string `json:"student_id"`
	Role       string `json:"role"`
	Department string `json:"department"`
	Grade      int    `json:"grade"`
	Comment    string `json:"comments"`
}

func (uc *userController) CreateUser(req *UpdateUserRequest) (res GetUserResponse, err error) {
	// 入力に対してのバリデーション
	if req.Password == "" {
		logger.Warn("CreateAccount: password is empty")
		return res, domain.BadRequest(errors.New("password is empty"))
	}
	if req.StudentID == "" {
		logger.Warn("CreateAccount: studentID is empty")
		return res, domain.BadRequest(errors.New("studentID is empty"))
	}

	// interactor
	user, err := uc.UserInteractor.AddUser(req.Name, req.Password, req.Role, req.StudentID, req.Department, req.Comment, req.Grade)
	if err != nil {
		return
	}

	// resをつくる
	res.Name = user.Name
	res.StudentID = user.StudentID
	res.Role = user.Role
	res.Department = user.Department
	res.Grade = user.Grade
	res.Comment = user.Comment
	return
}

type UpdateUserRequest struct {
	Name       string `json:"name"`
	StudentID  string `json:"student_id"`
	Password   string `json:"password"`
	Role       string `json:"role"`
	Department string `json:"department"`
	Grade      int    `json:"grade"`
	Comment    string `json:"comment"`
}

func (uc *userController) UpdateUser(userID int, req *UpdateUserRequest) (res GetUserResponse, err error) {
	user, err := uc.UserInteractor.UpdateUser(userID, req.Name, req.Password, req.Role, req.StudentID, req.Department, req.Comment, req.Grade)
	if err != nil {
		return
	}
	res.ID = userID
	res.Name = user.Name
	res.StudentID = user.StudentID
	res.Role = user.Role
	res.Department = user.Department
	res.Grade = user.Grade
	res.Comment = user.Comment
	return

}

func (uc *userController) DeleteUser(userID int) error {
	return uc.UserInteractor.DeleteUser(userID)
}
