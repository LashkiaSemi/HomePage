package service

type userService struct{}

type UserService interface {
}

func NewUserService() UserService {
	return &userService{}
}
