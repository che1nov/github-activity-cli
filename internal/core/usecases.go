package core

import "fmt"

// APIClient defines the methods for interacting with the GitHub API.
type APIClient interface {
	FetchEvents(username string) ([]Event, error)
}

// ActivityUsecase defines the business logic for processing GitHub activity.
type ActivityUsecase struct {
	apiClient APIClient
}

// NewActivityUsecase creates a new instance of ActivityUsecase.
func NewActivityUsecase(apiClient APIClient) *ActivityUsecase {
	return &ActivityUsecase{apiClient: apiClient}
}

// FetchAndDisplayActivity fetches and displays the user's activity.
func (uc *ActivityUsecase) FetchAndDisplayActivity(username string) error {
	// Fetch events using the APIClient
	events, err := uc.apiClient.FetchEvents(username)
	if err != nil {
		return err
	}

	// Display only the last 10 events
	for _, event := range events[:10] {
		switch event.Type {
		case "PushEvent":
			fmt.Printf("- Pushed %d commit(s) to %s\n", len(event.Payload.Commits), event.Repo.Name)
		case "IssuesEvent":
			fmt.Printf("- %s an issue in %s\n", capitalize(event.Payload.Action), event.Repo.Name)
		case "WatchEvent":
			fmt.Printf("- Starred %s\n", event.Repo.Name)
		default:
			fmt.Printf("- %s in %s\n", event.Type, event.Repo.Name)
		}
	}

	return nil
}

// capitalize converts the first character of a string to uppercase.
func capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	return string(s[0]-32) + s[1:]
}
