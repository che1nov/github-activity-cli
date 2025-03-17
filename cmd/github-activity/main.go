package main

import (
	"github-activity-cli/internal/adapters"
	"github-activity-cli/internal/core"
	"os"
)

func main() {
	// Create adapters
	apiClient := adapters.NewGitHubAPIAdapter()
	usecase := core.NewActivityUsecase(apiClient)
	cliHandler := adapters.NewCLIAdapter(usecase)

	// Start the CLI
	err := cliHandler.Run(os.Args)
	if err != nil {
		os.Exit(1)
	}
}
