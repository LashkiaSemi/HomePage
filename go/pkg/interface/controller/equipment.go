package controller

import (
	"homepage/pkg/entity"
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

	// admin
	AdminGetAll() ([]map[string]string, error)
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
		return &EquipmentsResponse{}, err
	}
	var res EquipmentsResponse
	for _, data := range datas {
		res.Equipments = append(res.Equipments, convertToEquipmentResponse(data))
	}
	return &res, nil
}

func (ec *equipmentController) AdminGetAll() ([]map[string]string, error) {
	var res []map[string]string
	datas, err := ec.EquipmentInteractor.GetAll()
	if err != nil {
		err = errors.Wrap(err, "AdminGetAll")
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
