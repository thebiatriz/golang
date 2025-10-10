package main

import (
	"fmt"
	"github.com/thebiatriz/golang/searcher/internal/fetcher"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	priceChannel := make(chan float64)
	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		priceChannel <- fetcher.FetchPriceFromSite1()

	}() //função anonima

	go func() {
		defer wg.Done()
		priceChannel <- fetcher.FetchPriceFromSite2()

	}()

	go func() {
		defer wg.Done()
		priceChannel <- fetcher.FetchPriceFromSite3()

	}()

	go func() {
		wg.Wait()
		close(priceChannel)
	}()

	var totalPrice float64
	var countPrices int

	for price := range priceChannel {
		totalPrice += price
		countPrices++
		avgPrice := totalPrice / float64(countPrices)
		fmt.Printf("Preço recebido: R$ %.2f | Preço médio até agora: R$ %.2f\n", price, avgPrice)
	}

	fmt.Printf("\nTempo total: %s", time.Since(start))
}
