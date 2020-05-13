package handler

import (
	"homepage/pkg/infrastructure/server/response"
	"net/http"
)

// createInfo 描画時の必須データを作成
func createInfo(r *http.Request, pageType, studentID string) *response.Info {
	return &response.Info{
		PageType:  pageType,
		StudentID: studentID,
	}
}

// createFormField フォームのフィールドを一個作る
func createFormField(name, value, label, formType string, props map[string]string) *FormField {
	return &FormField{
		Name:  name,
		Value: value,
		Label: label,
		Type:  formType,
		Props: props,
	}
}

// FormField adminサイトのフォーム
type FormField struct {
	Name  string            // htmlのname
	Value string            // 初期値
	Label string            // フォームの表示名
	Type  string            // htmlのtype
	Props map[string]string // select用
}
