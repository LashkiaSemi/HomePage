package response

import (
	"encoding/json"
	"log"
	"net/http"
)

func Success(w http.ResponseWriter, res interface{}) {
	data, err := json.Marshal(res)
	if err != nil {
		log.Println("failed to marshal json: %v", err)
	}
	w.Write(data)
}

func BadRequest() {

}

func UnAuthorized() {

}
