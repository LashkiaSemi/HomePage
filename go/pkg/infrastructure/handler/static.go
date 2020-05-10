package handler

import (
	"homepage/pkg/infrastructure/server/response"
	"net/http"
)

// IndexHandler home部分のhandler
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	info := createInfo(r, "")
	var body interface{}
	response.Success(w, "index.html", info, &body)
}
