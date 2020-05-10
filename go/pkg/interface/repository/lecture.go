package repository

import (
	"homepage/pkg/domain/model"
	"homepage/pkg/usecase/interactor"
	"log"
)

type lectureRepository struct {
	SQLHandler
}

// NewLectureRepository リポジトリの作成
func NewLectureRepository(sh SQLHandler) interactor.LectureRepository {
	return &lectureRepository{
		SQLHandler: sh,
	}
}

func (lr *lectureRepository) FindAll() ([]*model.Lecture, error) {
	rows, err := lr.SQLHandler.Query(`
		SELECT l.id, l.title, l.file, l.comments, l.activation, l.created_at, users.name
		FROM lectures as l
		INNER JOIN users
		ON user_id = users.id;
	`)
	if err != nil {
		log.Println("lectureRepository: FindAll: ", err)
		return []*model.Lecture{}, err
	}
	lectures := []*model.Lecture{}
	for rows.Next() {
		var lecture model.Lecture
		var user model.User
		if err = rows.Scan(&lecture.ID, &lecture.Title, &lecture.File, &lecture.Comment, &lecture.Activation, &lecture.CreatedAt, &user.Name); err != nil {
			log.Println("lectureRepository: FindAll: ", err)
			continue
		}
		lecture.Author = &user
		lectures = append(lectures, &lecture)
	}

	return lectures, nil
}
