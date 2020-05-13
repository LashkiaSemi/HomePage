package repository

import (
	"homepage/pkg/entity"
	"homepage/pkg/usecase/interactor"
	"log"

	"github.com/pkg/errors"
)

type tagRepository struct {
	SQLHandler
}

// NewTagRepository リポジトリの作成
func NewTagRepository(sh SQLHandler) interactor.TagRepository {
	return &tagRepository{
		SQLHandler: sh,
	}
}

func (tr *tagRepository) FindAll() ([]*entity.Tag, error) {
	rows, err := tr.SQLHandler.Query(`
		SELECT id, name
		FROM tags
	`)
	var res []*entity.Tag
	if err != nil {
		if err == tr.SQLHandler.ErrNoRows() {
			log.Printf("hit no data: %v", err)
			return res, nil
		}
		err = errors.Wrap(err, "failed to execute query")
		return res, err
	}
	for rows.Next() {
		var tag entity.Tag
		if err = rows.Scan(&tag.ID, &tag.Name); err != nil {
			log.Printf("rows.Scan skip: %v", err)
			continue
		}
		res = append(res, &tag)
	}
	return res, nil
}

func (tr *tagRepository) FindByID(id int) (*entity.Tag, error) {
	row := tr.SQLHandler.QueryRow(`
		SELECT id, name
		FROM tags
		WHERE id=?
	`, id)
	var tag entity.Tag
	if err := row.Scan(&tag.ID, &tag.Name); err != nil {
		err = errors.Wrap(err, "failed to bind data")
		return &tag, err
	}
	return &tag, nil
}

func (tr *tagRepository) Create(data *entity.Tag) (int, error) {
	result, err := tr.SQLHandler.Execute(`
		INSERT INTO tags(name, created_at, updated_at)
		VALUES (?,?,?)
	`, data.Name, data.CreatedAt, data.UpdatedAt)
	if err != nil {
		err = errors.Wrap(err, "failed to execute query")
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		err = errors.Wrap(err, "failed to get id")
		return 0, err
	}
	return int(id), err
}

func (tr *tagRepository) UpdateByID(data *entity.Tag) error {
	_, err := tr.SQLHandler.Execute(`
		UPDATE tags
		SET name=?, updated_at=?
		WHERE id=?
	`, data.Name, data.UpdatedAt, data.ID)
	if err != nil {
		err = errors.Wrap(err, "failed to execute query")
	}
	return err
}

func (tr *tagRepository) DeleteByID(id int) error {
	_, err := tr.SQLHandler.Execute(`
		DELETE FROM tags
		WHERE id=?
	`, id)
	if err != nil {
		err = errors.Wrap(err, "failed to execute query")
	}
	return err
}
