# CI/CD Reference Template For OSS Projects

This document is a reusable CI/CD reference for new open source projects. It is based on the PR Maven CLI pipeline and should be adapted to each project's language, package manager, release model, and contributor maturity.

Use it as a checklist and design baseline. Do not copy workflows blindly; keep the structure, replace the language-specific commands, and preserve the security posture.

## Goals

- Keep pull requests cheap to validate.
- Make the protected branch depend on one stable aggregate status.
- Separate product CI, security checks, contributor automation, and release packaging.
- Use minimum GitHub token permissions by default.
- Keep release artifacts verifiable through checksums, SBOMs, and attestations.
- Keep the first production MVP usable without private services or paid infrastructure.

## Recommended Workflow Files

Create these workflows first:

- `.github/workflows/ci.yml`
- `.github/workflows/security.yml`
- `.github/workflows/release.yml`
- `.github/workflows/thank-contributor.yml`

Add these repository support files:

- `.github/CODEOWNERS`
- `.github/dependabot.yml`
- `.github/contributor-thanks.md`
- `.github/ISSUE_TEMPLATE/*`
- `.github/PULL_REQUEST_TEMPLATE.md`
- `docs/ci.md`
- `docs/release.md`
- `docs/permissions.md`
- `docs/oss-guardrails.md`
- `SECURITY.md`
- `CONTRIBUTING.md`
- `MAINTAINERS.md`

## Default Permissions

Every workflow should start with read-only permissions unless a job has a specific reason to write:

```yaml
permissions:
  contents: read
```

Grant write permissions only at the job level. Examples:

- `contents: write` only for the release publishing job.
- `id-token: write` and `attestations: write` only for jobs that generate artifact or SBOM attestations.
- `issues: write` and `pull-requests: write` only for comment automation that needs to post thank-you messages.
- `security-events: write` only for CodeQL or scanners that upload security results.

## CI Workflow

Purpose: validate pull requests and pushes to the default branch.

Recommended triggers:

```yaml
on:
  pull_request:
  push:
    branches:
      - main
```

Recommended controls:

```yaml
permissions:
  contents: read

concurrency:
  group: ci-${{ github.workflow }}-${{ github.ref }}
  cancel-in-progress: true
```

Recommended jobs:

- `Quality gate`: formatting, linting, static checks, and focused unit tests.
- `Test matrix`: supported operating systems and supported language runtime versions.
- `Race or concurrency check`: when the language has a useful race/concurrency detector.
- `Coverage gate`: generate coverage and enforce the MVP coverage floor.
- `Build`: compile or package the production artifact for supported platforms.
- `Smoke test`: run the built artifact against demo fixtures or realistic examples.
- `All CI checks`: aggregate job that depends on all required CI jobs.

The protected branch should require only the aggregate job:

```yaml
all-checks:
  name: All CI checks
  runs-on: ubuntu-latest
  needs:
    - quality
    - test
    - coverage
    - build
    - smoke
  steps:
    - name: Report success
      run: echo "All CI checks passed."
```

Use a stable aggregate check name so branch protection does not need to change whenever the matrix expands.

## Language-Specific Substitutions

Replace the commands, not the structure:

| Stack | Quality gate | Test matrix | Build | Security scanner |
| --- | --- | --- | --- | --- |
| Go | `gofmt`, `go vet`, `go test` | `go test ./...` | `go build` | `govulncheck` |
| Java/Maven | `mvn -B verify` or split `checkstyle`/`test` | Maven on supported JDKs | `mvn -B package` | Dependency review, CodeQL |
| Node | `npm ci`, lint, typecheck | `npm test` | `npm run build` | `npm audit` or ecosystem scanner |
| Python | formatter/linter/typecheck | `pytest` | package build | `pip-audit` or ecosystem scanner |
| Rust | `cargo fmt --check`, `cargo clippy`, tests | `cargo test` | `cargo build --release` | `cargo audit` |

For the MVP, prefer deterministic commands that run without private tokens, paid services, cloud accounts, or remote fixtures.

## Security Workflow

Purpose: run security checks without making every external security tool a hard merge blocker too early.

Recommended triggers:

```yaml
on:
  pull_request:
  push:
    branches:
      - main
  schedule:
    - cron: "17 9 * * 1"
  workflow_dispatch:
```

Recommended jobs:

- language vulnerability scanner;
- CodeQL when available for the language and repository visibility;
- Dependency Review on pull requests.

Recommended policy:

- Keep security checks visible on pull requests.
- Keep the protected branch gate focused on `All CI checks` until the maintainer pool is large enough to handle scanner noise quickly.
- Keep scheduled security checks enabled so drift appears even when no PR is open.
- Skip CodeQL automatically when a private repository or plan does not support code scanning.
- Use `continue-on-error: true` only when the tool is advisory by design and the policy is documented.

## Contributor Thank-You Workflow

Purpose: every issue and PR receives a consistent welcome message.

Recommended behavior:

