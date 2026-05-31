# Maintainers

PR Maven CLI is founder-led.

## Principal Maintainer

- Will-thom, GitHub: [@Will-thom](https://github.com/Will-thom)
- Role: Founder and Principal Maintainer
- Scope: product direction, roadmap, releases, merge policy, maintainer access, and final project decisions.
- Repository-wide code owner: `@Will-thom`, through `.github/CODEOWNERS`.

## Merge Policy

Only maintainers with repository write, maintain, or admin permissions may merge pull requests.

External contributors are welcome to open issues and pull requests. A maintainer is responsible for checking:

- issue and pull request labels;
- roadmap stage or milestone;
- CI status;
- security checks when applicable;
- scope and reviewability;
- tests and documentation;
- JSON and CLI compatibility.

The default branch is `main`. It should stay protected and use the `All CI checks` status as the required CI gate.

GitHub repository permissions should keep merge rights limited to users with write, maintain, or admin access. External contributors can propose changes through issues and pull requests, but they should not have merge access.

Label policy is documented in [docs/labels.md](docs/labels.md). Pull requests should not be merged while unlabeled when they affect release notes, roadmap order, or contributor-facing backlog state.

## Maintainer Path

Additional maintainers may be invited after sustained, high-quality contributions and review participation.

Maintainer access should stay intentionally small while the project is young.
