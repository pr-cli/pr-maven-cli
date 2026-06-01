# MVP Hardening Plan

Stage 2.1 is the MVP hardening lane between Maven signal expansion and Stage 3 provider-context work.

The goal is to keep the production-usable local MVP stable while contributors continue to add validation, fixtures, documentation, and CI coverage.

This plan is tracked by issue [#65](https://github.com/pr-cli/pr-maven-cli/issues/65) and should be read together with the [MVP acceptance checklist](mvp-acceptance.md), [Final Testable Delivery](final-testable-delivery.md), [Testing](testing.md), [CI/CD](ci.md), and the [Project visual map](project-visual-map.md).

## Guardrails

- Keep the core analyzer local-first and provider-agnostic.
- Add validation before adding provider adapters.
- Add fixtures and deterministic snapshots before expanding public output.
- Keep Stage 3 implementation-heavy issues blocked until readiness gates are closed.
- Prefer small issues that can be reviewed and merged independently.
- Keep every planned issue tied to a stage label, area label, status label, and contributor-fit label.

## Issue Requirements

Every Stage 2.1 issue should include:

- `stage: 2.1`;
- at least one `area:*` label;
- exactly one current `status:*` label;
- a contributor-fit label such as `good first contribution`, `help wanted`, `need help`, `oss first friendly`, or `agent-friendly`;
- explicit dependency notes when the issue is blocked.

Blocked issues should stay blocked until their dependency is closed or explicitly waived by a maintainer. Once dependencies are closed, the issue can move to `status: ready`.

## Suggested Merge Order

1. Close the root planning and deterministic-docs lane: #61, #62, and #65.
2. Finish ready validation foundations: #70, #71, #69, and #72.
3. Finish work unblocked by existing checklists and E2E coverage: #67, #74, #76, and #77.
4. Move validated checks into CI: #73 after #70.
5. Close the Stage 3 readiness gate: #79 after #71 and the release/MVP checklist work.
6. Review Stage 3 blocked labels: #102 after #79.
7. Start Stage 3 design issues before implementation-heavy provider work.

## Dependency Map

| Issue | Purpose | Dependency | Current lane |
| --- | --- | --- | --- |
| #61 | Deterministic text and JSON terminology | None | Completed documentation foundation |
| #62 | Deterministic context principles | #61 wording alignment | Completed architecture documentation |
| #65 | MVP hardening plan | #61 and #62 context | Root Stage 2.1 plan |
| #66 | Final MVP acceptance checklist | None | Completed acceptance foundation |
| #67 | Documentation command smoke tests | #66 | Documentation validation |
| #68 | End-to-end CLI acceptance suite | #66 | Completed E2E foundation |
| #69 | Golden JSON output snapshots | #9 completed | Ready validation work |
| #70 | Fixture integrity validation | None | Ready validation work |
| #71 | JSON schema validation in CI | #9 completed | Ready CI validation work |
| #72 | No-failure regression matrix | #13 completed | Ready validation work |
| #73 | Run fixture integrity validation in CI | #70 | Blocked until fixture validation exists |
| #74 | Output file E2E matrix | #15 completed, #68 | E2E validation follow-up |
| #75 | Release readiness validation checklist | #66 | Completed release foundation |
| #76 | Module filter E2E matrix | #16 completed, #68 | E2E validation follow-up |
| #77 | PR template MVP validation checklist | #66 | Maintainer workflow follow-up |
| #78 | Issue dependency map in contributor backlog | None | Completed backlog foundation |
| #79 | Stage 3 readiness gate | #66, #68, #71, #75 | Stage 3 gate |
| #80 | No-network core analyzer regression guard | #68 | Completed safety guard |
| #102 | Stage 3 blocked label review | #79 | Post-gate label alignment |

## Stage 3 Gate

Stage 3 design work can continue while Stage 2.1 is in progress, but implementation-heavy provider adapters should wait until:

- the MVP acceptance checklist is still accurate;
- release readiness is documented;
- JSON/schema validation is wired into CI;
- no-network local analyzer behavior remains tested;
- blocked Stage 3 labels are reviewed after #79.

This keeps provider context additive. GitHub, GitLab, CI-provider, and agent-oriented features must enrich local Maven evidence without replacing it.
