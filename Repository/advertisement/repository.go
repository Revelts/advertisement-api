package advertisement

import (
	"advertisement-api/Connection"
	Query "advertisement-api/Constants/advertisement"
	CompanyQuery "advertisement-api/Constants/company"
	"advertisement-api/Controller/Dto"
)

func (a advertisement) CreateAdvertisement(params Dto.AdvertisementAttributes) (id int, err error) {
	err = Connection.PostgresConnection.QueryRow(Query.CreateAdvertisement, params.Name, params.Category, params.BasePrice, params.CreatedBy).Scan(&id)
	if err != nil {
		return
	}
	return
}

func (a advertisement) BuyAdvertisement(params Dto.BuyAdvertisement) (id int, err error) {
	var adsAttr Dto.AdvertisementAttributes
	var companyAttr Dto.GetCompanyProfile

	err = Connection.PostgresConnection.QueryRow(CompanyQuery.GetCompanyProfile, params.CompanyId).Scan(&companyAttr.Id, &companyAttr.Name, &companyAttr.Balance)
	if err != nil {
		return
	}

	err = Connection.PostgresConnection.QueryRow(Query.GetAdvertisementInfo, params.AdvertisementId).Scan(&adsAttr.Id, &adsAttr.OwnedBy, &adsAttr.Name, &adsAttr.Category, &adsAttr.BasePrice, &adsAttr.CreatedBy)
	if err != nil {
		return
	}

	if companyAttr.Balance < adsAttr.BasePrice {
		err = nil
		return
	}

	err = Connection.PostgresConnection.QueryRow(Query.BuyAdvertisement, params.CompanyId, params.AdvertisementId).Scan(&id)
	if err != nil {
		return
	}
	return
}

func (a advertisement) ViewAllAdvertisements() (resp []Dto.AdvertisementAttributes, err error) {
	rows, err := Connection.PostgresConnection.Query(Query.ViewAllAdvertisements)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var data Dto.AdvertisementAttributes
		err = rows.Scan(&data.Id, &data.OwnedBy, &data.Name, &data.Category, &data.BasePrice, &data.CreatedBy, &data.CreatedAt, &data.UpdatedAt)
		if err != nil {
			return
		}
		resp = append(resp, data)
	}
	return
}
