package repository

import (
	"homepage/pkg/entity"
	"homepage/pkg/usecase/interactor"
	"log"
)

type jobRepository struct {
	SQLHandler
}

// NewJobRepository リポジトリの作成
func NewJobRepository(sh SQLHandler) interactor.JobRepository {
	return &jobRepository{
		SQLHandler: sh,
	}
}

func (jr *jobRepository) FindAll() ([]*entity.Job, error) {
	rows, err := jr.SQLHandler.Query(`
		SELECT id, company, job
		FROM jobs
	`)
	if err != nil {
		log.Println("job: FindAll: ", err)
		return []*entity.Job{}, err
	}
	var jobs []*entity.Job
	for rows.Next() {
		var job entity.Job
		if err = rows.Scan(&job.ID, &job.Company, &job.Job); err != nil {
			log.Println("job: FindAll: ", err)
			continue
		}
		jobs = append(jobs, &job)
	}
	return jobs, nil
}
