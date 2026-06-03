package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestDocumentedCommandSmokeSuite(t *testing.T) {
	binary := buildAcceptanceBinary(t)
	repoRoot := filepath.Clean(filepath.Join("..", ".."))
	outputPath := filepath.Join(t.TempDir(), "prmaven-report.json")

	tests := []struct {
		name                 string
		documentedCommand    string
		args                 []string
		wantExitCode         int
		wantStdout           []string
		unwantedStdout       []string
		wantJSONFindingCount *int
		wantOutputPath       string
		wantOutputJSONCount  *int
	}{
		{
			name:              "help short flag",
			documentedCommand: "prmaven -h",
			args:              []string{"-h"},
			wantExitCode:      0,
			wantStdout: []string{
				"Usage:",
				"Commands:",
				"Exit codes:",
			},
		},
		{
			name:              "help command",
			documentedCommand: "prmaven help",
			args:              []string{"help"},
			wantExitCode:      0,
			wantStdout: []string{
				"Usage:",
				"Flags:",
				"Examples:",
			},
		},
		{
			name:              "version command",
			documentedCommand: "prmaven version",
			args:              []string{"version"},
			wantExitCode:      0,
			wantStdout:        []string{"dev"},
		},
		{
			name:              "failure demo text",
			documentedCommand: "prmaven fails -project demo/multi-module-failure",
			args:              []string{"fails", "-project", "demo/multi-module-failure"},
			wantExitCode:      1,
			wantStdout: []string{
				"PR Maven CLI - Maven failure context",
				"Modules: 3 | Reports: 2 | Findings: 2",
				"maven-surefire-plugin",
				"maven-failsafe-plugin",
			},
		},
		{
			name:                 "failure demo json",
			documentedCommand:    "prmaven why -project demo/multi-module-failure -format json",
			args:                 []string{"why", "-project", "demo/multi-module-failure", "-format", "json"},
			wantExitCode:         1,
			wantJSONFindingCount: intPtr(2),
		},
		{
			name:              "no failure demo text",
			documentedCommand: "prmaven fails -project demo/no-failure",
			args:              []string{"fails", "-project", "demo/no-failure"},
			wantExitCode:      0,
			wantStdout: []string{
				"Modules: 2 | Reports: 1 | Findings: 0",
				"No Maven test or quality failures found",
			},
		},
		{
			name:              "module filter demo",
			documentedCommand: "prmaven why -project demo/multi-module-failure -module payment-core",
			args:              []string{"why", "-project", "demo/multi-module-failure", "-module", "payment-core"},
			wantExitCode:      1,
			wantStdout: []string{
				"Findings: 1",
				"Module: payment-core (payment-core)",
				"maven-surefire-plugin",
			},
			unwantedStdout: []string{
				"payment-api",
			},
		},
		{
			name:              "json output file",
			documentedCommand: "prmaven why -project demo/multi-module-failure -format json -output prmaven-report.json",
			args: []string{
				"why",
				"-project", "demo/multi-module-failure",
				"-format", "json",
				"-output", outputPath,
			},
			wantExitCode:        1,
			wantOutputPath:      outputPath,
			wantOutputJSONCount: intPtr(2),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stdout, stderr, exitCode := runDocumentedCommand(t, repoRoot, binary, tt.args...)
			if exitCode != tt.wantExitCode {
				t.Fatalf("documented command %q exit code = %d, want %d\nstdout:\n%s\nstderr:\n%s", tt.documentedCommand, exitCode, tt.wantExitCode, stdout, stderr)
			}
			if stderr != "" {
				t.Fatalf("documented command %q stderr = %q, want empty", tt.documentedCommand, stderr)
			}

			for _, expected := range tt.wantStdout {
				if !strings.Contains(stdout, expected) {
					t.Fatalf("documented command %q stdout missing %q\n%s", tt.documentedCommand, expected, stdout)
				}
			}
			for _, unwanted := range tt.unwantedStdout {
				if strings.Contains(stdout, unwanted) {
					t.Fatalf("documented command %q stdout contains unwanted %q\n%s", tt.documentedCommand, unwanted, stdout)
				}
			}
			if tt.wantJSONFindingCount != nil {
				assertDocumentedJSONFindingCount(t, tt.documentedCommand, []byte(stdout), *tt.wantJSONFindingCount)
			}

			if tt.wantOutputPath != "" {
				if stdout != "" {
					t.Fatalf("documented command %q stdout = %q, want empty when -output is set", tt.documentedCommand, stdout)
				}
				output, err := os.ReadFile(tt.wantOutputPath)
				if err != nil {
					t.Fatalf("read documented output file for %q: %v", tt.documentedCommand, err)
				}
				if tt.wantOutputJSONCount != nil {
					assertDocumentedJSONFindingCount(t, tt.documentedCommand, output, *tt.wantOutputJSONCount)
				}
			}
		})
	}
}

func runDocumentedCommand(t *testing.T, dir, binary string, args ...string) (string, string, int) {
	t.Helper()

	command := exec.Command(binary, args...)
	command.Dir = dir
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	command.Stdout = &stdout
	command.Stderr = &stderr

	err := command.Run()
	if err == nil {
		return stdout.String(), stderr.String(), 0
	}

	var exitErr *exec.ExitError
	if errors.As(err, &exitErr) {
		return stdout.String(), stderr.String(), exitErr.ExitCode()
	}

	t.Fatalf("run documented command: %v\nstdout:\n%s\nstderr:\n%s", err, stdout.String(), stderr.String())
	return "", "", 1
}

func assertDocumentedJSONFindingCount(t *testing.T, command string, output []byte, want int) {
	t.Helper()

	var report struct {
		Summary struct {
			FindingCount int `json:"findingCount"`
		} `json:"summary"`
	}
	if err := json.Unmarshal(output, &report); err != nil {
		t.Fatalf("documented command %q produced invalid JSON: %v\n%s", command, err, string(output))
	}
	if report.Summary.FindingCount != want {
		t.Fatalf("documented command %q finding count = %d, want %d", command, report.Summary.FindingCount, want)
	}
}

func intPtr(value int) *int {
	return &value
}
