package main

import (
	"homepage/pkg/infrastructure/handler"
	"homepage/pkg/infrastructure/server/response"
	"log"
	"net/http"
	"text/template"
)

func main() {
	port := "8080"

	// cssの反映
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))
	http.HandleFunc("/health", healthHandler)

	http.HandleFunc("/", handler.IndexHandler)
	http.HandleFunc("/login", dummyHandler("login.html"))

	http.ListenAndServe(":"+port, nil)

}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Message string `json:"message"`
	}{
		Message: "health",
	}
	response.Success(w, &data)
}

func dummyHandler(file string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles(
			"template/"+file,
			"template/_footer.html",
			"template/_header.html",
		)
		if err != nil {
			log.Printf("failed to parse template: %v", err)
		}
		if err = t.Execute(w, struct{}{}); err != nil {
			log.Printf("failed to execute template: %v", err)
		}
	}
}
