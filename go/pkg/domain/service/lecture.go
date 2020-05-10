package service

type lectureService struct{}

// LectureService レクチャーのドメインロジックを実装
type LectureService interface {
}

// NewLectureService サービスを作成
func NewLectureService() LectureService {
	return &lectureService{}
}
