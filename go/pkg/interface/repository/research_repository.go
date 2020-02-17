package repository

import (
	"errors"
	"homepage/pkg/domain"
	"homepage/pkg/domain/logger"
	"homepage/pkg/usecase/interactor"
	"time"
)

type researchRepository struct {
	SQLHandler
}

// NewResearchRepository リポジトリを作成する
func NewResearchRepository(sh SQLHandler) interactor.ResearchRepository {
	return &researchRepository{
		SQLHandler: sh,
	}
}

func (rr *researchRepository) FindAll() (ress domain.Researches, err error) {
	rows, err := rr.SQLHandler.Query("SELECT id, title, author, file, comments, created_at FROM researches")
	if err != nil {
		return
	}
	for rows.Next() {
		var res domain.Research
		if err = rows.Scan(&res.ID, &res.Title, &res.Author, &res.File, &res.Comment, &res.CreatedAt); err != nil {
			continue
		}
		ress = append(ress, res)
	}
	return
}

func (rr *researchRepository) FindByID(resID int) (res domain.Research, err error) {
	row := rr.SQLHandler.QueryRow("SELECT id, title, author, file, comments, created_at FROM researches WHERE id=?", resID)
	if err = row.Scan(&res.ID, &res.Title, &res.Author, &res.File, &res.Comment, &res.CreatedAt); err != nil {
		if err == rr.SQLHandler.ErrNoRows() {
			logger.Warn("research findByID: no content")
			return res, domain.NotFound(errors.New("content not found"))
		}
		logger.Error("research findByID:", err)
		return res, domain.InternalServerError(err)
	}
	return
}

func (rr *researchRepository) Store(title, author, file, comment string, createdAt time.Time, isPublic int) (int, error) {
	result, err := rr.SQLHandler.Execute(
		"INSERT INTO researches(title, author, file, comments, created_at, updated_at, activation) VALUES (?,?,?,?,?,?)",
		title, author, file, comment, createdAt, createdAt, isPublic,
	)
	if err != nil {
		return 0, err
	}
	id, _ := result.LastInsertId()
	return int(id), nil
}

func (rr *researchRepository) Update(resID int, title, author, file, comment string, updatedAt time.Time, isPublic int) error {
	query, args, _ := makeUpdateQuery(
		"researches",
		map[string]interface{}{
			"title":      title,
			"author":     author,
			"file":       file,
			"comments":   comment,
			"updated_at": updatedAt,
			"activation": isPublic,
		},
		map[string]interface{}{
			"id": resID,
		},
	)
	_, err := rr.SQLHandler.Execute(query, args...)
	return err
}

func (rr *researchRepository) Delete(resID int) error {
	_, err := rr.SQLHandler.Execute("DELETE FROM researches WHERE id=?", resID)
	return err
}
