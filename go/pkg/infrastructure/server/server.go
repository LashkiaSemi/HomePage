package server

import (
	"fmt"
	"homepage/pkg/infrastructure/handler"
	"homepage/pkg/infrastructure/server/middleware"
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

type server struct {
	Host    string
	Port    string
	Handler *handler.AppHandler // アプリケーションハンドラ
}

// Server ルーティングとか全部やってくれる子
type Server interface {
	Serve()
}

// NewServer サーバを作るぞ！
func NewServer(host, port string, ah *handler.AppHandler) Server {
	return &server{
		Host:    host,
		Port:    port,
		Handler: ah,
	}
}

func (s *server) Serve() {
	r := mux.NewRouter()
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))

	//TODO: 環境変数とかのがいいかも。レクチャーの資料とかしまってある場所
	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("public/"))))

	r.HandleFunc("/health", healthHandler)

	// web site
	r.HandleFunc("/", s.Handler.StaticPageHandler.IndexHandler)
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
	r.HandleFunc("/admin", middleware.AdminAuthorized(s.Handler.StaticPageHandler.AdminIndexHandler))
	// TODO: ログイン限定にするほうがいいねこれ...
	r.HandleFunc("/admin/login", s.Handler.UserHandler.AdminLogin)
	r.HandleFunc("/admin/activities", middleware.AdminAuthorized(s.Handler.ActivityHandler.AdminGetAll))
	r.HandleFunc("/admin/societies", middleware.AdminAuthorized(s.Handler.SocietyHandler.AdminGetAll))
	r.HandleFunc("/admin/researches", middleware.AdminAuthorized(s.Handler.ResearchHandler.AdminGetAll))
	r.HandleFunc("/admin/jobs", middleware.AdminAuthorized(s.Handler.JobHandler.AdminGetAll))
	r.HandleFunc("/admin/members", middleware.AdminAuthorized(s.Handler.UserHandler.AdminGetAll))
	r.HandleFunc("/admin/lectures", middleware.AdminAuthorized(s.Handler.LectureHandler.AdminGetAll))
	r.HandleFunc("/admin/equipments", middleware.AdminAuthorized(s.Handler.EquipmentHandler.AdminGetAll))
	r.HandleFunc("/admin/tags", middleware.AdminAuthorized(s.Handler.TagHandler.AdminGetAll))

	r.HandleFunc("/admin/members/{id:[0-9]+}", middleware.AdminAuthorized(s.Handler.UserHandler.AdminGetByID))
	r.HandleFunc("/admin/activities/{id:[0-9]+}", middleware.AdminAuthorized(s.Handler.ActivityHandler.AdminGetByID))
	r.HandleFunc("/admin/societies/{id:[0-9]+}", middleware.AdminAuthorized(s.Handler.SocietyHandler.AdminGeByID))
	r.HandleFunc("/admin/jobs/{id:[0-9]+}", middleware.AdminAuthorized(s.Handler.JobHandler.AdminGetByID))
	r.HandleFunc("/admin/lectures/{id:[0-9]+}", middleware.AdminAuthorized(s.Handler.LectureHandler.AdminGetByID))
	r.HandleFunc("/admin/researches/{id:[0-9]+}", middleware.AdminAuthorized(s.Handler.ResearchHandler.AdminGetByID))
	r.HandleFunc("/admin/equipments/{id:[0-9]+}", middleware.AdminAuthorized(s.Handler.EquipmentHandler.AdminGetByID))
	r.HandleFunc("/admin/tags/{id:[0-9]+}", middleware.AdminAuthorized(s.Handler.TagHandler.AdminGetByID))

	r.HandleFunc("/admin/members/new", middleware.AdminAuthorized(s.Handler.UserHandler.AdminCreate))
	r.HandleFunc("/admin/members/{id}/edit", middleware.AdminAuthorized(s.Handler.UserHandler.AdminUpdateByID))
	r.HandleFunc("/admin/members/{id}/delete", middleware.AdminAuthorized(s.Handler.UserHandler.AdminDeleteByID))

	r.HandleFunc("/admin/activities/new", middleware.AdminAuthorized(s.Handler.ActivityHandler.AdminCreate))
	r.HandleFunc("/admin/activities/{id}/edit", middleware.AdminAuthorized(s.Handler.ActivityHandler.AdminUpdateByID))
	r.HandleFunc("/admin/activities/{id}/delete", middleware.AdminAuthorized(s.Handler.ActivityHandler.AdminDeleteByID))

	r.HandleFunc("/admin/societies/new", middleware.AdminAuthorized(s.Handler.SocietyHandler.AdminCreate))
	r.HandleFunc("/admin/societies/{id}/edit", middleware.AdminAuthorized(s.Handler.SocietyHandler.AdminUpdateByID))
	r.HandleFunc("/admin/societies/{id}/delete", middleware.AdminAuthorized(s.Handler.SocietyHandler.AdminDeleteByID))

	r.HandleFunc("/admin/jobs/new", middleware.AdminAuthorized(s.Handler.JobHandler.AdminCreate))
	r.HandleFunc("/admin/jobs/{id}/edit", middleware.AdminAuthorized(s.Handler.JobHandler.AdminUpdateByID))
	r.HandleFunc("/admin/jobs/{id}/delete", middleware.AdminAuthorized(s.Handler.JobHandler.AdminDeleteByID))

	r.HandleFunc("/admin/lectures/new", middleware.AdminAuthorized(s.Handler.LectureHandler.AdminCreate))
	r.HandleFunc("/admin/lectures/{id}/edit", middleware.AdminAuthorized(s.Handler.LectureHandler.AdminUpdateByID))
	r.HandleFunc("/admin/lectures/{id}/delete", middleware.AdminAuthorized(s.Handler.LectureHandler.AdminDeleteByID))

	r.HandleFunc("/admin/researches/new", middleware.AdminAuthorized(s.Handler.ResearchHandler.AdminCreate))
	r.HandleFunc("/admin/researches/{id}/edit", middleware.AdminAuthorized(s.Handler.ResearchHandler.AdminUpdateByID))
	r.HandleFunc("/admin/researches/{id}/delete", middleware.AdminAuthorized(s.Handler.ResearchHandler.AdminDeleteByID))

	r.HandleFunc("/admin/equipments/new", middleware.AdminAuthorized(s.Handler.EquipmentHandler.AdminCreate))
	r.HandleFunc("/admin/equipments/{id}/edit", middleware.AdminAuthorized(s.Handler.EquipmentHandler.AdminUpdateByID))
	r.HandleFunc("/admin/equipments/{id}/delete", middleware.AdminAuthorized(s.Handler.EquipmentHandler.AdminDeleteByID))

	r.HandleFunc("/admin/tags/new", middleware.AdminAuthorized(s.Handler.TagHandler.AdminCreate))
	r.HandleFunc("/admin/tags/{id}/edit", middleware.AdminAuthorized(s.Handler.TagHandler.AdminUpdateByID))
	r.HandleFunc("/admin/tags/{id}/delete", middleware.AdminAuthorized(s.Handler.TagHandler.AdminDeleteByID))

	log.Printf("[info] server running http://%v:%v", s.Host, s.Port)
	http.ListenAndServe(
		fmt.Sprintf("%s:%s", s.Host, s.Port),
		r,
	)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("health")
}

// dummyHandler web site用のダミーハンドラ
func dummyHandler(templateFile string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles(
			"template/"+templateFile,
			"template/_footer.html",
			"template/_header.html",
		)
		if err != nil {
			log.Printf("[error] failed to parse template: %v", err)
			return
		}
		if err = t.Execute(w, struct {
			Info *dummyInfo
		}{
			Info: &dummyInfo{
				StudentID: "dummy",
				PageType:  "dummy",
			},
		}); err != nil {
			log.Printf("[error] failed to execute template: %v", err)
			return
		}
	}
}

// adminDummyHandler admin site用のダミーハンドラ
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
			log.Printf("[error] failed to parse template: %v", err)
			return
		}
		if err = t.Execute(w, struct {
			Info *dummyInfo
		}{
			Info: &dummyInfo{
				StudentID: "dummy",
				PageType:  "dummy",
			},
		}); err != nil {
			log.Printf("[error] failed to execute template: %v", err)
			return
		}
	}
}

type dummyInfo struct {
	StudentID string
	PageType  string
	Errors    []string
}
