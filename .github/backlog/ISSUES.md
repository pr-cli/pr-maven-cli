# Contributor Backlog

This backlog mirrors the planned GitHub issues for Stage 2 and Stage 3.

Each item is intentionally scoped so a contributor or automated coding agent can implement it with clear acceptance criteria.

## Stage 2 - Contributor Growth and Maven Signal Expansion

### 1. Add Checkstyle XML parser support

Labels: `stage: 2`, `area: parser`, `help wanted`, `good first contribution`, `agent-friendly`

Acceptance criteria:

- Add a sanitized Checkstyle XML fixture.
- Produce a `Finding` with module path, report path, plugin name, message, and confidence reasons.
- Add unit tests for parser behavior.
- `go test ./...` passes.

### 2. Add SpotBugs XML parser support

Labels: `stage: 2`, `area: parser`, `help wanted`, `agent-friendly`

Acceptance criteria:

- Add a SpotBugs XML fixture.
- Map the bug instance to a Maven module.
- Include bug category/type and source file when available.
- `go test ./...` passes.

### 3. Extract Maven Enforcer failures from log fixtures

Labels: `stage: 2`, `area: parser`, `help wanted`, `agent-friendly`

Acceptance criteria:

- Add a sanitized Maven log fixture containing an Enforcer failure.
- Detect `maven-enforcer-plugin`.
- Emit phase/plugin/message context.
- Keep parsing deterministic and fixture-driven.

### 4. Extract JaCoCo threshold failures

Labels: `stage: 2`, `area: parser`, `help wanted`

Acceptance criteria:

- Add fixture coverage for a JaCoCo threshold failure.
- Detect `jacoco-maven-plugin`.
- Include threshold context in the finding message.
- Add focused tests.

### 5. Add parser registry abstraction

Labels: `stage: 2`, `area: architecture`, `help wanted`

Acceptance criteria:

- Introduce a small parser interface.
- Register Surefire and Failsafe through the same mechanism.
- Preserve current JSON output.
- Avoid speculative plugin framework design.

### 6. Add nested module discovery tests

Labels: `stage: 2`, `area: test`, `good first contribution`, `agent-friendly`

Acceptance criteria:

- Add a nested Maven module fixture.
- Verify recursive module discovery.
- Verify report-to-module mapping for nested paths.
- `go test ./...` passes.

### 7. Add Windows and Linux path normalization tests

Labels: `stage: 2`, `area: test`, `good first contribution`, `agent-friendly`

Acceptance criteria:

- Add tests for Windows-style and POSIX-style paths.
- Verify JSON output always uses slash-separated module/report paths.
- Avoid OS-specific test failures.

### 8. Document the JSON contract

Labels: `stage: 2`, `area: docs`, `good first contribution`, `agent-friendly`

Acceptance criteria:

- Add `docs/json-contract.md`.
- Document all fields in `Report`, `Module`, and `Finding`.
- Include one complete JSON example from the demo.
- Explain compatibility expectations.

### 9. Add JSON schema for report output

Labels: `stage: 2`, `area: docs`, `area: test`, `help wanted`

Acceptance criteria:

- Add `schema/prmaven-report.schema.json`.
- Ensure the schema matches the current `Report` contract.
- Add documentation showing how CI tools can validate output.

### 10. Add golden tests for text output

Labels: `stage: 2`, `area: test`, `good first contribution`, `agent-friendly`

Acceptance criteria:

- Add a golden fixture for the demo text output.
- Add a test that compares current output to the golden file.
- Keep output stable and intentional.

### 11. Add GitHub Actions CI for Go tests

Labels: `stage: 2`, `area: ci`, `good first contribution`, `agent-friendly`

Acceptance criteria:

- Add a workflow that runs `go test ./...`.
- Use a maintained Go version.
- Run on pull requests and pushes to `main`.
- Keep the workflow minimal.

### 12. Add release workflow for tagged builds

Labels: `stage: 2`, `area: ci`, `help wanted`

Acceptance criteria:

- Add a release workflow triggered by `v*` tags.
- Build binaries for Linux, macOS, and Windows.
- Attach checksums.
- Document the release process.

### 13. Add fixture contribution guide

Labels: `stage: 2`, `area: docs`, `good first contribution`, `agent-friendly`

