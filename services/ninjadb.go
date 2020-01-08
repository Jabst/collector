package api

import (
	model "consumer/collector/models"
	"consumer/collector/object"
)

const limit = 5

func (api *API) EvaluateExchangeData(data model.ResultExchangeData, products object.Target) []model.ResultItemData {

	return data.ResultData[:limit]
}
