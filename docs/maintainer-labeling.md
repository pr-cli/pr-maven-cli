# Maintainer Labeling Guide

This guide helps maintainers label issues consistently so contributors can find work, understand merge order, and choose tasks with the right amount of project context.

Use it together with [Label Guide](labels.md), which is the taxonomy reference.

## Minimum Label Set

Every planned issue should have:

- one type label, such as `bug`, `enhancement`, `documentation`, `maintenance`, or `security`;
- one or more area labels, such as `area: cli`, `area: parser`, `area: test`, `area: docs`, or `area: ci`;
- exactly one stage label, such as `stage: 2`, `stage: 2.1`, or `stage: 3`;
- one status label, usually `status: ready` or `status: blocked`;
- contributor-fit labels when the issue is intended for outside help.

Keep `good first issue` and `good first contribution` together when the issue is genuinely approachable. The first label helps GitHub discovery. The second label explains the project-specific expectation.

## Contributor-Fit Labels

### `good first contribution`

Use `good first contribution` when the issue is low-context, clearly scoped, and can be completed without understanding the full roadmap.

Good fits:

- documentation examples;
- focused CLI wording improvements;
- one parser fixture with a known report shape;
- one regression test for a known edge case;
- small compatibility notes with a clear source.

Avoid this label when the issue requires architecture decisions, new public contracts, provider integration design, broad refactors, or release policy decisions.

### `help wanted`

Use `help wanted` when maintainers actively want outside implementation help.

Good fits:

- parser coverage where the expected artifact format is already known;
- CI examples that can be validated locally or in GitHub Actions;
- test coverage gaps with a clear acceptance checklist;
- documentation improvements that need contributor examples.

Avoid this label when maintainers still need to decide the design. Use `need help` for design, validation, or outside context that is not ready for direct implementation.

### `agent-friendly`

Use `agent-friendly` when the issue is narrow enough for an automated coding agent to attempt and for a maintainer to review quickly.

Good fits:

- deterministic fixtures;
- additive documentation;
- focused test cases;
- isolated CLI behavior;
- changes with explicit files, commands, and acceptance criteria.

Avoid this label when the issue needs product judgment, ambiguous tradeoffs, private context, credentials, external service setup, or broad repository restructuring.

## Scoped Issue Template

Well-scoped issues should include:

- a short problem statement;
- expected files or surfaces to touch;
- explicit acceptance criteria;
- validation commands;
- dependency or blocked status when applicable.

Example:

```md
Title: Add Checkstyle fixture for empty violation file

Labels: documentation, area: parser, area: test, stage: 2, status: ready,
good first issue, good first contribution, agent-friendly

Problem:
The Checkstyle parser should keep reporting zero findings when the report file is present but has no violations.

Scope:
- Add one sanitized fixture under pkg/prmaven/testdata.
- Add one focused parser test.
- Do not change the JSON contract.

Acceptance:
- `go test ./pkg/prmaven` passes.
- The fixture is documented when it introduces a new report shape.
- The issue closes through one focused PR.
```

## Labeling Examples

| Issue shape | Suggested labels |
| --- | --- |
| Add one Maven report parser fixture | `enhancement`, `area: parser`, `area: test`, stage label, `status: ready`, `help wanted`, `agent-friendly` |
| Improve usage docs with one new example | `documentation`, `area: docs`, stage label, `status: ready`, `good first issue`, `good first contribution`, `oss first friendly` |
| Define a future provider interface | `enhancement`, `area: architecture`, provider area label, stage label, `status: blocked`, `help wanted`, `need help` |
| Investigate CI hardening policy | `maintenance`, `area: ci`, `area: validation`, stage label, `status: blocked`, `need help` |
| Release checklist cleanup | `maintenance`, `area: release`, stage label, `status: ready` |

## Maintainer Review Checklist

Before marking an issue `status: ready`, confirm:

- the issue belongs to the current roadmap stage or is intentionally independent;
- labels match type, area, stage, status, and contributor fit;
- blocked work names the blocking issue or decision;
- acceptance criteria can be verified by a contributor;
- the issue can be closed by one focused PR.

Before merging a PR, confirm:

- PR labels still match the issue labels;
- the PR does not expand beyond the scoped issue;
- docs or tests were updated when the public surface changed;
- status and stage labels still match the roadmap order.
