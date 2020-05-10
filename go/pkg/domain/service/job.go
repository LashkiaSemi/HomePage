package service

type jobService struct{}

type JobService interface {
}

func NewJobService() JobService {
	return &jobService{}
}
