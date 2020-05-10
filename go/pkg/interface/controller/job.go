package controller

import (
	"homepage/pkg/domain/model"
	"homepage/pkg/usecase/interactor"
)

type jobController struct {
	interactor.JobInteractor
}

type JobController interface {
	GetAll() (*JobsResponse, error)
}

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

type JobsResponse struct {
	Jobs []*JobResponse
}

type JobResponse struct {
	Company string
	Job     string
}

func convertToJobResponse(job *model.Job) *JobResponse {
	return &JobResponse{
		Company: job.Company,
		Job:     job.Job,
	}
}
