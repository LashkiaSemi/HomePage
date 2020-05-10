package repository

import (
	// TODO: domainに依存

	"homepage/pkg/domain/model"
	"homepage/pkg/usecase/interactor"
	"log"
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

func (sr *societyRepository) FindAll() ([]*model.Society, error) {
	rows, err := sr.SQLHandler.Query(`
		SELECT title, author, society, award, date
		FROM societies
		ORDER BY date DESC
	`)
	if err != nil {
		if err != sr.SQLHandler.ErrNoRows() {
			log.Println("sql error: ", err)
			return []*model.Society{}, err
		}
	}
	var datas []*model.Society
	for rows.Next() {
		var data model.Society
		if err = rows.Scan(&data.Title, &data.Author, &data.Society, &data.Award, &data.Date); err != nil {
			log.Println(err)
			continue
		}
		datas = append(datas, &data)
	}
	return datas, nil
}
