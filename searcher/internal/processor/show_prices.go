package processor

import (
	"fmt"
	"github.com/thebiatriz/golang/searcher/internal/models"
)

func ShowPriceAVG(priceChannel <-chan models.PriceDetails, done chan<- bool) {
	var totalPrice float64
	var countPrices int

	for price := range priceChannel {
		totalPrice += price.Value
		countPrices++
		avgPrice := totalPrice / float64(countPrices)
		fmt.Printf("[%s] Preço recebido de %s | R$ %.2f | Preço médio até agora: %.2f\n", price.Timestamp.Format("02-01 15:04:05"), price.StoreName, price.Value, avgPrice)
	}

	done <- true
}
