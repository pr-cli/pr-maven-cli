# CI/CD

PR Maven CLI uses a Stage 1 OSS-style GitHub pipeline.

The pipeline is intentionally dependency-light. Core checks do not require Maven, Docker, private credentials, hosted services, or external test infrastructure.

This is project CI, not a runtime provider adapter. The Stage 1 CLI does not call GitHub APIs. Runtime/provider integration scope is documented in [integrations.md](integrations.md).

## CI Workflow

File: `.github/workflows/ci.yml`

Runs on:

- pull requests;
- pushes to `main`.

Jobs:

- `Agent name guard`: validates that branch names, pull request titles, and commit messages do not include coding agent or tool names.
- `Quality gate`: `gofmt`, `go vet`, and unit tests.
- `Go tests`: Linux, Windows, macOS, Go 1.22.x, and current stable Go.
- `Race detector`: `go test -race ./...` on Linux.
- `Coverage gate`: coverage profile with a 70% total coverage floor.
- `JSON schema validation`: generates demo JSON reports and validates them against `schema/prmaven-report.schema.json` through Go tests.
- `Fixture integrity validation`: runs `TestFixtureIntegrity` so committed demo and testdata fixture drift is reported before merge.
- `Build`: cross-platform binary builds for Linux, macOS, and Windows.
- `CLI smoke test`: exercises the compiled binary against demo fixtures.
- `All CI checks`: stable aggregate job for future branch protection.

## Security Workflow

File: `.github/workflows/security.yml`

Runs on:

- pull requests;
- pushes to `main`;
- weekly schedule;
- manual dispatch.

Jobs:

- `Go vulnerability check`: runs `govulncheck`.
- `CodeQL`: static analysis for Go.
- `Dependency review`: reviews dependency changes on pull requests.

`CodeQL` runs while the repository is public. During private stabilization phases, the job is skipped unless code scanning is enabled for the private repository in GitHub settings.

`Dependency review` is advisory while the repository dependency graph is unavailable. It remains visible in the Security workflow, but the Stage 1 protected merge gate is `All CI checks`.

## Contributor Acknowledgement Workflow

File: `.github/workflows/thank-contributor.yml`

Runs on:

- new issues;
- new pull requests.

The workflow posts the standard thank-you message from `.github/contributor-thanks.md`. It uses a hidden marker to avoid duplicate comments if the workflow is re-run.

For pull requests from forks, the workflow uses `pull_request_target` without checking out or executing contributor code. It reads `.github/contributor-thanks.md` from the base repository through the GitHub API and writes a comment. It receives `issues: write` and `pull-requests: write` only for that comment flow.

## Release Workflow

File: `.github/workflows/release.yml`

Runs on:

- tags matching `v*`;
- manual dispatch for package validation.

Release artifacts:

- Linux amd64 and arm64 tarballs.
- macOS amd64 and arm64 tarballs.
- Windows amd64 zip.
- SHA-256 checksum files.
- SPDX JSON SBOM files for each package.
- GitHub artifact attestations for packages and checksum files.
- GitHub SBOM attestations binding each package archive to its SBOM.

SBOM generation uses `anchore/sbom-action` with Syft pinned in the workflow. Update the pinned Syft version intentionally when refreshing the release supply-chain toolchain.

The tag version is embedded in the CLI through:

```text
prmaven version
```

The workflow defaults to `contents: read`. Package jobs receive `id-token: write` and `attestations: write` only to generate GitHub artifact and SBOM attestations. Only the release publishing job receives `contents: write`.

Repository permission posture is documented in [permissions.md](permissions.md).

Reusable cross-project CI/CD templates are maintained outside this public repository so this document stays focused on the PR Maven CLI pipeline.

## Local Parity

Before opening a PR, contributors should run:

```bash
sh scripts/quality.sh
go test ./pkg/prmaven -run TestGeneratedJSONReportsValidateAgainstSchema -v
go test ./pkg/prmaven -run TestFixtureIntegrity -v
PRMAVEN_COVERAGE=1 sh scripts/test.sh
sh scripts/build.sh
```

On Windows PowerShell:

```powershell
.\scripts\quality.ps1
go test ./pkg/prmaven -run TestGeneratedJSONReportsValidateAgainstSchema -v
go test ./pkg/prmaven -run TestFixtureIntegrity -v
.\scripts\test.ps1 -Coverage
.\scripts\build.ps1
```

## Branch Protection Recommendation

The `main` branch should be protected.

Recommended required status:

- `All CI checks`

`All CI checks` includes the public metadata guard so pull requests cannot merge when branch names, pull request titles, or commit messages include blocked coding agent or tool names.

Keep security checks visible, but avoid making scheduled security tooling a blocker for focused contributor PRs until the project has more maintainers.

Only users with maintainer-level repository permissions should merge pull requests. See [MAINTAINERS.md](../MAINTAINERS.md).

Delete pull request head branches after merge so the repository does not accumulate stale contribution branches.

The default branch should require at least one approving pull request review and a code owner review. While `@Will-thom` is the only maintainer and code owner, a documented founder bypass may remain enabled to avoid deadlocking founder-authored maintenance PRs.
