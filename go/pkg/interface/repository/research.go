package repository

import (
	"homepage/pkg/entity"
	"homepage/pkg/usecase/interactor"
	"log"
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
	if err != nil {
		log.Println("researchRepository: FindAll: ", err)
		return []*entity.Research{}, err
	}
	var res []*entity.Research
	for rows.Next() {
		var data entity.Research
		if err = rows.Scan(&data.ID, &data.Title, &data.Author, &data.File, &data.Comment, &data.Activation, &data.CreatedAt); err != nil {
			log.Println("researchRepository: FindAll: ", err)
			continue
		}
		res = append(res, &data)
	}
	return res, nil
}
