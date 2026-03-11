package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"goinit/internal/scaffold"
)

func main() {
	nameFlag := flag.String("name", "", "Project name (required if not prompted)")
	moduleFlag := flag.String("module", "", "Go module name for generated project (defaults to project name)")
	pathFlag := flag.String("path", ".", "Base path where the project folder will be created")
	dryRunFlag := flag.Bool("dry-run", false, "Print planned files/directories without writing")
	forceFlag := flag.Bool("force", false, "Allow writing into a non-empty existing project folder")
	flag.Parse()

	projectName := strings.TrimSpace(*nameFlag)
	if projectName == "" {
		value, err := prompt("Enter project name (e.g., my-go-app): ")
		if err != nil {
			fail("read project name", err)
		}
		projectName = value
	}

	basePath, err := filepath.Abs(strings.TrimSpace(*pathFlag))
	if err != nil {
		fail("resolve base path", err)
	}

	plan, err := scaffold.BuildPlan(basePath, projectName, *moduleFlag)
	if err != nil {
		fail("build generation plan", err)
	}

	if *dryRunFlag {
		fmt.Printf("Dry run: %s\n", plan.Summary())
		for _, dir := range plan.Dirs {
			fmt.Printf("  dir  %s\n", filepath.Join(plan.ProjectDir, dir))
		}
		for _, path := range plan.SortedFilePaths() {
			fmt.Printf("  file %s\n", filepath.Join(plan.ProjectDir, path))
		}
		return
	}

	err = scaffold.ApplyPlan(plan, scaffold.Options{Force: *forceFlag})
	if err != nil {
		fail("apply generation plan", err)
	}

	fmt.Printf("Project structure created successfully in %s\n", plan.ProjectDir)
	fmt.Printf("Next: cd %s && go mod tidy\n", projectName)
}

func prompt(label string) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(label)
	value, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	value = strings.TrimSpace(value)
	if value == "" {
		return "", fmt.Errorf("value cannot be empty")
	}
	return value, nil
}

func fail(step string, err error) {
	fmt.Fprintf(os.Stderr, "Error: %s: %v\n", step, err)
	os.Exit(1)
}
