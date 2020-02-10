package repository

import (
	"errors"
	"homepage/pkg/domain"
	"homepage/pkg/domain/logger"
	"homepage/pkg/usecase/interactor"
	"time"
)

type lectureRepository struct {
	SQLHandler
}

// NewLectureHandler ハンドラの作成
func NewLectureHandler(sh SQLHandler) interactor.LectureRepository {
	return &lectureRepository{
		SQLHandler: sh,
	}
}

func (lr *lectureRepository) FindAll() (lecs domain.Lectures, err error) {
	rows, err := lr.SQLHandler.Query(
		`SELECT lec.id, lec.title, lec.file, lec.comments, lec.created_at, lec.updated_at, users.id, users.name, users.role, users.student_id, intr.comments, intr.department, intr.grade 
		FROM lectures as lec
		LEFT OUTER JOIN users
		ON lec.user_id = users.id
		INNER JOIN introductions as intr
		ON users.id = intr.user_id`,
	)
	if err != nil {
		return lecs, domain.InternalServerError(err)
	}

	for rows.Next() {
		var lec domain.Lecture
		if err = rows.Scan(&lec.ID, &lec.Title, &lec.File, &lec.Comment, &lec.CreatedAt, &lec.UpdatedAt, &lec.User.ID, &lec.User.Name, &lec.User.Role, &lec.User.StudentID, &lec.User.Comment, &lec.User.Department, &lec.User.Grade); err != nil {
			logger.Warn("lecture findAll: skip data")
			continue
		}
		lecs = append(lecs, lec)
	}
	return
}

func (lr *lectureRepository) FindByID(lecID int) (lec domain.Lecture, err error) {
	row := lr.SQLHandler.QueryRow(
		`SELECT lec.id, lec.title, lec.file, lec.comments, lec.created_at, lec.updated_at, users.id, users.name, users.role, users.student_id, intr.comments, intr.department, intr.grade 
		FROM lectures as lec
		LEFT OUTER JOIN users
		ON lec.user_id = users.id
		INNER JOIN introductions as intr
		ON users.id = intr.user_id
		WHERE lec.id=?`,
		lecID,
	)
	if err = row.Scan(&lec.ID, &lec.Title, &lec.File, &lec.Comment, &lec.CreatedAt, &lec.UpdatedAt, &lec.User.ID, &lec.User.Name, &lec.User.Role, &lec.User.StudentID, &lec.User.Comment, &lec.User.Department, &lec.User.Grade); err != nil {
		if err == lr.SQLHandler.ErrNoRows() {
			logger.Warn("lecture findByID: content not found")
			return lec, domain.NotFound(errors.New("content not found"))
		}
		logger.Error("lecture findByID: ", err)
		return lec, domain.InternalServerError(err)
	}
	return
}

func (lr *lectureRepository) Store(title, file, comment string, userID, isPublic int, createdAt time.Time) (int, error) {
	result, err := lr.SQLHandler.Execute(
		"INSERT INTO lectures(title, file, comments, user_id, activation, created_at, updated_at) VALUES (?,?,?,?,?,?,?)",
		title, file, comment, userID, isPublic, createdAt, createdAt,
	)
	if err != nil {
		return 0, err
	}
	id, _ := result.LastInsertId()
	return int(id), nil
}

func (lr *lectureRepository) Update(lecID int, title, file, comment string, userID, isPublic int, updatedAt time.Time) error {
	query, args, _ := makeUpdateQuery(
		"lectures",
		map[string]interface{}{
			"title":      title,
			"file":       file,
			"comments":   comment,
			"activation": isPublic,
			"user_id":    userID,
			"updated_at": updatedAt,
		},
		map[string]interface{}{
			"id": lecID,
		},
	)
	_, err := lr.SQLHandler.Execute(query, args...)
	return err
}

func (lr *lectureRepository) Delete(lecID int) error {
	_, err := lr.SQLHandler.Execute("DELETE FROM lectures WHERE id=?", lecID)
	return err
}
