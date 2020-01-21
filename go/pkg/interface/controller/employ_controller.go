package controller

import (
	"homepage/pkg/usecase/interactor"
)

// EmployController コントローラ
type EmployController interface {
	ShowAll() (GetCompaniesResponse, error)
	ShowByID(compID int) (GetCompanyResponse, error)
	Create(req *UpdateCompanyRequest) (GetCompanyResponse, error)
	Update(compID int, req *UpdateCompanyRequest) (GetCompanyResponse, error)
	Delete(compID int) error
}

type employController struct {
	EmployInteractor interactor.EmployInteractor
}

// NewEmployController コントローラの作成
func NewEmployController(ei interactor.EmployInteractor) EmployController {
	return &employController{
		EmployInteractor: ei,
	}
}

func (ec *employController) ShowAll() (res GetCompaniesResponse, err error) {
	comps, err := ec.EmployInteractor.FetchAll()
	if err != nil {
		return
	}
	for _, comp := range comps {
		res.Companies = append(res.Companies, GetCompanyResponse{
			ID:      comp.ID,
			Company: comp.Company,
		})
	}
	return
}

// GetCompaniesResponse 企業
type GetCompaniesResponse struct {
	Companies []GetCompanyResponse `json:"companies"`
}

// GetCompanyResponse 企業
type GetCompanyResponse struct {
	ID      int    `json:"id"`
	Company string `json:"company"`
}

func (ec *employController) ShowByID(compID int) (res GetCompanyResponse, err error) {
	comp, err := ec.EmployInteractor.FetchByID(compID)
	if err != nil {
		return
	}
	res.ID = comp.ID
	res.Company = comp.Company
	return
}

func (ec *employController) Create(req *UpdateCompanyRequest) (res GetCompanyResponse, err error) {
	comp, err := ec.EmployInteractor.Add(req.Company)
	if err != nil {
		return
	}
	res.ID = comp.ID
	res.Company = comp.Company
	return
}

// UpdateCompanyRequest 新規と更新のリクエスト
type UpdateCompanyRequest struct {
	Company string `json:"company"`
}

func (ec *employController) Update(compID int, req *UpdateCompanyRequest) (res GetCompanyResponse, err error) {
	comp, err := ec.EmployInteractor.Update(compID, req.Company)
	if err != nil {
		return
	}
	res.ID = comp.ID
	res.Company = comp.Company
	return
}

func (ec *employController) Delete(compID int) error {
	return ec.EmployInteractor.Delete(compID)
}