Acceptance criteria:

- Add `docs/fixtures.md`.
- Explain how to sanitize Maven reports and logs.
- Explain where fixtures should live.
- Include examples for Surefire and Failsafe.

### 14. Add no-failure demo fixture

Labels: `stage: 2`, `area: test`, `good first contribution`, `agent-friendly`

Acceptance criteria:

- Add a fixture project with passing reports.
- Verify CLI exits `0`.
- Verify text output says no failures were found.
- Verify JSON has `findingCount: 0`.

### 15. Add CLI module filter

Labels: `stage: 2`, `area: cli`, `help wanted`

Acceptance criteria:

- Add `-module` filter support.
- Limit findings to a selected Maven module path or module artifact ID.
- Add tests for matching and no-match cases.

### 16. Add output file option

Labels: `stage: 2`, `area: cli`, `good first contribution`, `agent-friendly`

Acceptance criteria:

- Add `-output <path>`.
- Write text or JSON output to the selected file.
- Keep stdout behavior unchanged when `-output` is absent.
- Add tests.

### 17. Add confidence documentation

Labels: `stage: 2`, `area: docs`, `good first contribution`

Acceptance criteria:

- Document what `high`, `medium`, and `low` confidence mean.
- Explain why Stage 1 currently reports high confidence for report-backed findings.
- Link the documentation from the README.

### 18. Add Maven 3.9 compatibility fixture notes

Labels: `stage: 2`, `area: docs`, `good first contribution`

Acceptance criteria:

- Document Maven 3.9.x as the production baseline.
- Mention Maven 3.9.16 as the documented current baseline.
- Explain why Maven 4 is tracked separately.

### 19. Improve CLI help output

Labels: `stage: 2`, `area: cli`, `good first contribution`, `agent-friendly`

Acceptance criteria:

- Add clear usage examples to `-h` output.
- Document commands and flags.
- Keep implementation dependency-free.

### 20. Add maintainer issue labeling guide

Labels: `stage: 2`, `area: docs`, `good first contribution`, `agent-friendly`

Acceptance criteria:

- Add a short maintainer doc for labels.
- Explain `good first contribution`, `help wanted`, and `agent-friendly`.
- Include examples of well-scoped issues.

## Stage 2.1 - MVP Hardening and Validation

Stage 2.1 is a validation lane between the implemented Stage 2 work and the optional provider-context work planned for Stage 3.

Root plan:

- #65 - `[Stage 2.1] MVP hardening plan`

Suggested merge order:

1. #66 - Add final MVP acceptance checklist.
2. Ready validation work: #70, #71, #69, #72, and #78.
3. Work unblocked by #66: #67, #68, #75, and #77.
4. Work unblocked by #70 and #68: #73, #76, #74, and #80.
5. Stage 3 gate: #79, after #66, #68, #71, and #75.
6. Stage 3 label unblock review: #102, after #79.

Additional ready documentation work:

- #61 - Clarify deterministic text and JSON terminology.
- #62 - Document deterministic context principles.

Additional guardrail and release-planning work:

- #100 - Define v0.3.0 release acceptance checklist.
- #101 - Add Stage 3 dependency map visual. Completed and maintained through `docs/project-visual-map.md`.
- #102 - Review Stage 3 blocked labels after the readiness gate.

### Documentation Polish Issues

#### D1. Clarify deterministic text and JSON terminology

Issue: #61

Labels: `stage: 2.1`, `area: docs`, `good first contribution`, `oss first friendly`, `agent-friendly`, `status: ready`

Acceptance criteria:

- Explain `Deterministic Text`.
- Explain `Deterministic JSON`.
- Keep wording aligned with README, JSON contract, and product principles.
- `go test ./...` passes.

#### D2. Document deterministic context principles

Issue: #62

Related wording issue: #61.

Labels: `stage: 2.1`, `area: architecture`, `area: docs`, `help wanted`, `need help`, `status: ready`

Acceptance criteria:

- Explain that deterministic output does not depend on an LLM, loose heuristics, random ordering, network access, or subjective interpretation.
- Explain the product principle: parsing plus rules, one package manager, one failure model, one deterministic contract.
- Link the explanation from `README.md`, `MANIFESTO.md`, or `docs/project-context.md` where it best fits.
- `go test ./...` passes.

