package Controller

import (
	AdsConstants "advertisement-api/Constants/advertisement"
	"advertisement-api/Controller/Dto"
	"advertisement-api/Library"
	"advertisement-api/Repository"
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateCompany(w http.ResponseWriter, r *http.Request) {
	var requestBody Dto.CreateCompany
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		Library.HttpResponseError(w, r, err.Error(), http.StatusBadRequest)
		return
	}
	repository := Repository.AllRepository
	id, err := repository.Company.CreateCompany(requestBody)
	if err != nil {
		Library.HttpResponseError(w, r, err.Error(), http.StatusInternalServerError)
		return
	}
	//butuh data sembako
	respSuccess := fmt.Sprintf("Success create a new company with id %d", id)
	Library.HttpResponseSuccess(w, r, respSuccess)
}

func ViewAllCompany(w http.ResponseWriter, r *http.Request) {
	repository := Repository.AllRepository.Company
	resp, err := repository.ViewAllCompany()
	if err != nil {
		Library.HttpResponseError(w, r, err.Error(), http.StatusInternalServerError)
		return
	}
	Library.HttpResponseSuccess(w, r, resp)
}

func ViewAllAdvertisement(w http.ResponseWriter, r *http.Request) {
	repository := Repository.AllRepository.Advertisement
	resp, err := repository.ViewAllAdvertisements()
	if err != nil {
		Library.HttpResponseError(w, r, err.Error(), http.StatusInternalServerError)
		return
	}
	Library.HttpResponseSuccess(w, r, resp)
}

func PlaceAdvertisement(w http.ResponseWriter, r *http.Request) {
	var requestBody Dto.CreateAdvertisement
	var err error
	err = json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		Library.HttpResponseError(w, r, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(requestBody)
	Company := Repository.AllRepository.Company
	Ads := Repository.AllRepository.Advertisement
	company, err := Company.GetCompanyProfile(requestBody.CreatedBy)
	if company.Balance < requestBody.BasePrice {
		Library.HttpResponseError(w, r, AdsConstants.FailToCreateAdvertisementNotEnoughMoney, http.StatusInternalServerError)
		return
	}

	var newCompanyBalance = company.Balance - requestBody.BasePrice
	var requestCompany Dto.UpdatedCompanyBalance

	requestCompany.OldBalance = company.Balance
	requestCompany.Name = company.Name
	requestCompany.NewBalance = newCompanyBalance
	requestCompany.Id = requestBody.CreatedBy

	_, err = Company.UpdateCompanyBalance(requestCompany)

	if err != nil {
		Library.HttpResponseError(w, r, err.Error(), http.StatusInternalServerError)
		return
	}

	advertisementId, err := Ads.CreateAdvertisement(requestBody)

	if err != nil {
		Library.HttpResponseError(w, r, err.Error(), http.StatusInternalServerError)
		return
	}

	var SendResponse Dto.SuccessCreateAdvertisement
	SendResponse.Message = fmt.Sprintf("Success creating new advertisement with id %d", advertisementId)
	SendResponse.CreateAdvertisement.CreatedBy = requestBody.CreatedBy
	SendResponse.CreateAdvertisement.Name = requestBody.Name
	SendResponse.CreateAdvertisement.BasePrice = requestBody.BasePrice
	SendResponse.CreateAdvertisement.Category = requestBody.Category

	SendResponse.UpdatedCompanyBalance.NewBalance = requestCompany.NewBalance
	SendResponse.UpdatedCompanyBalance.OldBalance = requestCompany.OldBalance
	SendResponse.UpdatedCompanyBalance.Name = requestCompany.Name
	SendResponse.UpdatedCompanyBalance.Id = requestCompany.Id

	if err != nil {
		Library.HttpResponseError(w, r, err.Error(), http.StatusInternalServerError)
		return
	}

	if err != nil {
		Library.HttpResponseError(w, r, err.Error(), http.StatusInternalServerError)
		return
	}
	Library.HttpResponseSuccess(w, r, SendResponse)
}
