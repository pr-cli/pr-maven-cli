## Summary

Explain what changed and why.

## MVP Validation Checklist

Use the [testing guide](https://github.com/pr-cli/pr-maven-cli/blob/main/docs/testing.md) and [MVP acceptance checklist](https://github.com/pr-cli/pr-maven-cli/blob/main/docs/mvp-acceptance.md) to choose the checks that apply to this change.

- [ ] Relevant local checks were run, or this is a documentation-only/no-runtime-impact PR and the reason is explained below
- [ ] User-facing CLI, output, or JSON changes include focused CLI, smoke, golden, schema, or compatibility validation
- [ ] Parser, fixture, or report changes include sanitized fixtures and fixture integrity validation when applicable
- [ ] Documentation, examples, compatibility, and local-first impact are updated or called out when applicable

## Validation

- [ ] `go test ./...`
- [ ] `go vet ./...`
- [ ] CLI behavior tested when user-facing output changes
- [ ] Golden files updated when text output intentionally changes
- [ ] Documentation updated when contributor or user behavior changes
- [ ] No GitHub, GitLab, AI provider, telemetry, or external-service dependency was added to core analysis
- [ ] JSON compatibility was preserved or the compatibility impact is explained
- [ ] Branch name, PR title, and commit messages avoid coding agent or tool names

## Scope

- [ ] This PR is focused on one issue or one clearly bounded change
- [ ] No unrelated refactors
- [ ] No generated binaries, private logs, or local cache files committed
- [ ] Issue and PR labels match the type, area, stage, and status guidance when applicable
