# MVP Acceptance Checklist

This checklist turns [Final Testable Delivery](final-testable-delivery.md) into a maintainer-facing acceptance gate for the production-usable local MVP.

Use it before release tags, large validation changes, and Stage 3 provider-context work.

For the Stage 2.1 issue order and dependency map, see the [MVP hardening plan](mvp-hardening-plan.md).

## Functional Acceptance

- [ ] `prmaven fails -project demo/multi-module-failure` reports fixture-backed Maven findings.
- [ ] `prmaven why -project demo/multi-module-failure -format json` emits parseable JSON.
- [ ] `prmaven fails -project demo/no-failure` exits successfully and reports zero findings.
- [ ] `prmaven why -project demo/multi-module-failure -module payment-core` limits output to the selected module.
- [ ] `prmaven why -project demo/multi-module-failure -format json -output prmaven-report.json` writes a report file.

## Validation Acceptance

- [ ] `go test ./...` passes.
- [ ] CLI end-to-end tests cover findings, no findings, invalid usage, JSON output, module filtering, and output files.
- [ ] Parser tests cover Surefire, Failsafe, Checkstyle, SpotBugs, Enforcer, and JaCoCo fixtures.
- [ ] Golden tests protect human-readable text output.
- [ ] JSON schema tests validate the generated report shape.
- [ ] Path normalization tests cover Windows-style and POSIX-style paths.

## Safety Acceptance

- [ ] Core analysis requires no network access.
- [ ] Core analysis requires no GitHub, GitLab, CI provider, or AI provider token.
- [ ] Core analysis does not upload reports or logs.
- [ ] Unsupported artifacts do not trigger speculative findings.
- [ ] Provider integrations remain optional and outside the core analyzer.
- [ ] Maven 4 is not declared production-ready before compatibility is validated.

## Documentation Acceptance

- [ ] README links to installation, usage, JSON contract, testing, CI, release, and contribution docs.
- [ ] Installation docs describe source builds and release artifacts.
- [ ] Usage docs include text, JSON, module filter, and output-file examples.
- [ ] JSON contract docs describe public report fields and compatibility expectations.
- [ ] Testing docs explain local, CI, fixture, and golden-test coverage.
- [ ] Integration docs state that Stage 1 has no native GitHub or GitLab runtime adapter.
- [ ] Permission docs explain the local-first/no-token posture.

## Release Readiness Acceptance

- [ ] `main` is synchronized with `origin/main`.
- [ ] Branch protection requires `All CI checks`.
- [ ] The release workflow exists and is triggered by `v*` tags.
- [ ] Release artifacts include Linux, macOS, and Windows packages.
- [ ] SHA-256 checksum files are generated for release packages.
- [ ] The tag version is embedded in `prmaven version`.
- [ ] The release notes are generated from GitHub metadata.

## Stage 3 Gate

Before Stage 3 implementation issues begin, confirm:

- [ ] the MVP acceptance checklist has been reviewed;
- [ ] the release readiness checklist has been reviewed;
- [ ] local-first behavior remains documented and tested;
- [ ] provider-context work is additive and optional.
