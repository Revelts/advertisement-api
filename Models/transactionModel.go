package Models

//
//import (
//	"advertisement-api/Connection"
//	AdsQuery "advertisement-api/Constants/advertisement"
//	CompanyQuery "advertisement-api/Constants/company"
//	TrxQuery "advertisement-api/Constants/transaction"
//	"advertisement-api/Controller/Dto"
//	"context"
//)
//
//package Models
//
//type AdsTrx interface {
//	CreateTrx() (Dto.AdvertisementAttributes, error)
//}
//
//type Transaction struct {
//	TrxId int `json:"transaction_id"`
//	AdvertisementId int `json:"advertisement_id"`
//	CompanyId int `json:"company_id"`
//	StartAdvertising string `json:"start_advertising"`
//	EndAdvertising string `json:"end_advertising"`
//	AdvertisementLocation string `json:"advertisement_location"`
//	AdvertisementFee float64 `json:"advertisement_fee"`
//}
//
//func (t Transaction) CreateTrx() (Trx Dto.SuccessBuyAdvertisement, err error) {
//
//	var adsData Dto.AdvertisementAttributes
//	var companyData Dto.GetCompanyProfile
//	var trxData Transaction
//
//	ctx := context.Background()
//	tx, err := Connection.PostgresConnection.BeginTx(ctx, nil)
//	if err != nil {
//		return
//	}
//	err = tx.QueryRowContext(ctx, AdsQuery.GetAdvertisementInfo, t.AdvertisementId).Scan(&adsData.Id, &adsData.OwnedBy, &adsData.Name, &adsData.Category, &adsData.BasePrice)
//	if err != nil {
//		tx.Rollback()
//		return
//	}
//
//	err = tx.QueryRowContext(ctx, CompanyQuery.GetCompanyProfile, t.CompanyId).Scan(&companyData.Id, &companyData.Name, &companyData.Balance)
//	if err != nil {
//		tx.Rollback()
//		return
//	}
//
//	if companyData.Balance < adsData.BasePrice {
//		err = nil
//		tx.Rollback()
//		return
//	}
//
//	err = tx.QueryRowContext(ctx, AdsQuery.BuyAdvertisement, t.CompanyId, t.AdvertisementId).Scan(&adsData.Id, &adsData.UpdatedAt, &adsData.OwnedBy)
//	if err != nil {
//		return
//	}
//
//	err = tx.QueryRowContext(ctx, AdsQuery.BuyAdvertisement, t.CompanyId, t.AdvertisementId).Scan(&adsData.Id, &adsData.UpdatedAt, &adsData.OwnedBy)
//	if err != nil {
//		return
//	}
//
//	err = tx.QueryRowContext(ctx, TrxQuery.CreateTransaction, t.AdvertisementId, t.CompanyId, t.StartAdvertising, t.EndAdvertising, t.AdvertisementLocation, t.AdvertisementFee).Send(
//		)
//
//	// Transaction > CompanyInfo
//	Trx.CompanyInfo.Id = companyData.Id
//	Trx.CompanyInfo.Name = companyData.Name
//	Trx.CompanyInfo.OldBalance = companyData.Balance
//	Trx.CompanyInfo.NewBalance = companyData.Balance - adsData.BasePrice
//
//	Trx.AdvertisementInfo.Id = adsData.Id
//	Trx.AdvertisementInfo.BasePrice = adsData.BasePrice
//	Trx.AdvertisementInfo.CreatedBy = adsData.CreatedBy
//	Trx.AdvertisementInfo.CreatedAt = adsData.CreatedAt
//	Trx.AdvertisementInfo.UpdatedAt = adsData.UpdatedAt
//	Trx.AdvertisementInfo.Name = adsData.Name
//	Trx.AdvertisementInfo.OwnedBy = adsData.OwnedBy
//	Trx.AdvertisementInfo.Category = adsData.Category
//
//
//
//	return Dto.AdvertisementAttributes{}
//}
//
//func TransactionDirector(adsId int, companyId int, ) AdsTrx {
//	return &Transaction{
//		AdvertisementId: adsId,
//		CompanyId: companyId,
//	}
//}
//
