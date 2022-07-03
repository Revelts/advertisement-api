package Repository

import (
	"advertisement-api/Repository/advertisement"
	"advertisement-api/Repository/company"
)

type repository struct {
	Company       company.CompanyInterface
	Advertisement advertisement.AdvertisementInterface
}

var AllRepository = repository{
	Company:       company.InitCompany(),
	Advertisement: advertisement.InitAdvertisement(),
}
