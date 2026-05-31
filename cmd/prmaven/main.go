package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path"
	"strings"

	"github.com/pr-cli/pr-maven-cli/pkg/prmaven"
)

var version = "dev"

func main() {
	os.Exit(run(os.Args[1:], os.Stdout, os.Stderr))
}

func run(args []string, stdout, stderr io.Writer) int {
	if wantsHelp(args) {
		writeUsage(stdout)
		return 0
	}

	command := "fails"
	if len(args) > 0 && args[0] != "" && args[0][0] != '-' {
		command = args[0]
		args = args[1:]
	}

	flags := flag.NewFlagSet("prmaven", flag.ContinueOnError)
	flags.SetOutput(stderr)
	flags.Usage = func() {
		writeUsage(stderr)
	}
	projectDir := flags.String("project", ".", "Maven project directory")
	format := flags.String("format", "text", "output format: text or json")
	moduleFilter := flags.String("module", "", "limit findings to a Maven module path or artifactId")
	outputPath := flags.String("output", "", "write output to file instead of stdout")

	if err := flags.Parse(args); err != nil {
		writeUsage(stderr)
		return 2
	}

	if flags.NArg() > 0 {
		command = flags.Arg(0)
	}

	switch command {
	case "fails", "why", "version", "help":
	default:
		fmt.Fprintf(stderr, "unknown command %q\n", command)
		writeUsage(stderr)
		return 2
	}

	if command == "help" {
		writeUsage(stdout)
		return 0
	}

	if command == "version" {
		fmt.Fprintln(stdout, version)
		return 0
	}

	report, err := prmaven.Analyze(prmaven.Options{ProjectDir: *projectDir})
	if err != nil {
		fmt.Fprintln(stderr, err)
		return 1
	}
	report = filterReportByModule(report, *moduleFilter)

	switch *format {
	case "json", "text":
	default:
		fmt.Fprintf(stderr, "unknown format %q\n", *format)
		fmt.Fprintln(stderr, "available formats: text, json")
		return 2
	}

	output := stdout
	var outputFile *os.File
	if *outputPath != "" {
		outputFile, err = os.Create(*outputPath)
		if err != nil {
			fmt.Fprintf(stderr, "create output file %q: %v\n", *outputPath, err)
			return 1
		}
		output = outputFile
	}

	var writeErr error
	switch *format {
	case "json":
		writeErr = prmaven.WriteJSON(output, report)
	case "text":
		writeErr = prmaven.WriteText(output, report)
	}

	if outputFile != nil {
		if closeErr := outputFile.Close(); writeErr == nil && closeErr != nil {
			writeErr = closeErr
		}
	}

	if writeErr != nil {
		fmt.Fprintln(stderr, writeErr)
		return 1
	}

	if report.Summary.FindingCount > 0 {
		return 1
	}
	return 0
}

func wantsHelp(args []string) bool {
	if len(args) > 0 && args[0] == "help" {
		return true
	}
	for _, arg := range args {
		switch arg {
		case "-h", "--help", "-help":
			return true
		}
	}
	return false
}

func writeUsage(w io.Writer) {
	fmt.Fprint(w, `PR Maven CLI - Maven-aware PR/CI failure context

Usage:
  prmaven [command] [flags]

Commands:
  fails     Analyze local Maven reports and print actionable failure context.
  why       Analyze local Maven reports; reserved for richer causality context.
  version   Print the CLI version.
  help      Print this help.

Flags:
  -project string
      Maven project directory to analyze. Defaults to ".".
  -format string
      Output format: text or json. Defaults to "text".
  -module string
      Limit findings to a Maven module path or artifactId.
  -output string
      Write output to a file instead of stdout.
  -h, -help, --help
      Print this help.

Examples:
  prmaven fails -project .
  prmaven fails -project . -format json
  prmaven why -project . -module payment-core
  prmaven why -project . -format json -output prmaven-report.json
  prmaven version

Exit codes:
  0  Analysis completed with no findings, or help/version completed.
  1  Analysis completed with Maven findings, or analysis/output failed.
  2  Invalid CLI usage.
`)
}

func filterReportByModule(report prmaven.Report, moduleFilter string) prmaven.Report {
	moduleFilter = strings.TrimSpace(moduleFilter)
	if moduleFilter == "" {
		return report
	}

	matchingModulePaths := map[string]bool{}
	cleanFilterPath := cleanModulePath(moduleFilter)
	for _, module := range report.Modules {
		if module.Name == moduleFilter || cleanModulePath(module.Path) == cleanFilterPath {
			matchingModulePaths[module.Path] = true
		}
	}

	filteredFindings := make([]prmaven.Finding, 0, len(report.Findings))
	for _, finding := range report.Findings {
		if matchingModulePaths[finding.ModulePath] {
			filteredFindings = append(filteredFindings, finding)
		}
	}

	report.Findings = filteredFindings
	report.Summary.FindingCount = len(filteredFindings)
	return report
}

func cleanModulePath(value string) string {
	value = strings.TrimSpace(value)
	value = strings.ReplaceAll(value, "\\", "/")
	return path.Clean(value)
}
