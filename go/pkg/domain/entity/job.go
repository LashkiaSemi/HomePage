package entity

import (
	"homepage/pkg/helper"
)

// Job 就職先
type Job struct {
	ID        int
	Company   string
	Job       string
	CreatedAt string
	UpdatedAt string
}

func NewJob(company, job string) *Job {
	now := helper.FormattedDateTimeNow()
	return &Job{
		Company:   company,
		Job:       job,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func (j Job) Update(company, job string) *Job {
	j.Company = company
	j.Job = job
	j.UpdatedAt = helper.FormattedDateTimeNow()
	return &j
}
