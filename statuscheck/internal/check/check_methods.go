package check

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/thebiatriz/golang/statuscheck/internal/model"
)

func AllowMethodsToURLs(c *gin.Context) {
	var req model.CheckRequest
	err := c.BindJSON(&req)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Erro ao receber o campo das URL's",
			"details": err.Error(),
		})
	}

	methodsChannel := make(chan model.ResponseMethods, len(req.URLs))
	var wg sync.WaitGroup
	wg.Add(len(req.URLs))

	for _, url := range req.URLs {
		go func(u string) {
			defer wg.Done()

			methods, err := checkURLMethods(u)

			finalMethods := methods

			if err != nil {
				finalMethods = nil
			}

			methodsChannel <- model.ResponseMethods{
				URL:     u,
				Methods: finalMethods,
			}
		}(url)
	}

	go func() {
		wg.Wait()
		close(methodsChannel)
	}()

	responseMap := make(map[string]model.ResponseMethods)
	for result := range methodsChannel {
		responseMap[result.URL] = result
	}

	c.IndentedJSON(http.StatusOK, responseMap)
}

func checkURLMethods(urlToCheck string) ([]string, error) {
	_, err := url.ParseRequestURI(urlToCheck)

	if err != nil {
		return nil, fmt.Errorf("formato da url inválido: %w", err)
	}

	client := http.Client{
		Timeout: 15 * time.Second,
	}

	newRequest, err := http.NewRequest("OPTIONS", urlToCheck, nil)
	if err != nil {
		return nil, fmt.Errorf("erro ao criar a nova requisiçaõ: %w", err)
	}

	resp, err := client.Do(newRequest)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, fmt.Errorf("tempo de requisição excedido: %w", err)
		}
		return nil, fmt.Errorf("erro ao executar a nova requisiçaõ: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return nil, fmt.Errorf("código de status inesperado: %d", resp.StatusCode)
	}

	if resp.StatusCode == http.StatusServiceUnavailable {
		return nil, fmt.Errorf("o servidor remoto está temporariamente indisponível (503)")
	}

	allowedMethodsHeader := resp.Header.Get("Allow")
	if strings.TrimSpace(allowedMethodsHeader) == "" {
		return []string{}, nil
	}

	methodsSlice := strings.Split(allowedMethodsHeader, ",")

	finalMethods := make([]string, 0, len(methodsSlice))

	for _, method := range methodsSlice {
		trimmedMethod := strings.TrimSpace(method)
		finalMethods = append(finalMethods, trimmedMethod)
	}

	return finalMethods, nil
}
