package main

import (
	"context"
	"fmt"
)

type metricsService struct{
	next PriceFetcher
}

func NewMetricsService(a PriceFetcher) PriceFetcher{
	return &metricsService{
		next: a,
	}
}

func (s *metricsService) FetchPrice(ctx context.Context, ticker string) (price float64, err error) {
	fmt.Println("Pushing metrics to Prometheus...")	
	//metrics. Push to prometheus (gauge,counters)
	return s.next.FetchPrice(ctx, ticker)

}