package repository

import (
	"homepage/pkg/domain/entity"
	"homepage/pkg/usecase/interactor"
	"log"

	"github.com/pkg/errors"
)

type equipmentRepository struct {
	SQLHandler
}

// NewEquipmentRepository リポジトリの作成
func NewEquipmentRepository(sh SQLHandler) interactor.EquipmentRepository {
	return &equipmentRepository{
		SQLHandler: sh,
	}
}

func (er *equipmentRepository) FindAll() ([]*entity.Equipment, error) {
	rows, err := er.SQLHandler.Query(`
		SELECT e.id, e.name, e.num, e.note, tags.name, DATE_FORMAT(e.created_at, '%Y/%m/%d') as date
		FROM equipments as e
		JOIN tags ON tags.id = tag_id
		ORDER BY date DESC
	`)
	var res []*entity.Equipment
	if err != nil {
		if err == er.SQLHandler.ErrNoRows() {
			log.Printf("[warn] not data hit: %v", err)
			return res, nil
		}
		err = errors.Wrap(err, "failed to execute query")
		return res, err
	}
	for rows.Next() {
		var data entity.Equipment
		var tag entity.Tag
		if err = rows.Scan(&data.ID, &data.Name, &data.Stock, &data.Comment, &tag.Name, &data.CreatedAt); err != nil {
			log.Printf("[warn] rows.Scan skip: %v", err)
			continue
		}
		data.Tag = &tag
		res = append(res, &data)
	}
	return res, nil
}

func (er *equipmentRepository) FindByID(id int) (*entity.Equipment, error) {
	row := er.SQLHandler.QueryRow(`
		SELECT e.id, e.name, e.num, e.note, tags.id, tags.name
		FROM equipments as e
		JOIN tags ON tags.id = tag_id
		WHERE e.id = ?
	`, id)
	var data entity.Equipment
	var tag entity.Tag
	if err := row.Scan(&data.ID, &data.Name, &data.Stock, &data.Comment, &tag.ID, &tag.Name); err != nil {
		err = errors.Wrap(err, "failed to bind data")
		return &data, err
	}
	data.Tag = &tag
	return &data, nil
}

func (er *equipmentRepository) Create(data *entity.Equipment) (int, error) {
	result, err := er.SQLHandler.Execute(`
		INSERT INTO equipments(name, num, note, tag_id, created_at, updated_at)
		VALUES (?,?,?,?,?,?)
	`, data.Name, data.Stock, data.Comment, data.Tag.ID, data.CreatedAt, data.UpdatedAt)
	if err != nil {
		err = errors.Wrap(err, "failed to execute query")
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		err = errors.Wrap(err, "failed to get new id")
		return 0, err
	}
	return int(id), nil
}

func (er *equipmentRepository) UpdateByID(data *entity.Equipment) error {
	_, err := er.SQLHandler.Execute(`
		UPDATE equipments
		SET name=?, num=?, note=?, tag_id=?, updated_at=?
		WHERE id=?
	`, data.Name, data.Stock, data.Comment, data.Tag.ID, data.UpdatedAt, data.ID)
	if err != nil {
		err = errors.Wrap(err, "failed to execute query")
		return err
	}
	return nil
}

func (er *equipmentRepository) DeleteByID(id int) error {
	_, err := er.SQLHandler.Execute(`
		DELETE FROM equipments
		WHERE id=?
	`, id)
	if err != nil {
		err = errors.Wrap(err, "failed to execute query")
	}
	return err
}
