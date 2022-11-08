package main

//business logic file
//needs to be clean, very important to not use any JSON representations,  
//we will do it somewhere else


import (
	"context"
	"fmt"
)

//Declaring interface with the function we will call FetchPrice(), to fetch a price
//context is used along many tools, it is common practive to pass it as first argument: 
//example: cancelling after 1 sec,
//function will return float and error
type PriceFetcher interface {
	FetchPrice(context.Context, string) (float64,error)
}

//priceFether implements the PriceFetcher interface
type priceFetcher struct {}

func (s *priceFetcher) FetchPrice(ctx context.Context, ticker string) (float64, error) {
	return MockPriceFetcher(ctx, ticker)
}


//mimic an API call
var priceMocks = map[string] float64 {
	"BTC": 20_000.0,
	"ETH": 200,
	"GG": 100_000.0,
}

func MockPriceFetcher(ctx context.Context, ticker string) (float64, error) {
	price, ok := priceMocks[ticker]
	if !ok {
		return price, fmt.Errorf("the given ticker (%s) is not supported", ticker)
	}
	return price, nil
}