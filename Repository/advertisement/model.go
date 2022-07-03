package advertisement

import "advertisement-api/Controller/Dto"

type AdvertisementInterface interface {
	CreateAdvertisement(params Dto.CreateAdvertisement) (id int, err error)
	BuyAdvertisement(params Dto.CreateAdvertisement) (id int, err error)
	ViewAllAdvertisements() (resp []Dto.ViewAllAdvertisements, err error)
}

type advertisement struct{}

func InitAdvertisement() AdvertisementInterface {
	return &advertisement{}
}
