package main

import (
	"fmt"
	"time"

	"github.com/thebiatriz/golang/searcher/internal/fetcher"
	"github.com/thebiatriz/golang/searcher/internal/processor"
)

func main() {
	start := time.Now()
	priceChannel := make(chan float64)

	go fetcher.FetchPrices(priceChannel)

	processor.ShowPriceAVG(priceChannel)

	fmt.Printf("\nTempo total: %s", time.Since(start))
}