### 21. Add final MVP acceptance checklist

Issue: #66

Labels: `stage: 2.1`, `area: docs`, `area: validation`, `good first contribution`, `oss first friendly`, `agent-friendly`, `status: ready`

Acceptance criteria:

- Add a checklist document or section that references `docs/final-testable-delivery.md`.
- Include functional, validation, safety, documentation, and release-readiness checks.
- Link the checklist from `README.md` or `docs/testing.md`.
- `go test ./...` passes.

### 22. Add documentation command smoke tests

Issue: #67

Depends on: #66.

Labels: `stage: 2.1`, `area: docs`, `area: test`, `area: validation`, `help wanted`, `need help`, `agent-friendly`, `status: blocked`

Acceptance criteria:

- Identify safe local commands to smoke test.
- Add a focused test or script for representative documented commands.
- Avoid publish, push, tag, or provider-token commands.
- Document how maintainers should update the smoke test.
- `go test ./...` passes.

### 23. Add end-to-end CLI acceptance suite

Issue: #68

Depends on: #66.

Labels: `stage: 2.1`, `area: test`, `area: validation`, `help wanted`, `need help`, `agent-friendly`, `status: blocked`

Acceptance criteria:

- Exercise `demo/multi-module-failure` through the CLI.
- Exercise `demo/no-failure` through the CLI.
- Verify documented exit codes.
- Verify JSON output is parseable.
- Keep tests local-first and provider-agnostic.
- `go test ./...` passes.

### 24. Add golden JSON output snapshots

Issue: #69

Depends on: #9 (completed).

Labels: `stage: 2.1`, `area: test`, `area: validation`, `good first contribution`, `oss first friendly`, `agent-friendly`, `status: ready`

Acceptance criteria:

- Add a golden JSON fixture for `demo/multi-module-failure`.
- Add a golden JSON fixture for `demo/no-failure`.
- Normalize machine-specific fields such as `projectRoot` before comparison.
- Document how to update snapshots intentionally.
- `go test ./...` passes.

### 25. Add fixture integrity validation

Issue: #70

Labels: `stage: 2.1`, `area: test`, `area: validation`, `good first contribution`, `oss first friendly`, `agent-friendly`, `status: ready`

Acceptance criteria:

- Validate that expected fixture report and log files exist.
- Validate that demo fixtures remain deterministic and self-contained.
- Validate intentionally committed `target` report files.
- Produce clear failures when a fixture is incomplete.
- Document the validation command.
- `go test ./...` passes.

### 26. Add JSON schema validation to CI

Issue: #71

Depends on: #9 (completed).

Labels: `stage: 2.1`, `area: ci`, `area: validation`, `help wanted`, `need help`, `status: ready`

Acceptance criteria:

- Generate JSON from a demo fixture during validation.
- Validate the generated JSON against the report schema.
- Keep the validation dependency-light for GitHub Actions.
- Document the validation command in `docs/testing.md` or `docs/ci.md`.
- `go test ./...` passes.

### 27. Add no-failure regression matrix for output modes

Issue: #72

Depends on: #13 (completed).

Labels: `stage: 2.1`, `area: test`, `area: validation`, `good first contribution`, `oss first friendly`, `agent-friendly`, `status: ready`

Acceptance criteria:

- Verify text no-failure output.
- Verify JSON no-failure output has `findingCount: 0`.
- Verify `-output` works with no findings.
- Verify exit code `0` remains stable.
- `go test ./...` passes.

### 28. Run fixture integrity validation in CI

Issue: #73

Depends on: #70.

Labels: `stage: 2.1`, `area: ci`, `area: validation`, `help wanted`, `need help`, `status: blocked`

Acceptance criteria:

- CI runs fixture integrity validation on pull requests.
- CI output clearly identifies fixture validation failures.
- `docs/ci.md` documents where this validation runs.
- `go test ./...` passes.

### 29. Add output file E2E matrix

Issue: #74

Depends on: #15 (completed), #68.

Labels: `stage: 2.1`, `area: test`, `area: validation`, `help wanted`, `need help`, `agent-friendly`, `status: blocked`

Acceptance criteria:

- Verify text output can be written to a selected file.
- Verify JSON output can be written to a selected file.
- Verify stdout behavior when `-output` is absent.
- Verify generated files are cleaned up in tests.
- `go test ./...` passes.

