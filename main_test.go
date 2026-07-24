package main

import (
	"testing"
)

func TestDatabasePathAndArgsUsesDefaultPath(t *testing.T) {
	path, args := databasePathAndArgs([]string{"vitalynq", "observations", "list"})

	if path != defaultDatabasePath {
		t.Fatalf("path = %q, want %q", path, defaultDatabasePath)
	}

	wantArgs := []string{"vitalynq", "observations", "list"}
	assertArgs(t, args, wantArgs)
}

func TestDatabasePathAndArgsUsesCustomPath(t *testing.T) {
	path, args := databasePathAndArgs([]string{"vitalynq", "--db", "test.db", "observations", "list"})

	if path != "test.db" {
		t.Fatalf("path = %q, want %q", path, "test.db")
	}

	wantArgs := []string{"vitalynq", "observations", "list"}
	assertArgs(t, args, wantArgs)
}

func assertArgs(t *testing.T, got []string, want []string) {
	t.Helper()

	if len(got) != len(want) {
		t.Fatalf("len(args) = %d, want %d", len(got), len(want))
	}

	for index := range want {
		if got[index] != want[index] {
			t.Fatalf("args[%d] = %q, want %q", index, got[index], want[index])
		}
	}
}
