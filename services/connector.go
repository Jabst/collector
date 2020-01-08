package api

import (
	model "consumer/collector/models"
)

func (api *API) ConnectorPostExchange(params [][]interface{}) (interface{}, error) {

	var x [][]string = make([][]string, len(params))

	for idx, elem := range params {
		x[idx] = make([]string, 0)
		for _, e := range elem {
			x[idx] = append(x[idx], e.(string))
		}
	}

	// TODO adicionar verificaçoes
	resultExchangeData, err := api.PostExchange(x[0], x[1])
	if err != nil {
		return model.ResultExchangeData{}, err
	}

	return resultExchangeData, nil

}
