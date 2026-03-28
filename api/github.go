package api

import (
	"fmt"
	"io"
	"net/http"
	"encoding/json"
	"github-activity/models"
)

func FetchUserActivity(username string) ([]models.Event, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s/events", username)

	response, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data: %w", err)
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("user not found: %s", username)
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed with status: %d", response.StatusCode)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var events []models.Event;

	err = json.Unmarshal(body, &events)
	if err != nil {
		return nil, fmt.Errorf("failed to parse json %s", err)
	}

	return events, nil
}