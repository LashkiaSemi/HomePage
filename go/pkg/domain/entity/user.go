package entity

import (
	"homepage/pkg/configs"
	"time"
)

type UserService interface {
	Update(name, studentID, department, comment string, grade int) *User
}

// User ユーザ
type User struct {
	ID         int
	StudentID  string
	Name       string
	Password   string
	Role       string
	Department string
	Grade      int
	Comment    string
	CreatedAt  string
	UpdatedAt  string
}

func (u *User) Create(name, studentID, password, role, department, comment string, grade int) {
	u.Name = name
	u.StudentID = studentID
	u.Password = password
	u.Role = role
	u.Department = department
	u.Comment = comment
	u.Grade = grade
	u.CreatedAt = time.Now().Format(configs.DateTimeFormat)
	u.UpdatedAt = u.CreatedAt
}

func (u User) AdminUpdate(name, studentID, department, comment, role string, grade int) *User {
	u.Name = name
	u.StudentID = studentID
	u.Role = role
	u.Department = department
	u.Comment = comment
	u.Grade = grade
	u.UpdatedAt = time.Now().Format(configs.DateTimeFormat)
	return &u
}

func (u User) Update(name, studentID, department, comment string, grade int) *User {
	res := u
	res.Name = name
	res.StudentID = studentID
	res.Department = department
	res.Comment = comment
	res.Grade = grade
	res.UpdatedAt = time.Now().Format(configs.DateTimeFormat)
	return &res
}
