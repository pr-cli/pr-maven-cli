# Implementation Status

This document records what is currently implemented, validated, and intentionally deferred.

For future scope and target release windows, see [ROADMAP.md](../ROADMAP.md).

## Current State

PR Maven CLI is a production-usable local MVP for Maven report triage.

The current implementation can inspect local Maven test, quality, and selected build log artifacts, map findings to Maven modules, and emit deterministic text or JSON output through both a CLI and a Go library.

The core analyzer remains local-first and provider-agnostic. GitHub has first-party project automation and examples, but the runtime analyzer does not require GitHub APIs or tokens.

## Implemented Commands and Interfaces

CLI entrypoint:

```bash
prmaven
```

Implemented commands:

- `fails`;
- `why`;
- `version`.

Implemented flags:

- `-project`;
- `-format text|json`;
- `-module`;
- `-output`.

Library interface:

- `pkg/prmaven.Analyze`;
- `pkg/prmaven.Options`;
- `pkg/prmaven.Report`;
- `pkg/prmaven.Summary`;
- `pkg/prmaven.Module`;
- `pkg/prmaven.Finding`.

## Implemented Packages and Files

Core code:

- `cmd/prmaven`: CLI entrypoint and command behavior.
- `pkg/prmaven/analyzer.go`: analysis orchestration.
- `pkg/prmaven/pom.go`: Maven module discovery.
- `pkg/prmaven/format.go`: text and JSON formatting.
- `pkg/prmaven/path.go`: path normalization helpers.
- `pkg/prmaven/types.go`: public report model.

Fixtures and examples:

- `demo/multi-module-failure`;
- `demo/no-failure`;
- `pkg/prmaven/testdata`;
- `examples/library`;
- `examples/github-actions`.

Automation:

- `.github/workflows/ci.yml`;
- `.github/workflows/release.yml`;
- `.github/workflows/security.yml`;
- `.github/workflows/thank-contributor.yml`;
- `scripts/test.*`;
- `scripts/build.*`;
- `scripts/quality.*`.

## Implemented Behavior

### Maven Module Discovery

The analyzer discovers Maven modules from `pom.xml`, including nested module layouts.

Validated by:

- unit tests;
- nested module fixtures;
- path normalization tests.

### Test Report Parsing

The analyzer parses:

- Surefire JUnit XML reports;
- Failsafe JUnit XML reports.

Validated by:

- demo fixtures;
- parser tests;
- CLI end-to-end tests;
- golden output tests.

### Quality Report Parsing

The analyzer parses:

- Checkstyle XML reports;
- SpotBugs XML reports.

Validated by:

- sanitized fixture projects;
- focused parser tests.

### Build Log Extraction

The analyzer extracts deterministic findings from:

- Maven Enforcer log fixtures;
- JaCoCo threshold log fixtures.

Validated by:

- sanitized log fixtures;
- focused tests.

### Output Formats

The CLI emits:

- human-readable text;
- JSON;
- optional file output.

Current findings emit `high` confidence with human-readable confidence reasons because Stage 1 findings are backed by deterministic Maven report or log artifacts. The `medium` and `low` levels are documented as reserved future vocabulary.

Validated by:

- CLI tests;
- JSON schema tests;
- golden text-output tests.

### Module Filtering

The CLI supports limiting findings through `-module`, matching Maven module path or artifact id.

Validated by:

- matching tests;
- no-match tests;
- CLI behavior tests.

### OSS Project Infrastructure

The repository includes:

- README;
- manifesto;
- roadmap;
- contribution guide;
- issue templates;
- pull request template;
- maintainer and governance docs;
- CI/CD documentation;
- release documentation;
- security and permission posture docs;
- contributor thank-you automation.

## Explicitly Out of Scope Right Now

The following remain intentionally deferred:

- native GitHub runtime adapter;
- native GitLab runtime adapter;
- PR changed-file analysis;
- check-run ingestion;
- baseline comparison against `main`;
- Markdown PR comment generation;
- SARIF output;
- GitHub workflow annotations;
- provider token management;
- hosted service behavior;
- telemetry;
- automatic code fixes;
- Maven 4 production support.

## Current Source Boundary

PR Maven CLI starts from files already present on disk:

```text
pom.xml files
target/surefire-reports
target/failsafe-reports
target/checkstyle-result.xml
target/spotbugs*.xml
selected target/*.log files
```

It does not currently fetch artifacts from CI providers, inspect remote pull request metadata, or call external APIs during core analysis.

## Next Implementation Priorities

Near-term priorities:

- keep parser fixtures easy for contributors to extend;
- document fixture contribution rules;
- harden output compatibility expectations;
- keep issue labels and backlog aligned with the roadmap.

Stage 3 priorities:

- design optional GitHub adapter interfaces;
- add provider mocks and fixtures before networked behavior;
- design PR-to-module relevance scoring;
- investigate SARIF and GitHub annotations;
- document privacy guidance for CI logs.

## Verification Commands

Standard local verification:

```bash
go test ./...
```

CLI smoke tests:

```bash
go run ./cmd/prmaven fails -project demo/multi-module-failure
go run ./cmd/prmaven why -project demo/multi-module-failure -format json
go run ./cmd/prmaven fails -project demo/no-failure
```

Scripted verification:

```bash
./scripts/test.sh
./scripts/quality.sh
```

On Windows PowerShell:

```powershell
.\scripts\test.ps1
.\scripts\quality.ps1
```
