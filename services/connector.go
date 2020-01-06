package api

import (
	model "consumer/collector/models"
	"log"
)

func (api *API) ConnectorPostExchange(params []interface{}) (interface{}, error) {

	log.Println(params)
	// TODO adicionar verificaçoes
	resultExchangeData, err := api.PostExchange(params[0].([]string), params[1].([]string))
	if err != nil {
		return model.ResultExchangeData{}, err
	}

	return resultExchangeData, nil

}
