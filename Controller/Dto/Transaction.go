package Dto

import "time"

type Transaction struct {
	TrxId                 int     `json:"transaction_id"`
	AdvertisementId       int     `json:"advertisement_id"`
	CompanyId             int     `json:"company_id"`
	StartAdvertising      string  `json:"start_advertising"`
	EndAdvertising        string  `json:"end_advertising"`
	AdvertisementLocation string  `json:"advertisement_location"`
	AdvertisementFee      float64 `json:"advertisement_fee"`
}

type TransactionTimeStamp struct {
	StartAdvertisement time.Time `json:"start_advertisement"`
	EndAdvertising     time.Time `json:"end_advertising"`
}

type TransactionStatus struct {
	TrxId     int    `json:"transaction_id"`
	Status    string `json:"transaction_status"`
	CreatedAt string `json:"transaction_created_at"`
}