### 30. Add release readiness validation checklist

Issue: #75

Depends on: #66.

Labels: `stage: 2.1`, `area: docs`, `area: validation`, `help wanted`, `need help`, `status: blocked`

Acceptance criteria:

- Add a release readiness checklist to `docs/release.md` or a linked document.
- Include tests, CI, schema, examples, docs, and artifact expectations.
- Keep the checklist suitable for v0.x releases.
- Link to `docs/final-testable-delivery.md`.
- `go test ./...` passes.

### 31. Add module filter E2E matrix

Issue: #76

Depends on: #16 (completed), #68.

Labels: `stage: 2.1`, `area: test`, `area: validation`, `help wanted`, `need help`, `agent-friendly`, `status: blocked`

Acceptance criteria:

- Verify filtering by module path.
- Verify filtering by module artifact id.
- Verify no-match behavior.
- Verify JSON summary and findings reflect the filtered view.
- `go test ./...` passes.

### 32. Update PR template with MVP validation checklist

Issue: #77

Depends on: #66.

Labels: `stage: 2.1`, `area: docs`, `area: validation`, `help wanted`, `need help`, `status: blocked`

Acceptance criteria:

- Add an MVP validation checklist section to `.github/PULL_REQUEST_TEMPLATE.md`.
- Keep the checklist short and contributor-friendly.
- Link to `docs/testing.md` or the acceptance checklist.
- Do not require irrelevant checks for documentation-only PRs.

### 33. Add issue dependency map to contributor backlog

Issue: #78

Labels: `stage: 2.1`, `area: docs`, `good first contribution`, `oss first friendly`, `agent-friendly`, `status: ready`

Acceptance criteria:

- Update `.github/backlog/ISSUES.md` with a Stage 2.1 section.
- Include issue numbers, dependency notes, and suggested merge order.
- Keep Stage 3 dependencies clear without over-constraining contributors.
- Ensure every listed issue has labels.

### 34. Define Stage 3 readiness gate

Issue: #79

Depends on: #66, #68, #71, #75.

Labels: `stage: 2.1`, `area: architecture`, `area: validation`, `help wanted`, `need help`, `status: blocked`

Acceptance criteria:

- Add a short Stage 3 readiness note to `ROADMAP.md` or `docs/integrations.md`.
- Define required validation gates before provider adapters land.
- Clarify that network adapters remain optional.
- Link relevant Stage 3 architecture issues.
- `go test ./...` passes.

### 35. Add no-network core analyzer regression guard

Issue: #80

Depends on: #68.

Labels: `stage: 2.1`, `area: test`, `area: validation`, `help wanted`, `need help`, `agent-friendly`, `status: blocked`

Acceptance criteria:

- Add a test or validation that exercises core analysis without provider tokens.
- Ensure GitHub/GitLab environment variables are not required for core tests.
- Document the local-first/no-network expectation in `docs/testing.md` or `docs/permissions.md`.
- `go test ./...` passes.

## Stage 3 - PR and CI Context Layer

Stage 3 adds optional PR and CI provider context without weakening the local-first Maven analyzer.

Implementation work in this stage should remain gated by the Stage 2.1 readiness gate:

- #79 - `[Stage 2.1] Define Stage 3 readiness gate`

Provider-context planning root:

- #82 - `[Stage 3] Provider context planning map`

Suggested planning merge order:

1. #100 - Define v0.3.0 release acceptance checklist.
2. Keep #101 current through the visual project map.
3. #102 - Review blocked labels after the Stage 3 readiness gate.
4. #95 - Add provider-context dependency map to contributor backlog.
5. #84 and #85 - Define adapter package boundaries and provider fixture/mock contracts.
6. #83 and #86 - Define provider error fallback behavior and GitHub permission expectations.
7. #87 and #89 - Define changed-files and check-runs fixture contracts.
8. #90, #88, and #91 - Define JSON, Markdown, and CLI UX contracts.
9. #93, #92, and #94 - Define CI artifact, GitLab parity, and agent evidence-bundle contracts.

### Provider-Context Planning Issues

#### P0. Define v0.3.0 release acceptance checklist

Issue: #100

Labels: `stage: 3`, `enhancement`, `area: docs`, `area: validation`, `help wanted`, `need help`, `status: ready`

