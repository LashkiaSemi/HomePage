package repository

import (
	// TODO: domainに依存

	"homepage/pkg/entity"
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

func (sr *societyRepository) FindAll() ([]*entity.Society, error) {
	rows, err := sr.SQLHandler.Query(`
		SELECT title, author, society, award, date
		FROM societies
		ORDER BY date DESC
	`)
	if err != nil {
		if err != sr.SQLHandler.ErrNoRows() {
			log.Println("sql error: ", err)
			return []*entity.Society{}, err
		}
	}
	var datas []*entity.Society
	for rows.Next() {
		var data entity.Society
		if err = rows.Scan(&data.Title, &data.Author, &data.Society, &data.Award, &data.Date); err != nil {
			log.Println(err)
			continue
		}
		datas = append(datas, &data)
	}
	return datas, nil
}
