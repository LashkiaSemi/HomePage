package service

type equipmentService struct {}

// EquipmentService 備品のドメインロジックを実装
type EquipmentService interface {}

// NewEquipmentService サービスの作成
func NewEquipmentService() EquipmentService {
	return &equipmentService{}
}