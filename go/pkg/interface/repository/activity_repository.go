package repository

import (
	"errors"
	"homepage/pkg/domain"
	"homepage/pkg/domain/logger"
	"homepage/pkg/usecase/interactor"
	"time"
)

type activityRepository struct {
	SQLHandler
}

func NewActivityRepository(sh SQLHandler) interactor.ActivityRepository {
	return &activityRepository{
		SQLHandler: sh,
	}
}

func (ar *activityRepository) FindActivities() (acts domain.Activities, err error) {
	rows, err := ar.SQLHandler.Query("SELECT id, date, activity, created_at, updated_at FROM activities")
	if err != nil {
		if err == ar.SQLHandler.ErrNoRows() {
			logger.Warn(err)
			return acts, domain.BadRequest(err)
		}
		logger.Error(err)
		return acts, err
	}

	for rows.Next() {
		var act domain.Activity
		if err = rows.Scan(&act.ID, &act.Date, &act.Activity, &act.CreatedAt, &act.UpdatedAt); err != nil {
			logger.Error("FindActivities: skip data.")
			continue
		}
		acts = append(acts, act)
	}
	return
}

func (ar *activityRepository) FindActivityByID(actID int) (act domain.Activity, err error) {
	row := ar.SQLHandler.QueryRow("SELECT id, date, activity, created_at, updated_at FROM activities WHERE id=?", actID)
	if err = row.Scan(&act.ID, &act.Date, &act.Activity, &act.CreatedAt, &act.UpdatedAt); err != nil {
		if err == ar.SQLHandler.ErrNoRows() {
			logger.Warn("FindActivityByID", err)
			return act, domain.NotFound(errors.New("FindActivityByID: Content not found"))
		}
		logger.Error("FindActivityByID", err)
		return act, domain.InternalServerError(err)
	}
	return
}

func (ar *activityRepository) StoreActivity(date time.Time, act string, createdAt time.Time) (int, error) {
	result, err := ar.SQLHandler.Execute(
		"INSERT INTO activities(date, activity, created_at, updated_at) VALUES (?,?,?,?)",
		date, act, createdAt, createdAt,
	)
	if err != nil {
		return 0, domain.InternalServerError(err)
	}
	id, _ := result.LastInsertId()
	return int(id), nil
}

func (ar *activityRepository) UpdateActivity(actID int, date time.Time, act string, updatedAt time.Time) error {
	query, args, _ := makeUpdateQuery(
		"activities",
		map[string]interface{}{
			"date":       date,
			"activity":   act,
			"updated_at": updatedAt,
		},
		map[string]interface{}{
			"id": actID,
		},
	)
	_, err := ar.SQLHandler.Execute(query, args...)
	return err
}

func (ar *activityRepository) DeleteActivity(actID int) error {
	_, err := ar.SQLHandler.Execute("DELETE FROM activities WHERE id=?", actID)
	return err
}
