package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

type watcherSettings struct {
	Status    bool          `json:"enable"`
	Interval  int           `json:"interval"`
	Params    []interface{} `json:"params"`
	Operation string        `json:"operation"`
}

// GetExchange fetches the data from poe api, processes it and sorts it
func (api *API) GetExchange(w http.ResponseWriter, r *http.Request) {

	want, ok := r.URL.Query()["want"]
	if !ok {
		w.Write([]byte("want is not valid"))
	}

	wantValue := strings.Split(want[0], ",")

	have, ok := r.URL.Query()["have"]
	if !ok {
		w.Write([]byte("have is not valid"))
	}

	haveValue := strings.Split(have[0], ",")

	data, err := api.PostExchange(wantValue, haveValue)
	if err != nil {
		w.Write([]byte(err.Error()))
		panic(err)
	}

	marshalledData, err := json.Marshal(data)

	w.Write(marshalledData)

}

func (_api *API) PostWatcher(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	var wSettings watcherSettings

	err := decoder.Decode(&wSettings)
	if err != nil {
		log.Println(err)
	}

	_api.Watcher.AddTask(TasksBundle{
		Signature:        _api.ConnectorPostExchange,
		Params:           []interface{}{wSettings.Params[0].([]string), wSettings.Params[1].([]string)},
		ReturnStructType: "exchange",
	})

	w.Write([]byte("Task Added"))

}
