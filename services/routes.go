package api

import (
	"encoding/json"
	"net/http"
	"strings"
)

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
