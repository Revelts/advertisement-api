package advertisement

import (
	"advertisement-api/Connection"
	Query "advertisement-api/Constants/advertisement"
	"advertisement-api/Controller/Dto"
)

func (a advertisement) CreateAdvertisement(params Dto.CreateAdvertisement) (id int, err error) {
	err = Connection.PostgresConnection.QueryRow(Query.CreateAdvertisement, params.Name, params.Category, params.BasePrice, params.CreatedBy).Scan(&id) //esekusi return single row value
	if err != nil {
		return
	}
	return
}

func (a advertisement) BuyAdvertisement(params Dto.CreateAdvertisement) (id int, err error) {

	return
}

func (a advertisement) ViewAllAdvertisements() (resp []Dto.ViewAllAdvertisements, err error) {
	rows, err := Connection.PostgresConnection.Query(Query.ViewAllAdvertisements)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var data Dto.ViewAllAdvertisements
		err = rows.Scan(&data.Id, &data.OwnedBy, &data.Name, &data.Category, &data.BasePrice, &data.CreatedBy, &data.CreatedAt, &data.UpdatedAt)
		if err != nil {
			return
		}
		resp = append(resp, data)
	}
	return
}
