package service

type userService struct{}

// UserService ユーザのドメインロジックを実装
type UserService interface {
}

// NewUserService サービスを作成
func NewUserService() UserService {
	return &userService{}
}
