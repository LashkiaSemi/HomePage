package service

type jobService struct{}

// JobService 就職先のドメインロジックを実装
type JobService interface {
}

// NewJobService サービスを作成
func NewJobService() JobService {
	return &jobService{}
}
