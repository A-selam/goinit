package scaffold

import "testing"

func TestBuildPlanDefaults(t *testing.T) {
	base := t.TempDir()
	plan, err := BuildPlan(base, "sample-app", "")
	if err != nil {
		t.Fatalf("BuildPlan returned error: %v", err)
	}

	if plan.ProjectDir == "" {
		t.Fatal("expected project dir to be set")
	}
	if len(plan.Dirs) == 0 {
		t.Fatal("expected directories in plan")
	}
	if _, ok := plan.Files["go.mod"]; !ok {
		t.Fatal("expected go.mod in plan files")
	}
	if _, ok := plan.Files["README.md"]; !ok {
		t.Fatal("expected README.md in plan files")
	}
}

func TestBuildPlanRejectsInvalidName(t *testing.T) {
	_, err := BuildPlan(t.TempDir(), "bad/name", "")
	if err == nil {
		t.Fatal("expected error for invalid project name")
	}
}
