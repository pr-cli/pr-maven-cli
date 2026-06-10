# v0.3.0 Release Acceptance Checklist

This checklist defines what must be true before tagging the Stage 3 target release, `v0.3.0`.

Use it together with the [release process](release.md), [MVP acceptance checklist](mvp-acceptance.md), [integration scope](integrations.md), [project visual map](project-visual-map.md), and [roadmap](../ROADMAP.md).

## Release Principle

Stage 3 may add pull request and CI context, but it must not weaken the local Maven analyzer.

Before `v0.3.0` is tagged:

- [ ] local Maven report parsing works without network access;
- [ ] local Maven report parsing works without GitHub, GitLab, CI provider, hosted-service, telemetry, or AI provider tokens;
- [ ] provider integrations are optional and additive;
- [ ] provider failures preserve local Maven findings instead of hiding them;
- [ ] every new provider-facing behavior is covered by fixtures, mocks, or deterministic local tests before relying on live APIs.

## Planning And Design Gates

Resolve or explicitly defer the relevant Stage 3 design issues before implementation-heavy issues are unblocked.

- [ ] [#82 Provider context planning map](https://github.com/pr-cli/pr-maven-cli/issues/82) defines the provider-context lane.
- [ ] [#83 Provider error and offline fallback taxonomy](https://github.com/pr-cli/pr-maven-cli/issues/83) defines graceful degradation.
- [ ] [#84 Provider adapter package boundaries](https://github.com/pr-cli/pr-maven-cli/issues/84) defines what can import provider-specific code.
- [ ] [#85 Provider fixture and mock contract](https://github.com/pr-cli/pr-maven-cli/issues/85) defines default test strategy.
- [ ] [#86 GitHub token and permission matrix](https://github.com/pr-cli/pr-maven-cli/issues/86) documents planned GitHub permissions.
- [ ] [#87 Changed-files fixture contract](https://github.com/pr-cli/pr-maven-cli/issues/87) defines PR file-change evidence.
- [ ] [#89 Check-runs fixture contract](https://github.com/pr-cli/pr-maven-cli/issues/89) defines CI check evidence.
- [ ] [#90 PR context JSON extension contract](https://github.com/pr-cli/pr-maven-cli/issues/90) defines additive machine-readable output.
- [ ] [#91 why, explain, and ci command UX boundaries](https://github.com/pr-cli/pr-maven-cli/issues/91) preserves command intent and exit-code behavior.
- [ ] [#92 GitLab parity boundary](https://github.com/pr-cli/pr-maven-cli/issues/92) keeps GitLab support optional and investigation-first.
- [ ] [#93 CI artifact directory fixture layout](https://github.com/pr-cli/pr-maven-cli/issues/93) defines local artifact-directory validation.
- [ ] [#94 Agent evidence bundle schema](https://github.com/pr-cli/pr-maven-cli/issues/94) stays additive and provider-neutral.

Implementation-heavy issues such as GitHub adapters, Markdown output, `explain`, `ci`, relevance scoring, baseline comparison, and evidence bundle output should remain blocked until the related design contract is merged or explicitly waived by a maintainer.

## Provider Adapter Acceptance

Provider adapters are release-ready only when:

- [ ] provider clients are behind interfaces and outside core Maven report parsing;
- [ ] default analysis does not create network clients or read provider tokens;
- [ ] missing-token, insufficient-permission, rate-limit, network-failure, not-found, and unsupported-state cases are tested;
- [ ] public and private repository permission expectations are documented;
- [ ] live-provider tests, if any, are opt-in and skipped by default;
- [ ] GitHub is documented as the first native provider target and GitLab remains optional until its parity boundary is accepted.

## Fixtures, Mocks, And CI Acceptance

Before `v0.3.0`:

- [ ] changed-files fixtures cover added, modified, deleted, renamed, and nested-module paths;
- [ ] check-runs fixtures cover success, failure, skipped, cancelled, timed out, and pending states;
- [ ] CI artifact-directory fixtures cover artifacts outside module `target` directories;
- [ ] provider fixtures are sanitized and safe to publish;
- [ ] default tests run without real provider credentials;
- [ ] `go test ./...` passes;
- [ ] `scripts/test.sh` or `scripts/test.ps1` passes when practical for the release environment;
- [ ] GitHub Actions `All CI checks` passes on the release commit.

## JSON And Compatibility Acceptance

New Stage 3 output is release-ready only when:

- [ ] new JSON fields are additive;
- [ ] existing `Report`, `Summary`, `Module`, and `Finding` fields remain compatible or the compatibility impact is documented;
- [ ] schema validation covers report-only and provider-context-enriched output;
- [ ] examples show both local-only and PR/CI-context output;
- [ ] consumers can safely ignore unknown future fields;
- [ ] Markdown or other human-readable outputs are deterministic and backed by examples or tests.

## CLI UX Acceptance

Before release:

- [ ] `fails` and `why` keep their Stage 1/2 behavior unless a compatibility note explains the change;
- [ ] any `explain` or `ci` command has documented intent, defaults, formats, and exit codes;
- [ ] CI-oriented behavior works with local report artifacts and optional provider context;
- [ ] user-facing command examples are covered by documented-command smoke tests when practical;
- [ ] README and usage docs clearly distinguish local analysis from optional provider enrichment.

## Documentation Acceptance

Before tagging `v0.3.0`:

- [ ] `README.md` links to installation, usage, testing, integration, release, and compatibility docs;
- [ ] `docs/integrations.md` documents current native provider support and planned boundaries;
- [ ] `docs/permissions.md` documents provider-token and repository-permission expectations;
- [ ] `docs/json-contract.md` documents any new public JSON fields;
- [ ] `docs/testing.md` explains provider fixtures, mocks, and opt-in live tests when they exist;
- [ ] `docs/release.md` links to this checklist;
- [ ] `docs/project-visual-map.md` remains aligned with issue labels and dependencies.

## Release Execution Acceptance

Before creating the tag:

- [ ] `main` is synchronized with `origin/main`;
- [ ] no release-blocking issue or pull request is open;
- [ ] the standard release readiness checklist in [release.md](release.md) is complete;
- [ ] release artifacts still include Linux, macOS, and Windows packages;
- [ ] checksums, SBOMs, and artifact attestations are generated;
- [ ] release notes identify Stage 3 scope and call out optional provider-context behavior.
