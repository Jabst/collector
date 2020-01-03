package main

import (
	"consumer/collector/object"
	api "consumer/collector/services"
	"net/http"
)

func main() {

	ninjaDB, err := object.InitializeInMemory()
	if err != nil {
		panic(err)
	}

	poeAPI, err := object.InitializePoeAPI()
	if err != nil {
		panic(err)
	}

	poeCurrencyDictionary, err := object.InitializePoeCurrency()
	if err != nil {
		panic(err)
	}

	accountInfo, err := object.InitializeAccounts()
	if err != nil {
		panic(err)
	}

	API := api.NewAPI(accountInfo, poeAPI, ninjaDB, poeCurrencyDictionary)

	/*

		err = API.DumpData()
		if err != nil {
			panic(err)
		}*/

	http.HandleFunc("/exchange", API.GetExchange)
	http.ListenAndServe(":8080", nil)

	return
}
