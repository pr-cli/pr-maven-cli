# Provider Adapter Package Boundaries

Stage 3 provider adapters must enrich local Maven findings without coupling provider-specific clients to the core analyzer.

This document defines the intended package boundaries before any native GitHub or GitLab adapter lands.

## Boundary Goals

- Preserve `pkg/prmaven` as the local Maven report analyzer.
- Keep provider network clients out of core parsing, mapping, formatting, and JSON report generation.
- Make provider context optional and injectable.
- Keep default tests runnable without provider tokens, network access, hosted services, or CI APIs.
- Allow provider implementations to evolve without changing the local report contract.

## Proposed Package Shape

The exact names can evolve during implementation, but the ownership boundary should stay stable:

```text
pkg/prmaven
  Core analyzer, Maven report parsers, report model, formatters, module mapping.

pkg/prmaven/provider
  Provider-neutral interfaces, request/response contracts, errors, and fake clients.

pkg/prmaven/provider/github
  Optional GitHub implementation behind provider interfaces.

pkg/prmaven/provider/gitlab
  Optional GitLab implementation, if the GitLab parity boundary is accepted.

cmd/prmaven
  CLI wiring, flags, command UX, dependency injection, and output selection.
```

## Import Rules

Allowed:

- `cmd/prmaven` may import `pkg/prmaven` and optional provider packages for command wiring.
- Provider implementations may import provider-neutral interfaces.
- Provider implementations may import their own client dependencies when implementation issues explicitly allow them.
- Tests may import fakes and fixtures for deterministic provider behavior.

Not allowed:

- `pkg/prmaven` must not import `pkg/prmaven/provider/github` or `pkg/prmaven/provider/gitlab`.
- Core parser packages must not import GitHub, GitLab, HTTP client, CI provider, telemetry, or AI provider code.
- Core analyzer execution must not read provider tokens or create network clients.
- JSON/text report formatting must not require provider context to render report-only findings.

## Provider Interface Expectations

Provider-neutral interfaces should be narrow and evidence-oriented. They should represent data PR Maven CLI needs, not the full provider API.

Expected interface families:

- changed files for a pull request or merge request;
- check runs, jobs, or pipeline summaries;
- provider identity and permission diagnostics;
- provider error classification.

Interfaces should return deterministic structures that can be backed by:

- sanitized JSON fixtures;
- fake clients in tests;
- optional live clients in manually enabled integration checks.

## Testability Expectations

Every provider implementation should have:

- fake-client tests for success and failure paths;
- fixture tests for sanitized changed-files and check-run payloads;
- missing-token and insufficient-permission tests;
- rate-limit, network-failure, not-found, and unsupported-state tests;
- report-only regression tests proving provider failures do not erase local Maven findings.

Default `go test ./...` must not require live provider credentials or network access.

## Local-First Preservation

The local analyzer remains the base layer. Provider context can add relevance, grouping, or explanation, but it cannot become required to:

- discover Maven modules;
- parse Surefire, Failsafe, Checkstyle, SpotBugs, Enforcer, or JaCoCo artifacts;
- emit text output;
- emit JSON output;
- write output files;
- filter by module.

When provider context is unavailable, PR Maven CLI should still return the local report-backed result and include provider diagnostics only when the user opted into provider-aware behavior.
