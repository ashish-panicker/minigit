package plumbing

import (
	"fmt"
	"os"
	"path/filepath"
)

// CheckRepoState verifies whether a .minigit repository already exists.
// Returns an error if:
//
// - .minigit exists (file or directory)
// - filesystem check failed
// - nil if repo does NOT exist (meaning it's safe to initialize)
func CheckRepoState() error {

	info, err := os.Stat(RepoDir)
	if err == nil {
		if !info.IsDir() {
			return fmt.Errorf("'.minigit' exists but is not a directory")
		}
		return fmt.Errorf("repository already exists")
	}

	if !os.IsNotExist(err) {
		return fmt.Errorf("failed to check .minigit directory: %w", err)

	}
	return nil
}

// CreateRepoStructure creates the necessary directory structure for a new repository.
func CreateRepoStructure() error {
	if err := CreateDir(RepoDir); err != nil {
		return err
	}
	if err := CreateDir(filepath.Join(RepoDir, ObjectsDir)); err != nil {
		return err
	}
	if err := CreateDir(filepath.Join(RepoDir, RefsDir, HeadsDir)); err != nil {
		return err
	}
	if err := CreateHeadFile(); err != nil {
		return err
	}
	return nil
}

func CreateDir(path string) error {
	if err := os.MkdirAll(path, 0755); err != nil {
		return fmt.Errorf("failed to create dir %s: %w", path, err)
	}
	return nil
}

// CreateHeadFile creates the HEAD file inside the .minigit directory.
func CreateHeadFile() error {
	err := os.WriteFile(filepath.Join(RepoDir, HeadFile), []byte(HeadFileContent), 0644)
	if err != nil {
		return err
	}
	return nil
}
