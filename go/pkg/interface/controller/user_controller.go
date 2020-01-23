package controller

import (
	"errors"
	"homepage/pkg/domain"
	"homepage/pkg/domain/logger"
	"homepage/pkg/usecase/interactor"
)

// UserController コントローラ
type UserController interface {
	ShowAll() (GetUsersResponse, error)
	ShowByID(userID int) (GetUserResponse, error)
	Create(req *UpdateUserRequest) (GetUserResponse, error)
	Update(userID int, req *UpdateUserRequest) (GetUserResponse, error)
	Delete(userID int) error
}

type userController struct {
	UserInteractor interactor.UserInteractor
}

// NewUserController コントローラの作成
func NewUserController(ui interactor.UserInteractor) UserController {
	return &userController{
		UserInteractor: ui,
	}
}

func (uc *userController) ShowAll() (res GetUsersResponse, err error) {
	users, err := uc.UserInteractor.FetchAll()
	if err != nil {
		return
	}
	for _, user := range users {
		res.Users = append(res.Users, convertUserToResponse(&user))
	}
	return
}

// GetUsersResponse 複数
type GetUsersResponse struct {
	Users []GetUserResponse `json:"users"`
}

func (uc *userController) ShowByID(userID int) (res GetUserResponse, err error) {
	if userID == 0 {
		logger.Warn("ShowUserByUserID: userID is empty")
		return res, domain.BadRequest(errors.New("userID is empty"))
	}

	user, err := uc.UserInteractor.FetchByID(userID)
	if err != nil {
		return
	}
	return convertUserToResponse(&user), nil
}

// GetUserResponse 一件データ
type GetUserResponse struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	StudentID  string `json:"student_id"`
	Role       string `json:"role"`
	Department string `json:"department"`
	Grade      int    `json:"grade"`
	Comment    string `json:"comments"`
}

func (uc *userController) Create(req *UpdateUserRequest) (res GetUserResponse, err error) {
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
	user, err := uc.UserInteractor.Add(req.Name, req.Password, req.Role, req.StudentID, req.Department, req.Comment, req.Grade)
	if err != nil {
		return
	}

	// resをつくる
	return convertUserToResponse(&user), nil
}

// UpdateUserRequest 新規、更新時リクエスト
type UpdateUserRequest struct {
	Name       string `json:"name"`
	StudentID  string `json:"student_id"`
	Password   string `json:"password"`
	Role       string `json:"role"`
	Department string `json:"department"`
	Grade      int    `json:"grade"`
	Comment    string `json:"comment"`
}

func (uc *userController) Update(userID int, req *UpdateUserRequest) (res GetUserResponse, err error) {
	user, err := uc.UserInteractor.Update(userID, req.Name, req.Password, req.Role, req.StudentID, req.Department, req.Comment, req.Grade)
	if err != nil {
		return
	}
	return convertUserToResponse(&user), nil

}

func (uc *userController) Delete(userID int) error {
	return uc.UserInteractor.Delete(userID)
}

func convertUserToResponse(user *domain.User) GetUserResponse {
	return GetUserResponse{
		ID:         user.ID,
		Name:       user.Name,
		StudentID:  user.StudentID,
		Role:       user.Role,
		Department: user.Department,
		Grade:      user.Grade,
		Comment:    user.Comment,
	}
}
