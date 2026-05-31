# Release Process

PR Maven CLI releases are driven by Git tags.

Use this process together with the [MVP acceptance checklist](mvp-acceptance.md) for `v0.x` releases.

For the first MVP release record, see the [v0.1.0 release snapshot](release-snapshot-v0.1.0.md).

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
- [ ] SPDX JSON SBOM generation is enabled for every release package.
- [ ] GitHub artifact attestation generation is enabled for release packages and checksums.
- [ ] GitHub SBOM attestation generation is enabled for release packages.
- [ ] The tag version will be embedded in `prmaven version`.

After the release workflow completes:

- [ ] The GitHub release exists for the pushed tag.
- [ ] Linux amd64 and arm64 tarballs are attached.
- [ ] macOS amd64 and arm64 tarballs are attached.
- [ ] Windows amd64 zip is attached.
- [ ] A `.sha256` file exists for every package.
- [ ] A `.sbom.spdx.json` file exists for every package.
- [ ] A `.sbom.spdx.json.sha256` file exists for every package SBOM.
- [ ] GitHub artifact attestations exist for every package and checksum file.
- [ ] GitHub SBOM attestations bind package archives to their SBOM files.
- [ ] Release notes were generated.

## Create a Release

1. Ensure `main` is green.
2. Choose a semantic version such as `v0.1.0`.
3. Create and push a signed annotated tag:

```bash
git tag -s v0.1.0 -m "PR Maven CLI v0.1.0"
git tag -v v0.1.0
git push origin v0.1.0
```

Signed tags require a local GPG, SSH, or S/MIME signing key configured in Git and added to GitHub. Do not store release signing private keys in repository secrets.

## Signed Tag Policy

Release tags should be signed annotated tags created from a local maintainer machine.

Preferred command:

```bash
git tag -s v0.1.0 -m "PR Maven CLI v0.1.0"
git tag -v v0.1.0
```

Operational rules:

- The signing private key must stay outside the repository and outside GitHub Actions secrets.
- The release owner should verify the tag locally before pushing it.
- GitHub should show the pushed tag as verified when the signing public key is configured on the maintainer account.
- If an emergency `v0.x` release must be cut without a signing key, document the exception in the release notes and replace it with signed tags for normal releases.

The release workflow will:

- build Linux, macOS, and Windows binaries;
- package archives;
- generate SHA-256 checksums;
- generate SPDX JSON SBOM files and SBOM checksums;
- generate GitHub artifact attestations for package archives and checksum files;
- generate GitHub SBOM attestations that bind package archives to their SBOM files;
- create a GitHub release;
- generate release notes from GitHub metadata.

## Verify Release Provenance

Release artifacts are attested with GitHub artifact attestations.

After downloading a package or checksum file, verify its provenance with:

```bash
gh attestation verify prmaven-v0.1.0-linux-amd64.tar.gz -R pr-cli/pr-maven-cli
gh attestation verify prmaven-v0.1.0-linux-amd64.tar.gz.sha256 -R pr-cli/pr-maven-cli
gh attestation verify prmaven-v0.1.0-linux-amd64.sbom.spdx.json -R pr-cli/pr-maven-cli
```

Use the matching file name for the target platform and release version.

Validate the SBOM checksum with:

```bash
sha256sum -c prmaven-v0.1.0-linux-amd64.sbom.spdx.json.sha256
```

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
