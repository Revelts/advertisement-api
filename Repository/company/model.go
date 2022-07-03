package company

import "advertisement-api/Controller/Dto"

type CompanyInterface interface {
	CreateCompany(params Dto.CreateCompany) (id int, err error)
	ViewAllCompany() (resp []Dto.ViewAllCompany, err error)
	GetCompanyProfile(id int) (resp Dto.GetCompanyProfile, err error)
	UpdateCompanyBalance(params Dto.UpdatedCompanyBalance) (resp Dto.GetCompanyProfile, err error)
}

type company struct{}

func InitCompany() CompanyInterface {
	return &company{}
}
