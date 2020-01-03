package operator

import (
	model "consumer/collector/models"
)

// CalculateRatio will calculate the ratio between the search/seller
func CalculateRatio(listings *model.ResultExchangeData) {
	for index, item := range listings.ResultData {

		var sellPrice float64 = float64(item.Listing.Price.ExchangePriceListing.Amount)
		var buyPrice float64 = float64(item.Listing.Price.ExchangeItemListing.Amount)

		listings.ResultData[index].Listing.RatioValue = buyPrice / sellPrice
	}

	return
}
