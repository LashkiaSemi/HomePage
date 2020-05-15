package controller

import (
	"homepage/pkg/entity"
	"homepage/pkg/usecase/interactor"
	"strconv"

	"github.com/pkg/errors"
)

type userController struct {
	interactor.UserInteractor
}

// UserController ユーザの入出力を変換
type UserController interface {
	GetAll() (*UsersResponse, error)
	GetAllGroupByGrade() (*UsersGroupByGradeResponse, error)
	GetByID(userID int) (*UserResponse, error)
	GetByStudentID(studentID string) (*UserResponse, error)
	UpdateByID(userID int, name, studentID, department, comment string, grade int) (*UserResponse, error)
	UpdatePasswordByStudentID(studentID, oldPassword, newPassword string) error

	DeleteByID(id int) error

	Login(studentID, password string) error

	// admin
	AdminLogin(studentID, password string) error
	AdminGetAll() ([]map[string]string, error)
	AdminGetByID(userID int) (*FieldsResponse, error)
	AdminCreate(name, studentID, department, comment, password, role string, grade int) (int, error)
	AdminUpdateByID(userID int, name, studentID, department, comment, role string, grade int) error
}

// NewUserController コントローラの作成
func NewUserController(ui interactor.UserInteractor) UserController {
	return &userController{
		UserInteractor: ui,
	}
}

func (uc *userController) GetAll() (*UsersResponse, error) {
	users, err := uc.UserInteractor.GetAll()
	if err != nil {
		err = errors.Wrap(err, "failed to original data for response")
		return &UsersResponse{}, err
	}
	var res UsersResponse
	for _, user := range users {
		res.Users = append(res.Users, convertToUserResponse(user))
	}
	return &res, err
}

func (uc *userController) GetAllGroupByGrade() (*UsersGroupByGradeResponse, error) {
	users, err := uc.UserInteractor.GetAll()
	if err != nil {
		err = errors.Wrap(err, "failed to original data for response")
		return &UsersGroupByGradeResponse{}, err
	}
	var res = make(map[int][]*UserResponse)
	for _, user := range users {
		res[user.Grade] = append(res[user.Grade], convertToUserResponse(user))
	}

	return &UsersGroupByGradeResponse{GradeUsers: res}, nil
}

func (uc *userController) GetByID(userID int) (*UserResponse, error) {
	user, err := uc.UserInteractor.GetByID(userID)
	if err != nil {
		err = errors.Wrap(err, "failed to original data for response")
		return &UserResponse{}, err
	}
	return convertToUserResponse(user), err
}

func (uc *userController) GetByStudentID(studentID string) (*UserResponse, error) {
	user, err := uc.UserInteractor.GetByStudentID(studentID)
	if err != nil {
		err = errors.Wrap(err, "failed to original data for response")
		return &UserResponse{}, err
	}
	return convertToUserResponse(user), err
}

func (uc *userController) UpdateByID(userID int, name, studentID, department, comment string, grade int) (*UserResponse, error) {
	user, err := uc.UserInteractor.UpdateByID(userID, name, studentID, department, comment, grade)
	if err != nil {
		err = errors.Wrap(err, "failed to original data for response")
		return &UserResponse{}, err
	}
	return convertToUserResponse(user), err
}

func (uc *userController) UpdatePasswordByStudentID(studentID, oldPassword, newPassword string) error {
	return uc.UserInteractor.UpdatePasswordByStudentID(studentID, oldPassword, newPassword)
}

func (uc *userController) DeleteByID(id int) error {
	return uc.UserInteractor.DeleteByID(id)
}

// session
func (uc *userController) Login(studentID, password string) error {
	return uc.UserInteractor.AuthenticationByStudentID(studentID, password)
}

func (uc *userController) AdminLogin(studentID, password string) error {
	return uc.UserInteractor.AdminAuthorizationByStudentID(studentID, password)
}

// admin
func (uc *userController) AdminGetAll() ([]map[string]string, error) {
	var res []map[string]string
	users, err := uc.UserInteractor.GetAll()
	if err != nil {
		err = errors.Wrap(err, "failed to original data for response")
		return res, err
	}
	for _, user := range users {
		res = append(res, map[string]string{
			"id":    strconv.Itoa(user.ID),
			"title": user.Name,
		})
	}

	return res, err
}

func (uc *userController) AdminGetByID(userID int) (*FieldsResponse, error) {
	var res FieldsResponse

	user, err := uc.UserInteractor.GetByID(userID)
	if err != nil {
		err = errors.Wrap(err, "failed to original data for response")
		return &res, err
	}

	res.Fields = append(res.Fields,
		&Field{Key: "ID", Value: user.ID},
		&Field{Key: "名前", Value: user.Name},
		&Field{Key: "学籍番号", Value: user.StudentID},
		&Field{Key: "学部", Value: user.Department},
		&Field{Key: "学年", Value: convertGradeFromIntToString(user.Grade)},
		&Field{Key: "コメント", Value: user.Comment},
		&Field{Key: "権限", Value: user.Role},
	)
	res.ID = user.ID
	return &res, err
}

func (uc *userController) AdminCreate(name, studentID, department, comment, password, role string, grade int) (int, error) {
	return uc.UserInteractor.AdminCreate(name, studentID, password, role, department, comment, grade)
}

func (uc *userController) AdminUpdateByID(userID int, name, studentID, department, comment, role string, grade int) error {
	return uc.UserInteractor.AdminUpdate(userID, name, studentID, role, department, comment, grade)
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
	Grade      string
	Comment    string
}

func convertToUserResponse(user *entity.User) *UserResponse {
	return &UserResponse{
		ID:         user.ID,
		StudentID:  user.StudentID,
		Name:       user.Name,
		Department: user.Department,
		Grade:      convertGradeFromIntToString(user.Grade),
		Comment:    user.Comment,
	}
}

// convertGradeFromIntToString 学年を文字列に変換してくれる
func convertGradeFromIntToString(grade int) string {
	switch grade {
	case 2:
		return "学部2年"
	case 3:
		return "学部3年"
	case 4:
		return "学部4年"
	case 5:
		return "大学院1年"
	case 6:
		return "大学院2年"
	default:
		return "卒業生"
	}
}
