package response

import (
	"encoding/json"
	"homepage/pkg/domain"
	"homepage/pkg/domain/logger"
	"net/http"
)

// Success 200の処理
func Success(w http.ResponseWriter, data interface{}) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		logger.Error(err)
		HTTPError(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

// NoContent 204の処理
func NoContent(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}

// HTTPError error全般
func HTTPError(w http.ResponseWriter, err error) {
	e, ok := err.(domain.Error)
	if !ok {
		e = domain.InternalServerError(err)
	}
	jsonData, _ := json.Marshal(&errorResponse{
		Code:    e.GetStatusCode(),
		Message: e.Error(),
	})
	w.WriteHeader(e.GetStatusCode())
	w.Write(jsonData)
}

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
