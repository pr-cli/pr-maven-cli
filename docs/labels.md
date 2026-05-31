# Label Guide

This project uses labels as release-note, triage, and contributor-navigation metadata.

Every public issue and pull request should have at least:

- one type label;
- one area label;
- one stage label when the work belongs to the roadmap;
- one status label when it is part of the planned backlog.

## Type Labels

Use GitHub-standard type labels where possible:

- `bug`: incorrect behavior, regressions, or broken examples.
- `enhancement`: user-facing or maintainer-facing improvements.
- `documentation`: documentation-only changes.
- `maintenance`: dependency, CI, release, or repository upkeep.
- `security`: security controls, vulnerability handling, or supply-chain hardening.

This mirrors the mature OSS convention used by Apache Maven, where issue and pull request labels classify work for release notes and contributor discovery.

## Area Labels

Use area labels to show the affected surface:

- `area: cli`
- `area: parser`
- `area: test`
- `area: docs`
- `area: ci`
- `area: release`
- `area: architecture`
- `area: validation`
- `area: github`
- `area: gitlab`
- `area: maven`
- `area: agent`

## Contributor Labels

Use these labels to signal contribution fit:

- `good first contribution`: low-context contribution with clear acceptance criteria.
- `oss first friendly`: suitable for a first OSS pull request.
- `help wanted`: maintainers are actively inviting implementation help.
- `need help`: maintainers need outside context, testing, or design review.
- `waiting help`: blocked on contributor, maintainer, or external feedback.
- `agent-friendly`: well-scoped enough for an automated coding agent to attempt.

## Status Labels

Use status labels to manage ordering:

- `status: ready`: contributors can start.
- `status: blocked`: wait for a listed dependency, design issue, or readiness gate.

Blocked issues should name their dependency in the issue body.

## Stage Labels

Use stage labels to keep merges aligned with the roadmap:

- `stage: 1`
- `stage: 2`
- `stage: 2.1`
- `stage: 3`

Every planned roadmap issue should carry exactly one stage label.

## Maintainer Checks

Before merging a pull request:

- confirm the PR has at least one type label;
- confirm it has the relevant area label;
- confirm roadmap work has the matching stage label;
- confirm blocked issues were unblocked before implementation started;
- add or update a milestone when the release target is known.
