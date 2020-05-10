package handler

import (
	"homepage/pkg/infrastructure/authentication"
	"homepage/pkg/infrastructure/server/response"
	"net/http"
)

// createInfo 描画時の必須データを作成
func createInfo(r *http.Request, pageType string) *response.Info {
	return &response.Info{
		IsLogin:  authentication.CheckIsLogin(r),
		PageType: pageType,
	}
}
