package controller

import (
	"errors"
	"homepage/pkg/domain"
	"homepage/pkg/domain/logger"
	"homepage/pkg/usecase/interactor"
)

// AccountController リクエストを受け取ってレスポンスを返却
type AccountController interface {
	ShowByID(userID int) (GetAccountResponse, error)
	ShowByStudentID(studentID string) (GetAccountResponse, error)
	Create(req *UpdateAccountRequest) (GetAccountResponse, error)
	Update(userID int, req *UpdateAccountRequest) (GetAccountResponse, error)
	UpdatePassword(userID int, req *UpdatePasswordRequest) (GetAccountResponse, error)
	Delete(userID int) error

	Login(req *LoginRequest) (LoginResponse, domain.Session, error)
}

type accountController struct {
	AccountInteractor interactor.AccountInteractor
}

// NewAccountController アカウントコントローラの作成
func NewAccountController(ai interactor.AccountInteractor) AccountController {
	return &accountController{
		AccountInteractor: ai,
	}
}

func (ac *accountController) ShowByID(userID int) (res GetAccountResponse, err error) {
	user, err := ac.AccountInteractor.FetchByID(userID)
	if err != nil {
		return
	}
	return convertAccountToResponse(&user), nil
}

func (ac *accountController) ShowByStudentID(studentID string) (res GetAccountResponse, err error) {
	user, err := ac.AccountInteractor.FetchByStudentID(studentID)
	if err != nil {
		return
	}
	return convertAccountToResponse(&user), nil
}

// GetAccountResponse アカウント情報の返却
type GetAccountResponse struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	StudentID  string `json:"student_id"`
	Role       string `json:"role"`
	Department string `json:"department"`
	Grade      int    `json:"grade"`
	Comment    string `json:"comment"`
}

func (ac *accountController) Create(req *UpdateAccountRequest) (res GetAccountResponse, err error) {
	// TODO: リクエストのバリデーションチェック
	if req.Password == "" {
		logger.Warn("CreateAccount: password is empty")
		return res, domain.BadRequest(errors.New("password is empty"))
	}
	if req.StudentID == "" {
		logger.Warn("CreateAccount: studentID is empty")
		return res, domain.BadRequest(errors.New("studentID is empty"))
	}

	// interactor
	user, err := ac.AccountInteractor.Add(req.Name, req.Password, req.Role, req.StudentID, req.Department, req.Comment, req.Grade)
	if err != nil {
		return
	}

	// resをつくる
	return convertAccountToResponse(&user), nil
}

// UpdateAccountRequest アカウントの作成、更新のリクエスト
type UpdateAccountRequest struct {
	Name       string `json:"name"`
	StudentID  string `json:"student_id"`
	Password   string `json:"password"`
	Role       string `json:"role"`
	Department string `json:"department"`
	Grade      int    `json:"grade"`
	Comment    string `json:"comment"`
}

func (ac *accountController) Update(userID int, req *UpdateAccountRequest) (res GetAccountResponse, err error) {
	user, err := ac.AccountInteractor.Update(userID, req.Name, req.Password, req.Role, req.StudentID, req.Department, req.Comment, req.Grade)
	if err != nil {
		return res, err
	}

	return convertAccountToResponse(&user), nil
}

func (ac *accountController) UpdatePassword(userID int, req *UpdatePasswordRequest) (res GetAccountResponse, err error) {
	if req.NewPassword == "" {
		logger.Warn("update password: NewPassword is empty")
		return res, domain.BadRequest(errors.New("password is empty"))
	}

	user, err := ac.AccountInteractor.UpdatePassword(userID, req.OldPassword, req.NewPassword)
	if err != nil {
		return res, err
	}
	return convertAccountToResponse(&user), nil
}

// UpdatePasswordRequest パスワード更新のリクエスト
type UpdatePasswordRequest struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

func (ac *accountController) Delete(userID int) error {
	err := ac.AccountInteractor.Delete(userID)
	return err
}

func (ac *accountController) Login(req *LoginRequest) (res LoginResponse, sess domain.Session, err error) {
	// バリデーションチェック
	if req.StudentID == "" {
		logger.Warn("login controller: studentID is empty")
		return res, sess, domain.BadRequest(errors.New("学籍番号は必須です"))
	}
	if req.Password == "" {
		logger.Warn("login controller: password is empty")
		return res, sess, domain.BadRequest(errors.New("パスワードは必須です"))
	}

	// いんたらくた
	sess, err = ac.AccountInteractor.Login(req.StudentID, req.Password)
	if err != nil {
		return res, sess, err
	}

	// れすぽんす
	res.UserID = sess.UserID
	res.SessionID = sess.SessionID
	res.StudentID = req.StudentID
	res.Role = sess.Role
	return
}

// LoginRequest ログインリクエスト
type LoginRequest struct {
	StudentID string `json:"student_id"`
	Password  string `json:"password"`
}

// LoginResponse ログイン時のレスポンス
type LoginResponse struct {
	UserID    int    `json:"user_id"`
	StudentID string `json:"student_id"`
	SessionID string `json:"session_id"`
	Role      string `json:"role"`
}

func convertAccountToResponse(user *domain.User) GetAccountResponse {
	return GetAccountResponse{
		ID:         user.ID,
		Name:       user.Name,
		StudentID:  user.StudentID,
		Role:       user.Role,
		Department: user.Department,
		Grade:      user.Grade,
		Comment:    user.Comment,
	}
}
