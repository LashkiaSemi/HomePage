package service

type researchService struct {
}

// ResearchService 卒業研究のドメインロジックを実装
type ResearchService interface {
}

// NewResearchService サービスの作成
func NewResearchService() ResearchService {
	return &researchService{}
}
