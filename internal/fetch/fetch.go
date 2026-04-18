package fetch

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

type Fetcher struct {
	httpClient *http.Client
}

func NewFetcher() *Fetcher {
	return &Fetcher{
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (f *Fetcher) Fetch(url string) ([]byte, error) {
	response, err := f.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code %d for %s", response.StatusCode, url)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
