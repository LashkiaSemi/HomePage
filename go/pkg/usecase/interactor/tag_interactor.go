package interactor

import (
	"homepage/conf"
	"homepage/pkg/domain"
	"time"
)

// TagInteractor インタラクタ
type TagInteractor interface {
	FetchAll() (domain.Tags, error)
	FetchByID(tagID int) (domain.Tag, error)
	Add(name string) (domain.Tag, error)
	Update(tagID int, name string) (domain.Tag, error)
	Delete(tagID int) error
}

type tagInteractor struct {
	TagRepository
}

// NewTagInteractor インタラクタの作成
func NewTagInteractor(tr TagRepository) TagInteractor {
	return &tagInteractor{
		TagRepository: tr,
	}
}

func (ti *tagInteractor) FetchAll() (domain.Tags, error) {
	return ti.TagRepository.FindAll()
}

func (ti *tagInteractor) FetchByID(tagID int) (domain.Tag, error) {
	return ti.TagRepository.FindByID(tagID)
}

func (ti *tagInteractor) Add(name string) (tag domain.Tag, err error) {
	createdAt := time.Now()
	id, err := ti.TagRepository.Store(name, createdAt)
	if err != nil {
		return
	}
	tag.ID = id
	tag.Name = name
	tag.CreatedAt = createdAt.Format(conf.DatetimeFormat)
	tag.UpdatedAt = tag.CreatedAt
	return
}

func (ti *tagInteractor) Update(tagID int, name string) (tag domain.Tag, err error) {
	updatedAt := time.Now()
	err = ti.TagRepository.Update(tagID, name, updatedAt)
	if err != nil {
		return
	}
	return ti.TagRepository.FindByID(tagID)
}

func (ti *tagInteractor) Delete(tagID int) error {
	return ti.TagRepository.Delete(tagID)
}
