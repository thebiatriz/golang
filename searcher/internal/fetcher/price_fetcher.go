package fetcher

import (
	"math/rand"
	"sync"
	"time"
)

// vai colocar o tipo de dado (float64) no canal
func FetchPrices(priceChannel chan<- float64) {
	var wg sync.WaitGroup
	wg.Add(4)

	go func() {
		defer wg.Done()
		priceChannel <- FetchPriceFromSite1()

	}()

	go func() {
		defer wg.Done()
		priceChannel <- FetchPriceFromSite2()

	}()

	go func() {
		defer wg.Done()
		priceChannel <- FetchPriceFromSite3()

	}()

	go func() {
		defer wg.Done()
		FetchAndSendMultiplePrices(priceChannel)
	}()

	go func() {
		wg.Wait()
		close(priceChannel)
	}()

}

// buscar preços de diferentes sites
func FetchPriceFromSite1() float64 {
	time.Sleep(1 * time.Second) //demora um segundo
	return rand.Float64() * 100 // retorna um valor aleatório
}

func FetchPriceFromSite2() float64 {
	time.Sleep(3 * time.Second) //demora três segundos
	return rand.Float64() * 100
}

func FetchPriceFromSite3() float64 {
	time.Sleep(2 * time.Second) //demora dois segundos
	return rand.Float64() * 100
}

func FetchAndSendMultiplePrices(priceChannel chan<- float64) {
	time.Sleep(6 * time.Second)

	prices := []float64{
		rand.Float64() * 100,
		rand.Float64() * 100,
		rand.Float64() * 100,
	}

	for _, price := range prices {
		priceChannel <- price
	}
}
