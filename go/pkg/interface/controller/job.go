package controller

import (
	"homepage/pkg/entity"
	"homepage/pkg/usecase/interactor"
)

type jobController struct {
	interactor.JobInteractor
}

// JobController 就職先の入出力を変換
type JobController interface {
	GetAll() (*JobsResponse, error)
}

// NewJobController コントローラの作成
func NewJobController(ji interactor.JobInteractor) JobController {
	return &jobController{
		JobInteractor: ji,
	}
}

func (jc *jobController) GetAll() (*JobsResponse, error) {
	jobs, err := jc.JobInteractor.GetAll()
	if err != nil {
		return &JobsResponse{}, err
	}
	var res JobsResponse
	for _, job := range jobs {
		res.Jobs = append(res.Jobs, convertToJobResponse(job))
	}
	return &res, nil
}

// JobsResponse 就職先のレスポンス
type JobsResponse struct {
	Jobs []*JobResponse
}

// JobResponse 就職先のレスポンス
type JobResponse struct {
	Company string
	Job     string
}

func convertToJobResponse(job *entity.Job) *JobResponse {
	return &JobResponse{
		Company: job.Company,
		Job:     job.Job,
	}
}
