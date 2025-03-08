package main

import (
	"fmt"
	"go_crud/api"
	"sync"
)

func main() {
	currency_list := []string{"BTC", "ETH", "XRP"}

	var wg sync.WaitGroup
	for _, currency := range currency_list {
		wg.Add(1)
		go func() {
			getCurrencyRate(currency)
			wg.Done()
		}()
	}
	wg.Wait()
}

func getCurrencyRate(currency string) {
	rate, err := api.GetRate(currency)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Rate for %v: %v\n", rate.Currency, rate.Price)
	}
}
