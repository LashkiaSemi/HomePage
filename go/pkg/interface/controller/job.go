package controller

import (
	"fmt"
	"homepage/pkg/domain/entity"
	"homepage/pkg/usecase/interactor"
	"strconv"

	"github.com/pkg/errors"
)

type jobController struct {
	interactor.JobInteractor
}

// JobController 就職先の入出力を変換
type JobController interface {
	GetAll() (*JobsResponse, error)
	GetByID(id int) (*JobResponse, error)

	Create(company, job string) (int, error)
	UpdateByID(id int, company, job string) error

	DeleteByID(id int) error

	// admin
	AdminGetAll() ([]map[string]string, error)
	AdminGetByID(id int) (*FieldsResponse, error)
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
		err = errors.Wrap(err, "failed to original data for response")
		return &JobsResponse{}, err
	}
	var res JobsResponse
	for _, job := range jobs {
		res.Jobs = append(res.Jobs, convertToJobResponse(job))
	}
	return &res, nil
}

func (jc *jobController) GetByID(id int) (*JobResponse, error) {
	data, err := jc.JobInteractor.GetByID(id)
	if err != nil {
		err = errors.Wrap(err, "failed to original data for response")
		return &JobResponse{}, err
	}
	return convertToJobResponse(data), nil
}

func (jc *jobController) Create(company, job string) (int, error) {
	return jc.JobInteractor.Create(company, job)
}

func (jc *jobController) UpdateByID(id int, company, job string) error {
	return jc.JobInteractor.UpdateByID(id, company, job)
}

func (jc *jobController) DeleteByID(id int) error {
	return jc.JobInteractor.DeleteByID(id)
}

// admin
func (jc *jobController) AdminGetAll() ([]map[string]string, error) {
	var res []map[string]string
	datas, err := jc.JobInteractor.GetAll()
	if err != nil {
		err = errors.Wrap(err, "failed to original data for response")
		return res, err
	}
	for _, data := range datas {
		res = append(res, map[string]string{
			"id":    strconv.Itoa(data.ID),
			"title": fmt.Sprintf("%s / %s", data.Company, data.Job),
		})
	}
	return res, nil
}

func (jc *jobController) AdminGetByID(id int) (*FieldsResponse, error) {
	var res FieldsResponse
	data, err := jc.JobInteractor.GetByID(id)
	if err != nil {
		err = errors.Wrap(err, "failed to original data for response")
		return &res, err
	}
	res.Fields = append(res.Fields,
		&Field{Key: "ID", Value: data.ID},
		&Field{Key: "企業名", Value: data.Company},
		&Field{Key: "仕事", Value: data.Job},
	)
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
