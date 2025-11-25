package commands

import (
	"fmt"

	"github.com/ashish-panicker/minigit/internal/plumbing"
)

// InitRepo initializes a new repository at the given path.
// It returns an error if the repository already exists.
// This will create necessary directories and files for the repository.
func InitRepo() error {

	// Step 1: Ensure the .minigit directory state
	if err := plumbing.CheckRepoState(); err != nil {
		return err
	}
	// Step 2: Create the .minigit directory and HEAD file
	if err := plumbing.CreateRepoStructure(); err != nil {
		return err
	}
	fmt.Printf("Repository initialized successfully.\n")
	return nil
}
