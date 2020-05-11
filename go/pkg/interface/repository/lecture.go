package repository

import (
	"homepage/pkg/entity"
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

func (lr *lectureRepository) FindAll() ([]*entity.Lecture, error) {
	rows, err := lr.SQLHandler.Query(`
		SELECT l.id, l.title, l.file, l.comments, l.activation, l.created_at, users.name, users.student_id
		FROM lectures as l
		INNER JOIN users
		ON user_id = users.id;
	`)
	if err != nil {
		log.Println("lectureRepository: FindAll: ", err)
		return []*entity.Lecture{}, err
	}
	lectures := []*entity.Lecture{}
	for rows.Next() {
		var lecture entity.Lecture
		var user entity.User
		if err = rows.Scan(&lecture.ID, &lecture.Title, &lecture.File, &lecture.Comment, &lecture.Activation, &lecture.CreatedAt, &user.Name, &user.StudentID); err != nil {
			log.Println("lectureRepository: FindAll: ", err)
			continue
		}
		lecture.Author = &user
		lectures = append(lectures, &lecture)
	}

	return lectures, nil
}
