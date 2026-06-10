# GitHub Provider Permissions

This document defines the planned permission posture for the optional Stage 3 GitHub provider adapter.

PR Maven CLI remains local-first. The core Maven report analyzer does not need a GitHub token, does not call GitHub APIs, and must continue to work from local report artifacts only.

Use this document with the [Provider Context Plan](provider-context-plan.md), [Provider Adapter Package Boundaries](provider-adapter-boundaries.md), [Provider Fixtures And Mocks](provider-fixtures-and-mocks.md), and [Provider Errors And Offline Fallbacks](provider-errors-and-fallbacks.md).

Reference GitHub documentation:

- [Permissions required for fine-grained personal access tokens](https://docs.github.com/en/rest/authentication/permissions-required-for-fine-grained-personal-access-tokens)
- [Use GITHUB_TOKEN for authentication in workflows](https://docs.github.com/en/actions/tutorials/authenticate-with-github_token)

## Permission Principles

- Prefer read-only permissions.
- Keep provider access optional and explicit.
- Never require provider credentials for report-only Maven analysis.
- Do not request broad repository write permissions for initial provider context.
- Keep token handling out of core parser, mapper, formatter, and report packages.
- Degrade gracefully when provider context is unavailable.

## Planned Operations

| Planned operation | Purpose | Fine-grained token permission | GitHub Actions permission | Required for public repos | Required for private repos |
| --- | --- | --- | --- | --- | --- |
| Read pull request metadata | Identify PR number, base/head refs, and changed-file pagination context. | `Pull requests: read` | `pull-requests: read` | Optional when public API access is enough, but useful for rate limits and consistency. | Yes. |
| Read changed files for a pull request | Map changed paths back to Maven modules and explain relevance. | `Pull requests: read` | `pull-requests: read` | Optional when public API access is enough. | Yes. |
| Read check runs or check suites | Correlate local Maven findings with failing CI checks. | `Checks: read` | `checks: read` | Optional when public API access is enough. | Yes, when check-run context is requested. |
| Read workflow runs, jobs, logs, or artifacts | Find CI artifact locations or failing workflow context when that behavior is explicitly added. | `Actions: read` | `actions: read` | Optional when public API access is enough. | Yes, when workflow context is requested. |
| Read commit statuses | Support repositories that still expose status contexts outside Checks API. | `Commit statuses: read` | `statuses: read` | Optional when public API access is enough. | Yes, when status context is requested. |
| Read repository contents | Load provider-side config or map files only if a future issue explicitly scopes that behavior. | `Contents: read` | `contents: read` | Optional and not part of the initial changed-files adapter. | Only if that behavior is implemented. |
| Read repository metadata | Resolve repository identity and basic visibility details. | `Metadata: read` | Covered by the workflow token baseline. | Usually available for public repositories. | Yes. |

The initial GitHub adapter should not create comments, annotations, statuses, labels, branches, commits, releases, or merges.

## Public Repository Behavior

For public repositories, provider context should support a no-token or anonymous mode when GitHub permits the requested read operation.

No-token public behavior must still be treated as best effort:

- rate limits may be lower;
- some endpoints may require authentication even for public data;
- missing provider context must not hide local Maven findings;
- diagnostics should explain that provider context was unavailable or partial.

The local analyzer path remains fully supported without provider context.

## Private Repository Behavior

For private repositories, GitHub provider context requires an explicit token with the narrowest read permissions needed for the requested operation.

Recommended minimums:

- changed-files context: `Pull requests: read` and `Metadata: read`;
- check-run context: `Checks: read` and `Metadata: read`;
- workflow-run or artifact context: `Actions: read` and `Metadata: read`;
- legacy status context: `Commit statuses: read` and `Metadata: read`;
- repository-content reads: `Contents: read` only when the feature explicitly needs repository files.

Private repository support must not make provider credentials mandatory for local report analysis.

## No-Token Local Analyzer Behavior

These commands must continue to work without GitHub credentials, provider environment variables, network access, or live API fixtures:

- report discovery;
- Maven module discovery;
- Surefire and Failsafe parsing;
- Checkstyle, SpotBugs, and JaCoCo parsing;
- text output;
- JSON output;
- report-only no-failure behavior;
- output file writing.

When provider context is requested but no usable token is available, the adapter should follow the fallback rules in [Provider Errors And Offline Fallbacks](provider-errors-and-fallbacks.md):

- keep local Maven findings;
- mark provider context as unavailable or partial when the output contract supports it;
- emit a concise diagnostic for provider-aware modes;
- avoid converting provider errors into Maven findings.

## Token Handling Expectations

Future GitHub provider implementation should:

- accept credentials through explicit provider wiring, CLI configuration, or well-documented CI environment behavior;
- never log token values;
- avoid storing tokens in fixtures, snapshots, or generated output;
- keep default tests credential-free;
- avoid reading provider environment variables from core analyzer packages;
- keep live-provider checks opt-in and skipped by default.

## Non-Goals For The Initial Adapter

The initial GitHub provider adapter should not:

- post pull request comments;
- update commit statuses;
- create check-run annotations;
- approve, review, label, close, reopen, or merge pull requests;
- upload artifacts;
- read repository secrets;
- require write-scoped tokens;
- require a GitHub token for local Maven analysis.

Any write operation must be introduced by a separate issue with explicit security, permission, test, and maintainer-review criteria.
