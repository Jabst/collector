package main

import (
	"consumer/collector/object"
	api "consumer/collector/services"
	"log"
	"net/http"
)

const secondsIn5Minute = 5

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

	discordBots, err := object.InitializeDiscordBots()
	if err != nil {
		panic(err)
	}

	watcher := api.InitWatcher(api.WatcherSettings{Active: true, Interval: secondsIn5Minute, Tasks: []api.TasksBundle{}})

	API := api.NewAPI(accountInfo, poeAPI, ninjaDB, poeCurrencyDictionary, discordBots, watcher)

	go API.StartWatcher()

	http.HandleFunc("/exchange", API.GetExchange)
	http.HandleFunc("/task", API.PostWatcher)
	http.ListenAndServe(":8080", nil)

	return
}
