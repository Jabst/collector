package api

import (
	model "consumer/collector/models"
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type TasksBundle struct {
	Signature        interface{}
	Params           [][]interface{}
	ReturnStructType string
}

type WatcherSettings struct {
	Active   bool
	Interval int
	Tasks    []TasksBundle
}

func InitWatcher(settings WatcherSettings) WatcherSettings {
	return WatcherSettings{
		Active:   settings.Active,
		Interval: settings.Interval,
	}
}

func (watcher *WatcherSettings) AddTask(task TasksBundle) {
	watcher.Tasks = append(watcher.Tasks, task)
}

func (watcher *WatcherSettings) UpdateWatcher(settings WatcherSettings) {
	watcher.Active = settings.Active
	watcher.Interval = settings.Interval
}

func (api *API) StartWatcher() {
	for api.Watcher.Active {
		log.Println(fmt.Sprintf("Starting watcher with %d seconds interval", api.Watcher.Interval))
		for index := range api.Watcher.Tasks {
			log.Println(fmt.Sprintf("Working on task #%d", index))
			retVal, _ := api.Watcher.Tasks[index].Signature.(func([][]interface{}) (interface{}, error))(api.Watcher.Tasks[index].Params)

			switch api.Watcher.Tasks[index].ReturnStructType {
			case "exchange":
				value := retVal.(model.ResultExchangeData)

				filteredData := api.EvaluateExchangeData(value, api.NinjaDB.Products)
				byteValueData, err := json.Marshal(filteredData)
				if err != nil {
					panic(err)
				}
				api.ShipData(byteValueData)
			}
		}
		log.Println(fmt.Sprintf("sleeping for: %d", api.Watcher.Interval))
		time.Sleep(time.Duration(api.Watcher.Interval) * time.Second)
	}
}
