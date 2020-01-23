package repository

import (
	"errors"
	"homepage/pkg/domain"
	"homepage/pkg/domain/logger"
	"homepage/pkg/usecase/interactor"
	"time"
)

type societyRepository struct {
	SQLHandler
}

// NewSocietyRepository リポジトリを作成する
func NewSocietyRepository(sh SQLHandler) interactor.SocietyRepository {
	return &societyRepository{
		SQLHandler: sh,
	}
}

func (sr *societyRepository) FindAll() (socs domain.Societies, err error) {
	rows, err := sr.SQLHandler.Query("SELECT id, title, author, society, award, date FROM societies")
	if err != nil {
		return
	}
	for rows.Next() {
		var soc domain.Society
		if err = rows.Scan(&soc.ID, &soc.Title, &soc.Author, &soc.Society, &soc.Award, &soc.Date); err != nil {
			continue
		}
		socs = append(socs, soc)
	}
	return
}

func (sr *societyRepository) FindByID(socID int) (soc domain.Society, err error) {
	row := sr.SQLHandler.QueryRow("SELECT id, title, author, society, award, date FROM societies WHERE id=?", socID)
	if err = row.Scan(&soc.ID, &soc.Title, &soc.Author, &soc.Society, &soc.Award, &soc.Date); err != nil {
		if err == sr.SQLHandler.ErrNoRows() {
			logger.Warn("findSocietyByID: ", err)
			return soc, domain.NotFound(errors.New("content not found"))
		}
		logger.Error("findSocietyByID: ", err)
		return soc, domain.InternalServerError(err)
	}
	return
}

func (sr *societyRepository) Store(title, author, society, award string, date, createdAt time.Time) (int, error) {
	result, err := sr.SQLHandler.Execute(
		"INSERT INTO societies(title, author, society, award, date, created_at, updated_at) VALUES (?,?,?,?,?,?,?)",
		title, author, society, award, date, createdAt, createdAt,
	)
	if err != nil {
		return 0, err
	}
	id, _ := result.LastInsertId()
	return int(id), nil
}

func (sr *societyRepository) Update(socID int, title, author, society, award string, date, updatedAt time.Time) error {
	query, args, _ := makeUpdateQuery(
		"societies",
		map[string]interface{}{
			"title":      title,
			"author":     author,
			"society":    society,
			"award":      award,
			"date":       date,
			"updated_at": updatedAt,
		},
		map[string]interface{}{
			"id": socID,
		},
	)
	_, err := sr.SQLHandler.Execute(query, args...)
	return err
}

func (sr *societyRepository) Delete(socID int) error {
	_, err := sr.SQLHandler.Execute("DELETE FROM societies WHERE id=?", socID)
	return err
}
