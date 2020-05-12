package handler

import (
	"homepage/pkg/infrastructure/auth"
	"homepage/pkg/infrastructure/server/response"
	"net/http"
)

// IndexHandler home部分のhandler
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	info := createInfo(r, "", auth.GetStudentIDFromCookie(r))
	var body interface{}
	response.Success(w, "index.html", info, body)
}

// AdminIndexHandler adminのトップページ
func AdminIndexHandler(w http.ResponseWriter, r *http.Request) {
	info := createInfo(r, "", auth.GetStudentIDFromCookie(r))
	var body interface{}
	response.AdminRender(w, "index.html", info, body)
}
