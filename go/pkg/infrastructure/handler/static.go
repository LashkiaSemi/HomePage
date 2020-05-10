package handler

import (
	"homepage/pkg/infrastructure/server/response"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	header := &response.HeaderData{
		IsLogin:  false,
		PageType: "",
	}
	var body interface{}
	response.Success(w, "index.html", header, &body)
}
