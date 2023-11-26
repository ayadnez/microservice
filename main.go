package main

import (
	"flag"
)

func main() {
	// client := client.New("http://localhost:3000")
	// price, err := client.FetchPrice(context.Background(), "ETH")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%v+\n", price)

	// return

	listenAddr := flag.String("listenaddr", ":3000", "listen address the service is running")
	flag.Parse()
	svc := NewloggingService(NewmetricsService(&priceFetcher{}))
	server := NewJSONAPISever(*listenAddr, svc)
	//price, err := svc.FetchPrice(context.Background(), "ETH")
	server.Run()

}
