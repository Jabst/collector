package model

type ExchangeBodyRequest struct {
	Exchange Exchange `json:"exchange"`
}

type Exchange struct {
	Status Status   `json:"status"`
	Have   []string `json:"have"`
	Want   []string `json:"want"`
}

type Status struct {
	Option string `json:"option"`
}

// ResultExchangeResponse is the struct of the exchange response
type ResultExchangeResponse struct {
	Result []ResultItemResponse `json:"result"`
}

type ResultItemResponse struct {
	ID      string                `json:"id"`
	Listing ResultListingResponse `json:"listing"`
}

type ResultListingResponse struct {
	Indexed    string `json:"indexed"`
	WhisperMsg string `json:"whisper"`
	Price      Price  `json:"price"`
}

type Price struct {
	ExchangePrice ExchangePrice `json:"exchange"`
	ExchangeItem  ExchangeItem  `json:"item"`
}

type ExchangePrice struct {
	Currency string  `json:"currency"`
	Amount   float64 `json:"amount"`
}

type ExchangeItem struct {
	Currency string  `json:"currency"`
	Amount   float64 `json:"amount"`
	Stock    float64 `json:"stock"`
	ID       string  `json:"id"`
}

// ResultExchangeData is the internal model for the exchange data response
type ResultExchangeData struct {
	ResultData []ResultItemData
}

type ResultItemData struct {
	ID      string
	Listing ResultListingData
}

type ResultListingData struct {
	Indexed    string
	WhisperMsg string
	Price      PriceListing
	RatioValue float64
}

type PriceListing struct {
	ExchangePriceListing ExchangePriceListing
	ExchangeItemListing  ExchangeItemListing
}

type ExchangePriceListing struct {
	Currency string
	Amount   float64
}

type ExchangeItemListing struct {
	Currency string
	Amount   float64
	Stock    float64
	ID       string
}
