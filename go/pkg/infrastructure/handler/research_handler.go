package handler

import (
	"encoding/json"
	"errors"
	"homepage/conf"
	"homepage/pkg/domain"
	"homepage/pkg/domain/logger"
	"homepage/pkg/infrastructure/server/response"
	"homepage/pkg/interface/controller"
	"homepage/pkg/interface/repository"
	"homepage/pkg/usecase/interactor"
	"net/http"
)

// ResearchHandler ハンドラ
type ResearchHandler interface {
	GetAll(w http.ResponseWriter, r *http.Request)
	GetByID(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type researchHandler struct {
	ResearchController controller.ResearchController
}

// NewResearchHandler ハンドラを作成する
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
	res, err := rh.ResearchController.ShowAll()
	if err != nil {
		response.HTTPError(w, err)
		return
	}
	response.Success(w, res)
}

func (rh *researchHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	resID, err := getIntParameter(r.URL.Path, "/researches/", "")
	if err != nil {
		logger.Warn("research get: params error. userID parameter is not int. researchID is ", resID)
		response.HTTPError(w, domain.BadRequest(errors.New("params error")))
		return
	}

	res, err := rh.ResearchController.ShowByID(resID)
	if err != nil {
		response.HTTPError(w, err)
		return
	}
	response.Success(w, res)
}

func (rh *researchHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req controller.UpdateResearchRequest
	body := []byte(r.FormValue("body"))
	err := json.Unmarshal(body, &req)
	if err != nil {
		logger.Error(err)
		response.HTTPError(w, domain.InternalServerError(err))
		return
	}
	file, reader, err := r.FormFile("file")
	if err != nil {
		logger.Error("lecture handler: ", err)
		response.HTTPError(w, domain.InternalServerError(err))
		return
	}
	defer file.Close()
	req.File = reader.Filename

	// save file
	err = saveFile(file, conf.FileDir+"/researches/", req.File)
	if err != nil {
		logger.Error("research create handler: ", err)
		response.HTTPError(w, domain.InternalServerError(err))
		return
	}

	// save db
	res, err := rh.ResearchController.Create(&req)
	if err != nil {
		response.HTTPError(w, err)
		return
	}
	response.Success(w, res)
}

func (rh *researchHandler) Update(w http.ResponseWriter, r *http.Request) {
	resID, err := getIntParameter(r.URL.Path, "/researches/", "")
	if err != nil {
		logger.Warn("research get: params error. userID parameter is not int. researchID is ", resID)
		response.HTTPError(w, domain.BadRequest(errors.New("params error")))
		return
	}

	// read data
	var req controller.UpdateResearchRequest
	body := []byte(r.FormValue("body"))
	err = json.Unmarshal(body, &req)
	if err != nil {
		logger.Error(err)
		response.HTTPError(w, domain.InternalServerError(err))
		return
	}
	file, reader, err := r.FormFile("file")
	if file != nil {
		defer file.Close()
		req.File = reader.Filename
	}
	if err != nil {
		if file != nil {
			logger.Error("research update handler: ", err)
			response.HTTPError(w, domain.InternalServerError(err))
			return
		}
	}

	// file save
	if file != nil {
		err = saveFile(file, conf.FileDir+"/researches/", req.File)
		if err != nil {
			logger.Error("research update handler: ", err)
			response.HTTPError(w, domain.InternalServerError(err))
			return
		}
	}
	// save db
	res, err := rh.ResearchController.Update(resID, &req)
	if err != nil {
		response.HTTPError(w, err)
		return
	}
	response.Success(w, res)
}

func (rh *researchHandler) Delete(w http.ResponseWriter, r *http.Request) {
	resID, err := getIntParameter(r.URL.Path, "/researches/", "")
	if err != nil {
		logger.Warn("research get: params error. userID parameter is not int. researchID is ", resID)
		response.HTTPError(w, domain.BadRequest(errors.New("params error")))
		return
	}

	err = rh.ResearchController.Delete(resID)
	if err != nil {
		response.HTTPError(w, err)
		return
	}
	response.NoContent(w)
}
