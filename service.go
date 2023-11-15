package main

import (
	"context"
	"fmt"
)

type PriceFetcher interface {
	FetchPrice(context.Context, string) (float64, error)
}

type priceFetcher struct{}

func (s *priceFetcher) FetchPrice(ctx context.Context, ticker string) (float64, error) {
	// write business logic here
	return MockpriceFetcher(ctx, ticker)
}

var priceMocks = map[string]float64{
	"BTC": 20_000.0,
	"ETH": 23_000.0,
	"GG":  25_0000.0,
}

func MockpriceFetcher(ctx context.Context, ticker string) (float64, error) {
	price, ok := priceMocks[ticker]
	if !ok {
		return price, fmt.Errorf("the given ticker (%s) is not supported", ticker)
	}
	return price, nil
}
