package server

import (
	"fmt"
	"homepage/pkg/configs"
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
	Router  *mux.Router
	Handler *handler.AppHandler // アプリケーションハンドラ
}

// Server ルーティングとか全部やってくれる子
type Server interface {
	Serve()
	HandleFunc(endpoint string, appFunc http.HandlerFunc) *mux.Route
}

// NewServer サーバを作るぞ！
func NewServer(host, port string, ah *handler.AppHandler) Server {
	return &server{
		Host:    host,
		Port:    port,
		Router:  mux.NewRouter(),
		Handler: ah,
	}
}

func (s *server) Serve() {
	// r := mux.NewRouter()
	s.Router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(configs.StaticDir))))
	//http.Dirの部分を絶対パスに変えればええねんな...
	s.Router.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir(configs.PublicDir))))

	s.HandleFunc("/health", healthHandler)

	// web site
	s.HandleFunc("/", s.Handler.StaticPageHandler.IndexHandler)
	s.HandleFunc("/login", s.Handler.UserHandler.Login)
	s.HandleFunc("/logout", middleware.Authorized(s.Handler.UserHandler.Logout))
	s.HandleFunc("/activities", s.Handler.ActivityHandler.GetActivities)
	s.HandleFunc("/societies", s.Handler.SocietyHandler.GetAll)
	s.HandleFunc("/researches", s.Handler.ResearchHandler.GetAll)
	s.HandleFunc("/jobs", s.Handler.JobHandler.GetAll)
	s.HandleFunc("/members", s.Handler.UserHandler.GetAllGroupByGrade)
	s.HandleFunc("/members/{id}", s.Handler.UserHandler.GetByID)
	s.HandleFunc("/members/edit/profile", middleware.Authorized(s.Handler.UserHandler.UpdateByID))
	s.HandleFunc("/members/edit/password", middleware.Authorized(s.Handler.UserHandler.UpdatePasswordByStudentID))
	s.HandleFunc("/links", handler.LinkHandler)
	s.HandleFunc("/equipments", middleware.Authorized(s.Handler.EquipmentHandler.GetAll))
	s.HandleFunc("/lectures", middleware.Authorized(s.Handler.LectureHandler.GetAll))
	s.HandleFunc("/lectures/new", middleware.Authorized(s.Handler.LectureHandler.Create))
	s.HandleFunc("/lectures/{id}/edit", middleware.Authorized(s.Handler.LectureHandler.UpdateByID))
	s.HandleFunc("/lectures/{id}/delete", middleware.Authorized(s.Handler.LectureHandler.DeleteByID))

	// admin site
	s.HandleFunc("/admin/login", middleware.Authorized(s.Handler.UserHandler.AdminLogin))
	s.HandleFunc("/admin", middleware.AdminAuthorized(s.Handler.StaticPageHandler.AdminIndexHandler))
	s.HandleFunc("/admin/activities", middleware.AdminAuthorized(s.Handler.ActivityHandler.AdminGetAll))
	s.HandleFunc("/admin/societies", middleware.AdminAuthorized(s.Handler.SocietyHandler.AdminGetAll))
	s.HandleFunc("/admin/researches", middleware.AdminAuthorized(s.Handler.ResearchHandler.AdminGetAll))
	s.HandleFunc("/admin/jobs", middleware.AdminAuthorized(s.Handler.JobHandler.AdminGetAll))
	s.HandleFunc("/admin/members", middleware.AdminAuthorized(s.Handler.UserHandler.AdminGetAll))
	s.HandleFunc("/admin/lectures", middleware.AdminAuthorized(s.Handler.LectureHandler.AdminGetAll))
	s.HandleFunc("/admin/equipments", middleware.AdminAuthorized(s.Handler.EquipmentHandler.AdminGetAll))
	s.HandleFunc("/admin/tags", middleware.AdminAuthorized(s.Handler.TagHandler.AdminGetAll))

	s.HandleFunc("/admin/members/{id:[0-9]+}", middleware.AdminAuthorized(s.Handler.UserHandler.AdminGetByID))
	s.HandleFunc("/admin/activities/{id:[0-9]+}", middleware.AdminAuthorized(s.Handler.ActivityHandler.AdminGetByID))
	s.HandleFunc("/admin/societies/{id:[0-9]+}", middleware.AdminAuthorized(s.Handler.SocietyHandler.AdminGeByID))
	s.HandleFunc("/admin/jobs/{id:[0-9]+}", middleware.AdminAuthorized(s.Handler.JobHandler.AdminGetByID))
	s.HandleFunc("/admin/lectures/{id:[0-9]+}", middleware.AdminAuthorized(s.Handler.LectureHandler.AdminGetByID))
	s.HandleFunc("/admin/researches/{id:[0-9]+}", middleware.AdminAuthorized(s.Handler.ResearchHandler.AdminGetByID))
	s.HandleFunc("/admin/equipments/{id:[0-9]+}", middleware.AdminAuthorized(s.Handler.EquipmentHandler.AdminGetByID))
	s.HandleFunc("/admin/tags/{id:[0-9]+}", middleware.AdminAuthorized(s.Handler.TagHandler.AdminGetByID))

	s.HandleFunc("/admin/members/new", middleware.AdminAuthorized(s.Handler.UserHandler.AdminCreate))
	s.HandleFunc("/admin/members/{id}/edit", middleware.AdminAuthorized(s.Handler.UserHandler.AdminUpdateByID))
	s.HandleFunc("/admin/members/{id}/delete", middleware.AdminAuthorized(s.Handler.UserHandler.AdminDeleteByID))

	s.HandleFunc("/admin/activities/new", middleware.AdminAuthorized(s.Handler.ActivityHandler.AdminCreate))
	s.HandleFunc("/admin/activities/{id}/edit", middleware.AdminAuthorized(s.Handler.ActivityHandler.AdminUpdateByID))
	s.HandleFunc("/admin/activities/{id}/delete", middleware.AdminAuthorized(s.Handler.ActivityHandler.AdminDeleteByID))

	s.HandleFunc("/admin/societies/new", middleware.AdminAuthorized(s.Handler.SocietyHandler.AdminCreate))
	s.HandleFunc("/admin/societies/{id}/edit", middleware.AdminAuthorized(s.Handler.SocietyHandler.AdminUpdateByID))
	s.HandleFunc("/admin/societies/{id}/delete", middleware.AdminAuthorized(s.Handler.SocietyHandler.AdminDeleteByID))

	s.HandleFunc("/admin/jobs/new", middleware.AdminAuthorized(s.Handler.JobHandler.AdminCreate))
	s.HandleFunc("/admin/jobs/{id}/edit", middleware.AdminAuthorized(s.Handler.JobHandler.AdminUpdateByID))
	s.HandleFunc("/admin/jobs/{id}/delete", middleware.AdminAuthorized(s.Handler.JobHandler.AdminDeleteByID))

	s.HandleFunc("/admin/lectures/new", middleware.AdminAuthorized(s.Handler.LectureHandler.AdminCreate))
	s.HandleFunc("/admin/lectures/{id}/edit", middleware.AdminAuthorized(s.Handler.LectureHandler.AdminUpdateByID))
	s.HandleFunc("/admin/lectures/{id}/delete", middleware.AdminAuthorized(s.Handler.LectureHandler.AdminDeleteByID))

	s.HandleFunc("/admin/researches/new", middleware.AdminAuthorized(s.Handler.ResearchHandler.AdminCreate))
	s.HandleFunc("/admin/researches/{id}/edit", middleware.AdminAuthorized(s.Handler.ResearchHandler.AdminUpdateByID))
	s.HandleFunc("/admin/researches/{id}/delete", middleware.AdminAuthorized(s.Handler.ResearchHandler.AdminDeleteByID))

	s.HandleFunc("/admin/equipments/new", middleware.AdminAuthorized(s.Handler.EquipmentHandler.AdminCreate))
	s.HandleFunc("/admin/equipments/{id}/edit", middleware.AdminAuthorized(s.Handler.EquipmentHandler.AdminUpdateByID))
	s.HandleFunc("/admin/equipments/{id}/delete", middleware.AdminAuthorized(s.Handler.EquipmentHandler.AdminDeleteByID))

	s.HandleFunc("/admin/tags/new", middleware.AdminAuthorized(s.Handler.TagHandler.AdminCreate))
	s.HandleFunc("/admin/tags/{id}/edit", middleware.AdminAuthorized(s.Handler.TagHandler.AdminUpdateByID))
	s.HandleFunc("/admin/tags/{id}/delete", middleware.AdminAuthorized(s.Handler.TagHandler.AdminDeleteByID))

	log.Printf("[info] server running http://%v:%v", s.Host, s.Port)
	if err := http.ListenAndServe(
		fmt.Sprintf("%s:%s", s.Host, s.Port),
		s.Router,
	); err != nil {
		log.Printf("[error] failed to running server: %v", err)
		return
	}
}

func (s *server) HandleFunc(endpoint string, appFunc http.HandlerFunc) *mux.Route {
	return s.Router.HandleFunc(endpoint, httpHandler(appFunc))
}

func httpHandler(appFunc http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Add("Access-Control-Allow-Headers", "Content-Type,Accept,Origin")
		writer.Header().Add("Access-Control-Allow-Credentials", "true")
		writer.Header().Add("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE")

		if request.Method == http.MethodOptions {
			return
		}

		appFunc(writer, request)
	}
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
