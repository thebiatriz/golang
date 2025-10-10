package processor

import "fmt"

func ShowPriceAVG(priceChannel <-chan float64, done chan<- bool) {
	var totalPrice float64
	var countPrices int

	for price := range priceChannel {
		totalPrice += price
		countPrices++
		avgPrice := totalPrice / float64(countPrices)
		fmt.Printf("Preço recebido: R$ %.2f | Preço médio até agora: R$ %.2f\n", price, avgPrice)
	}

	done <- true
}
