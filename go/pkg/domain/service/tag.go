package service

import (
	"homepage/pkg/domain/entity"
	"homepage/pkg/domain/repository"
)

type Tag interface {
	GetAll() ([]*entity.Tag, error)
	GetByID(id int) (*entity.Tag, error)
	Create(name string) (int, error)
	UpdateByID(id int, name string) error
	DeleteByID(id int) error
}

type tag struct {
	repo repository.TagRepository
}

func NewTag(repo repository.TagRepository) Tag {
	return &tag{
		repo: repo,
	}
}

func (t *tag) GetAll() ([]*entity.Tag, error) {
	return t.repo.FindAll()
}

func (t *tag) GetByID(id int) (*entity.Tag, error) {
	return t.repo.FindByID(id)
}

func (t *tag) Create(name string) (int, error) {
	ta := entity.NewTag(name)
	id, err := t.repo.Create(ta)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (t *tag) UpdateByID(id int, name string) error {
	tag, err := t.repo.FindByID(id)
	if err != nil {
		return err
	}
	newTag := tag.Update(name)
	return t.repo.UpdateByID(newTag)
}

func (t *tag) DeleteByID(id int) error {
	return t.repo.DeleteByID(id)
}
