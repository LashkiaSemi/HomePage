package repository

import (
	"homepage/pkg/entity"
	"homepage/pkg/usecase/interactor"
	"log"

	"github.com/pkg/errors"
)

type activityRepository struct {
	SQLHandler
}

// NewActivityRepository リポジトリの作成
func NewActivityRepository(sh SQLHandler) interactor.ActivityRepository {
	return &activityRepository{
		SQLHandler: sh,
	}
}

func (ar *activityRepository) FindAll() ([]*entity.Activity, error) {
	rows, err := ar.SQLHandler.Query(`
		SELECT id, activity, show_date, date, annotation, is_important, is_notify
		FROM activities
		ORDER BY date DESC
	`)
	var acts []*entity.Activity
	if err != nil {
		if err == ar.SQLHandler.ErrNoRows() {
			log.Printf("[warn] not data hit: %v", err)
			return acts, nil
		}
		err = errors.Wrap(err, "failed to execute query")
		return acts, err
	}
	for rows.Next() {
		var act entity.Activity
		if err = rows.Scan(&act.ID, &act.Activity, &act.ShowDate, &act.Date, &act.Annotation, &act.IsImportant, &act.IsNotify); err != nil {
			log.Printf("[warn] rows.Scan skip: %v", err)
			continue
		}
		acts = append(acts, &act)
	}
	return acts, nil
}

func (ar *activityRepository) FindByID(id int) (*entity.Activity, error) {
	row := ar.SQLHandler.QueryRow(`
		SELECT id, activity, show_date, date, annotation, is_important, is_notify
		FROM activities
		WHERE id=?
	`, id)
	var data entity.Activity
	if err := row.Scan(&data.ID, &data.Activity, &data.ShowDate, &data.Date, &data.Annotation, &data.IsImportant, &data.IsNotify); err != nil {
		err = errors.Wrap(err, "failed to bind data")
		return &data, err
	}
	return &data, nil
}

// TODO: 修正？お知らせのためのやつだからなこれ。もしくは新しく作るか
// ロジック的には、dateが現在時刻の翌日より前の時刻、のデータを持ってくる感じ
func (ar *activityRepository) FindUpcoming() ([]*entity.Activity, error) {
	rows, err := ar.SQLHandler.Query(`
		SELECT id, activity, show_date, date, annotation, is_important, is_notify
		FROM activities
		WHERE date > DATE_SUB(CURRENT_DATE(), INTERVAL 1 DAY)
		ORDER BY date ASC
	`)
	var acts []*entity.Activity
	if err != nil {
		if err == ar.SQLHandler.ErrNoRows() {
			log.Printf("[warn] not data hit: %v", err)
			return acts, nil
		}
		err = errors.Wrap(err, "failed to execute query")
		return acts, err
	}
	for rows.Next() {
		var act entity.Activity
		if err = rows.Scan(&act.ID, &act.Activity, &act.ShowDate, &act.Date, &act.Annotation, &act.IsImportant, &act.IsNotify); err != nil {
			log.Printf("[warn] rows.Scan skip: %v", err)
			continue
		}
		acts = append(acts, &act)
	}
	return acts, nil
}

func (ar *activityRepository) FindByNotify() ([]*entity.Activity, error) {
	rows, err := ar.SQLHandler.Query(`
		SELECT id, activity, show_date, date, annotation, is_important, is_notify
		FROM activities
		WHERE is_notify = 1
		ORDER BY date ASC
	`)
	var datas []*entity.Activity
	if err != nil {
		if err == ar.SQLHandler.ErrNoRows() {
			log.Printf("[warn] not data hit: %v", err)
			return datas, nil
		}
		err = errors.Wrap(err, "failed to execute query")
		return datas, err
	}
	for rows.Next() {
		var data entity.Activity
		if err = rows.Scan(&data.ID, &data.Activity, &data.ShowDate, &data.Date, &data.Annotation, &data.IsImportant, &data.IsNotify); err != nil {
			log.Printf("[warn] rows.Scan skip: %v", err)
			continue
		}
		datas = append(datas, &data)
	}
	return datas, nil
}

func (ar *activityRepository) Create(data *entity.Activity) (int, error) {
	result, err := ar.SQLHandler.Execute(`
		INSERT INTO activities(show_date, date, activity, annotation, is_important, is_notify, created_at, updated_at)
		VALUES (?,?,?,?,?,?,?,?)
	`, data.ShowDate, data.Date, data.Activity, data.Annotation, data.IsImportant, data.IsNotify, data.CreatedAt, data.UpdatedAt)
	if err != nil {
		err = errors.Wrap(err, "failed to execute query")
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		err = errors.Wrap(err, "failed to id new data")
		return 0, err
	}
	return int(id), nil
}

func (ar *activityRepository) UpdateByID(data *entity.Activity) error {
	_, err := ar.SQLHandler.Execute(`
		UPDATE activities
		SET show_date=?, date=?, activity=?, annotation=?, is_important=?, is_notify=?, updated_at=?
		WHERE id=?
	`, data.ShowDate, data.Date, data.Activity, data.Annotation, data.IsImportant, data.IsNotify, data.UpdatedAt, data.ID)
	if err != nil {
		err = errors.Wrap(err, "failed to execute query")
		return err
	}
	return nil
}

func (ar *activityRepository) DeleteByID(id int) error {
	_, err := ar.SQLHandler.Execute(`
		DELETE FROM activities
		WHERE id=?
	`, id)
	if err != nil {
		err = errors.Wrap(err, "failed to execute query")
	}
	return err
}
