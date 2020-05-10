package repository

import (
	"homepage/pkg/domain/model"
	"homepage/pkg/usecase/interactor"
	"log"
)

type jobRepository struct {
	SQLHandler
}

func NewJobRepository(sh SQLHandler) interactor.JobRespotiroy {
	return &jobRepository{
		SQLHandler: sh,
	}
}

func (jr *jobRepository) FindAll() ([]*model.Job, error) {
	rows, err := jr.SQLHandler.Query(`
		SELECT company, job
		FROM jobs
	`)
	if err != nil {
		log.Println("job: FindAll: ", err)
		return []*model.Job{}, err
	}
	var jobs []*model.Job
	for rows.Next() {
		var job model.Job
		if err = rows.Scan(&job.Company, &job.Job); err != nil {
			log.Println("job: FindAll: ", err)
			continue
		}
		jobs = append(jobs, &job)
	}
	return jobs, nil
}
