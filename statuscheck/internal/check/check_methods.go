package check

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/thebiatriz/golang/statuscheck/internal/model"
)

func AllowMethodsToURL(c *gin.Context) {
	url := c.Query("url")

	if strings.TrimSpace(url) == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Url está vazia"})
		return
	}

	methods, err := checkURLMethods(url)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Erro ao verificar métodos para a URL",
			"details": err.Error(),
		})
		return
	}

	response := model.ResponseMethods{
		URL:     url,
		Methods: methods,
	}

	c.IndentedJSON(http.StatusOK, response)
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
