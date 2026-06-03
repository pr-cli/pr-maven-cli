package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func TestCLIAcceptanceSuite(t *testing.T) {
	binary := buildAcceptanceBinary(t)

	tests := []struct {
		name             string
		args             []string
		wantExitCode     int
		wantStdout       []string
		wantFindingCount int
		wantJSON         bool
	}{
		{
			name:         "failure text workflow",
			args:         []string{"fails", "-project", "../../demo/multi-module-failure"},
			wantExitCode: 1,
			wantStdout: []string{
				"PR Maven CLI - Maven failure context",
				"Modules: 3 | Reports: 2 | Findings: 2",
				"maven-surefire-plugin",
				"maven-failsafe-plugin",
			},
		},
		{
			name:         "no failure text workflow",
			args:         []string{"fails", "-project", "../../demo/no-failure"},
			wantExitCode: 0,
			wantStdout: []string{
				"Modules: 2 | Reports: 1 | Findings: 0",
				"No Maven test or quality failures found",
			},
		},
		{
			name:             "failure json workflow",
			args:             []string{"why", "-project", "../../demo/multi-module-failure", "-format", "json"},
			wantExitCode:     1,
			wantFindingCount: 2,
			wantJSON:         true,
		},
		{
			name:             "no failure json workflow",
			args:             []string{"why", "-project", "../../demo/no-failure", "-format", "json"},
			wantExitCode:     0,
			wantFindingCount: 0,
			wantJSON:         true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stdout, stderr, exitCode := runAcceptanceBinary(t, binary, tt.args...)
			if exitCode != tt.wantExitCode {
				t.Fatalf("exit code = %d, want %d\nstdout:\n%s\nstderr:\n%s", exitCode, tt.wantExitCode, stdout, stderr)
			}
			if stderr != "" {
				t.Fatalf("stderr = %q, want empty", stderr)
			}

			for _, expected := range tt.wantStdout {
				if !strings.Contains(stdout, expected) {
					t.Fatalf("stdout missing %q\n%s", expected, stdout)
				}
			}

			if tt.wantJSON {
				var report struct {
					Summary struct {
						FindingCount int `json:"findingCount"`
					} `json:"summary"`
				}
				if err := json.Unmarshal([]byte(stdout), &report); err != nil {
					t.Fatalf("stdout is not valid JSON: %v\n%s", err, stdout)
				}
				if report.Summary.FindingCount != tt.wantFindingCount {
					t.Fatalf("finding count = %d, want %d", report.Summary.FindingCount, tt.wantFindingCount)
				}
			}
		})
	}
}

type acceptanceOutputCase struct {
	name             string
	format           string
	outputFile       string
	wantExitCode     int
	wantStdout       []string
	wantFileContains []string
	wantFindingCount int
	wantJSON         bool
}

func TestCLIOutputFileMatrix(t *testing.T) {
	binary := buildAcceptanceBinary(t)

	tests := []acceptanceOutputCase{
		{
			name:         "text stdout without output file",
			format:       "text",
			wantExitCode: 1,
			wantStdout: []string{
				"PR Maven CLI - Maven failure context",
				"Modules: 3 | Reports: 2 | Findings: 2",
				"maven-surefire-plugin",
				"maven-failsafe-plugin",
			},
		},
		{
			name:             "json stdout without output file",
			format:           "json",
			wantExitCode:     1,
			wantJSON:         true,
			wantFindingCount: 2,
		},
		{
			name:         "text output file",
			format:       "text",
			outputFile:   "prmaven-report.txt",
			wantExitCode: 1,
			wantFileContains: []string{
				"PR Maven CLI - Maven failure context",
				"Modules: 3 | Reports: 2 | Findings: 2",
				"maven-surefire-plugin",
				"maven-failsafe-plugin",
			},
		},
		{
			name:             "json output file",
			format:           "json",
			outputFile:       "prmaven-report.json",
			wantExitCode:     1,
			wantJSON:         true,
			wantFindingCount: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			outputDir := t.TempDir()
			outputPath := filepath.Join(outputDir, "unused-report")
			args := []string{"why", "-project", "../../demo/multi-module-failure"}
			if tt.format != "text" {
				args = append(args, "-format", tt.format)
			}
			if tt.outputFile != "" {
				outputPath = filepath.Join(outputDir, tt.outputFile)
				args = append(args, "-output", outputPath)
			}

			stdout, stderr, exitCode := runAcceptanceBinary(t, binary, args...)
			if exitCode != tt.wantExitCode {
				t.Fatalf("exit code = %d, want %d\nstdout:\n%s\nstderr:\n%s", exitCode, tt.wantExitCode, stdout, stderr)
			}
			if stderr != "" {
				t.Fatalf("stderr = %q, want empty", stderr)
			}

			if tt.outputFile == "" {
				if stdout == "" {
					t.Fatal("stdout is empty, want report output when -output is absent")
				}
				if _, err := os.Stat(outputPath); !errors.Is(err, os.ErrNotExist) {
					t.Fatalf("unexpected output file at %s when -output is absent", outputPath)
				}
				assertAcceptanceOutput(t, stdout, tt)
				return
			}

			if stdout != "" {
				t.Fatalf("stdout = %q, want empty when -output is set", stdout)
			}

			output, err := os.ReadFile(outputPath)
			if err != nil {
				t.Fatalf("read output file: %v", err)
			}
			assertAcceptanceOutput(t, string(output), tt)
		})
	}
}

func assertAcceptanceOutput(t *testing.T, output string, tt acceptanceOutputCase) {
	t.Helper()

	for _, expected := range append(tt.wantStdout, tt.wantFileContains...) {
		if !strings.Contains(output, expected) {
			t.Fatalf("output missing %q\n%s", expected, output)
		}
	}

	if tt.wantJSON {
		var report struct {
			Summary struct {
				FindingCount int `json:"findingCount"`
			} `json:"summary"`
		}
		if err := json.Unmarshal([]byte(output), &report); err != nil {
			t.Fatalf("output is not valid JSON: %v\n%s", err, output)
		}
		if report.Summary.FindingCount != tt.wantFindingCount {
			t.Fatalf("finding count = %d, want %d", report.Summary.FindingCount, tt.wantFindingCount)
		}
	}
}

func buildAcceptanceBinary(t *testing.T) string {
	t.Helper()

	binary := filepath.Join(t.TempDir(), "prmaven")
	if runtime.GOOS == "windows" {
		binary += ".exe"
	}

	command := exec.Command("go", "build", "-o", binary, ".")
	output, err := command.CombinedOutput()
	if err != nil {
		t.Fatalf("build acceptance binary: %v\n%s", err, string(output))
	}
	return binary
}

func runAcceptanceBinary(t *testing.T, binary string, args ...string) (string, string, int) {
	t.Helper()

	command := exec.Command(binary, args...)
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

	t.Fatalf("run acceptance binary: %v\nstdout:\n%s\nstderr:\n%s", err, stdout.String(), stderr.String())
	return "", "", 1
}
