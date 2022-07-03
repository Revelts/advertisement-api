package Dto

type SuccessCreateAdvertisement struct {
	Message               string                  `json:"message"`
	UpdatedCompanyBalance UpdatedCompanyBalance   `json:"company_info"`
	CreateAdvertisement   AdvertisementAttributes `json:"created_advertisement"`
}

type SuccessBuyAdvertisement struct {
	Message           string                  `json:"message"`
	CompanyInfo       UpdatedCompanyBalance   `json:"company_info"`
	AdvertisementInfo AdvertisementAttributes `json:"advertisement_info"`
	//TransactionInfo   Models.Transaction      `json:"transaction_info"`
}

type BuyAdvertisement struct {
	AdvertisementId int `json:"advertisement_id"`
	CompanyId       int `json:"company_id"`
}

type RespBuyAdvertisement struct {
	Id      int    `json:"advertisement_id"`
	Message string `json:"message"`
}

type AdvertisementAttributes struct {
	Id        int     `json:"id"`
	Name      string  `json:"name"`
	BasePrice float64 `json:"base_price"`
	Category  int     `json:"category"`
	OwnedBy   int     `json:"owned_by"`
	CreatedBy int     `json:"created_by"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}
