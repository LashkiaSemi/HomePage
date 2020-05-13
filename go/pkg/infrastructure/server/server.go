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
	r.HandleFunc("/lectures/{id}/delete", middleware.Authorized(s.Handler.LectureHandler.DeleteByID))

	// admin site
	// TODO: middleware
	r.HandleFunc("/admin", handler.AdminIndexHandler)
	r.HandleFunc("/admin/activities", s.Handler.ActivityHandler.AdminGetAll)
	r.HandleFunc("/admin/societies", s.Handler.SocietyHandler.AdminGetAll)
	r.HandleFunc("/admin/researches", s.Handler.ResearchHandler.AdminGetAll)
	r.HandleFunc("/admin/jobs", s.Handler.JobHandler.AdminGetAll)
	r.HandleFunc("/admin/members", s.Handler.UserHandler.AdminGetAll)
	r.HandleFunc("/admin/lectures", s.Handler.LectureHandler.AdminGetAll)
	r.HandleFunc("/admin/equipments", s.Handler.EquipmentHandler.AdminGetAll)
	r.HandleFunc("/admin/tags", s.Handler.TagHandler.AdminGetAll)

	r.HandleFunc("/admin/members/{id:[0-9]+}", s.Handler.UserHandler.AdminGetByID)
	r.HandleFunc("/admin/activities/{id:[0-9]+}", s.Handler.ActivityHandler.AdminGetByID)
	r.HandleFunc("/admin/societies/{id:[0-9]+}", s.Handler.SocietyHandler.AdminGeByID)
	r.HandleFunc("/admin/jobs/{id:[0-9]+}", s.Handler.JobHandler.AdminGetByID)
	r.HandleFunc("/admin/lectures/{id:[0-9]+}", s.Handler.LectureHandler.AdminGetByID)
	r.HandleFunc("/admin/researches/{id:[0-9]+}", s.Handler.ResearchHandler.AdminGetByID)
	r.HandleFunc("/admin/equipments/{id:[0-9]+}", s.Handler.EquipmentHandler.AdminGetByID)
	r.HandleFunc("/admin/tags/{id:[0-9]+}", s.Handler.TagHandler.AdminGetByID)

	r.HandleFunc("/admin/members/new", s.Handler.UserHandler.AdminCreate)
	r.HandleFunc("/admin/members/{id}/edit", s.Handler.UserHandler.AdminUpdateByID)
	r.HandleFunc("/admin/members/{id}/delete", s.Handler.UserHandler.AdminDeleteByID)

	r.HandleFunc("/admin/activities/new", s.Handler.ActivityHandler.Create)
	r.HandleFunc("/admin/activities/{id}/edit", s.Handler.ActivityHandler.UpdateByID)
	r.HandleFunc("/admin/activities/{id}/delete", s.Handler.ActivityHandler.AdminDeleteByID)

	r.HandleFunc("/admin/societies/new", s.Handler.SocietyHandler.Create)
	r.HandleFunc("/admin/societies/{id}/edit", s.Handler.SocietyHandler.UpdateByID)
	r.HandleFunc("/admin/societies/{id}/delete", s.Handler.SocietyHandler.AdminDeleteByID)

	r.HandleFunc("/admin/jobs/new", s.Handler.JobHandler.Create)
	r.HandleFunc("/admin/jobs/{id}/edit", s.Handler.JobHandler.UpdateByID)
	r.HandleFunc("/admin/jobs/{id}/delete", s.Handler.JobHandler.AdminDeleteByID)

	r.HandleFunc("/admin/lectures/new", s.Handler.LectureHandler.AdminCreate)
	r.HandleFunc("/admin/lectures/{id}/edit", s.Handler.LectureHandler.AdminUpdateByID)
	r.HandleFunc("/admin/lectures/{id}/delete", s.Handler.LectureHandler.AdminDeleteByID)

	r.HandleFunc("/admin/researches/new", s.Handler.ResearchHandler.Create)
	r.HandleFunc("/admin/researches/{id}/edit", s.Handler.ResearchHandler.UpdateByID)
	r.HandleFunc("/admin/researches/{id}/delete", s.Handler.ResearchHandler.AdminDeleteByID)

	r.HandleFunc("/admin/equipments/new", s.Handler.EquipmentHandler.Create)
	r.HandleFunc("/admin/equipments/{id}/edit", s.Handler.EquipmentHandler.UpdateByID)
	r.HandleFunc("/admin/equipments/{id}/delete", s.Handler.EquipmentHandler.AdminDeleteByID)

	r.HandleFunc("/admin/tags/new", s.Handler.TagHandler.Create)
	r.HandleFunc("/admin/tags/{id}/edit", s.Handler.TagHandler.UpdateByID)
	r.HandleFunc("/admin/tags/{id}/delete", s.Handler.TagHandler.AdminDeleteByID)

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

func adminDummyHandler(templateFile string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		funcMap := template.FuncMap{"convPageType": func(p string) string { return "dummy" }}
		t := template.New(templateFile).Funcs(funcMap)
		t, err := t.ParseFiles(
			"template/admin/"+templateFile,
			"template/admin/_footer.html",
			"template/admin/_header.html",
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
	Errors    []string
}
