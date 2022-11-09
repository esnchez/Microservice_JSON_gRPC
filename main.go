package main

import (
	"context"
	"fmt"
	"log"
)



func main() {
	pf := priceFetcher{} 
	service := NewLoggingService(&pf)
	price, err := service.FetchPrice(context.Background(), "GG")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(price)
}