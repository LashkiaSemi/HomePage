package controller

import (
	"homepage/pkg/usecase/interactor"
)

// EquipmentController コントローラ
type EquipmentController interface {
	ShowAll() (GetEquipmentsResponse, error)
	ShowByID(equID int) (GetEquipmentResponse, error)
	Create(req *UpdateEquipmentRequest) (GetEquipmentResponse, error)
	Update(equID int, req *UpdateEquipmentRequest) (GetEquipmentResponse, error)
	Delete(equID int) error
}

type equipmentController struct {
	EquipmentInteractor interactor.EquipmentInteractor
}

// NewEquipmentController コントローラの作成
func NewEquipmentController(ei interactor.EquipmentInteractor) EquipmentController {
	return &equipmentController{
		EquipmentInteractor: ei,
	}
}

func (ec *equipmentController) ShowAll() (res GetEquipmentsResponse, err error) {
	equs, err := ec.EquipmentInteractor.FetchAll()
	if err != nil {
		return
	}
	// TODO: tags
	for _, equ := range equs {
		res.Equipments = append(res.Equipments, GetEquipmentResponse{
			ID:    equ.ID,
			Name:  equ.Name,
			Stock: equ.Stock,
			Note:  equ.Note,
			// Tag:
		})
	}
	return
}

func (ec *equipmentController) ShowByID(equID int) (res GetEquipmentResponse, err error) {
	equ, err := ec.EquipmentInteractor.FetchByID(equID)
	if err != nil {
		return
	}
	res.ID = equ.ID
	res.Name = equ.Name
	res.Stock = equ.Stock
	res.Note = equ.Note
	// TODO: tag
	return
}

func (ec *equipmentController) Create(req *UpdateEquipmentRequest) (res GetEquipmentResponse, err error) {
	// TODO: バリデーション。あれば
	if req.TagID == 0 {
		req.TagID = 6
	}

	equ, err := ec.EquipmentInteractor.Add(req.Name, req.Note, req.Stock, req.TagID)
	if err != nil {
		return
	}
	res.ID = equ.ID
	res.Name = equ.Name
	res.Stock = equ.Stock
	res.Note = equ.Note
	// TODO: tag
	return
}

func (ec *equipmentController) Update(equID int, req *UpdateEquipmentRequest) (res GetEquipmentResponse, err error) {
	// TODO: バリデーション。あれば
	if req.TagID == 0 {
		req.TagID = 6
	}

	equ, err := ec.EquipmentInteractor.Update(equID, req.Name, req.Note, req.Stock, req.TagID)
	if err != nil {
		return
	}
	res.ID = equ.ID
	res.Name = equ.Name
	res.Stock = equ.Stock
	res.Note = equ.Note
	// TODO: tag
	return
}

func (ec *equipmentController) Delete(equID int) error {
	return ec.EquipmentInteractor.Delete(equID)
}

// GetEquipmentsResponse 複数データのレスポンス
type GetEquipmentsResponse struct {
	Equipments []GetEquipmentResponse `json:"equipments"`
}

// GetEquipmentResponse 一件取得
type GetEquipmentResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Stock int    `json:"stock"`
	Note  string `json:"note"`
	// TODO: tags
	// Tag GetTagResponse `json:"tag"`
}

// UpdateEquipmentRequest 作成、更新リクエスト
type UpdateEquipmentRequest struct {
	Name  string `json:"name"`
	Stock int    `json:"stock"`
	Note  string `json:"note"`
	TagID int    `json:"tag_id"`
}
