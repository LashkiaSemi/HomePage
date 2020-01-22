package controller

import (
	"homepage/pkg/usecase/interactor"
)

// EmployController コントローラ
type EmployController interface {
	ShowAll() (GetJobsResponse, error)
	ShowByID(jobID int) (GetJobResponse, error)
	Create(req *UpdateJobRequest) (GetJobResponse, error)
	Update(jobID int, req *UpdateJobRequest) (GetJobResponse, error)
	Delete(jobID int) error
}

type employController struct {
	EmployInteractor interactor.EmployInteractor
}

// NewEmployController コントローラの作成
func NewEmployController(ei interactor.EmployInteractor) EmployController {
	return &employController{
		EmployInteractor: ei,
	}
}

func (ec *employController) ShowAll() (res GetJobsResponse, err error) {
	jobs, err := ec.EmployInteractor.FetchAll()
	if err != nil {
		return
	}
	for _, job := range jobs {
		res.Jobs = append(res.Jobs, GetJobResponse{
			ID:      job.ID,
			Company: job.Company,
			Job:     job.Job,
		})
	}
	return
}

// GetJobsResponse 企業
type GetJobsResponse struct {
	Jobs []GetJobResponse `json:"jobs"`
}

// GetJobResponse 企業
type GetJobResponse struct {
	ID      int    `json:"id"`
	Company string `json:"company"`
	Job     string `json:"job"`
}

func (ec *employController) ShowByID(jobID int) (res GetJobResponse, err error) {
	job, err := ec.EmployInteractor.FetchByID(jobID)
	if err != nil {
		return
	}
	res.ID = job.ID
	res.Company = job.Company
	res.Job = job.Job
	return
}

func (ec *employController) Create(req *UpdateJobRequest) (res GetJobResponse, err error) {
	job, err := ec.EmployInteractor.Add(req.Company, req.Job)
	if err != nil {
		return
	}
	res.ID = job.ID
	res.Company = job.Company
	res.Job = job.Job
	return
}

// UpdateJobRequest 新規と更新のリクエスト
type UpdateJobRequest struct {
	Company string `json:"company"`
	Job     string `json:"job"`
}

func (ec *employController) Update(jobID int, req *UpdateJobRequest) (res GetJobResponse, err error) {
	job, err := ec.EmployInteractor.Update(jobID, req.Company, req.Job)
	if err != nil {
		return
	}
	res.ID = job.ID
	res.Company = job.Company
	res.Job = job.Job
	return
}

func (ec *employController) Delete(jobID int) error {
	return ec.EmployInteractor.Delete(jobID)
}
