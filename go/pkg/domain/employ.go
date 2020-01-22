package domain

// Job 就職先企業
type Job struct {
	ID        int
	Company   string
	Job       string
	CreatedAt string
	UpdatedAt string
}

// Jobs 就職先企業
type Jobs []Job
