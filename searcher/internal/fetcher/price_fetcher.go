package fetcher

import (
	"math/rand"
	"sync"
	"time"
	"github.com/thebiatriz/golang/searcher/internal/models"
)

// vai colocar o tipo de dado (float64) no canal
func FetchPrices(priceChannel chan<- models.PriceDetails) {
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
func FetchPriceFromSite1() models.PriceDetails {
	time.Sleep(1 * time.Second) //demora um segundo
	return models.PriceDetails{
		StoreName: "Loja Lua",
		Value:     rand.Float64() * 100,
		Timestamp: time.Now(),
	}
}

func FetchPriceFromSite2() models.PriceDetails {
	time.Sleep(3 * time.Second) //demora três segundos
	return models.PriceDetails{
		StoreName: "Loja Mágica",
		Value:     rand.Float64() * 100,
		Timestamp: time.Now(),
	}
}

func FetchPriceFromSite3() models.PriceDetails {
	time.Sleep(2 * time.Second) //demora dois segundos
	return models.PriceDetails{
		StoreName: "Loja Chique",
		Value:     rand.Float64() * 100,
		Timestamp: time.Now(),
	}
}

func FetchAndSendMultiplePrices(priceChannel chan<- models.PriceDetails) {
	time.Sleep(6 * time.Second)

	prices := []float64{
		rand.Float64() * 100,
		rand.Float64() * 100,
		rand.Float64() * 100,
	}

	for _, price := range prices {
		priceChannel <- models.PriceDetails{
			StoreName: "Loja Tem-Tudo",
			Value:     price,
			Timestamp: time.Now(),
		}
	}
}
