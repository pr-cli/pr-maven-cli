# Provider Fixtures And Mocks

Stage 3 provider behavior must be testable without live GitHub, GitLab, CI provider, or hosted-service access.

This document defines the fixture and fake-client contract for optional provider context.

Changed-files fixtures have a dedicated contract in [Changed-Files Fixture Contract](changed-files-fixture-contract.md). Check-run fixtures have a dedicated contract in [Check-Runs Fixture Contract](check-runs-fixture-contract.md).

## Fixture Principles

Provider fixtures should be:

- sanitized and safe to publish;
- stable across machines and CI workers;
- small enough to review in pull requests;
- close enough to provider payloads to preserve useful field names;
- normalized where provider APIs include volatile IDs, timestamps, URLs, or actor metadata;
- stored under a provider-specific `testdata` directory once implementation begins.

Default tests must not call live provider APIs.

## Suggested Fixture Shape

Provider fixtures should use JSON unless a provider artifact is naturally another format.

Recommended layout:

```text
pkg/prmaven/provider/testdata
  github
    changed-files
      added-modified-renamed.json
      deleted-files.json
      nested-modules.json
    check-runs
      success.json
      failure.json
      skipped-cancelled-timed-out.json
      pending.json
    errors
      missing-token.json
      rate-limit.json
      permission-denied.json
      not-found.json
  gitlab
    merge-request-files
    pipeline-jobs
```

Each fixture should include a short README or inline test name that explains:

- the scenario being modeled;
- the expected normalized provider-neutral result;
- any fields intentionally omitted or redacted.

## Fake Client Contract

Provider fake clients should support deterministic responses for:

- successful changed-file lookup;
- successful check-run or pipeline summary lookup;
- missing token;
- insufficient permissions;
- rate limit;
- network failure;
- not found;
- unsupported provider state;
- partial provider data.

Fakes should not read environment variables unless a test explicitly validates token-resolution behavior. Tests that validate token behavior should set variables inside the test and restore them before returning.

## Error Scenario Expectations

Provider failures should degrade gracefully.

The canonical error categories are defined in [Provider Errors And Offline Fallbacks](provider-errors-and-fallbacks.md).

Expected behavior:

- missing token: provider context is unavailable, local Maven findings remain;
- insufficient permissions: include a provider diagnostic, local Maven findings remain;
- rate limit: include retry/context information when available, local Maven findings remain;
- network failure: do not retry indefinitely in default tests, local Maven findings remain;
- not found: treat provider context as unavailable for that target, local Maven findings remain;
- unsupported state: report a bounded diagnostic rather than inventing findings.

## Test Layers

Recommended test layers:

1. Provider-neutral unit tests over fake clients.
2. Fixture parser/normalizer tests using sanitized provider payloads.
3. CLI wiring tests that prove report-only mode still works without provider credentials.
4. Optional live-provider checks that are skipped by default and never required for `go test ./...`.

## Compatibility With Local-First Core

Provider fixtures and mocks must not make provider context part of the core analyzer contract.

The following must continue to work without provider fixtures, mocks, credentials, or network access:

- Maven module discovery;
- report parsing;
- text output;
- JSON output;
- output-file writing;
- module filtering;
- no-failure behavior.
