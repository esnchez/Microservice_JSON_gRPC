package main

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
)

//Logging service file
//Separated from service business logic


//this struct is going to hold a next object, which is going to be a PriceFetcher interface
type loggingService struct {
	next PriceFetcher
}

func NewLoggingService(a PriceFetcher) PriceFetcher {
	return &loggingService{
		next: a,
	}
}

//Same decorator pattern
// func NewLoggingService(a PriceFetcher) *loggingService {
// 	return &loggingService{a}
// }

func (s *loggingService) FetchPrice(ctx context.Context, ticker string) (price float64, err error) {
	defer func(begin time.Time) {
		logrus.WithFields(logrus.Fields{
			"requestID": ctx.Value("requestID"),
			"took": time.Since(begin),
			"err": err,
			"price":price,
		}).Info("fetchPrice")
	}(time.Now())

	return s.next.FetchPrice(ctx, ticker)

}
