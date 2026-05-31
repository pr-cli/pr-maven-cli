# Release Snapshot: v0.1.0

This snapshot records what shipped in PR Maven CLI `v0.1.0`, what was validated, and what should happen next.

It is based on the repository documentation, the `v0.1.0` GitHub release, the release workflow run, and the project issue backlog.

## Release Metadata

- Version: `v0.1.0`
- Target roadmap stage: Stage 1 production-usable local MVP
- Release URL: <https://github.com/pr-cli/pr-maven-cli/releases/tag/v0.1.0>
- Tag commit: `0ba04b29d190b00617b8909a3f85b64bbf514668`
- Release PR: [#98 Prepare MVP release readiness](https://github.com/pr-cli/pr-maven-cli/pull/98)
- Release workflow run: <https://github.com/pr-cli/pr-maven-cli/actions/runs/26701293106>
- Published: 2026-05-31 02:40 UTC
- Local project posture: public OSS repository, local-first runtime, provider-agnostic core

## What Shipped

### Product Scope

`v0.1.0` ships a production-usable local CLI and Go library for Maven failure triage.

The release can:

- discover Maven modules from `pom.xml` files, including nested module layouts;
- parse Surefire JUnit XML reports;
- parse Failsafe JUnit XML reports;
- parse Checkstyle XML reports;
- parse SpotBugs XML reports;
- extract Maven Enforcer failures from local Maven log fixtures;
- extract JaCoCo threshold failures from local Maven log fixtures;
- map findings back to Maven modules;
- identify Maven plugin and phase when supported by deterministic evidence;
- generate conservative Maven reproduction commands;
- emit human-readable text output;
- emit stable JSON output;
- write text or JSON output to a selected file;
- filter output by Maven module path or artifact id;
- expose the analyzer as `pkg/prmaven` for Go callers.

### CLI and Library Surface

Implemented CLI commands:

- `prmaven fails`
- `prmaven why`
- `prmaven version`

Implemented CLI flags:

- `-project`
- `-format text|json`
- `-module`
- `-output`

Implemented public Go API:

- `prmaven.Analyze`
- `prmaven.Options`
- `prmaven.Report`
- `prmaven.Summary`
- `prmaven.Module`
- `prmaven.Finding`

### Runtime Boundary

The release starts from local files already available in a Maven workspace or CI artifact directory.

Supported local inputs include:

- `pom.xml`
- `target/surefire-reports`
- `target/failsafe-reports`
- `target/checkstyle-result.xml`
- `target/spotbugsXml.xml`
- `target/spotbugs.xml`
- selected local Maven log files such as `target/maven-enforcer.log`, `target/jacoco.log`, and `target/maven.log`

The core analyzer does not:

- call GitHub or GitLab APIs;
- require provider tokens;
- download CI artifacts;
- upload logs or reports;
- require AI providers;
- require telemetry;
- run Maven on behalf of the user.

GitHub is currently the only platform with first-party project automation and a copyable CI example. There is no native GitHub or GitLab runtime adapter in `v0.1.0`.

## Release Artifacts

The release workflow published Linux, macOS, and Windows packages with SHA-256 checksums.

Artifacts:

- `prmaven-v0.1.0-linux-amd64.tar.gz`
- `prmaven-v0.1.0-linux-amd64.tar.gz.sha256`
- `prmaven-v0.1.0-linux-arm64.tar.gz`
- `prmaven-v0.1.0-linux-arm64.tar.gz.sha256`
- `prmaven-v0.1.0-darwin-amd64.tar.gz`
- `prmaven-v0.1.0-darwin-amd64.tar.gz.sha256`
- `prmaven-v0.1.0-darwin-arm64.tar.gz`
- `prmaven-v0.1.0-darwin-arm64.tar.gz.sha256`
- `prmaven-v0.1.0-windows-amd64.zip`
- `prmaven-v0.1.0-windows-amd64.zip.sha256`

## Validation Snapshot

Validation completed for the release:

- `go test ./...`
- CLI smoke test for `demo/multi-module-failure`
- CLI JSON smoke test for `demo/multi-module-failure`
- CLI no-failure smoke test for `demo/no-failure`
- PR CI for [#98](https://github.com/pr-cli/pr-maven-cli/pull/98)
- release workflow for `v0.1.0`

The project CI includes:

- formatting and vet checks;
- Go tests across Linux, Windows, macOS, Go 1.22.x, and current stable Go;
- race detector;
- coverage gate;
- cross-platform build jobs;
- compiled CLI smoke test;
- aggregate `All CI checks` job for branch protection;
- security workflow with govulncheck, CodeQL, and dependency review;
- release workflow that builds packages and checksums from `v*` tags.

## Issues Completed Into This Release Line

The release includes the original Stage 1 MVP plus the completed Stage 2 and Stage 2.1 hardening work that landed before the `v0.1.0` tag.

Parser and analyzer expansion:

- [#1 Add Checkstyle XML parser support](https://github.com/pr-cli/pr-maven-cli/issues/1)
- [#2 Add SpotBugs XML parser support](https://github.com/pr-cli/pr-maven-cli/issues/2)
- [#3 Extract Maven Enforcer failures from log fixtures](https://github.com/pr-cli/pr-maven-cli/issues/3)
- [#4 Extract JaCoCo threshold failures](https://github.com/pr-cli/pr-maven-cli/issues/4)
- [#5 Add parser registry abstraction](https://github.com/pr-cli/pr-maven-cli/issues/5)

Validation and fixture coverage:

- [#6 Add nested module discovery tests](https://github.com/pr-cli/pr-maven-cli/issues/6)
- [#7 Add Windows and Linux path normalization tests](https://github.com/pr-cli/pr-maven-cli/issues/7)
- [#9 Add JSON schema for report output](https://github.com/pr-cli/pr-maven-cli/issues/9)
- [#10 Add golden tests for text output](https://github.com/pr-cli/pr-maven-cli/issues/10)
- [#13 Add no-failure demo fixture](https://github.com/pr-cli/pr-maven-cli/issues/13)

Documentation and contributor readiness:

- [#8 Document the JSON contract](https://github.com/pr-cli/pr-maven-cli/issues/8)
- [#11 Add fixture contribution guide](https://github.com/pr-cli/pr-maven-cli/issues/11)
- [#63 After renaming the organization, update the repository remote](https://github.com/pr-cli/pr-maven-cli/issues/63)
- [#66 Add final MVP acceptance checklist](https://github.com/pr-cli/pr-maven-cli/issues/66)
- [#75 Add release readiness validation checklist](https://github.com/pr-cli/pr-maven-cli/issues/75)
- [#78 Add issue dependency map to contributor backlog](https://github.com/pr-cli/pr-maven-cli/issues/78)
- [#95 Add provider-context dependency map to contributor backlog](https://github.com/pr-cli/pr-maven-cli/issues/95)

CLI and automation:

- [#12 Add GitHub Actions CI for Go tests](https://github.com/pr-cli/pr-maven-cli/issues/12)
- [#14 Add release workflow for tagged builds](https://github.com/pr-cli/pr-maven-cli/issues/14)
- [#15 Add output file option](https://github.com/pr-cli/pr-maven-cli/issues/15)
- [#16 Add CLI module filter](https://github.com/pr-cli/pr-maven-cli/issues/16)

## Source Documents

Use these files as the living source of truth after this snapshot:

- [Roadmap](../ROADMAP.md)
- [Final testable delivery](final-testable-delivery.md)
- [MVP acceptance checklist](mvp-acceptance.md)
- [Implementation status](implementation-status.md)
- [Release process](release.md)
- [Testing](testing.md)
- [CI/CD](ci.md)
- [Usage](usage.md)
- [Integrations](integrations.md)
- [Permission posture](permissions.md)

## Immediate Next Steps

These are the next actions after `v0.1.0`. They should be handled before broad Stage 3 implementation, because they harden the MVP, improve contributor confidence, or keep the release line easy to maintain.

Recommended order:

1. Finish the remaining ready Stage 2 documentation and ergonomics issues:
   - [#17 Add Maven 3.9 compatibility fixture notes](https://github.com/pr-cli/pr-maven-cli/issues/17)
   - [#18 Add confidence documentation](https://github.com/pr-cli/pr-maven-cli/issues/18)
   - [#19 Improve CLI help output](https://github.com/pr-cli/pr-maven-cli/issues/19)
   - [#20 Add maintainer issue labeling guide](https://github.com/pr-cli/pr-maven-cli/issues/20)
2. Close the ready Stage 2.1 documentation and deterministic-context issues:
   - [#61 Clarify deterministic text and JSON terminology](https://github.com/pr-cli/pr-maven-cli/issues/61)
   - [#62 Document deterministic context principles](https://github.com/pr-cli/pr-maven-cli/issues/62)
   - [#65 MVP hardening plan](https://github.com/pr-cli/pr-maven-cli/issues/65)
3. Add focused validation fixtures before adding provider-context behavior:
   - [#69 Add golden JSON output snapshots](https://github.com/pr-cli/pr-maven-cli/issues/69)
   - [#70 Add fixture integrity validation](https://github.com/pr-cli/pr-maven-cli/issues/70)
   - [#72 Add no-failure regression matrix for output modes](https://github.com/pr-cli/pr-maven-cli/issues/72)
4. Move validation into CI once the local fixtures are stable:
   - [#71 Add JSON schema validation to CI](https://github.com/pr-cli/pr-maven-cli/issues/71)
   - [#73 Run fixture integrity validation in CI](https://github.com/pr-cli/pr-maven-cli/issues/73)
5. Expand end-to-end MVP acceptance coverage:
   - [#67 Add documentation command smoke tests](https://github.com/pr-cli/pr-maven-cli/issues/67)
   - [#68 Add end-to-end CLI acceptance suite](https://github.com/pr-cli/pr-maven-cli/issues/68)
   - [#74 Add output file E2E matrix](https://github.com/pr-cli/pr-maven-cli/issues/74)
   - [#76 Add module filter E2E matrix](https://github.com/pr-cli/pr-maven-cli/issues/76)
6. Update contributor and maintainer gates:
   - [#77 Update PR template with MVP validation checklist](https://github.com/pr-cli/pr-maven-cli/issues/77)
   - [#79 Define Stage 3 readiness gate](https://github.com/pr-cli/pr-maven-cli/issues/79)
   - [#80 Add no-network core analyzer regression guard](https://github.com/pr-cli/pr-maven-cli/issues/80)

## Future Next Steps

Future work should preserve the local-first analyzer while adding optional PR and CI context.

Stage 3 design and contracts:

- [#82 Provider context planning map](https://github.com/pr-cli/pr-maven-cli/issues/82)
- [#83 Define provider error and offline fallback taxonomy](https://github.com/pr-cli/pr-maven-cli/issues/83)
- [#84 Define provider adapter package boundaries](https://github.com/pr-cli/pr-maven-cli/issues/84)
- [#85 Define provider fixture and mock contract](https://github.com/pr-cli/pr-maven-cli/issues/85)
- [#86 Document GitHub token and permission matrix](https://github.com/pr-cli/pr-maven-cli/issues/86)
- [#87 Design changed-files fixture contract](https://github.com/pr-cli/pr-maven-cli/issues/87)
- [#89 Design check-runs fixture contract](https://github.com/pr-cli/pr-maven-cli/issues/89)
- [#90 Design PR context JSON extension contract](https://github.com/pr-cli/pr-maven-cli/issues/90)
- [#92 Define GitLab parity boundary for provider adapters](https://github.com/pr-cli/pr-maven-cli/issues/92)
- [#93 Design CI artifact directory fixture layout](https://github.com/pr-cli/pr-maven-cli/issues/93)
- [#94 Draft agent evidence bundle schema](https://github.com/pr-cli/pr-maven-cli/issues/94)

Stage 3 implementation candidates:

- [#21 Add GitHub changed files adapter](https://github.com/pr-cli/pr-maven-cli/issues/21)
- [#22 Design GitHub adapter interface](https://github.com/pr-cli/pr-maven-cli/issues/22)
- [#23 Add GitHub check runs adapter](https://github.com/pr-cli/pr-maven-cli/issues/23)
- [#24 Design baseline comparison model](https://github.com/pr-cli/pr-maven-cli/issues/24)
- [#25 Add PR-to-module relevance scoring](https://github.com/pr-cli/pr-maven-cli/issues/25)
- [#26 Implement confidence model v2](https://github.com/pr-cli/pr-maven-cli/issues/26)
- [#27 Add prmaven explain command](https://github.com/pr-cli/pr-maven-cli/issues/27)
- [#28 Add prmaven ci command](https://github.com/pr-cli/pr-maven-cli/issues/28)
- [#29 Add Markdown PR summary output](https://github.com/pr-cli/pr-maven-cli/issues/29)
- [#34 Add agent evidence bundle output](https://github.com/pr-cli/pr-maven-cli/issues/34)
- [#36 Add CI artifact directory option](https://github.com/pr-cli/pr-maven-cli/issues/36)
- [#88 Define Markdown PR summary content contract](https://github.com/pr-cli/pr-maven-cli/issues/88)
- [#91 Define why, explain, and ci command UX boundaries](https://github.com/pr-cli/pr-maven-cli/issues/91)

Research, distribution, and platform guidance:

- [#30 Investigate SARIF output](https://github.com/pr-cli/pr-maven-cli/issues/30)
- [#31 Investigate GitHub annotation output](https://github.com/pr-cli/pr-maven-cli/issues/31)
- [#32 Add Maven 4 compatibility investigation](https://github.com/pr-cli/pr-maven-cli/issues/32)
- [#33 Research GitLab merge request support](https://github.com/pr-cli/pr-maven-cli/issues/33)
- [#35 Add privacy guide for CI logs](https://github.com/pr-cli/pr-maven-cli/issues/35)
- [#37 Add internal engineering platform integration guide](https://github.com/pr-cli/pr-maven-cli/issues/37)
- [#38 Add GitHub Action usage example](https://github.com/pr-cli/pr-maven-cli/issues/38)
- [#39 Add stable output compatibility policy](https://github.com/pr-cli/pr-maven-cli/issues/39)
- [#40 Add package manager distribution research](https://github.com/pr-cli/pr-maven-cli/issues/40)

## Guardrails For The Next Release

Do not treat Stage 3 as permission to weaken the MVP contract.

The following should remain true:

- local Maven report analysis works without network access;
- provider integrations are optional;
- tests use fixtures and mocks before live provider calls;
- JSON compatibility is documented before new public fields become expected by users;
- Maven 4 support remains an investigation until validated against stable Maven 4 behavior;
- release artifacts continue to include checksums for every package.

