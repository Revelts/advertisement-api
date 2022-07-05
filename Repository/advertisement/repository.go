package advertisement

import (
	"advertisement-api/Connection"
	Query "advertisement-api/Constants/advertisement"
	CompanyQuery "advertisement-api/Constants/company"
	"advertisement-api/Constants/transaction"
	"advertisement-api/Controller/Dto"
	"errors"
	"time"
)

func (a advertisement) CreateAdvertisement(params Dto.AdvertisementAttributes) (id int, err error) {
	err = Connection.PostgresConnection.QueryRow(Query.CreateAdvertisement, params.Name, params.Category, params.BasePrice, params.CreatedBy).Scan(&id)
	if err != nil {
		return
	}
	return
}

func (a advertisement) BuyAdvertisement(params Dto.BuyAdvertisement) (ads Dto.SuccessBuyAdvertisement, err error) {
	var adsAttr Dto.AdvertisementAttributes
	var companyAttr Dto.GetCompanyProfile
	var queryString string

	tx, err := Connection.PostgresConnection.Begin()
	if err != nil {
		return
	}
	err = tx.QueryRow(CompanyQuery.GetCompanyProfile, params.CompanyId).Scan(&companyAttr.Id, &companyAttr.Name, &companyAttr.Balance)
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			err = rollbackErr
		}
		return
	}

	err = tx.QueryRow(Query.GetAdvertisementInfo, params.AdvertisementId).Scan(&adsAttr.Id, &adsAttr.OwnedBy, &adsAttr.Name, &adsAttr.Category, &adsAttr.BasePrice, &adsAttr.CreatedBy)
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			err = rollbackErr
		}
		return
	}

	if companyAttr.Balance < adsAttr.BasePrice {
		err = errors.New("Not enough money!")
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			err = rollbackErr
		}
		return
	}

	if adsAttr.OwnedBy == 0 {
		_, err = tx.Exec(Query.BuyAdvertisement, params.CompanyId, params.AdvertisementId)
		if err != nil {
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				err = rollbackErr
			}
			return
		}
	}

	queryString = `SELECT start_advertising, end_advertising FROM public.transaction
						WHERE advertisement_id = $1`
	rows, err := tx.Query(queryString, adsAttr.Id)
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			err = rollbackErr
		}
		return
	}
	defer rows.Close()

	var startNewAdvertising time.Time
	var trxStatus int
	var transactionTimeStamp []Dto.TransactionTimeStamp
	for rows.Next() {
		var data Dto.TransactionTimeStamp
		err = rows.Scan(&data.StartAdvertisement, &data.EndAdvertising)
		if err != nil {
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				err = rollbackErr
			}
			return
		}
		transactionTimeStamp = append(transactionTimeStamp, data)
	}
	for index := 0; index < len(transactionTimeStamp); index++ {
		today := time.Now()
		startNewAdvertising = transactionTimeStamp[index].EndAdvertising
		sameDay := startNewAdvertising.Add(-24 * time.Hour)
		if today == sameDay {
			trxStatus = 1
		} else {
			trxStatus = 0
		}
	}

	if len(transactionTimeStamp) == 0 {
		startNewAdvertising = time.Now()
		trxStatus = 1
	}
	var endNewAdvertising = startNewAdvertising.AddDate(0, 0, params.DayDuration)

	var transactionData Dto.Transaction
	queryString = `INSERT INTO public.transaction 
	(advertisement_id, company_id, start_advertising, end_advertising, advertisement_location, fee) 
		VALUES ($1, $2, $3, $4, $5, $6) RETURNING trx_id, advertisement_id, company_id, TO_CHAR(start_advertising,'YYYY-MM-DD HH24:MI:SS') , TO_CHAR(end_advertising, 'YYYY-MM-DD HH24:MI:SS'), advertisement_location, fee`
	err = tx.QueryRow(queryString, params.AdvertisementId, params.CompanyId, startNewAdvertising, endNewAdvertising, params.Location, adsAttr.BasePrice).Scan(&transactionData.TrxId, &transactionData.AdvertisementId, &transactionData.CompanyId, &transactionData.StartAdvertising, &transactionData.EndAdvertising, &transactionData.AdvertisementLocation, &transactionData.AdvertisementFee)

	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			err = rollbackErr
		}
		return
	}

	var transactionStatus Dto.TransactionStatus
	var statusTransaction int
	queryString = `INSERT INTO public.transaction_status 
	(trx_id, status) 
		VALUES ($1, $2) RETURNING trx_id, status, TO_CHAR(created_at,'YYYY-MM-DD HH24:MI:SS')`
	err = tx.QueryRow(queryString, transactionData.TrxId, trxStatus).Scan(&transactionStatus.TrxId, &statusTransaction, &transactionStatus.CreatedAt)

	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			err = rollbackErr
		}
		return
	}

	ads.TransactionInfo = transactionData
	ads.TransactionInfoStatus.TrxId = transactionStatus.TrxId
	ads.TransactionInfoStatus.CreatedAt = transactionStatus.CreatedAt
	ads.TransactionInfoStatus.Status = transaction.TransactionType(statusTransaction).Transaction()

	tx.Commit()
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
