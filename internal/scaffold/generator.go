package scaffold

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

// ApplyPlan creates directories and files from a plan.
func ApplyPlan(plan Plan, options Options) error {
	if plan.ProjectDir == "" {
		return errors.New("project dir cannot be empty")
	}

	if err := ensureProjectDir(plan.ProjectDir, options.Force); err != nil {
		return err
	}

	for _, dir := range plan.Dirs {
		fullPath := filepath.Join(plan.ProjectDir, dir)
		if err := os.MkdirAll(fullPath, 0o755); err != nil {
			return fmt.Errorf("create directory %s: %w", dir, err)
		}
	}

	for _, path := range plan.SortedFilePaths() {
		fullPath := filepath.Join(plan.ProjectDir, path)
		if err := os.MkdirAll(filepath.Dir(fullPath), 0o755); err != nil {
			return fmt.Errorf("prepare directory for %s: %w", path, err)
		}
		if err := os.WriteFile(fullPath, []byte(plan.Files[path]), 0o644); err != nil {
			return fmt.Errorf("create file %s: %w", path, err)
		}
	}

	return nil
}

func ensureProjectDir(projectDir string, force bool) error {
	info, err := os.Stat(projectDir)
	if err != nil {
		if os.IsNotExist(err) {
			return os.MkdirAll(projectDir, 0o755)
		}
		return fmt.Errorf("check project directory: %w", err)
	}
	if !info.IsDir() {
		return fmt.Errorf("project path exists and is not a directory: %s", projectDir)
	}

	entries, err := os.ReadDir(projectDir)
	if err != nil {
		return fmt.Errorf("read project directory: %w", err)
	}
	if len(entries) > 0 && !force {
		return fmt.Errorf("project directory is not empty: %s (use -force to overwrite)", projectDir)
	}
	return nil
}
