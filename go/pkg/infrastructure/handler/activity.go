package handler

import (
	"homepage/pkg/infrastructure/server/response"
	"net/http"
	"time"
)

type activityHandler struct {
}

type ActivityHandler interface {
	GetActivities(w http.ResponseWriter, r *http.Request)
}

func NewActivityHandler() ActivityHandler {
	return &activityHandler{}
}

func (ah *activityHandler) GetActivities(w http.ResponseWriter, r *http.Request) {
	info := createInfo(r, "activity")

	// TODO: get data
	datas := []*activity{
		&activity{
			Date:     time.Now(),
			Activity: "れくちゃー",
		},
		&activity{
			Date:     time.Now(),
			Activity: "ぜみ見学",
		},
		&activity{
			Date:     time.Now(),
			Activity: "選考",
		},
	}

	body := struct {
		Activities []*activity
	}{
		Activities: datas,
	}
	response.Success(w, "activity/index.html", info, body)
}

// TODO: move handler to controller
type activity struct {
	Date     time.Time
	Activity string
}
