package main

import (
	"consumer/collector/object"
	api "consumer/collector/services"
	"log"
	"net/http"
)

func Printer(data []interface{}) {
	log.Println(data[0])
}

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

	watcher := api.InitWatcher(api.WatcherSettings{Active: true, Interval: 5, Tasks: []api.TasksBundle{}})

	API := api.NewAPI(accountInfo, poeAPI, ninjaDB, poeCurrencyDictionary, watcher)

	API.Watcher.StartWatcher()

	/*

		err = API.DumpData()
		if err != nil {
			panic(err)
		}*/

	http.HandleFunc("/exchange", API.GetExchange)
	go http.ListenAndServe(":8080", nil)

	return
}
