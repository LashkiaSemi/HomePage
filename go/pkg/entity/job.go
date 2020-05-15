package entity

import (
	"homepage/pkg/configs"
	"time"
)

// Job 就職先
type Job struct {
	ID        int
	Company   string
	Job       string
	CreatedAt string
	UpdatedAt string
}

func (j *Job) Create(company, job string) {
	j.Company = company
	j.Job = job
	j.CreatedAt = time.Now().Format(configs.DateTimeFormat)
	j.UpdatedAt = j.CreatedAt
}

func (j Job) Update(company, job string) *Job {
	j.Company = company
	j.Job = job
	j.UpdatedAt = time.Now().Format(configs.DateTimeFormat)
	return &j
}
