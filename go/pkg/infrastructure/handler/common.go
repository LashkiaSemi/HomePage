package handler

import (
	"homepage/pkg/infrastructure/auth"
	"homepage/pkg/infrastructure/server/response"
	"net/http"
)

// createInfo 描画時の必須データを作成
func createInfo(r *http.Request, pageType string) *response.Info {
	return &response.Info{
		IsLogin:  auth.CheckIsLogin(r),
		PageType: pageType,
	}
}
