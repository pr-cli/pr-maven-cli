# Testing

PR Maven CLI treats tests as part of the public product contract.

The Stage 1 test system is designed for:

- contributors working locally;
- maintainers reviewing pull requests;
- automated coding agents proposing scoped changes;
- CI runs on Linux, Windows, and macOS.

## Local Commands

For release and Stage 3 readiness reviews, use the [MVP acceptance checklist](mvp-acceptance.md) alongside the commands below.

Run the standard suite:

```bash
go test ./...
```

On Windows PowerShell:

```powershell
.\scripts\test.ps1
.\scripts\test.ps1 -Race -Coverage
```

On Unix-like shells:

```bash
./scripts/test.sh
PRMAVEN_RACE=1 PRMAVEN_COVERAGE=1 ./scripts/test.sh
```

With `make`:

```bash
make test
make test-race
make coverage
make coverage-check
make ci
```

Run the focused JSON schema validation:

```bash
go test ./pkg/prmaven -run TestGeneratedJSONReportsValidateAgainstSchema -v
```

## Test Layers

### Library Tests

Package: `pkg/prmaven`.

Coverage includes:

- local-first regression guards that exercise core analysis without provider tokens;
- static guards against core network/provider-client imports and provider environment reads;
- Maven module discovery from `pom.xml`;
- Surefire report parsing;
- Failsafe report parsing;
- Checkstyle report parsing;
- SpotBugs report parsing;
- Maven Enforcer log parsing;
- JaCoCo threshold log parsing;
- report-to-module mapping;
- slash-separated path normalization for JSON output;
- reproduction command generation;
- JSON output contract;
- generated demo JSON validation against `schema/prmaven-report.schema.json`;
- text output snapshots;
- missing project error behavior.

### CLI End-to-End Tests

Package: `cmd/prmaven`.

Coverage includes:

- building the real CLI binary in a temporary directory for acceptance testing;
- running the failure and no-failure demo workspaces through the compiled CLI;
- `fails` text output;
- `why` JSON output;
- non-zero exit when findings exist;
- zero exit when no findings exist;
- parseable JSON output for both finding and no-finding workflows;
- output-file behavior for text and JSON with temporary test files;
- stdout behavior when `-output` is absent;
- module filtering by artifactId, module path, no-match behavior, and filtered JSON findings;
- invalid command and invalid format handling.

### Documented Command Smoke Tests

Documented command smoke tests live in `cmd/prmaven/doc_commands_test.go`.

Run the focused suite after changing README command examples, usage docs, installation docs, or CLI flags:

```bash
go test ./cmd/prmaven -run TestDocumentedCommandSmokeSuite -v
```

The suite uses a compiled local `prmaven` binary and runs from the repository root so demo paths such as `demo/multi-module-failure` match the public docs. It should cover representative safe local commands only.

When maintainers update documented commands:

- add or update the matching table entry in `TestDocumentedCommandSmokeSuite`;
- keep commands local-first and dependency-light;
- use demo fixtures or temporary output files instead of a real user workspace;
- do not include publish, push, tag, release, provider-token, package-upload, or remote API commands in this smoke suite.

### Demo Fixtures

Fixtures live under `demo/`.

- `demo/multi-module-failure`: Maven aggregator with Surefire and Failsafe findings.
- `demo/no-failure`: Maven aggregator with passing Surefire report output.
- `pkg/prmaven/testdata/checkstyle-project`: Maven aggregator with a sanitized Checkstyle report fixture.
- `pkg/prmaven/testdata/spotbugs-project`: Maven aggregator with a sanitized SpotBugs report fixture.
- `pkg/prmaven/testdata/enforcer-project`: Maven aggregator with a sanitized Maven Enforcer log fixture.
- `pkg/prmaven/testdata/jacoco-project`: Maven aggregator with a sanitized JaCoCo threshold log fixture.
- `pkg/prmaven/testdata/nested-module-project`: Maven aggregator with a nested module and Surefire report fixture.

The `target/*-reports` directories and selected `target/*.log` files are intentionally versioned because they are stable test fixtures, not local build output.

Fixture compatibility notes, including the Maven 3.9.x production baseline and Maven 4 tracking boundary, are documented in [Fixture Notes](fixtures.md).

### Fixture Integrity

Fixture integrity is validated by `TestFixtureIntegrity` in `pkg/prmaven`.

Run the focused validation after adding, removing, or renaming committed fixture files:

```bash
go test ./pkg/prmaven -run TestFixtureIntegrity -v
```

The test checks that expected demo and `testdata` files exist, that intentionally committed `target` report and log artifacts remain present, and that unexpected generated files are not added under fixture `target` directories.

### Golden Files

Golden files live under:

```text
pkg/prmaven/testdata/golden
```

They protect human-readable text output and machine-readable JSON output from accidental changes.

Run the focused golden snapshot validation:

```bash
go test ./pkg/prmaven -run TestWrite.*GoldenFiles -v
```

When output changes intentionally, update the affected golden file in the same PR and explain the reason. JSON snapshots must normalize `projectRoot` to `<PROJECT_ROOT>/demo/...` so the files remain stable across machines and CI workers.

## CI

GitHub Actions runs:

- Go tests on Linux, Windows, and macOS;
- Go 1.22.x and the current stable Go release;
- race detector on Linux;
- coverage generation on Linux;
- fixture integrity validation for committed demo and testdata artifacts;
- a minimum total coverage gate of 70%.

The CI workflow is intentionally dependency-light. Core tests do not require Maven, network services, GitHub tokens, Docker, or external APIs. The `pkg/prmaven` local-first guard fails if production core code starts importing network/provider clients or reading provider environment variables.

## Contributor Expectations

For parser changes:

- add a sanitized fixture;
- add a focused unit test;
- update golden files if text output changes;
- keep JSON fields stable unless the issue explicitly allows a compatibility change.

For CLI changes:

- add or update end-to-end tests;
- document exit code changes;
- update README examples if user-facing commands change.

For documentation-only changes:

- tests are usually not required;
- keep examples aligned with the current CLI behavior.
