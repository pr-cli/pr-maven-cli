# Provider Context Plan

This document organizes Stage 3 provider-context work before implementation-heavy adapters begin.

Provider context means optional pull request, changed-file, check-run, CI artifact, and summary metadata that can enrich local Maven findings. It must remain additive: the local analyzer continues to work from local Maven report artifacts without network access or provider tokens.

Use this plan with the [project visual map](project-visual-map.md), [v0.3.0 release acceptance checklist](release-acceptance-v0.3.0.md), [integration scope](integrations.md), and [roadmap](../ROADMAP.md).

## Readiness Gates

Stage 3 provider-context work starts after these gates are available:

- [#79 Stage 3 readiness gate](https://github.com/pr-cli/pr-maven-cli/issues/79)
- [#100 v0.3.0 release acceptance checklist](https://github.com/pr-cli/pr-maven-cli/issues/100)
- [#101 visual dependency map](https://github.com/pr-cli/pr-maven-cli/issues/101)
- [#102 post-gate label review](https://github.com/pr-cli/pr-maven-cli/issues/102)

Implementation-heavy issues remain blocked until their related design contracts, fixtures, and mocks land.

## Design Contract Order

The first ready Stage 3 lane is design and validation, not live provider implementation.

Recommended order:

1. [#84 Provider adapter package boundaries](https://github.com/pr-cli/pr-maven-cli/issues/84) and [#85 provider fixture and mock contract](https://github.com/pr-cli/pr-maven-cli/issues/85)
2. [#83 provider error and offline fallback taxonomy](https://github.com/pr-cli/pr-maven-cli/issues/83) and [#86 GitHub token and permission matrix](https://github.com/pr-cli/pr-maven-cli/issues/86)
3. [#87 changed-files fixture contract](https://github.com/pr-cli/pr-maven-cli/issues/87) and [#89 check-runs fixture contract](https://github.com/pr-cli/pr-maven-cli/issues/89)
4. [#90 PR context JSON extension contract](https://github.com/pr-cli/pr-maven-cli/issues/90), [#88 Markdown PR summary content contract](https://github.com/pr-cli/pr-maven-cli/issues/88), and [#91 why, explain, and ci command UX boundaries](https://github.com/pr-cli/pr-maven-cli/issues/91)
5. [#93 CI artifact directory fixture layout](https://github.com/pr-cli/pr-maven-cli/issues/93), [#92 GitLab parity boundary](https://github.com/pr-cli/pr-maven-cli/issues/92), and [#94 agent evidence bundle schema](https://github.com/pr-cli/pr-maven-cli/issues/94)

## Implementation Issues That Stay Blocked

The following issues should not start until the relevant design contract is merged or explicitly waived by a maintainer:

- [#21 GitHub changed files adapter](https://github.com/pr-cli/pr-maven-cli/issues/21)
- [#22 GitHub adapter interface](https://github.com/pr-cli/pr-maven-cli/issues/22)
- [#23 GitHub check runs adapter](https://github.com/pr-cli/pr-maven-cli/issues/23)
- [#24 baseline comparison model](https://github.com/pr-cli/pr-maven-cli/issues/24)
- [#25 PR-to-module relevance scoring](https://github.com/pr-cli/pr-maven-cli/issues/25)
- [#26 confidence model v2](https://github.com/pr-cli/pr-maven-cli/issues/26)
- [#27 prmaven explain command](https://github.com/pr-cli/pr-maven-cli/issues/27)
- [#28 prmaven ci command](https://github.com/pr-cli/pr-maven-cli/issues/28)
- [#29 Markdown PR summary output](https://github.com/pr-cli/pr-maven-cli/issues/29)
- [#34 agent evidence bundle output](https://github.com/pr-cli/pr-maven-cli/issues/34)
- [#36 CI artifact directory option](https://github.com/pr-cli/pr-maven-cli/issues/36)

## Contributor Slicing Rules

Each provider-context contribution should:

- change one contract, fixture shape, or documented boundary at a time;
- preserve local report-only behavior;
- avoid live API calls in default tests;
- prefer sanitized fixtures and fake clients over real provider data;
- document permission and offline behavior before adding token-aware code;
- keep JSON additions additive and compatible with existing consumers.

## Label Expectations

Every provider-context issue should keep:

- one `stage:*` label;
- at least one `area:*` label;
- exactly one current `status:*` label;
- at least one contributor-fit label such as `help wanted`, `need help`, `good first contribution`, `oss first friendly`, or `agent-friendly`.

Ready design issues signal that contributors may start drafting contracts. Blocked implementation issues signal that the repo is still waiting for the prerequisite contract, fixture, or mock design.

Detailed package-boundary guidance lives in [Provider Adapter Package Boundaries](provider-adapter-boundaries.md).
Fixture and fake-client guidance lives in [Provider Fixtures And Mocks](provider-fixtures-and-mocks.md).
Provider failure behavior lives in [Provider Errors And Offline Fallbacks](provider-errors-and-fallbacks.md).
GitHub read-token expectations live in [GitHub Provider Permissions](github-provider-permissions.md).
Changed-file fixture expectations live in [Changed-Files Fixture Contract](changed-files-fixture-contract.md).
Check-run fixture expectations live in [Check-Runs Fixture Contract](check-runs-fixture-contract.md).
PR context JSON extension expectations live in [PR Context JSON Extension Contract](pr-context-json-extension.md).
Markdown PR summary expectations live in [Markdown PR Summary Contract](markdown-pr-summary-contract.md).
Command UX boundaries live in [Command UX Boundaries](command-ux-boundaries.md).
