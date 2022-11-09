package types

//used in api.go but also to build client.go, so we have to decouple it from api.go file

type PriceResponse struct {
	Price float64 `json:"price"`
	Ticker string `json:"ticker"`
}