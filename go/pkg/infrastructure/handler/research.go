package handler

import (
	"homepage/pkg/infrastructure/auth"
	"homepage/pkg/infrastructure/server/response"
	"homepage/pkg/interface/controller"
	"homepage/pkg/interface/repository"
	"homepage/pkg/usecase/interactor"
	"net/http"
)

type researchHandler struct {
	controller.ResearchController
}

// ResearchHandler 卒業研究の入出力を受け付け
type ResearchHandler interface {
	GetAll(w http.ResponseWriter, r *http.Request)
}

// NewResearchHandler ハンドラの作成
func NewResearchHandler(sh repository.SQLHandler) ResearchHandler {
	return &researchHandler{
		ResearchController: controller.NewResearchController(
			interactor.NewResearchInteractor(
				repository.NewResearchRepository(sh),
			),
		),
	}
}

func (rh *researchHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	info := createInfo(r, "research", auth.GetStudentIDFromCookie(r))

	res, err := rh.ResearchController.GetAll()
	if err != nil {
		// log.Println("researchHandler: GetAll: ", err)
		response.InternalServerError(w, info)
	}
	response.Success(w, "research/index.html", info, res)
}