Acceptance criteria:

- Define the minimum release acceptance checklist for `v0.3.0`.
- Tie the checklist to the Stage 3 readiness gate and provider-context contracts.
- Keep provider adapter work optional and additive.
- Link from the roadmap or visual project map.

#### P0.1. Add Stage 3 dependency map visual

Issue: #101

Status: completed by the visual project map.

Labels: `stage: 3`, `documentation`, `area: docs`, `good first contribution`, `oss first friendly`, `agent-friendly`, `status: ready`

Maintenance criteria:

- Keep `docs/project-visual-map.md` aligned with the GitHub issue state.
- Update it when Stage 3 dependencies or blocked labels change.

#### P0.2. Review Stage 3 blocked labels after readiness gate

Issue: #102

Depends on: #79.

Labels: `stage: 2.1`, `enhancement`, `area: validation`, `area: docs`, `help wanted`, `need help`, `status: blocked`

Acceptance criteria:

- Review Stage 3 issues currently labeled `status: blocked`.
- Unblock only design issues whose dependencies are closed.
- Preserve implementation-heavy issues as blocked until contracts and mocks exist.
- Update the visual project map when labels change.

#### P1. Add provider-context dependency map to contributor backlog

Issue: #95

Labels: `stage: 3`, `documentation`, `area: docs`, `good first contribution`, `oss first friendly`, `agent-friendly`, `status: ready`

Acceptance criteria:

- Add the new Stage 3 planning issues to the backlog document.
- Include dependency notes and suggested merge order.
- Keep the existing Stage 3 implementation issues visible.
- Ensure every listed issue has labels.

#### P2. Define provider adapter package boundaries

Issue: #84

Depends on: #79, #22.

Labels: `stage: 3`, `area: architecture`, `help wanted`, `need help`, `status: blocked`

Acceptance criteria:

- Propose package boundaries for core analyzer, provider interfaces, and provider implementations.
- Explain what cannot import provider-specific code.
- Include testability expectations for adapters.
- Document how this preserves local-first behavior.

#### P3. Define provider fixture and mock contract

Issue: #85

Depends on: #79, #22.

Labels: `stage: 3`, `area: architecture`, `area: test`, `help wanted`, `need help`, `agent-friendly`, `status: blocked`

Acceptance criteria:

- Define fixture formats for provider responses.
- Define fake client behavior for tests.
- Include success, missing token, rate limit, permission denied, and not found scenarios.
- Keep tests runnable without provider credentials.

#### P4. Define provider error and offline fallback taxonomy

Issue: #83

Depends on: #79, #84.

Labels: `stage: 3`, `area: architecture`, `area: docs`, `help wanted`, `need help`, `status: blocked`

Acceptance criteria:

- Define errors for missing token, insufficient permissions, network failure, API rate limit, not found, and unsupported provider state.
- Explain which errors should degrade gracefully.
- Preserve local Maven findings when provider context fails.
- Include JSON/reporting expectations for adapter failures.

#### P5. Document GitHub token and permission matrix

Issue: #86

Depends on: #79, #22.

Labels: `stage: 3`, `area: docs`, `area: github`, `good first contribution`, `oss first friendly`, `agent-friendly`, `status: blocked`

Acceptance criteria:

- List planned GitHub operations and required permissions.
- Distinguish public repository use from private repository use.
- Explain no-token local analyzer behavior.
- Link from `docs/integrations.md` or `docs/permissions.md`.

#### P6. Design changed-files fixture contract

Issue: #87

Depends on: #21, #84, #85.

Labels: `stage: 3`, `area: github`, `area: test`, `help wanted`, `need help`, `agent-friendly`, `status: blocked`

Acceptance criteria:

- Define a sanitized fixture shape for changed files.
- Cover added, modified, deleted, renamed, and nested module files.
- Include mapping expectations from changed files to Maven modules.
- Keep fixture tests provider-token-free.

#### P7. Design check-runs fixture contract

Issue: #89

Depends on: #23, #84, #85.

Labels: `stage: 3`, `area: github`, `area: test`, `help wanted`, `need help`, `agent-friendly`, `status: blocked`

Acceptance criteria:

