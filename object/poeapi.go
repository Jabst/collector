package object

import (
	"consumer/collector/utils"
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
)

type PoEAPIEndpoints struct {
	GetStash     string `json:"getStash"`
	GetExchange  string `json:"getExchange"`
	PostExchange string `json:"postExchange"`
}

type PoEAPICurrency struct {
	Currency []string `json:"currency"`
}

type PoEAPIGateway struct {
	GetStashURL     string
	GetExchangeURL  string
	PostExchangeURL string
}

type PoECurrency struct {
	Currency []string
}

const poeFilename = "./config/poe_endpoint.json"
const currencyPoeFilename = "./config/currency_dictionary.json"

// InitializePoeAPI initalizes the queriable API
func InitializePoeAPI() (PoEAPIGateway, error) {
	poeEndpoints, err := utils.OpenFile(poeFilename)
	if err != nil {
		return PoEAPIGateway{}, err
	}

	defer poeEndpoints.Close()

	poeResource, err := bootstrapPoeData(poeEndpoints)
	if err != nil {
		return PoEAPIGateway{}, err
	}

	return PoEAPIGateway{
		GetStashURL:     poeResource.GetStash,
		GetExchangeURL:  poeResource.GetExchange,
		PostExchangeURL: poeResource.PostExchange,
	}, nil
}

func InitializePoeCurrency() (PoECurrency, error) {
	currencyDictionary, err := utils.OpenFile(currencyPoeFilename)
	if err != nil {
		return PoECurrency{}, err
	}

	defer currencyDictionary.Close()

	poeCurrency, err := bootstrapCurrencyData(currencyDictionary)
	if err != nil {
		return PoECurrency{}, err
	}

	return PoECurrency{
		Currency: poeCurrency.Currency,
	}, nil

}

// PrepareURL formats the URL with provided account name, league, tab index and stash status
func (api *PoEAPIGateway) PrepareURL(accName string, league string, tabIndex, public string) string {

	var ret string

	ret = strings.Replace(api.GetStashURL, "$$ACCNAME$$", accName, -1)
	ret = strings.Replace(ret, "$$LEAGUE$$", league, -1)
	ret = strings.Replace(ret, "$$TABINDEX$$", tabIndex, -1)
	ret = strings.Replace(ret, "$$PUBLIC$$", public, -1)

	return ret
}

func (api *PoEAPIGateway) PrepareExchangeUrl(trades []string, requestID string) string {

	var ret string

	ret = strings.Replace(api.GetExchangeURL, "$$EXCHANGEID$$", strings.Join(trades, ","), -1)
	ret = strings.Replace(ret, "$$REQUESTID$$", requestID, -1)

	return ret
}

func bootstrapPoeData(fileDescriptor *os.File) (PoEAPIEndpoints, error) {
	var poeAPI PoEAPIEndpoints

	byteValue, err := ioutil.ReadAll(fileDescriptor)
	if err != nil {
		return PoEAPIEndpoints{}, nil
	}

	if err := json.Unmarshal(byteValue, &poeAPI); err != nil {
		return PoEAPIEndpoints{}, nil
	}

	return poeAPI, nil

}

func bootstrapCurrencyData(fileDescriptor *os.File) (PoEAPICurrency, error) {
	var poeAPI PoEAPICurrency

	byteValue, err := ioutil.ReadAll(fileDescriptor)
	if err != nil {
		return PoEAPICurrency{}, nil
	}

	if err := json.Unmarshal(byteValue, &poeAPI); err != nil {
		return PoEAPICurrency{}, nil
	}

	return poeAPI, nil

}
