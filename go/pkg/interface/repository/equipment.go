package repository

import (
	"homepage/pkg/domain/model"
	"homepage/pkg/usecase/interactor"
	"log"
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

func (er *equipmentRepository) FindAll() ([]*model.Equipment, error) {
	rows, err := er.SQLHandler.Query(`
		SELECT e.id, e.name, e.num, e.note, tags.name, e.created_at
		FROM equipments as e
		JOIN tags ON tags.id = tag_id
		ORDER BY created_at DESC
	`)
	if err != nil {
		log.Println("equipmentRepository: FindAll: ", err)
		return []*model.Equipment{}, err
	}
	var res []*model.Equipment
	for rows.Next() {
		var data model.Equipment
		var tag model.Tag
		if err = rows.Scan(&data.ID, &data.Name, &data.Stock, &data.Comment, &tag.Name, &data.CreatedAt); err != nil {
			log.Println("equipmentRepository: FindAll: ", err)
			continue
		}
		data.Tag = &tag
		res = append(res, &data)
	}
	return res, nil
}
