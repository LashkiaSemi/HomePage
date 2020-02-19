package repository

import (
	"errors"
	"homepage/pkg/domain"
	"homepage/pkg/domain/logger"
	"homepage/pkg/usecase/interactor"
	"time"
)

type equipmentRepository struct {
	SQLHandler
}

// NewEquipmentRepository リポジトリを作成する
func NewEquipmentRepository(sh SQLHandler) interactor.EquipmentRepository {
	return &equipmentRepository{
		SQLHandler: sh,
	}
}

func (er *equipmentRepository) FindAll() (equs domain.Equipments, err error) {
	rows, err := er.SQLHandler.Query(
		`SELECT equipments.id, equipments.name, num, note, equipments.created_at, equipments.updated_at, tags.id, tags.name
		FROM equipments
		LEFT OUTER JOIN tags
		ON tags.id = equipments.tag_id`)
	if err != nil {
		return
	}
	for rows.Next() {
		var equ domain.Equipment
		if err = rows.Scan(&equ.ID, &equ.Name, &equ.Stock, &equ.Note, &equ.CreatedAt, &equ.UpdatedAt, &equ.Tag.ID, &equ.Tag.Name); err != nil {
			logger.Warn("equipment findAll: skip data. equipment.name=", equ.Name)
			err = nil
			equs = append(equs, equ)
			continue
		}
		equs = append(equs, equ)
	}
	return
}

func (er *equipmentRepository) FindByID(equID int) (equ domain.Equipment, err error) {
	row := er.SQLHandler.QueryRow(
		`SELECT equipments.id, equipments.name, num, note, equipments.created_at, equipments.updated_at, tags.id, tags.name
		FROM equipments
		LEFT OUTER JOIN tags
		ON tags.id = equipments.tag_id
		WHERE equipments.id=?`,
		equID,
	)
	if err = row.Scan(&equ.ID, &equ.Name, &equ.Stock, &equ.Note, &equ.CreatedAt, &equ.UpdatedAt, &equ.Tag.ID, &equ.Tag.Name); err != nil {
		if err == er.SQLHandler.ErrNoRows() {
			logger.Warn("equipment findByID: no content")
			return equ, domain.NotFound(errors.New("content not found"))
		}
		logger.Error("equipment findByID: ", err)
		return equ, domain.InternalServerError(err)
	}
	return
}

func (er *equipmentRepository) Store(name, note string, stock, tagID int, createdAt time.Time) (int, error) {
	result, err := er.SQLHandler.Execute(
		"INSERT INTO equipments(name, num, note, tag_id, created_at, updated_at) VALUES (?,?,?,?,?,?)",
		name, stock, note, tagID, createdAt, createdAt,
	)
	if err != nil {
		return 0, err
	}
	id, _ := result.LastInsertId()
	return int(id), err
}

func (er *equipmentRepository) Update(equID int, name, note string, stock, tagID int, updatedAt time.Time) error {
	query, args, _ := makeUpdateQuery(
		"equipments",
		map[string]interface{}{
			"name":       name,
			"note":       note,
			"num":        stock,
			"tag_id":     tagID,
			"updated_at": updatedAt,
		},
		map[string]interface{}{
			"id": equID,
		},
	)
	_, err := er.SQLHandler.Execute(query, args...)
	return err
}

func (er *equipmentRepository) Delete(equID int) error {
	_, err := er.SQLHandler.Execute("DELETE FROM equipments WHERE id=?", equID)
	return err
}
