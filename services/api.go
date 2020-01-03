package api

import (
	model "consumer/collector/models"
	"consumer/collector/object"
	operator "consumer/collector/operators"
	"consumer/collector/utils"
	"encoding/json"
	"time"
)

const breakdownSliceQuantity int = 14

type API struct {
	Accounts               object.TrackAccounts
	PoeAPI                 object.PoEAPIGateway
	NinjaDB                object.NinjaDatabase
	PoeCurrency            object.PoECurrency
	CachedExchangeSearches []CachedExchangeSearch
}

type CachedExchangeSearch struct {
	Data      model.ResultExchangeData
	Timestamp time.Time
}

type ExchangeResponse struct {
	Result    []string `json:"result"`
	RequestID string   `json:"id"`
	Total     int64    `json:"total"`
}

func NewAPI(accounts object.TrackAccounts, poeAPI object.PoEAPIGateway, ninjaDB object.NinjaDatabase, poeCurrencyDictionary object.PoECurrency) API {
	return API{
		Accounts:               accounts,
		PoeAPI:                 poeAPI,
		NinjaDB:                ninjaDB,
		PoeCurrency:            poeCurrencyDictionary,
		CachedExchangeSearches: make([]CachedExchangeSearch, 0),
	}
}

// Gets account information that is stored in memory
func (api *API) getAccountInfo(accName string) object.TrackAccount {

	for _, acc := range api.Accounts.Accounts {
		if acc.AccountName == accName {
			return acc
		}
	}

	return object.TrackAccount{}
}

func (api *API) ProduceRatio(listings *model.ResultExchangeData) error {
	operator.CalculateRatio(listings)
	return nil
}

func (api *API) SortListings(listings *model.ResultExchangeData) error {
	operator.SortListing(listings)
	return nil
}

func (api *API) DumpData() error {
	byteValues, err := json.Marshal(api.CachedExchangeSearches)
	if err != nil {
		return err
	}

	fileDescriptor, err := utils.CreateFile("./dump.json")
	if err != nil {
		return err
	}

	defer fileDescriptor.Close()

	err = utils.WriteToFile(fileDescriptor, byteValues)
	if err != nil {
		return err
	}

	return nil
}

// PostExchange queries the bulk item api
func (api *API) PostExchange(want []string, have []string) (model.ResultExchangeData, error) {

	bodyReq := model.ExchangeBodyRequest{
		Exchange: model.Exchange{
			Status: model.Status{
				Option: "online",
			},
			Have: have,
			Want: want,
		},
	}

	var resp ExchangeResponse

	value, err := utils.DoPostHTTPRequest(api.PoeAPI.PostExchangeURL, bodyReq)
	if err != nil {
		return model.ResultExchangeData{}, err
	}

	if err := json.Unmarshal(value, &resp); err != nil {
		return model.ResultExchangeData{}, err
	}

	var ret model.ResultExchangeData = model.ResultExchangeData{
		ResultData: make([]model.ResultItemData, 0),
	}

	// breaks the big slice container into smaller slices in a 2D array
	preparedContainer := utils.BreakdownInto(resp.Result, breakdownSliceQuantity)

	for _, content := range preparedContainer {
		requestURL := api.PoeAPI.PrepareExchangeUrl(content, resp.RequestID)
		var resultListing model.ResultExchangeResponse

		tradeInfo, err := utils.DoHTTPRequest(requestURL)
		if err != nil {
			return model.ResultExchangeData{}, err
		}

		if err := json.Unmarshal(tradeInfo, &resultListing); err != nil {
			return model.ResultExchangeData{}, err
		}

		// appends each listing into the return variable
		for _, item := range resultListing.Result {
			ret.ResultData = append(ret.ResultData, model.ResultItemData{
				ID: item.ID,
				Listing: model.ResultListingData{
					Indexed:    item.Listing.Indexed,
					WhisperMsg: item.Listing.WhisperMsg,
					RatioValue: -1,
					Price: model.PriceListing{
						ExchangePriceListing: model.ExchangePriceListing{
							Currency: item.Listing.Price.ExchangePrice.Currency,
							Amount:   item.Listing.Price.ExchangePrice.Amount,
						},
						ExchangeItemListing: model.ExchangeItemListing{
							Currency: item.Listing.Price.ExchangeItem.Currency,
							Amount:   item.Listing.Price.ExchangeItem.Amount,
							Stock:    item.Listing.Price.ExchangeItem.Stock,
							ID:       item.Listing.Price.ExchangeItem.ID,
						},
					},
				},
			})
		}
	}

	api.ProduceRatio(&ret)
	api.SortListings(&ret)

	api.CachedExchangeSearches = append(api.CachedExchangeSearches, CachedExchangeSearch{
		Data:      ret,
		Timestamp: time.Now(),
	})

	return ret, nil
}
