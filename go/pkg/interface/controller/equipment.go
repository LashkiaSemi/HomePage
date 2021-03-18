package controller

import (
	"homepage/pkg/domain/entity"
	"homepage/pkg/usecase/interactor"
	"strconv"

	"github.com/pkg/errors"
)

type equipmentController struct {
	interactor.EquipmentInteractor
}

// EquipmentController 備品の入出力を変換
type EquipmentController interface {
	GetAll() (*EquipmentsResponse, error)
	GetByID(id int) (*EquipmentResponse, error)

	Create(name, comment string, stock, tagID int) (int, error)
	UpdateByID(id int, name, comment string, stock, tagID int) error

	DeleteByID(id int) error

	// admin
	AdminGetAll() ([]map[string]string, error)
	AdminGetByID(id int) (*FieldsResponse, error)
}

// NewEquipmentController コントローラの作成
func NewEquipmentController(ei interactor.EquipmentInteractor) EquipmentController {
	return &equipmentController{
		EquipmentInteractor: ei,
	}
}

func (ec *equipmentController) GetAll() (*EquipmentsResponse, error) {
	datas, err := ec.EquipmentInteractor.GetAll()
	if err != nil {
		err = errors.Wrap(err, "failed to original data for response")
		return &EquipmentsResponse{}, err
	}
	var res EquipmentsResponse
	for _, data := range datas {
		res.Equipments = append(res.Equipments, convertToEquipmentResponse(data))
	}
	return &res, nil
}

func (ec *equipmentController) GetByID(id int) (*EquipmentResponse, error) {
	data, err := ec.EquipmentInteractor.GetByID(id)
	if err != nil {
		err = errors.Wrap(err, "failed to original data for response")
		return &EquipmentResponse{}, err
	}
	return convertToEquipmentResponse(data), nil
}

func (ec *equipmentController) Create(name, comment string, stock, tagID int) (int, error) {
	return ec.EquipmentInteractor.Create(name, comment, stock, tagID)
}

func (ec *equipmentController) UpdateByID(id int, name, comment string, stock, tagID int) error {
	return ec.EquipmentInteractor.UpdateByID(id, name, comment, stock, tagID)
}

func (ec *equipmentController) DeleteByID(id int) error {
	return ec.EquipmentInteractor.DeleteByID(id)
}

// admin
func (ec *equipmentController) AdminGetAll() ([]map[string]string, error) {
	var res []map[string]string
	datas, err := ec.EquipmentInteractor.GetAll()
	if err != nil {
		err = errors.Wrap(err, "failed to original data for response")
		return res, err
	}
	for _, data := range datas {
		res = append(res, map[string]string{
			"id":    strconv.Itoa(data.ID),
			"title": data.Name,
		})
	}
	return res, nil
}

func (ec *equipmentController) AdminGetByID(id int) (*FieldsResponse, error) {
	var res FieldsResponse
	data, err := ec.EquipmentInteractor.GetByID(id)
	if err != nil {
		err = errors.Wrap(err, "failed to original data for response")
		return &res, err
	}
	res.Fields = append(res.Fields,
		&Field{Key: "ID", Value: data.ID},
		&Field{Key: "備品名", Value: data.Name},
		&Field{Key: "在庫", Value: data.Stock},
		&Field{Key: "コメント", Value: data.Comment},
		&Field{Key: "タグ", Value: data.Tag.Name},
	)
	return &res, nil
}

// EquipmentsResponse 備品のレスポンス
type EquipmentsResponse struct {
	Equipments []*EquipmentResponse
}

// EquipmentResponse 備品のレスポンス
type EquipmentResponse struct {
	ID        int
	Name      string
	Stock     int
	Comment   string
	Tag       *TagResponse
	CreatedAt string
}

func convertToEquipmentResponse(data *entity.Equipment) *EquipmentResponse {
	return &EquipmentResponse{
		ID:        data.ID,
		Name:      data.Name,
		Stock:     data.Stock,
		Comment:   data.Comment,
		Tag:       convertToTagResponse(data.Tag),
		CreatedAt: data.CreatedAt,
	}
}
