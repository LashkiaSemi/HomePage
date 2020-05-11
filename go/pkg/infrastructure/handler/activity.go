package handler

import (
	"homepage/pkg/infrastructure/auth"
	"homepage/pkg/infrastructure/server/response"
	"net/http"
	"time"
)

type activityHandler struct {
}

// ActivityHandler 活動内容の入出力を受付
type ActivityHandler interface {
	GetActivities(w http.ResponseWriter, r *http.Request)
}

// NewActivityHandler ハンドラの作成
func NewActivityHandler() ActivityHandler {
	return &activityHandler{}
}

func (ah *activityHandler) GetActivities(w http.ResponseWriter, r *http.Request) {
	info := createInfo(r, "activity", auth.GetStudentIDFromCookie(r))

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
