package server

import (
	"homepage/pkg/infrastructure/handler"
	"homepage/pkg/infrastructure/server/middleware"
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
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
	r := mux.NewRouter()
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))

	//TODO: 環境変数とかのがいいかも。レクチャーの資料とかしまってある場所
	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("public/"))))

	// r.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))
	r.HandleFunc("/health", healthHandler)

	r.HandleFunc("/", handler.IndexHandler)
	r.HandleFunc("/login", s.Handler.UserHandler.Login)
	r.HandleFunc("/logout", middleware.Authorized(s.Handler.UserHandler.Logout))
	r.HandleFunc("/activities", s.Handler.ActivityHandler.GetActivities)
	r.HandleFunc("/societies", s.Handler.SocietyHandler.GetAll)
	r.HandleFunc("/researches", s.Handler.ResearchHandler.GetAll)
	r.HandleFunc("/jobs", s.Handler.JobHandler.GetAll)
	r.HandleFunc("/members", s.Handler.UserHandler.GetAllGroupByGrade)
	r.HandleFunc("/members/{id}", s.Handler.UserHandler.GetByID)
	r.HandleFunc("/members/edit/profile", middleware.Authorized(s.Handler.UserHandler.UpdateByID))
	r.HandleFunc("/members/edit/password", middleware.Authorized(s.Handler.UserHandler.UpdatePasswordByStudentID))
	r.HandleFunc("/links", handler.LinkHandler)
	r.HandleFunc("/equipments", middleware.Authorized(s.Handler.EquipmentHandler.GetAll))
	r.HandleFunc("/lectures", middleware.Authorized(s.Handler.LectureHandler.GetAll))
	r.HandleFunc("/lectures/new", middleware.Authorized(s.Handler.LectureHandler.Create))
	r.HandleFunc("/lectures/{id}/edit", middleware.Authorized(s.Handler.LectureHandler.UpdateByID))

	// TODO: フォーム形式にして！
	r.HandleFunc("/lectures/{id}/delete", s.Handler.LectureHandler.DeleteByID)

	log.Println("server running http://localhost:8080")
	http.ListenAndServe(":"+s.Port, r)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("health")
}

func dummyHandler(templateFile string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles(
			"template/"+templateFile,
			"template/_footer.html",
			"template/_header.html",
		)
		if err != nil {
			log.Printf("failed to parse template: %v", err)
		}
		if err = t.Execute(w, struct {
			Info *dummyInfo
		}{
			Info: &dummyInfo{
				StudentID: "dummy",
				PageType:  "",
			},
		}); err != nil {
			log.Printf("failed to execute template: %v", err)
		}
	}
}

type dummyInfo struct {
	StudentID string
	PageType  string
}
