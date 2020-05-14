package repository

import (
	"homepage/pkg/entity"
	"homepage/pkg/usecase/interactor"
	"log"

	"github.com/pkg/errors"
)

type societyRepository struct {
	SQLHandler
}

// NewSocietyRepository リポジトリの作成
func NewSocietyRepository(sh SQLHandler) interactor.SocietyRepository {
	return &societyRepository{
		SQLHandler: sh,
	}
}

func (sr *societyRepository) FindAll() ([]*entity.Society, error) {
	rows, err := sr.SQLHandler.Query(`
		SELECT id, title, author, society, award, date
		FROM societies
		ORDER BY date DESC
	`)
	var datas []*entity.Society
	if err != nil {
		if err != sr.SQLHandler.ErrNoRows() {
			log.Printf("[warn] hit no data: %v", err)
			return datas, nil
		}
		err = errors.Wrap(err, "failed to execute query")
		return datas, err
	}
	for rows.Next() {
		var data entity.Society
		if err = rows.Scan(&data.ID, &data.Title, &data.Author, &data.Society, &data.Award, &data.Date); err != nil {
			log.Printf("[warn] rows.Scan skip: %v", err)
			continue
		}
		datas = append(datas, &data)
	}
	return datas, nil
}

func (sr *societyRepository) FindByID(id int) (*entity.Society, error) {
	row := sr.SQLHandler.QueryRow(`
		SELECT id, title, author, society, award, date
		FROM societies
		WHERE id=?
	`, id)
	var data entity.Society
	if err := row.Scan(&data.ID, &data.Title, &data.Author, &data.Society, &data.Award, &data.Date); err != nil {
		err = errors.Wrap(err, "failed to bind data")
		return &data, err
	}
	return &data, nil
}

func (sr *societyRepository) Create(data *entity.Society) (int, error) {
	result, err := sr.SQLHandler.Execute(`
		INSERT INTO societies(title, author, soiety, award, date, created_at, updated_at)
		VALUES (?,?,?,?,?,?,?)
	`, data.Title, data.Author, data.Society, data.Award, data.Date, data.CreatedAt, data.UpdatedAt)
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

func (sr *societyRepository) UpdateByID(data *entity.Society) error {
	_, err := sr.SQLHandler.Execute(`
		UPDATE societies
		SET title=?, author=?, society=?, award=?, date=?, updated_at=?
		WHERE id=?
	`, data.Title, data.Author, data.Society, data.Award, data.Date, data.UpdatedAt, data.ID)
	if err != nil {
		err = errors.Wrap(err, "failed to execute query")
	}
	return err
}

func (sr *societyRepository) DeleteByID(id int) error {
	_, err := sr.SQLHandler.Execute(`
		DELETE FROM societies
		WHERE id=?
	`, id)
	if err != nil {
		err = errors.Wrap(err, "failed to execute query")
	}
	return err
}
