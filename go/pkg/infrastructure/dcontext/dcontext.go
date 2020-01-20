package dcontext

import (
	"context"
)

type key string

const (
	studentIDKey key = "studentID"
	userIDKey    key = "userID"
)

// SetStudentID studentIDをcontextに保存
func SetStudentID(ctx context.Context, studentID string) context.Context {
	return context.WithValue(ctx, studentIDKey, studentID)
}

// GetStudentIDFromContext contextからstudentIDを取得
func GetStudentIDFromContext(ctx context.Context) string {
	var studentID string
	if ctx.Value(studentIDKey) != nil {
		studentID = ctx.Value(studentIDKey).(string)
	}
	return studentID
}

// SetUserID userIDをcontextに保存
func SetUserID(ctx context.Context, userID int) context.Context {
	return context.WithValue(ctx, userIDKey, userID)
}

// GetUserIDFromContext contextからstudentIDを取得
func GetUserIDFromContext(ctx context.Context) int {
	var userID int
	if ctx.Value(userIDKey) != nil {
		userID = ctx.Value(userIDKey).(int)
	}
	return userID
}
