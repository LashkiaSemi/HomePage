package server

import (
	"homepage/pkg/infrastructure/handler"
	"log"
	"net/http"
	"text/template"
)

type server struct {
	Port    string
	Handler *handler.AppHandler
}

// Server ルーティングとか全部やってくれる子
type Server interface {
	Serve()
}

// NewServer サーバを作るぞ！
func NewServer(port string, ah *handler.AppHandler) Server {
	return &server{
		Port:    port,
		Handler: ah,
	}
}

func (s *server) Serve() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))
	http.HandleFunc("/health", healthHandler)

	http.HandleFunc("/", handler.IndexHandler)
	http.HandleFunc("/login", s.Handler.AuthHandler.Login)
	http.HandleFunc("/activities", s.Handler.ActivityHandler.GetActivities)
	http.HandleFunc("/societies", s.Handler.SocietyHandler.GetAll)
	http.HandleFunc("/researches", dummyHandler("research/index.html"))
	http.HandleFunc("/jobs", s.Handler.JobHandler.GetAll)
	http.HandleFunc("/members", dummyHandler("member/index.html"))
	http.HandleFunc("/links", handler.LinkHandler)
	http.HandleFunc("/equipments", dummyHandler("equipment/index.html"))
	http.HandleFunc("/lectures", dummyHandler("lecture/index.html"))

	log.Println("server running http://localhost:8080")
	http.ListenAndServe(":"+s.Port, nil)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	// data := struct {
	// 	Message string `json:"message"`
	// }{
	// 	Message: "health",
	// }
	// response.Success(w, &data)
	println("health")
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
