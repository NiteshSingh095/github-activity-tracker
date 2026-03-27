package api

import (
	"fmt"
	"io"
	"net/http"
)

func FetchUserActivity(username string) ([]byte, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s/events", username)

	response, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data: %w", err)
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed with status: %d", response.StatusCode)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return body, nil
}