- Run on new issues and pull requests.
- Read the message from `.github/contributor-thanks.md`.
- Add a hidden marker so reruns do not duplicate comments.
- For fork PRs, use `pull_request_target` only if the workflow does not check out or execute contributor code.
- Never run contributor code in a `pull_request_target` workflow.

## Release Workflow

Purpose: build release artifacts from tags and optional manual dry runs.

Recommended triggers:

```yaml
on:
  push:
    tags:
      - "v*"
  workflow_dispatch:
```

Recommended release outputs:

- package artifacts for supported operating systems and architectures;
- SHA-256 checksum files;
- SPDX JSON SBOM files for each package;
- artifact attestations for release packages;
- provenance attestations for checksum files;
- SBOM attestations binding each package archive to its SBOM;
- GitHub release created only for version tags.

Recommended release permissions:

```yaml
permissions:
  contents: read
```

Package jobs:

```yaml
permissions:
  contents: read
  id-token: write
  attestations: write
```

Release publishing job:

```yaml
permissions:
  contents: write
```

Recommended release policy:

- Use signed annotated tags for real releases.
- Keep private signing keys outside repository secrets.
- Use manual dispatch only for package validation or dry-run builds.
- Pin the SBOM generator version intentionally.
- Attach checksums and SBOM files to every release.
- Document how users verify checksums and attestations.

## Branch Protection

Protect `main` before inviting external contributors.

Recommended rules:

- Require pull requests before merge.
- Require at least one approving review.
- Require CODEOWNER review.
- Require conversation resolution.
- Require `All CI checks`.
- Block force pushes.
- Block branch deletion.
- Delete pull request head branches after merge.
- Allow maintainer-only merges.

If the founder is the only maintainer, document any temporary founder bypass so maintenance PRs do not deadlock. Remove or narrow the bypass after adding another trusted maintainer.

## Repository Security Settings

Enable everything available on the free plan when possible:

- secret scanning;
- secret scanning push protection;
- Dependabot security updates;
- Dependabot version updates for GitHub Actions and the project package manager;
- dependency graph;
- branch protection;
- code owner file.

Defer these until the project has enough maintainer capacity:

- GitHub Actions SHA pinning;
- organization-level action allowlists;
- enforced signed-tag rulesets;
- SBOM vulnerability severity gates;
- artifact signing beyond GitHub attestations.

## Release Tag Policy

Use signed annotated tags:

```bash
git tag -s v0.1.0 -m "Project Name v0.1.0"
git tag -v v0.1.0
git push origin v0.1.0
```

If GPG signing is not ready, document the release as unsigned and do not pretend otherwise. Prefer solving maintainer signing before the first public stability push.

## Required Documentation

Each project should document:

- what each workflow does;
- what runs locally before opening a PR;
- which status check protects `main`;
- which jobs are advisory;
- which jobs require paid-plan features, if any;
- how releases are cut;
- how artifacts, checksums, SBOMs, and attestations are verified;
- who can merge;
- how CODEOWNERS works;
- what contributors should do when CI fails.

## MVP Acceptance Checklist

Use this before calling the first version production-usable:

- [ ] `ci.yml` runs on PRs and pushes to `main`.
- [ ] CI has quality, test, build, smoke, and aggregate jobs.
- [ ] The aggregate job is named `All CI checks`.
- [ ] `main` requires `All CI checks`.
- [ ] `main` requires pull request review and CODEOWNER review.
- [ ] Workflow default permissions are read-only.
- [ ] Security workflow runs on PRs, `main`, schedule, and manual dispatch.
- [ ] Dependabot is configured.
- [ ] Contributor thank-you workflow is safe for forks.
- [ ] Release workflow creates checksums.
- [ ] Release workflow creates SBOMs.
- [ ] Release workflow creates artifact attestations.
- [ ] Release docs explain signed tags.
- [ ] Local parity commands are documented.
- [ ] CI and release docs avoid promising unsupported providers or paid features.

## What Not To Do In The First MVP

- Do not require private tokens for core CI.
- Do not make flaky external services part of the protected merge gate.
- Do not put broad write permissions at the workflow level.
- Do not execute fork code from `pull_request_target`.
- Do not require every security scanner to block merges before the maintainer workflow can absorb the noise.
- Do not promise signed releases, SBOM gates, or package-manager distribution before the implementation exists.
- Do not make the branch protection depend on many matrix job names when one aggregate check is enough.

## Copy Strategy For New Projects

1. Copy this document into the new repository.
2. Copy the PR Maven CLI workflow structure.
3. Replace Go-specific commands with the target project's language commands.
4. Keep the aggregate job and permission model.
5. Add project-specific demos or fixtures for the smoke test.
6. Create `docs/ci.md`, `docs/release.md`, and `docs/permissions.md`.
7. Enable free GitHub security settings.
8. Protect `main`.
9. Open a setup PR and let CI prove the baseline.
10. Merge only after `All CI checks` is green.
