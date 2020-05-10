package server

import (
	"homepage/pkg/infrastructure/handler"
	"homepage/pkg/infrastructure/server/response"
	"log"
	"net/http"
	"text/template"
)

type server struct {
	Port string
}

// Server ルーティングとか全部やってくれる子
type Server interface {
	Serve()
}

// NewServer サーバを作るぞ！
func NewServer(port string) Server {
	return &server{
		Port: port,
	}
}

func (s *server) Serve() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))
	http.HandleFunc("/health", healthHandler)

	http.HandleFunc("/", handler.IndexHandler)
	http.HandleFunc("/login", dummyHandler("login.html"))
	http.HandleFunc("/activities", dummyHandler("activity/index.html"))
	http.HandleFunc("/societies", dummyHandler("society/index.html"))
	http.HandleFunc("/researches", dummyHandler("research/index.html"))
	http.HandleFunc("/jobs", dummyHandler("job/index.html"))
	http.HandleFunc("/members", dummyHandler("member/index.html"))
	http.HandleFunc("/links", dummyHandler("link/index.html"))
	http.HandleFunc("/equipments", dummyHandler("equipment/index.html"))
	http.HandleFunc("/lectures", dummyHandler("lecture/index.html"))

	http.ListenAndServe(":"+s.Port, nil)
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
