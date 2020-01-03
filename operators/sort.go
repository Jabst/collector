package operator

import (
	model "consumer/collector/models"
	"sort"
)

// SortListing sorts the listing by the ratios
func SortListing(listings *model.ResultExchangeData) {
	sort.SliceStable(listings.ResultData, func(a, b int) bool {
		return listings.ResultData[a].Listing.RatioValue < listings.ResultData[b].Listing.RatioValue
	})

	return
}
