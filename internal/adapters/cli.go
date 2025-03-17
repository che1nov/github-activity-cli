package adapters

import (
	"fmt"
	"github-activity-cli/internal/core"
	"os"
)

// CLIAdapter implements the interface for handling command-line interactions.
type CLIAdapter struct {
	usecase *core.ActivityUsecase
}

// NewCLIAdapter creates a new instance of CLIAdapter.
func NewCLIAdapter(usecase *core.ActivityUsecase) *CLIAdapter {
	return &CLIAdapter{usecase: usecase}
}

// Run starts the CLI application.
func (c *CLIAdapter) Run(args []string) error {
	// Check if the username argument is provided
	if len(args) != 2 {
		fmt.Println("Usage: github-activity <username>")
		os.Exit(1)
	}

	username := args[1]
	fmt.Printf("Loading activity for user %s...\n\n", username)

	// Fetch and display user activity
	err := c.usecase.FetchAndDisplayActivity(username)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return nil
}
