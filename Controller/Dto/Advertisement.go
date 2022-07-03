package Dto

type CreateAdvertisement struct {
	Name      string  `json:"name"`
	BasePrice float64 `json:"base_price"`
	Category  int     `json:"category"`
	CreatedBy int     `json:"created_by"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}

type SuccessCreateAdvertisement struct {
	Message               string                `json:"message"`
	UpdatedCompanyBalance UpdatedCompanyBalance `json:"company_info"`
	CreateAdvertisement   CreateAdvertisement   `json:"created_advertisement"`
}

type BuyAdvertisement struct {
	Id        int `json:"id"`
	CompanyId int `json:"company_id"`
}

type ViewAllAdvertisements struct {
	Id        int     `json:"id"`
	Name      string  `json:"name"`
	BasePrice float64 `json:"base_price"`
	Category  int     `json:"category"`
	OwnedBy   int     `json:"owned_by"`
	CreatedBy int     `json:"created_by"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}
