package handler

import (
	"html/template"
	"log"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(
		"template/index.html",
		"template/_footer.html",
		"template/_header.html",
	)
	if err != nil {
		log.Printf("failed to parse template: %v", err)
	}
	if err = t.Execute(w, struct {
		IsLogin  bool
		PageType string
	}{
		IsLogin:  true,
		PageType: "index",
	}); err != nil {
		log.Printf("failed to execute template: %v", err)
	}
}
