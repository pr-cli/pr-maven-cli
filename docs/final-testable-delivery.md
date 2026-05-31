# Final Testable Delivery

This document defines what "production-usable MVP" means for PR Maven CLI.

For staged dates and future scope, see [ROADMAP.md](../ROADMAP.md). This document focuses on the observable delivery target and its acceptance checks.

## Target Workflow

```text
Maven reports/logs exist locally
  -> prmaven scans a Maven workspace
  -> modules are discovered from pom.xml files
  -> supported reports and logs are parsed
  -> findings are mapped back to Maven modules
  -> reproduction commands are generated
  -> text or JSON output is emitted
  -> CI or a coding agent can consume the result
```

## Source Boundary

PR Maven CLI starts after Maven, CI, or a developer has produced local artifacts.

The MVP does not require PR Maven CLI to:

- run Maven;
- download CI artifacts;
- call GitHub or GitLab APIs;
- access remote repositories;
- upload logs or reports.

## Required Capabilities

### Discover Maven Modules

Acceptance expectations:

- discover the root Maven project;
- discover direct and nested modules from `pom.xml`;
- preserve module path and artifact id;
- keep path output stable across operating systems.

### Parse Supported Failure Artifacts

Acceptance expectations:

- parse Surefire JUnit XML failures;
- parse Failsafe JUnit XML errors;
- parse Checkstyle XML violations;
- parse SpotBugs XML bug instances;
- extract Maven Enforcer failures from sanitized log fixtures;
- extract JaCoCo threshold failures from sanitized log fixtures;
- ignore unsupported files without failing the full analysis.

### Map Findings to Maven Context

Acceptance expectations:

- map findings to the closest Maven module;
- include the source report path;
- identify the Maven plugin and phase when known;
- include confidence reasons based on deterministic evidence.

### Generate Reproduction Commands

Acceptance expectations:

- generate module-scoped Maven commands for Surefire and Failsafe findings;
- use appropriate properties such as `-Dtest` and `-Dit.test`;
- keep commands conservative for quality plugin findings.

### Emit Human-Readable Output

Acceptance expectations:

- include module, plugin, phase, source, confidence, and reproduction command;
- keep no-failure output clear;
- protect intentional wording through golden tests.

### Emit Machine-Readable Output

Acceptance expectations:

- emit stable JSON for `Report`, `Summary`, `Module`, and `Finding`;
- document all public fields;
- validate the schema against generated demo output;
- keep path separators stable.

### Support Focused CLI Ergonomics

Acceptance expectations:

- support `fails` and `why`;
- support `version`;
- support `-project`;
- support `-format text|json`;
- support `-module`;
- support `-output`;
- return documented exit codes.

### Support Library Usage

Acceptance expectations:

- expose a Go package through `pkg/prmaven`;
- allow callers to run analysis without invoking the CLI;
- keep public structures simple and documented.

### Support OSS Maintenance

Acceptance expectations:

- provide setup and usage documentation;
- provide contribution guidance;
- provide issue and PR templates;
- run CI on pull requests;
- run release automation from tags;
- document security and permission posture.

## Acceptance Criteria

### Functional Acceptance

1. `go run ./cmd/prmaven fails -project demo/multi-module-failure` reports fixture-backed Maven findings.
2. `go run ./cmd/prmaven why -project demo/multi-module-failure -format json` emits parseable JSON.
3. `go run ./cmd/prmaven fails -project demo/no-failure` exits successfully and reports zero findings.
4. `go run ./cmd/prmaven why -project demo/multi-module-failure -module payment-core` limits output to the selected module.
5. `go run ./cmd/prmaven why -project demo/multi-module-failure -format json -output prmaven-report.json` writes the report to the selected path.

### Validation Acceptance

1. `go test ./...` passes.
2. CLI end-to-end tests cover success, findings, invalid usage, JSON, module filtering, and output files.
3. Parser tests cover all supported artifact kinds with sanitized fixtures.
4. Golden tests protect text output.
5. JSON schema tests validate generated report shape.
6. Path normalization tests cover Windows-style and POSIX-style paths.

### Safety Acceptance

1. Core analysis requires no network access.
2. Core analysis requires no provider token.
3. Core analysis does not upload reports or logs.
4. Unsupported artifacts do not trigger speculative findings.
5. Provider integrations remain optional and outside the core analyzer.
6. Maven 4 is not declared production-ready before compatibility is validated.

## End-to-End Test Scenario

The MVP should keep an end-to-end test path that proves:

```text
temporary CLI execution
  -> demo Maven workspace analyzed
  -> expected findings returned
  -> expected exit code observed
  -> JSON output parsed
  -> optional file output written
```

The scenario should include both:

- a fixture with findings;
- a fixture with no findings.

## Manual Smoke Test

From the repository root:

```bash
go test ./...
go run ./cmd/prmaven fails -project demo/multi-module-failure
go run ./cmd/prmaven why -project demo/multi-module-failure -format json
go run ./cmd/prmaven fails -project demo/no-failure
```

Optional file-output smoke test:

```bash
go run ./cmd/prmaven why -project demo/multi-module-failure -format json -output prmaven-report.json
```

Remove generated smoke-test files before committing.

## Out of Scope for This Delivery

- hosted service operation;
- native GitHub adapter implementation;
- native GitLab adapter implementation;
- PR comment publishing;
- SARIF output;
- GitHub annotation output;
- baseline comparison against `main`;
- automatic code fixes;
- telemetry;
- Maven 4 production support.

## Release Readiness Checklist

- [x] CLI entrypoint exists.
- [x] Go library package exists.
- [x] Maven module discovery is implemented.
- [x] Surefire and Failsafe parsing are implemented.
- [x] Checkstyle and SpotBugs parsing are implemented.
- [x] Enforcer and JaCoCo log extraction are implemented.
- [x] Text output exists.
- [x] JSON output exists.
- [x] JSON schema exists.
- [x] Demo fixtures exist.
- [x] No-failure fixture exists.
- [x] End-to-end CLI tests exist.
- [x] Golden text-output tests exist.
- [x] CI and release workflows exist.
- [x] Installation, usage, testing, CI, integration, release, and permission docs exist.
- [ ] Release tag is cut.
- [ ] Release artifacts are published.