- Define a sanitized fixture shape for check runs.
- Cover success, failure, skipped, cancelled, timed out, and pending states.
- Define how check-run metadata enriches Maven findings without replacing report evidence.
- Keep fixture tests provider-token-free.

#### P8. Design PR context JSON extension contract

Issue: #90

Depends on: #25, #26, #39, #79.

Labels: `stage: 3`, `area: architecture`, `area: docs`, `help wanted`, `need help`, `status: blocked`

Acceptance criteria:

- Propose additive JSON fields for PR context.
- Preserve current `Report`, `Module`, and `Finding` compatibility expectations.
- Explain how consumers should ignore unknown future fields.
- Include examples for report-only and PR-context-enriched output.

#### P9. Define Markdown PR summary content contract

Issue: #88

Depends on: #29, #90.

Labels: `stage: 3`, `area: cli`, `area: docs`, `help wanted`, `need help`, `agent-friendly`, `status: blocked`

Acceptance criteria:

- Define required sections for findings, reproduction commands, confidence, and provider context.
- Include examples for one finding, multiple findings, and no findings.
- Keep output suitable for PR comments.
- Document what must remain deterministic.

#### P10. Define `why`, `explain`, and `ci` command UX boundaries

Issue: #91

Depends on: #27, #28, #79.

Labels: `stage: 3`, `area: cli`, `area: docs`, `help wanted`, `need help`, `oss first friendly`, `status: blocked`

Acceptance criteria:

- Define user intent for `why`, `explain`, and `ci`.
- Clarify default output formats and exit-code expectations.
- Avoid breaking Stage 1/2 command behavior.
- Include examples for local developer and CI usage.

#### P11. Design CI artifact directory fixture layout

Issue: #93

Depends on: #36, #79.

Labels: `stage: 3`, `area: ci`, `area: test`, `help wanted`, `need help`, `agent-friendly`, `status: blocked`

Acceptance criteria:

- Define fixture layouts for source root plus artifact directory.
- Cover module mapping when artifacts are outside module `target` directories.
- Document failure modes when mapping is ambiguous.
- Keep tests local and fixture-driven.

#### P12. Define GitLab parity boundary for provider adapters

Issue: #92

Depends on: #33, #84, #85.

Labels: `stage: 3`, `area: gitlab`, `area: architecture`, `help wanted`, `need help`, `waiting help`, `status: blocked`

Acceptance criteria:

- Identify shared adapter concepts across GitHub and GitLab.
- Identify GitLab-specific merge request and pipeline metadata differences.
- Keep GitLab support optional and investigation-first.
- Recommend whether implementation should wait for GitHub adapter stabilization.

#### P13. Draft agent evidence bundle schema

Issue: #94

Depends on: #34, #90, #79.

Labels: `stage: 3`, `area: agent`, `area: docs`, `help wanted`, `need help`, `waiting help`, `status: blocked`

Acceptance criteria:

- Define additive JSON bundle fields for findings, commands, confidence reasons, relevant files, and provider context.
- Include report-only and PR-context examples.
- Avoid prompt-specific instructions or provider lock-in.
- Link to JSON compatibility expectations.

### Stage 3 Implementation Issues

### 36. Design GitHub adapter interface

Labels: `stage: 3`, `area: architecture`, `help wanted`

Acceptance criteria:

- Add a design document for optional GitHub adapters.
- Define interfaces for changed files and check runs.
- Keep network behavior outside the core analyzer.

### 37. Add GitHub changed files adapter

Labels: `stage: 3`, `area: github`, `help wanted`

Acceptance criteria:

- Add an optional adapter for PR changed files.
- Use dependency injection so tests can run without GitHub.
- Add fixture/mocked tests.

### 38. Add GitHub check runs adapter

Labels: `stage: 3`, `area: github`, `help wanted`

Acceptance criteria:

- Add an optional adapter for check run metadata.
- Avoid requiring tokens for local analyzer tests.
- Add mock tests for success and failure states.

### 39. Add PR-to-module relevance scoring

Labels: `stage: 3`, `area: github`, `area: architecture`, `help wanted`

Acceptance criteria:

- Combine changed files with Maven module paths.
- Emit relevance score or reasons.
- Keep scoring explainable and deterministic.

### 40. Design baseline comparison model

Labels: `stage: 3`, `area: architecture`, `help wanted`

