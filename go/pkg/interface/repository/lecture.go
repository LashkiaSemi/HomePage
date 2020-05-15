package repository

import (
	"homepage/pkg/entity"
	"homepage/pkg/usecase/interactor"
	"log"

	"github.com/pkg/errors"
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
		SELECT l.id, l.title, l.file, l.comments, l.activation, DATE_FORMAT(l.created_at, '%Y/%m/%d'), users.name, users.student_id
		FROM lectures as l
		INNER JOIN users
		ON user_id = users.id;
	`)
	lectures := []*entity.Lecture{}
	if err != nil {
		if err == lr.SQLHandler.ErrNoRows() {
			log.Printf("[warn] hit no data: %v", err)
			return lectures, nil
		}
		err = errors.Wrap(err, "failed to execute query")
		return lectures, err
	}
	for rows.Next() {
		var lecture entity.Lecture
		var user entity.User
		if err = rows.Scan(&lecture.ID, &lecture.Title, &lecture.File, &lecture.Comment, &lecture.Activation, &lecture.CreatedAt, &user.Name, &user.StudentID); err != nil {
			log.Printf("[warn] rows.Scan skip: %v", err)
			continue
		}
		lecture.Author = &user
		lectures = append(lectures, &lecture)
	}

	return lectures, nil
}

func (lr *lectureRepository) FindByID(id int) (*entity.Lecture, error) {
	row := lr.SQLHandler.QueryRow(`
		SELECT l.id, l.title, l.file, l.comments, l.activation, DATE_FORMAT(l.created_at, '%Y/%m/%d'), users.id, users.name, users.student_id
		FROM lectures as l
		INNER JOIN users
		ON user_id = users.id
		WHERE l.id = ?
	`, id)
	var lec entity.Lecture
	var user entity.User
	if err := row.Scan(&lec.ID, &lec.Title, &lec.File, &lec.Comment, &lec.Activation, &lec.CreatedAt, &user.ID, &user.Name, &user.StudentID); err != nil {
		err = errors.Wrap(err, "failed to bind data")
		return &entity.Lecture{}, err
	}
	lec.Author = &user
	return &lec, nil
}

func (lr *lectureRepository) FindAuthorByStudentID(studentID string) (*entity.User, error) {
	row := lr.SQLHandler.QueryRow(`
		SELECT users.id, users.name, users.student_id, intr.department, intr.grade, intr.comments
		FROM users
		INNER JOIN introductions as intr
		ON intr.user_id = users.id
		WHERE users.student_id = ?
	`, studentID)
	var user entity.User
	if err := row.Scan(&user.ID, &user.Name, &user.StudentID, &user.Department, &user.Grade, &user.Comment); err != nil {
		err = errors.Wrap(err, "failed to bind data")
		return &entity.User{}, err
	}
	return &user, nil
}

func (lr *lectureRepository) Create(lec *entity.Lecture) (int, error) {
	result, err := lr.SQLHandler.Execute(`
		INSERT INTO lectures(user_id, title, file, comments, activation, created_at, updated_at)
		VALUES (?,?,?,?,?,?,?)
	`, lec.Author.ID, lec.Title, lec.File, lec.Comment, lec.Activation, lec.CreatedAt, lec.UpdatedAt)
	if err != nil {
		err = errors.Wrap(err, "failed to execute query")
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		err = errors.Wrap(err, "failed to get id")
		return 0, err
	}
	return int(id), err
}

func (lr *lectureRepository) UpdateByID(lec *entity.Lecture) error {
	_, err := lr.SQLHandler.Execute(`
		UPDATE lectures
		SET title=?, file=?, comments=?, activation=?, user_id=?
		WHERE id=?
	`, lec.Title, lec.File, lec.Comment, lec.Activation, lec.Author.ID, lec.ID)
	if err != nil {
		err = errors.Wrap(err, "failed to execute query")
		return err
	}
	return nil
}

func (lr *lectureRepository) DeleteByID(id int) error {
	_, err := lr.SQLHandler.Execute(`
		DELETE FROM lectures WHERE id = ?
	`, id)
	if err != nil {
		err = errors.Wrap(err, "failed to execute query")
		return err
	}
	return nil
}
