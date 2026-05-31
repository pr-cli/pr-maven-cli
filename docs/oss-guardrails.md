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
- `main` requires at least one approving pull request review before merge.
- `main` requires code owner review before merge.
- `main` blocks force pushes and branch deletion.
- Pull request head branches are deleted after merge.
- Conversation resolution is required before merge.
- Repository workflow token permissions default to `contents: read`.
- `All CI checks` includes a public metadata guard that blocks branch names, pull request titles, and commit messages containing coding agent or tool names.
- Secret scanning, secret scanning push protection, and Dependabot security updates are enabled for the public repository.
- Core analysis must not require network access, provider tokens, telemetry, AI services, or hosted APIs.

## Contributor Safety

- Use [Label Guide](labels.md) for all planned issues and pull requests.
- Keep issue bodies explicit about dependencies and acceptance criteria.
- Keep `status: blocked` issues blocked until their dependency is closed or intentionally waived.
- Do not merge contributor PRs without a relevant type label, area label, and stage label when the PR belongs to the roadmap.
- Keep public metadata product-focused: branch names, pull request titles, commit messages, and merge messages should not include coding agent or tool names.
- Prefer fixture-based tests for parser and CLI behavior.

## Workflow Safety

- Workflows that run contributor code must use minimal permissions.
- `pull_request_target` workflows must never check out or execute contributor code.
- Comment-only automations may read base-repository templates and write comments.
- Release package jobs may receive `id-token: write` and `attestations: write` only to generate GitHub artifact and SBOM attestations.
- Release publishing jobs may receive `contents: write`; other jobs should stay read-only unless a specific need is documented.
- Required CI should keep the public metadata guard in the aggregate `All CI checks` gate.

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

- confirm required pull request review and code owner review are still enabled;
- remove any founder-only review bypass once another trusted maintainer or code owner can review founder-authored PRs;
- confirm only maintainers can merge;
- confirm `All CI checks` remains the required protected status;
- confirm no unexpected secrets, variables, webhooks, environments, deployments, or Pages configuration exists.

## Release Guardrails

- Keep release tags aligned with the roadmap release line.
- Prefer signed annotated release tags created by the founder or a trusted release maintainer.
- Preserve checksum generation for binary artifacts.
- Preserve SPDX JSON SBOM generation for release packages.
- Preserve GitHub artifact attestation generation for release packages and checksums.
- Preserve GitHub SBOM attestation generation for release packages.
- Keep JSON compatibility changes additive unless a breaking change is explicitly planned.
- Document release readiness in `docs/release.md` and the MVP/release checklists.

## Deferred Hardening

These controls are intentionally deferred until they fit the maintainer model:

- GitHub Actions allowlist or SHA pinning;
- enforced signed-tag rulesets;
- SBOM vulnerability policy and minimum severity gates.

They should be reconsidered before adding more maintainers or before a larger public contribution push.
