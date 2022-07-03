package company

import (
	"advertisement-api/Connection"
	Query "advertisement-api/Constants/company"
	"advertisement-api/Controller/Dto"
)

func (c company) UpdateCompanyBalance(params Dto.UpdatedCompanyBalance) (resp Dto.GetCompanyProfile, err error) {
	_, err = Connection.PostgresConnection.Exec(Query.UpdateCompanyBalance, params.NewBalance, params.Id)
	if err != nil {
		return
	}
	return
}

func (c company) GetCompanyProfile(id int) (resp Dto.GetCompanyProfile, err error) {
	err = Connection.PostgresConnection.QueryRow(Query.GetCompanyProfile, id).Scan(&resp.Id, &resp.Name, &resp.Balance)
	if err != nil {
		return
	}
	return
}

func (c company) CreateCompany(params Dto.CreateCompany) (id int, err error) {

	err = Connection.PostgresConnection.QueryRow(Query.CreateCompanyQuery, params.Name, params.Balance).Scan(&id) //esekusi return single row value
	if err != nil {
		return
	}
	return
}

func (c company) ViewAllCompany() (resp []Dto.ViewAllCompany, err error) {
	rows, err := Connection.PostgresConnection.Query(Query.ViewAllCompanies) //esekusi return single row value
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var data Dto.ViewAllCompany
		err = rows.Scan(&data.Id, &data.Name, &data.Balance)
		if err != nil {
			return
		}
		resp = append(resp, data)
	}
	return
}
