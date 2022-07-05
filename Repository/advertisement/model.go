package advertisement

import "advertisement-api/Controller/Dto"

type AdvertisementInterface interface {
	CreateAdvertisement(params Dto.AdvertisementAttributes) (id int, err error)
	BuyAdvertisement(params Dto.BuyAdvertisement) (ads Dto.SuccessBuyAdvertisement, err error)
	ViewAllAdvertisements() (resp []Dto.AdvertisementAttributes, err error)
}

type advertisement struct{}

func InitAdvertisement() AdvertisementInterface {
	return &advertisement{}
}
