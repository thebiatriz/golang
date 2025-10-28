package check

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/thebiatriz/golang/statuscheck/internal/model"
)

func CheckMultipleURLs(c *gin.Context) {
	var req model.CheckRequest

	err := c.BindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Erro ao ler os dados da requisição"})
	}

	resultChannel := make(chan model.ResponseCheck, len(req.URLs))

	var wg sync.WaitGroup
	wg.Add(len(req.URLs))

	for _, url := range req.URLs {
		go func(u string) {
			defer wg.Done()

			status, err := checkSite(u)

			finalStatus := status

			if err != nil {
				finalStatus = fmt.Sprintf("%s - %s", status, err)
			}

			resultChannel <- model.ResponseCheck{
				URL:    u,
				Status: finalStatus,
			}
		}(url)
	}

	go func() {
		wg.Wait()
		close(resultChannel)
	}()

	finalResultsMap := make(map[string]string)
	for result := range resultChannel {
		finalResultsMap[result.URL] = result.Status
	}

	c.IndentedJSON(http.StatusOK, finalResultsMap)
}

func checkSite(urlToCheck string) (string, error) {
	resp, err := http.Head(urlToCheck)

	if err != nil {
		return "ERROR", fmt.Errorf("falha ao conectar: %w", err)
	}

	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		return "UP", nil
	} else {
		return "DOWN", fmt.Errorf("código de status inesperado: %d", resp.StatusCode)
	}
}
