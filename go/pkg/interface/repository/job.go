package repository

import (
	"homepage/pkg/domain/entity"
	"homepage/pkg/domain/repository"
	"log"

	"github.com/pkg/errors"
)

type jobRepository struct {
	SQLHandler
}

// NewJobRepository リポジトリの作成
func NewJobRepository(sh SQLHandler) repository.JobRepository {
	return &jobRepository{
		SQLHandler: sh,
	}
}

func (jr *jobRepository) FindAll() ([]*entity.Job, error) {
	rows, err := jr.SQLHandler.Query(`
		SELECT id, company, job
		FROM jobs
		ORDER BY id DESC
	`)
	var jobs []*entity.Job
	if err != nil {
		if err == jr.SQLHandler.ErrNoRows() {
			log.Printf("[warn] hit no data: %v", err)
			return jobs, nil
		}
		err = errors.Wrap(err, "failed to execute query")
		return jobs, err
	}
	for rows.Next() {
		var job entity.Job
		if err = rows.Scan(&job.ID, &job.Company, &job.Job); err != nil {
			log.Println("[warn] rows.Scan skip: ", err)
			continue
		}
		jobs = append(jobs, &job)
	}
	return jobs, nil
}

func (jr *jobRepository) FindByID(id int) (*entity.Job, error) {
	row := jr.SQLHandler.QueryRow(`
		SELECT id, company, job
		FROM jobs
		WHERE id=?
	`, id)
	var data entity.Job
	if err := row.Scan(&data.ID, &data.Company, &data.Job); err != nil {
		err = errors.Wrap(err, "failed to bind data")
		return &data, err
	}
	return &data, nil
}

func (jr *jobRepository) Create(data *entity.Job) (int, error) {
	result, err := jr.SQLHandler.Execute(`
		INSERT INTO jobs(company, job, created_at, updated_at)
		VALUES (?,?,?,?)
	`, data.Company, data.Job, data.CreatedAt, data.UpdatedAt)
	if err != nil {
		err = errors.Wrap(err, "failed to execute query")
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		err = errors.Wrap(err, "failed to get id")
		return 0, err
	}
	return int(id), nil
}

func (jr *jobRepository) UpdateByID(data *entity.Job) error {
	_, err := jr.SQLHandler.Execute(`
		UPDATE jobs
		SET company=?, job=?, updated_at=?
		WHERE id=?
	`, data.Company, data.Job, data.UpdatedAt, data.ID)
	if err != nil {
		err = errors.Wrap(err, "failed to execute query")
		return err
	}
	return nil
}

func (jr *jobRepository) DeleteByID(id int) error {
	_, err := jr.SQLHandler.Execute(`
		DELETE FROM jobs
		WHERE id=?
	`, id)
	if err != nil {
		err = errors.Wrap(err, "failed to execute query")
	}
	return err
}
