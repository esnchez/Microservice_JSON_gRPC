package main

import (
	"flag"
	
	// "fmt"
	// "context"
	// "log"
	// "github.com/esnchez/client"
)

func main() {

	//testing the client, uncover the following code after booting up server, also imports
	
	// c := client.New("http://localhost:3000/")
	// price, err := c.FetchPrice(context.Background(), "ET")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("%+v\n", price)
	// return
	//

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