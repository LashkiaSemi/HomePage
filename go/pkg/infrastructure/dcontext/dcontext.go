package dcontext

import "context"

type key string

const (
	studentIDKey key = "studentID"
)

// SetStudentID contextに学籍番号を突っ込む
func SetStudentID(ctx context.Context, studentID string) context.Context {
	return context.WithValue(ctx, studentIDKey, studentID)
}

// GetStudentIDFromContext contextから学籍番号を取得
func GetStudentIDFromContext(ctx context.Context) string {
	var studentID string
	if ctx.Value(studentIDKey) != nil {
		studentID = ctx.Value(studentIDKey).(string)
	}
	return studentID
}
