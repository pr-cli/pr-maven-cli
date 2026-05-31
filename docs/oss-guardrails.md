# OSS Guardrails

This document summarizes the guardrails maintainers should preserve as PR Maven CLI grows.

The project follows a conservative OSS posture inspired by mature Maven ecosystem practices:

- every issue and pull request should be labeled for type and triage;
- changes should flow through pull requests;
- the default branch should be protected from force pushes and deletion;
- tests and documentation should accompany behavior changes;
- contributor tasks should stay scoped and reviewable.

## Current Hard Gates

- `main` is the default branch.
- `main` requires the `All CI checks` status before merge.
- `main` blocks force pushes and branch deletion.
- Conversation resolution is required before merge.
- Repository workflow token permissions default to `contents: read`.
- Core analysis must not require network access, provider tokens, telemetry, AI services, or hosted APIs.

## Contributor Safety

- Use [Label Guide](labels.md) for all planned issues and pull requests.
- Keep issue bodies explicit about dependencies and acceptance criteria.
- Keep `status: blocked` issues blocked until their dependency is closed or intentionally waived.
- Do not merge contributor PRs without a relevant type label, area label, and stage label when the PR belongs to the roadmap.
- Prefer fixture-based tests for parser and CLI behavior.

## Workflow Safety

- Workflows that run contributor code must use minimal permissions.
- `pull_request_target` workflows must never check out or execute contributor code.
- Comment-only automations may read base-repository templates and write comments.
- Release publishing jobs may receive `contents: write`; other jobs should stay read-only unless a specific need is documented.

## Local-First Product Guardrail

Core analysis must remain usable without:

- GitHub or GitLab tokens;
- CI provider APIs;
- network access;
- telemetry;
- AI provider calls;
- external services.

Stage 3 provider adapters must be additive. Provider failures should not erase local Maven evidence.

## Maintainer Expansion Gate

Before adding maintainers or collaborators with write access, maintainers should:

- require at least one pull request review;
- consider requiring code owner review for production code and workflow changes;
- confirm only maintainers can merge;
- confirm `All CI checks` remains the required protected status;
- confirm no unexpected secrets, variables, webhooks, environments, deployments, or Pages configuration exists.

## Release Guardrails

- Keep release tags aligned with the roadmap release line.
- Preserve checksum generation for binary artifacts.
- Keep JSON compatibility changes additive unless a breaking change is explicitly planned.
- Document release readiness in `docs/release.md` and the MVP/release checklists.

## Deferred Hardening

These controls are intentionally deferred until they fit the maintainer model:

- required pull request review while the project has only one maintainer;
- code owner review requirement;
- GitHub Actions allowlist or SHA pinning;
- signed tags, provenance, or SBOM generation.

They should be reconsidered before adding more maintainers or before a larger public contribution push.
