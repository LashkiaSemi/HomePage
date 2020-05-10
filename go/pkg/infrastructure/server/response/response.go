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
func Success(w http.ResponseWriter, templateFile string, header *HeaderData, body interface{}) {
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
		Header *HeaderData
		Data   interface{}
	}{
		Header: header,
		Data:   body,
	}); err != nil {
		// TODO: redirect internal server error
		log.Printf("failed to execute template: %v", err)
	}
}

func BadRequest() {

}

func UnAuthorized() {

}

func InternalServerError(w http.ResponseWriter, header *HeaderData) {
	t, _ := template.ParseFiles(
		"template/error.html",
		"template/_footer.html",
		"template/_header.html",
	)
	t.Execute(w, struct {
		Header *HeaderData
	}{
		Header: header,
	})
}

// HeaderData ヘッダー描画用のデータ
type HeaderData struct {
	IsLogin  bool
	PageType string
}
