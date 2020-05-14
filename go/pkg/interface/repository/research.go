package repository

import (
	"homepage/pkg/entity"
	"homepage/pkg/usecase/interactor"
	"log"

	"github.com/pkg/errors"
)

type researchRepository struct {
	SQLHandler
}

// NewResearchRepository リポジトリの作成
func NewResearchRepository(sh SQLHandler) interactor.ResearchRepository {
	return &researchRepository{
		SQLHandler: sh,
	}
}

func (rr *researchRepository) FindAll() ([]*entity.Research, error) {
	rows, err := rr.SQLHandler.Query(`
		SELECT id, title, author, file, comments,  activation, created_at
		FROM researches
		ORDER BY created_at DESC
	`)
	var res []*entity.Research
	if err != nil {
		if err == rr.SQLHandler.ErrNoRows() {
			log.Printf("[warn] no hit data: %v", err)
			return res, nil
		}
		err = errors.Wrap(err, "failed to execute query")
		return res, err
	}
	for rows.Next() {
		var data entity.Research
		if err = rows.Scan(&data.ID, &data.Title, &data.Author, &data.File, &data.Comment, &data.Activation, &data.CreatedAt); err != nil {
			log.Printf("[warn] rows.Scan skip: %v", err)
			continue
		}
		res = append(res, &data)
	}
	return res, nil
}

func (rr *researchRepository) FindByID(id int) (*entity.Research, error) {
	row := rr.SQLHandler.QueryRow(`
		SELECT id, title, author, file, comments, activation
		FROM researches
		WHERE id=?
	`, id)
	var data entity.Research
	if err := row.Scan(&data.ID, &data.Title, &data.Author, &data.File, &data.Comment, &data.Activation); err != nil {
		err = errors.Wrap(err, "failed to bind data")
		return &data, err
	}
	return &data, nil
}

func (rr *researchRepository) Create(data *entity.Research) (int, error) {
	result, err := rr.SQLHandler.Execute(`
		INSERT INTO researches(title, author, file, comments, activation, created_at, updated_at)
		VALUES (?,?,?,?,?,?,?)
	`, data.Title, data.Author, data.File, data.Comment, data.Activation, data.CreatedAt, data.UpdatedAt)
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

func (rr *researchRepository) UpdateByID(data *entity.Research) error {
	_, err := rr.SQLHandler.Execute(`
		UPDATE researches
		SET title=?, author=?, file=?, comments=?, activation=?, updated_at=?
		WHERE id=?
	`, data.Title, data.Author, data.File, data.Comment, data.Activation, data.UpdatedAt, data.ID)
	if err != nil {
		err = errors.Wrap(err, "failed to execute query")
		return err
	}
	return nil

}

func (rr *researchRepository) DeleteByID(id int) error {
	_, err := rr.SQLHandler.Execute(`
		DELETE FROM researches
		WHERE id=?
	`, id)
	if err != nil {
		err = errors.Wrap(err, "failed to execute query")
	}
	return err
}