Acceptance criteria:

- Add a design document for comparing PR findings against main/baseline findings.
- Define required inputs and failure modes.
- Avoid implementation until the model is reviewed.

### 41. Implement confidence model v2

Labels: `stage: 3`, `area: architecture`, `help wanted`

Acceptance criteria:

- Add confidence levels based on multiple evidence sources.
- Preserve confidence reasons in JSON.
- Add tests for report-only and PR-context-backed findings.

### 42. Add Markdown PR summary output

Labels: `stage: 3`, `area: cli`, `help wanted`, `agent-friendly`

Acceptance criteria:

- Add `-format markdown`.
- Include findings and reproduction commands.
- Keep output suitable for PR comments.

### 43. Add `prmaven explain` command

Labels: `stage: 3`, `area: cli`, `help wanted`

Acceptance criteria:

- Add `explain` as a richer diagnostic command.
- Preserve `fails` and `why` behavior.
- Add CLI tests.

### 44. Add `prmaven ci` command

Labels: `stage: 3`, `area: cli`, `area: ci`, `help wanted`

Acceptance criteria:

- Add a command optimized for CI workspaces.
- Output JSON by default or document the chosen behavior.
- Include examples for GitHub Actions.

### 45. Investigate SARIF output

Labels: `stage: 3`, `area: ci`, `help wanted`

Acceptance criteria:

- Add a short research document.
- Map `Finding` fields to SARIF concepts.
- Recommend whether SARIF belongs in the project.

### 46. Investigate GitHub annotation output

Labels: `stage: 3`, `area: ci`, `help wanted`

Acceptance criteria:

- Document whether GitHub workflow annotations are useful for Maven findings.
- Include examples and limitations.
- Recommend next implementation steps.

### 47. Research GitLab merge request support

Labels: `stage: 3`, `area: gitlab`, `help wanted`

Acceptance criteria:

- Add a design note for GitLab MR support.
- Identify APIs needed for changed files and pipeline jobs.
- Keep GitLab support optional.

### 48. Add Maven 4 compatibility investigation

Labels: `stage: 3`, `area: maven`, `help wanted`

Acceptance criteria:

- Document Maven 4 report compatibility risks.
- Add fixtures only when Maven 4 behavior is stable enough.
- Do not declare production support prematurely.

### 49. Add agent evidence bundle output

Labels: `stage: 3`, `area: agent`, `help wanted`

Acceptance criteria:

- Design a JSON bundle for coding agents.
- Include findings, commands, confidence reasons, and relevant files when available.
- Avoid prompt-specific coupling.

### 50. Add CI artifact directory option

Labels: `stage: 3`, `area: cli`, `area: ci`, `help wanted`

Acceptance criteria:

- Add a flag for scanning a CI artifact directory separate from the source root.
- Preserve module mapping behavior where possible.
- Add tests with fixture artifact layout.

### 51. Add privacy guide for CI logs

Labels: `stage: 3`, `area: docs`, `good first contribution`, `agent-friendly`

Acceptance criteria:

- Document how to sanitize logs before opening issues.
- Explain what data should never be pasted publicly.
- Link from `SECURITY.md` and `CONTRIBUTING.md`.

### 52. Add GitHub Action usage example

Labels: `stage: 3`, `area: docs`, `area: ci`, `help wanted`

Acceptance criteria:

- Add a documented example workflow using PR Maven CLI.
- Show text and JSON output modes.
- Avoid requiring a token for local report parsing.

### 53. Add internal engineering platform integration guide

Labels: `stage: 3`, `area: docs`, `help wanted`

Acceptance criteria:

- Document how platform teams can consume JSON output.
- Include examples for local/self-hosted execution.
- Explain privacy and no-telemetry behavior.

### 54. Add package manager distribution research

Labels: `stage: 3`, `area: release`, `help wanted`

Acceptance criteria:

- Research Homebrew, Scoop, npm wrapper, and direct binary releases.
- Recommend a first distribution path.
- Keep release complexity proportional to project maturity.

### 55. Add stable output compatibility policy

Labels: `stage: 3`, `area: docs`, `help wanted`

Acceptance criteria:

- Define compatibility rules for CLI output and JSON fields.
- Explain deprecation expectations before `v1.0.0`.
- Link from README and JSON contract docs.
