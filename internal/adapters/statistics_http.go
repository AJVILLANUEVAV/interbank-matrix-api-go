package adapters

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/AJVILLANUEVAV/interbank-matrix-api-go/internal/ports"
)

type StatisticsHTTPClient struct {
	baseURL string
	client  *http.Client
}

func NewStatisticsHTTPClient(baseURL string) *StatisticsHTTPClient {
	return &StatisticsHTTPClient{baseURL: baseURL, client: &http.Client{Timeout: 5 * time.Second}}
}

func (client *StatisticsHTTPClient) Calculate(rotated, q, r [][]float64) (map[string]any, error) {
	payload, err := json.Marshal(map[string]any{"matrices": map[string]any{"rotatedMatrix": rotated, "q": q, "r": r}})
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequestWithContext(context.Background(), http.MethodPost, client.baseURL+"/internal/v1/statistics", bytes.NewReader(payload))
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")
	response, err := client.client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ports.ErrStatisticsUnavailable, err)
	}
	defer response.Body.Close()
	if response.StatusCode >= http.StatusBadRequest {
		return nil, fmt.Errorf("%w: returned status %d", ports.ErrStatisticsUnavailable, response.StatusCode)
	}
	var result struct {
		Statistics map[string]any `json:"statistics"`
	}
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		return nil, err
	}
	return result.Statistics, nil
}
