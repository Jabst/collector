package api

import (
	model "consumer/collector/models"
	"fmt"
	"log"
	"time"
)

type TasksBundle struct {
	Signature        interface{}
	Params           []interface{}
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

func (watcher *WatcherSettings) StartWatcher() {
	for watcher.Active {
		log.Println(fmt.Sprintf("Starting watcher with %d as interval", watcher.Interval))
		for index := range watcher.Tasks {
			retVal, _ := watcher.Tasks[index].Signature.(func([]interface{}) (interface{}, error))(watcher.Tasks[index].Params)

			switch watcher.Tasks[index].ReturnStructType {
			case "exchange":
				value := retVal.(model.ResultExchangeData)

				log.Println(value.ResultData[0].ID)

			}
			log.Println(watcher.Tasks[index].Signature)

		}
		log.Println(fmt.Sprintf("sleeping for: %d", watcher.Interval))
		time.Sleep(time.Duration(watcher.Interval) * time.Second)
	}
}
