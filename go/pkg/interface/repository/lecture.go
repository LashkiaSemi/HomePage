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

func (lr *lectureRepository) FindByID(id int) (*entity.Lecture, error) {
	row := lr.SQLHandler.QueryRow(`
		SELECT l.id, l.title, l.file, l.comments, l.activation, l.created_at, users.id, users.name, users.student_id
		FROM lectures as l
		INNER JOIN users
		ON user_id = users.id
		WHERE l.id = ?
	`, id)

	var lec entity.Lecture
	var user entity.User
	if err := row.Scan(&lec.ID, &lec.Title, &lec.File, &lec.Comment, &lec.Activation, &lec.CreatedAt, &user.ID, &user.Name, &user.StudentID); err != nil {
		log.Println("lectureRepository: FindByID: ", err)
		return &entity.Lecture{}, err
	}
	lec.Author = &user
	return &lec, nil
}

func (lr *lectureRepository) UpdateByID(lec *entity.Lecture) error {
	_, err := lr.SQLHandler.Execute(`
		UPDATE lectures
		SET title=?, comments=?, activation=?
		WHERE id=?
	`, lec.Title, lec.Comment, lec.Activation, lec.ID)
	if err != nil {
		log.Println("lectureRepository: UpdateByID: ", err)
		return err
	}
	return nil
}
