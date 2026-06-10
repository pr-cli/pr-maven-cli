# Check-Runs Fixture Contract

This document defines the Stage 3 fixture contract for CI check-run metadata.

The contract exists before a live GitHub check-runs adapter is implemented so contributors can add parser, normalizer, fake-client, and reporting tests without provider tokens or network access.

Use this document with the [Provider Context Plan](provider-context-plan.md), [Provider Fixtures And Mocks](provider-fixtures-and-mocks.md), [Provider Adapter Package Boundaries](provider-adapter-boundaries.md), [GitHub Provider Permissions](github-provider-permissions.md), and [Changed-Files Fixture Contract](changed-files-fixture-contract.md).

Reference GitHub documentation:

- [REST API endpoints for check runs](https://docs.github.com/en/rest/checks/runs)

## Goals

- Define a sanitized JSON shape for check-run fixtures.
- Cover success, failure, skipped, cancelled, timed out, and pending states.
- Preserve enough provider-like fields to test normalizers without storing volatile provider data.
- Define how check-run metadata enriches Maven findings without replacing report-backed evidence.
- Keep default tests provider-token-free and network-free.

## Fixture Location

When implementation begins, check-run fixtures should live under:

```text
pkg/prmaven/provider/testdata/github/check-runs
```

Recommended first fixtures:

```text
success.json
failure.json
skipped-cancelled-timed-out.json
pending.json
mixed-checks.json
```

Each fixture should be small, reviewable, and safe to publish.

## Sanitized Fixture Shape

Check-run fixtures should use JSON.

Recommended top-level shape:

```json
{
  "provider": "github",
  "resource": "check_runs",
  "schemaVersion": 1,
  "scenario": "mixed check runs for a pull request head sha",
  "repository": {
    "owner": "example-org",
    "name": "example-maven-project",
    "visibility": "public"
  },
  "pullRequest": {
    "number": 123,
    "baseRef": "main",
    "headRef": "feature/module-change",
    "headSha": "0000000000000000000000000000000000000000"
  },
  "checkRuns": [
    {
      "name": "Go tests",
      "status": "completed",
      "conclusion": "success",
      "workflowName": "CI",
      "startedAt": "2026-01-01T00:00:00Z",
      "completedAt": "2026-01-01T00:02:00Z",
      "detailsUrl": "https://example.invalid/checks/go-tests",
      "summary": "All tests passed"
    },
    {
      "name": "Maven test",
      "status": "completed",
      "conclusion": "failure",
      "workflowName": "CI",
      "startedAt": "2026-01-01T00:00:00Z",
      "completedAt": "2026-01-01T00:03:00Z",
      "detailsUrl": "https://example.invalid/checks/maven-test",
      "summary": "Tests failed"
    },
    {
      "name": "Optional docs",
      "status": "completed",
      "conclusion": "skipped",
      "workflowName": "Docs",
      "startedAt": null,
      "completedAt": "2026-01-01T00:01:00Z",
      "detailsUrl": "https://example.invalid/checks/docs",
      "summary": "Skipped by path filter"
    },
    {
      "name": "Integration",
      "status": "completed",
      "conclusion": "cancelled",
      "workflowName": "CI",
      "startedAt": "2026-01-01T00:00:00Z",
      "completedAt": "2026-01-01T00:01:00Z",
      "detailsUrl": "https://example.invalid/checks/integration",
      "summary": "Cancelled by newer run"
    },
    {
      "name": "Long running test",
      "status": "completed",
      "conclusion": "timed_out",
      "workflowName": "CI",
      "startedAt": "2026-01-01T00:00:00Z",
      "completedAt": "2026-01-01T00:30:00Z",
      "detailsUrl": "https://example.invalid/checks/long-running-test",
      "summary": "Timed out"
    },
    {
      "name": "Queued quality gate",
      "status": "queued",
      "conclusion": null,
      "workflowName": "CI",
      "startedAt": null,
      "completedAt": null,
      "detailsUrl": "https://example.invalid/checks/quality-gate",
      "summary": "Waiting for runner"
    }
  ],
  "expected": {
    "total": 6,
    "completed": 5,
    "pending": 1,
    "failing": ["Maven test", "Long running test"],
    "nonBlocking": ["Optional docs", "Integration"],
    "providerContextState": "partial"
  }
}
```

## Required Fields

Each fixture must include:

- `provider`: provider name, initially `github`;
- `resource`: stable fixture resource name, initially `check_runs`;
- `schemaVersion`: integer contract version;
- `scenario`: short human-readable scenario;
- `checkRuns`: check-run entries;
- `expected`: normalized expectations used by tests.

Each check-run entry must include:

- `name`: stable display name;
- `status`: provider status such as `queued`, `in_progress`, `completed`, `waiting`, `requested`, or `pending`;
- `conclusion`: final result for completed checks, or `null` while pending;
- `workflowName`: workflow or grouping label when available;
- `startedAt`: timestamp or `null`;
- `completedAt`: timestamp or `null`;
- `detailsUrl`: sanitized URL or `null`;
- `summary`: short sanitized summary or `null`.

## Required Scenarios

The first check-runs fixture set should cover:

| Scenario | Required status/conclusion coverage | Expected behavior |
| --- | --- | --- |
| Passing checks | `completed` + `success` | Provider context can confirm CI health but should not create Maven findings. |
| Failing checks | `completed` + `failure` | Provider context can point to failing CI, but Maven findings still require local report evidence. |
| Skipped checks | `completed` + `skipped` | Skipped checks should be reported as non-blocking context unless a future issue defines stricter behavior. |
| Cancelled checks | `completed` + `cancelled` | Cancelled checks should be diagnostic context, not Maven evidence. |
| Timed-out checks | `completed` + `timed_out` | Timed-out checks can explain incomplete CI, but should not invent test failures. |
| Pending checks | `queued`, `in_progress`, or `pending` with `conclusion: null` | Pending checks should mark provider context as partial. |

## Sanitization Rules

Fixtures must not include:

- real private repository names;
- real branch names from private repositories;
- real commit SHAs, check-run IDs, node IDs, or database IDs;
- real actor names, emails, tokens, or account IDs;
- signed URLs;
- artifact URLs;
- volatile API URLs;
- raw log output containing private code or secrets.

Use stable placeholder values such as `example-org`, `example-maven-project`, `feature/module-change`, and zeroed SHAs.

## Maven Evidence Boundary

Check-run metadata enriches context. It does not replace Maven report evidence.

Provider-aware output may use check-run metadata to say:

- which CI checks are failing, pending, skipped, cancelled, or timed out;
- whether provider context is complete or partial;
- which check name or workflow likely corresponds to Maven validation;
- where a user might inspect CI details.

Provider-aware output must not:

- create a Maven finding from a failing check alone;
- change the source report path of a local Maven finding;
- hide local Maven findings when check-run context is missing;
- change report-only exit behavior unless a future issue explicitly scopes that behavior.

## Fake Client Expectations

Fake clients for check runs should support:

- successful response with completed checks;
- successful response with pending checks;
- empty check-run response;
- paginated response behavior without live network calls;
- missing token;
- insufficient permissions;
- rate limit;
- not found;
- unsupported status or conclusion.

Error behavior should follow [Provider Errors And Offline Fallbacks](provider-errors-and-fallbacks.md).

## Test Expectations

Default tests should:

- load committed check-run fixtures from disk;
- validate required fixture fields;
- normalize statuses and conclusions;
- assert expected provider context state;
- assert that failed check runs do not create Maven findings by themselves;
- run without GitHub credentials;
- run without network access;
- avoid depending on real GitHub API responses.

Optional live-provider tests may exist later, but they must be skipped by default and must not be required by `go test ./...`.

## Compatibility With Future JSON Output

Check-run context should be additive.

Future JSON fields should preserve existing local report fields and add provider context under a separate provider-aware extension. Missing, partial, pending, cancelled, or timed-out check-run data should be represented as provider diagnostics or metadata rather than Maven findings.
