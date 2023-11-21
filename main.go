package main

import (
	"flag"
)

func main() {
	listenAddr := flag.String("listenaddr", ":3000", "listen address the service is running")
	flag.Parse()
	svc := NewloggingService(NewmetricsService(&priceFetcher{}))
	sever := NewJSONAPISever(*listenAddr, svc)
	//price, err := svc.FetchPrice(context.Background(), "ETH")
	sever.Run()

}
