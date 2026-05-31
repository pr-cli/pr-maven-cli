# Release Process

PR Maven CLI releases are driven by Git tags.

Use this process together with the [MVP acceptance checklist](mvp-acceptance.md) for `v0.x` releases.

## Release Readiness Checklist

Before creating a release tag:

- [ ] `main` is synchronized with `origin/main`.
- [ ] No release-blocking pull request is open.
- [ ] `go test ./...` passes locally.
- [ ] `scripts/test.sh` or `scripts/test.ps1` passes when practical for the release environment.
- [ ] CI has passed on the release commit, including `All CI checks`.
- [ ] JSON output from `demo/multi-module-failure` is parseable.
- [ ] `schema/prmaven-report.schema.json` matches the generated report contract.
- [ ] README examples, usage docs, installation docs, and release docs point at the current repository.
- [ ] Demo and fixture docs describe committed Maven report artifacts.
- [ ] Release workflow exists for Linux, macOS, and Windows packages.
- [ ] SHA-256 checksum generation is enabled.
- [ ] The tag version will be embedded in `prmaven version`.

After the release workflow completes:

- [ ] The GitHub release exists for the pushed tag.
- [ ] Linux amd64 and arm64 tarballs are attached.
- [ ] macOS amd64 and arm64 tarballs are attached.
- [ ] Windows amd64 zip is attached.
- [ ] A `.sha256` file exists for every package.
- [ ] Release notes were generated.

## Create a Release

1. Ensure `main` is green.
2. Choose a semantic version such as `v0.1.0`.
3. Create and push the tag:

```bash
git tag v0.1.0
git push origin v0.1.0
```

The release workflow will:

- build Linux, macOS, and Windows binaries;
- package archives;
- generate SHA-256 checksums;
- create a GitHub release;
- generate release notes from GitHub metadata.

## Validate a Local Build

```bash
sh scripts/build.sh dist dev
./dist/prmaven version
```

On Windows PowerShell:

```powershell
.\scripts\build.ps1 -Version dev
.\dist\prmaven.exe version
```

## Version Contract

Release builds embed the tag in the CLI:

```bash
prmaven version
```

Development builds report `dev` unless a script or workflow passes a specific version.
