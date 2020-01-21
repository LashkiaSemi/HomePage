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

func (er *employRepository) FindAll() (comps domain.Companies, err error) {
	rows, err := er.SQLHandler.Query("SELECT id, company FROM companies")
	for rows.Next() {
		var comp domain.Company
		if err = rows.Scan(&comp.ID, &comp.Company); err != nil {
			continue
		}
		comps = append(comps, comp)
	}
	return
}

func (er *employRepository) FindByID(compID int) (comp domain.Company, err error) {
	row := er.SQLHandler.QueryRow("SELECT id, company FROM companies WHERE id=?", compID)
	if row.Scan(&comp.ID, &comp.Company); err != nil {
		if err == er.SQLHandler.ErrNoRows() {
			logger.Warn("employ findByID: ", err)
			return comp, domain.NotFound(errors.New("content not found"))
		}
		logger.Error("employ findByID: ", err)
		return comp, domain.InternalServerError(err)
	}
	return
}

func (er *employRepository) Store(company string, createdAt time.Time) (int, error) {
	result, err := er.SQLHandler.Execute(
		"INSERT INTO companies(company, created_at, updated_at) VALUES (?,?,?)",
		company, createdAt, createdAt,
	)
	if err != nil {
		return 0, err
	}
	id, _ := result.LastInsertId()
	return int(id), err
}

func (er *employRepository) Update(compID int, company string, updatedAt time.Time) error {
	query, args, _ := makeUpdateQuery(
		"companies",
		map[string]interface{}{
			"company":    company,
			"updated_at": updatedAt,
		},
		map[string]interface{}{
			"id": compID,
		},
	)
	_, err := er.SQLHandler.Execute(query, args...)
	return err
}

func (er *employRepository) Delete(compID int) error {
	_, err := er.SQLHandler.Execute("DELETE FROM companies WHERE id=?", compID)
	return err
}
