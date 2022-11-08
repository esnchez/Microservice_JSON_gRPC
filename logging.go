package main

import (

	// "context"
	// "github.com/sirupsen/logrus"


)
//this struct is going to hold a next object, which is going to be a PriceFetcher interface
type loggingService struct {
	next PriceFetcher
}

// func (s *loggingService) FetchPrice(ctx context.Context, ticker string) (price float64, err error) {


// }
