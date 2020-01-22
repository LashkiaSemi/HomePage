package repository

import (
	"errors"
	"homepage/pkg/domain"
	"homepage/pkg/domain/logger"
	"homepage/pkg/usecase/interactor"
	"time"
)

type employRepository struct {
	SQLHandler
}

// NewEmployRepository リポジトリを作成
func NewEmployRepository(sh SQLHandler) interactor.EmployRepository {
	return &employRepository{
		SQLHandler: sh,
	}
}

func (er *employRepository) FindAll() (jobs domain.Jobs, err error) {
	rows, err := er.SQLHandler.Query("SELECT id, company, job FROM jobs")
	for rows.Next() {
		var job domain.Job
		if err = rows.Scan(&job.ID, &job.Company, &job.Job); err != nil {
			continue
		}
		jobs = append(jobs, job)
	}
	return
}

func (er *employRepository) FindByID(jobID int) (job domain.Job, err error) {
	row := er.SQLHandler.QueryRow("SELECT id, company, job FROM jobs WHERE id=?", jobID)
	if row.Scan(&job.ID, &job.Company, &job.Job); err != nil {
		if err == er.SQLHandler.ErrNoRows() {
			logger.Warn("employ findByID: ", err)
			return job, domain.NotFound(errors.New("content not found"))
		}
		logger.Error("employ findByID: ", err)
		return job, domain.InternalServerError(err)
	}
	return
}

func (er *employRepository) Store(company, job string, createdAt time.Time) (int, error) {
	result, err := er.SQLHandler.Execute(
		"INSERT INTO jobs(company, job, created_at, updated_at) VALUES (?,?,?,?)",
		company, job, createdAt, createdAt,
	)
	if err != nil {
		return 0, err
	}
	id, _ := result.LastInsertId()
	return int(id), err
}

func (er *employRepository) Update(jobID int, company, job string, updatedAt time.Time) error {
	query, args, _ := makeUpdateQuery(
		"jobs",
		map[string]interface{}{
			"company":    company,
			"job":        job,
			"updated_at": updatedAt,
		},
		map[string]interface{}{
			"id": jobID,
		},
	)
	_, err := er.SQLHandler.Execute(query, args...)
	return err
}

func (er *employRepository) Delete(jobID int) error {
	_, err := er.SQLHandler.Execute("DELETE FROM jobs WHERE id=?", jobID)
	return err
}
