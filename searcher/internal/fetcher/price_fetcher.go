package fetcher

import (
	"math/rand"
	"time"
)

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