package response

import (
	"log"
	"net/http"
	"text/template"
)

// func Success(w http.ResponseWriter, res interface{}) {
// 	data, err := json.Marshal(res)
// 	if err != nil {
// 		log.Println("failed to marshal json: %v", err)
// 	}
// 	w.Write(data)
// }

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

func BadRequest() {

}

func UnAuthorized() {

}

func InternalServerError(w http.ResponseWriter, info *Info) {
	t, _ := template.ParseFiles(
		"template/error.html",
		"template/_footer.html",
		"template/_header.html",
	)
	t.Execute(w, struct {
		Info *Info
	}{
		Info: info,
	})
}

// HeaderData ヘッダー描画用のデータ
type Info struct {
	IsLogin  bool
	PageType string
}
