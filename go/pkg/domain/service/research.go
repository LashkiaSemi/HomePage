package service

import (
	"homepage/pkg/domain/entity"
	"homepage/pkg/domain/repository"
)

type Research interface {
	GetAll() ([]*entity.Research, error)
	GetByID(id int) (*entity.Research, error)
	Create(title, author, file, comment string, activation int) (int, error)
	UpdateByID(id int, title, author, file, comment string, activation int) error
	DeleteByID(id int) error
}

type research struct {
	repo repository.ResearchRepository
}

func NewResearch(repo repository.ResearchRepository) Research {
	return &research{
		repo: repo,
	}
}

func (r *research) GetAll() ([]*entity.Research, error) {
	return r.repo.FindAll()
}

func (r *research) GetByID(id int) (*entity.Research, error) {
	return r.repo.FindByID(id)
}

func (r *research) Create(title, author, file, comment string, activation int) (int, error) {
	res := entity.NewResearch(title, author, file, comment, activation)
	return r.repo.Create(res)
}

func (r *research) UpdateByID(id int, title, author, file, comment string, activation int) error {
	res, err := r.repo.FindByID(id)
	if err != nil {
		return err
	}
	newRes := res.Update(title, author, file, comment, activation)
	return r.repo.UpdateByID(newRes)
}

func (r *research) DeleteByID(id int) error {
	return r.repo.DeleteByID(id)
}
