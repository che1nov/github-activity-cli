package core

// Event represents a GitHub event.
type Event struct {
	Type    string `json:"type"`
	Repo    Repo   `json:"repo"`
	Payload Payload
}

// Repo represents a GitHub repository.
type Repo struct {
	Name string `json:"name"`
}

// Payload contains additional data related to the event.
type Payload struct {
	Commits []Commit `json:"commits"`
	Action  string   `json:"action"`
}

// Commit represents a GitHub commit.
type Commit struct {
	Message string `json:"message"`
}
