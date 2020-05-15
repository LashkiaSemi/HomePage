package response

import (
	"log"
	"net/http"
	"text/template"
)

// Render テンプレートファイルを指定して描画
func Render(w http.ResponseWriter, templateFile string, info *Info, body interface{}) {
	t, err := template.ParseFiles(
		"template/"+templateFile,
		"template/_footer.html",
		"template/_header.html",
	)
	if err != nil {
		InternalServerError(w, info)
		return
	}

	if err = t.Execute(w, struct {
		Info *Info
		Data interface{}
	}{
		Info: info,
		Data: body,
	}); err != nil {
		log.Printf("[error] failed to execute template: %v", err)
		InternalServerError(w, info)
		return
	}
}

// Forbidden 403の時のアレとか
func Forbidden(w http.ResponseWriter) {
	t, _ := template.ParseFiles(
		"template/error.html",
		"template/_footer.html",
		"template/_header.html",
	)
	err := &ErrorData{
		Title:   "Forbidden",
		Message: `<a href="/login">ログイン</a>してください。`,
	}
	t.Execute(w, struct {
		Info  *Info
		Error *ErrorData
	}{
		Info:  &Info{},
		Error: err,
	})
}

// NotFound 404の時のアレとか
func NotFound(w http.ResponseWriter, info *Info) {
	t, _ := template.ParseFiles(
		"template/error.html",
		"template/_footer.html",
		"template/_header.html",
	)
	err := &ErrorData{
		Title:   "Not Found",
		Message: "お探しのページは見つかりませんでした。",
	}
	t.Execute(w, struct {
		Info  *Info
		Error *ErrorData
	}{
		Info:  info,
		Error: err,
	})
}

// InternalServerError サーバのエラー
func InternalServerError(w http.ResponseWriter, info *Info) {
	t, _ := template.ParseFiles(
		"template/error.html",
		"template/_footer.html",
		"template/_header.html",
	)
	err := &ErrorData{
		Title:   "Internal Server Error",
		Message: "サーバでエラーが発生しました。",
	}
	t.Execute(w, struct {
		Info  *Info
		Error *ErrorData
	}{
		Info:  info,
		Error: err,
	})
}

// AdminRender テンプレートファイルを指定して描画
func AdminRender(w http.ResponseWriter, templateFile string, info *Info, body interface{}) {
	funcMap := template.FuncMap{"convPageType": convertPageType}
	t := template.New(templateFile).Funcs(funcMap)
	t, err := t.ParseFiles(
		"template/admin/"+templateFile,
		"template/admin/_footer.html",
		"template/admin/_header.html",
	)
	if err != nil {
		log.Printf("[error] failed to parse template: %v", err)
		InternalServerError(w, info)
		return
	}

	if err = t.Execute(w, struct {
		Info *Info
		Data interface{}
	}{
		Info: info,
		Data: body,
	}); err != nil {
		log.Printf("[error] failed to execute template: %v", err)
		InternalServerError(w, info)
		return
	}
}

// Info ヘッダー描画用のデータ
type Info struct {
	PageType  string
	StudentID string
	Errors    []string
}

// ErrorData エラー描画用
type ErrorData struct {
	Title   string
	Message string
}

func convertPageType(pageType string) string {
	switch pageType {
	case "activities":
		return "活動内容管理"
	case "societies":
		return "学会発表管理"
	case "researches":
		return "卒業研究管理"
	case "jobs":
		return "就職先管理"
	case "members":
		return "メンバー管理"
	case "links":
		return "外部リンク管理"
	case "equipments":
		return "研究室備品管理"
	case "lectures":
		return "レクチャー管理"
	case "tags":
		return "タグ管理"
	default:
		return "default"
	}
}
