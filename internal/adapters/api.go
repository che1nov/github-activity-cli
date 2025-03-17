package adapters

import (
	"encoding/json"
	"fmt"
	"github-activity-cli/internal/core"
	"io/ioutil"
	"net/http"
)

// GitHubAPIAdapter implements the APIClient interface to work with the GitHub API.
type GitHubAPIAdapter struct{}

// NewGitHubAPIAdapter creates a new instance of GitHubAPIAdapter.
func NewGitHubAPIAdapter() *GitHubAPIAdapter {
	return &GitHubAPIAdapter{}
}

// FetchEvents retrieves GitHub events for the specified user.
func (a *GitHubAPIAdapter) FetchEvents(username string) ([]core.Event, error) {
	// Construct the API URL for the given username
	url := fmt.Sprintf("https://api.github.com/users/%s/events", username)

	// Send an HTTP GET request to the GitHub API
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error connecting to the API: %v", err)
	}
	defer resp.Body.Close()

	// Check the HTTP response status code
	if resp.StatusCode == 404 {
		return nil, fmt.Errorf("user '%s' not found", username)
	} else if resp.StatusCode != 200 {
		return nil, fmt.Errorf("API failure (status code %d)", resp.StatusCode)
	}

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading the response body: %v", err)
	}

	// Parse the JSON response into a slice of core.Event
	var events []core.Event
	err = json.Unmarshal(body, &events)
	if err != nil {
		return nil, fmt.Errorf("error parsing JSON: %v", err)
	}

	// Return the list of events
	return events, nil
}
