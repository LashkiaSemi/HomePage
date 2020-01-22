package repository

import (
	"errors"
	"homepage/pkg/domain"
	"homepage/pkg/domain/logger"
	"homepage/pkg/usecase/interactor"
	"time"
)

type tagRepository struct {
	SQLHandler
}

// NewTagRepository リポジトリを作成
func NewTagRepository(sh SQLHandler) interactor.TagRepository {
	return &tagRepository{
		SQLHandler: sh,
	}
}

func (tr *tagRepository) FindAll() (tags domain.Tags, err error) {
	rows, err := tr.SQLHandler.Query("SELECT id, name, created_at, updated_at FROM tags")
	if err != nil {
		return
	}
	for rows.Next() {
		var tag domain.Tag
		if err = rows.Scan(&tag.ID, &tag.Name, &tag.CreatedAt, &tag.UpdatedAt); err != nil {
			logger.Warn("tag find all: skip data")
			continue
		}
		tags = append(tags, tag)
	}
	return
}

func (tr *tagRepository) FindByID(tagID int) (tag domain.Tag, err error) {
	row := tr.SQLHandler.QueryRow(
		"SELECT id, name, created_at, updated_at FROM tags WHERE id=?",
		tagID,
	)
	if err := row.Scan(&tag.ID, &tag.Name, &tag.CreatedAt, &tag.UpdatedAt); err != nil {
		if err == tr.SQLHandler.ErrNoRows() {
			logger.Warn("tag findByID: content not found")
			return tag, domain.NotFound(errors.New("content not found"))
		}
		logger.Error("tag findByID: ", err)
		return tag, domain.InternalServerError(err)
	}
	return
}

func (tr *tagRepository) Store(name string, createdAt time.Time) (int, error) {
	result, err := tr.SQLHandler.Execute(
		"INSERT INTO tags(name, created_at, updated_at) VALUES (?,?,?)",
		name, createdAt, createdAt,
	)
	if err != nil {
		return 0, err
	}
	id, _ := result.LastInsertId()
	return int(id), nil
}

func (tr *tagRepository) Update(tagID int, name string, updatedAt time.Time) error {
	query, args, _ := makeUpdateQuery(
		"tags",
		map[string]interface{}{
			"name":       name,
			"updated_at": updatedAt,
		},
		map[string]interface{}{
			"id": tagID,
		},
	)
	_, err := tr.SQLHandler.Execute(query, args...)
	return err
}

func (tr *tagRepository) Delete(tagID int) error {
	_, err := tr.SQLHandler.Execute("DELETE FROM tags WHERE id=?", tagID)
	return err
}
