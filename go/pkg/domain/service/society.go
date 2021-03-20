package service

import (
	"homepage/pkg/domain/entity"
	"homepage/pkg/domain/repository"
)

type Society interface {
	GetAll() ([]*entity.Society, error)
	GetByID(id int) (*entity.Society, error)
	Create(title, author, society, award, date string) (int, error)
	UpdateByID(id int, title, author, society, award, date string) error
	DeleteByID(id int) error
}

type society struct {
	repo repository.SocietyRepository
}

func NewSociety(repo repository.SocietyRepository) Society {
	return &society{
		repo: repo,
	}
}

func (s *society) GetAll() ([]*entity.Society, error) {
	return s.repo.FindAll()
}

func (s *society) GetByID(id int) (*entity.Society, error) {
	return s.repo.FindByID(id)
}

func (s *society) Create(title, author, society, award, date string) (int, error) {
	soc := entity.NewSociety(title, author, society, award, date)
	return s.repo.Create(soc)
}

func (s *society) UpdateByID(id int, title, author, society, award, date string) error {
	soc, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	return s.repo.UpdateByID(soc)
}

func (s *society) DeleteByID(id int) error {
	return s.repo.DeleteByID(id)
}
