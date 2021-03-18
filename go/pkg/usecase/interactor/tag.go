package interactor

import (
	"homepage/pkg/domain/entity"

	"github.com/pkg/errors"
)

type tagInteractor struct {
	TagRepository
}

// TagInteractor タグ関連のユースケースを実現
type TagInteractor interface {
	GetAll() ([]*entity.Tag, error)
	GetByID(id int) (*entity.Tag, error)

	Create(name string) (int, error)
	UpdateByID(id int, name string) error

	DeleteByID(id int) error
}

// NewTagInteractor インタラクタの作成
func NewTagInteractor(tr TagRepository) TagInteractor {
	return &tagInteractor{
		TagRepository: tr,
	}
}

func (ti *tagInteractor) GetAll() ([]*entity.Tag, error) {
	return ti.TagRepository.FindAll()
}

func (ti *tagInteractor) GetByID(id int) (*entity.Tag, error) {
	return ti.TagRepository.FindByID(id)
}

func (ti *tagInteractor) Create(name string) (int, error) {
	tag := entity.Tag{}
	tag.Create(name)

	id, err := ti.TagRepository.Create(&tag)
	if err != nil {
		err = errors.Wrap(err, "failed to insert db")
	}
	return id, err
}

func (ti *tagInteractor) UpdateByID(id int, name string) error {
	tag, err := ti.TagRepository.FindByID(id)
	if err != nil {
		err = errors.Wrap(err, "failed to get original data")
		return err
	}
	newTag := tag.Update(name)

	err = ti.TagRepository.UpdateByID(newTag)
	if err != nil {
		err = errors.Wrap(err, "failed to update db")
	}
	return err
}

func (ti *tagInteractor) DeleteByID(id int) error {
	return ti.TagRepository.DeleteByID(id)
}
