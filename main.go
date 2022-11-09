package main

import (
	// "context"
	"flag"
	// "fmt"
	// "log"
)

func main() {
	listenAddress := flag.String("listenAddress", ":3000", "listen address the service is running")
	flag.Parse()

	pf := priceFetcher{} 
	service := NewLoggingService(NewMetricsService(&pf))

	server := NewJSONAPIService( *listenAddress , service)
	server.Run()

	// price, err := service.FetchPrice(context.Background(), "GG")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(price)
}