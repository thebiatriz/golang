package main

import (
	"fmt"
	"time"

	"github.com/thebiatriz/golang/searcher/internal/fetcher"
	"github.com/thebiatriz/golang/searcher/internal/models"
	"github.com/thebiatriz/golang/searcher/internal/processor"
)

func main() {
	start := time.Now()
	priceChannel := make(chan models.PriceDetails)

	done := make(chan bool)

	go fetcher.FetchPrices(priceChannel)
	go processor.ShowPriceAVG(priceChannel, done)

	<-done // mecanismo de sincronização

	fmt.Printf("\nTempo total: %s", time.Since(start))
}
