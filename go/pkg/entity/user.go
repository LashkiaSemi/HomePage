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
