# Provider Errors And Offline Fallbacks

Optional provider adapters must fail predictably and preserve local Maven evidence.

This taxonomy defines deterministic provider error categories and fallback behavior for Stage 3 design and implementation work.

## Error Categories

| Category | Meaning | Default behavior |
| --- | --- | --- |
| `missing_token` | Provider context was requested, but no required token or credential was configured. | Degrade gracefully and keep local Maven findings. |
| `insufficient_permissions` | A token exists but cannot read the requested PR, merge request, check run, pipeline, or artifact metadata. | Degrade gracefully, keep local Maven findings, and report the permission gap. |
| `network_failure` | Provider API or CI service could not be reached. | Degrade gracefully and avoid unbounded retries in default commands/tests. |
| `rate_limited` | Provider returned a rate-limit response or equivalent throttling state. | Degrade gracefully and include retry context when available. |
| `not_found` | The requested repository, PR, merge request, check run, pipeline, job, or artifact was not found. | Degrade gracefully for optional context; do not invent findings. |
| `unsupported_state` | Provider returned a valid state the adapter does not support yet. | Emit a bounded diagnostic and keep report-only output usable. |
| `invalid_fixture` | A committed provider fixture is malformed or does not match the contract. | Fail the focused fixture test. |
| `invalid_configuration` | Mutually incompatible provider flags or options were supplied. | Return a user-facing CLI/configuration error before provider calls. |

## Graceful Degradation Rules

Provider failures should not erase local Maven findings.

When provider context fails:

- keep parsed Maven findings, modules, summary counts, report paths, and reproduction commands;
- mark provider context as unavailable or partial only when that output contract exists;
- include a concise diagnostic when the user opted into provider-aware behavior;
- avoid converting missing provider context into a Maven finding;
- avoid changing report-only exit codes unless the user explicitly requested provider context as a required input.

## Reporting Expectations

Future JSON extensions should represent provider failures as diagnostics or metadata, not as Maven findings.

Expected fields should be additive and may include:

- provider name;
- operation name;
- normalized error category;
- whether local Maven findings were preserved;
- whether provider context is unavailable or partial;
- optional retry or permission hint when safe to expose.

Human-readable output should keep local findings first. Provider diagnostics should appear as a short secondary section or explicit command output, never as a replacement for local Maven evidence.

## Test Expectations

The [Provider Fixtures And Mocks](provider-fixtures-and-mocks.md) contract should cover at least:

- missing token;
- insufficient permissions;
- network failure;
- rate limit;
- not found;
- unsupported state;
- successful provider context;
- partial provider context.

Default `go test ./...` must not require live provider credentials or network access.

## Local-First Rule

The local analyzer remains authoritative for Maven report-backed findings.

Provider context may add relevance and explanation, but optional provider errors must not prevent:

- report parsing;
- JSON report generation;
- text report generation;
- module filtering;
- output-file writing;
- no-failure output.
