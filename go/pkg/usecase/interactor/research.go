package interactor

import (
	"homepage/pkg/entity"

	"github.com/pkg/errors"
)

type researchInteractor struct {
	ResearchRepository
}

// ResearchInteractor 卒業研究のユースケースを実装
type ResearchInteractor interface {
	GetAll() ([]*entity.Research, error)
	GetByID(id int) (*entity.Research, error)

	Create(title, author, file, comment string, activation int) (int, error)
	UpdateByID(id int, title, author, file, comment string, activation int) error
	DeleteByID(id int) error
}

// NewResearchInteractor インタラクタを作成
func NewResearchInteractor(rr ResearchRepository) ResearchInteractor {
	return &researchInteractor{
		ResearchRepository: rr,
	}
}

func (ri *researchInteractor) GetAll() ([]*entity.Research, error) {
	return ri.ResearchRepository.FindAll()
}

func (ri *researchInteractor) GetByID(id int) (*entity.Research, error) {
	return ri.ResearchRepository.FindByID(id)

}

func (ri *researchInteractor) Create(title, author, file, comment string, activation int) (int, error) {
	// create obj
	data := entity.Research{}
	data.Create(title, author, file, comment, activation)

	// insert db
	id, err := ri.ResearchRepository.Create(&data)
	if err != nil {
		err = errors.Wrap(err, "failed to insert db")
		return 0, err
	}
	return id, nil
}

func (ri *researchInteractor) UpdateByID(id int, title, author, file, comment string, activation int) error {
	data, err := ri.ResearchRepository.FindByID(id)
	if err != nil {
		err = errors.Wrap(err, "failed to get original data")
		return err
	}
	newData := data.Update(title, author, file, comment, activation)

	// update db
	err = ri.ResearchRepository.UpdateByID(newData)
	if err != nil {
		err = errors.Wrap(err, "failed to update db")
		return err
	}
	return nil
}

func (ri *researchInteractor) DeleteByID(id int) error {
	return ri.ResearchRepository.DeleteByID(id)
}
