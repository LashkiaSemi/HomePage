package entity

import (
	"homepage/pkg/helper"
)

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

func NewUser(name, studentID, password, role, department, comment string, grade int) *User {
	now := helper.FormattedDateTimeNow()
	return &User{
		StudentID:  studentID,
		Name:       name,
		Password:   password,
		Role:       role,
		Department: department,
		Grade:      grade,
		Comment:    comment,
		CreatedAt:  now,
		UpdatedAt:  now,
	}
}

func (u User) AdminUpdate(name, studentID, department, comment, role string, grade int) *User {
	u.Name = name
	u.StudentID = studentID
	u.Role = role
	u.Department = department
	u.Comment = comment
	u.Grade = grade
	u.UpdatedAt = helper.FormattedDateTimeNow()
	return &u
}

func (u User) Update(name, studentID, department, comment string, grade int) *User {
	res := u
	res.Name = name
	res.StudentID = studentID
	res.Department = department
	res.Comment = comment
	res.Grade = grade
	res.UpdatedAt = helper.FormattedDateTimeNow()
	return &res
}

// IsAdmin check user role is admin or owner
func (u *User) IsAdmin() bool {
	return u.Role == "admin" || u.Role == "owner"
}
