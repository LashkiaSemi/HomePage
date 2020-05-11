package response

import (
	"log"
	"net/http"
	"text/template"
)

// Success テンプレートファイルを指定して描画
func Success(w http.ResponseWriter, templateFile string, info *Info, body interface{}) {
	t, err := template.ParseFiles(
		"template/"+templateFile,
		"template/_footer.html",
		"template/_header.html",
	)
	if err != nil {
		// TODO: redirect internal server error
		log.Printf("failed to parse template: %v", err)
	}

	if err = t.Execute(w, struct {
		Info *Info
		Data interface{}
	}{
		Info: info,
		Data: body,
	}); err != nil {
		// TODO: redirect internal server error
		log.Printf("failed to execute template: %v", err)
	}
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

// Info ヘッダー描画用のデータ
type Info struct {
	PageType  string
	StudentID string
}

// ErrorData エラー描画用
type ErrorData struct {
	Title   string
	Message string
}
