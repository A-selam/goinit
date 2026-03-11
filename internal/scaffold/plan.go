package scaffold

import (
	"errors"
	"fmt"
	"path/filepath"
	"sort"
	"strings"
)

// Plan defines the structure and files that should be generated.
type Plan struct {
	ProjectDir string
	Dirs       []string
	Files      map[string]string
}

// Options control how the plan is applied on disk.
type Options struct {
	Force bool
}

// BuildPlan creates a generation plan for a clean architecture Go project.
func BuildPlan(baseDir, projectName, moduleName string) (Plan, error) {
	projectName = strings.TrimSpace(projectName)
	if projectName == "" {
		return Plan{}, errors.New("project name cannot be empty")
	}
	if strings.ContainsAny(projectName, "\\/") {
		return Plan{}, errors.New("project name must not contain path separators")
	}

	moduleName = strings.TrimSpace(moduleName)
	if moduleName == "" {
		moduleName = projectName
	}

	projectDir := filepath.Join(baseDir, projectName)
	dirs := []string{
		"cmd/api",
		"config",
		"delivery/controller",
		"delivery/route",
		"domain",
		"infrastructure/middleware",
		"repository",
		"usecase",
		"utils",
		"tests",
		".github/workflows",
	}

	files := map[string]string{
		"cmd/api/main.go":                                cmdMainTemplate(),
		"config/config.go":                               "package config\n",
		"delivery/controller/user_controller.go":         "package controller\n",
		"delivery/controller/order_controller.go":        "package controller\n",
		"delivery/route/routes.go":                       "package route\n",
		"domain/user.go":                                 "package domain\n",
		"domain/repository.go":                           "package domain\n",
		"infrastructure/middleware/auth_middleware.go":   "package middleware\n",
		"infrastructure/middleware/logger_middleware.go": "package middleware\n",
		"infrastructure/database.go":                     "package infrastructure\n",
		"infrastructure/logger.go":                       "package infrastructure\n",
		"infrastructure/twilio.go":                       "package infrastructure\n",
		"repository/user_repository.go":                  "package repository\n",
		"usecase/user_usecase.go":                        "package usecase\n",
		"utils/utils.go":                                 "package utils\n",
		"utils/Dockerfile":                               dockerfileTemplate(),
		"utils/openapi.yaml":                             openAPITemplate(),
		".github/workflows/ci.yml":                       ciTemplate(),
		"tests/user_usecase_test.go":                     "package tests\n",
		".gitignore":                                     gitignoreTemplate(),
		"README.md":                                      readmeTemplate(projectName),
		"render.yaml":                                    renderTemplate(projectName),
		"go.mod":                                         goModTemplate(moduleName),
	}

	return Plan{
		ProjectDir: projectDir,
		Dirs:       dirs,
		Files:      files,
	}, nil
}

// SortedFilePaths returns deterministic file paths for display and writes.
func (p Plan) SortedFilePaths() []string {
	paths := make([]string, 0, len(p.Files))
	for path := range p.Files {
		paths = append(paths, path)
	}
	sort.Strings(paths)
	return paths
}

// Summary returns human-readable plan details.
func (p Plan) Summary() string {
	return fmt.Sprintf("project=%s dirs=%d files=%d", p.ProjectDir, len(p.Dirs), len(p.Files))
}
