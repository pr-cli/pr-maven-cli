# CI Log Privacy

PR Maven CLI is designed for local-first Maven failure analysis. Public issues,
pull requests, fixtures, and examples should use the smallest sanitized evidence
needed to reproduce or explain a failure.

Use this guide before pasting CI logs, Maven reports, workflow output, or provider
metadata into a public issue, pull request, fixture, or discussion.

## Sanitization Workflow

1. Start from the local Maven artifact when possible, such as a Surefire XML
   report, Failsafe XML report, Checkstyle report, SpotBugs report, JaCoCo log,
   or Maven Enforcer log. Prefer the report slice that proves the failure over a
   full CI console log.
2. Trim the log to the smallest useful excerpt. Keep the command, plugin name,
   module path, failing test or rule name, and the relevant error text.
3. Replace private values with stable placeholders:
   - repository, organization, customer, and service names: `<ORG>`, `<REPO>`,
     `<CUSTOMER>`, `<SERVICE>`;
   - users and email addresses: `<USER>`, `<EMAIL>`;
   - internal hosts, IPs, and URLs: `<HOST>`, `<IP>`, `<URL>`;
   - local paths: `<PROJECT_ROOT>`, `<HOME>`, `<WORKSPACE>`;
   - credentials and identifiers: `<TOKEN>`, `<SECRET>`, `<KEY_ID>`;
   - timestamps or run IDs when not needed for reproduction: `<TIMESTAMP>`,
     `<RUN_ID>`.
4. Re-read the sanitized excerpt before posting it. Check both the visible text
   and any attached files.
5. If a failure can only be explained with sensitive data, do not post it
   publicly. Follow the private reporting path in `SECURITY.md`.

## Never Paste Publicly

Do not paste or commit:

- API keys, access tokens, refresh tokens, session cookies, SSH keys, private
  keys, signing keys, or cloud credentials;
- GitHub, GitLab, CI, artifact registry, package registry, or cloud provider
  tokens;
- customer names, tenant names, account IDs, invoice data, billing exports, or
  production identifiers;
- proprietary source code, private dependency coordinates, private repository
  names, internal hostnames, internal URLs, or private IP ranges;
- full CI logs when a short Maven report excerpt is enough;
- `.env` files, secret manager output, kubeconfig files, cloud CLI profiles, or
  downloaded artifacts from private CI runs.

## Good Public Evidence

Good public evidence is:

- small enough for a maintainer to review quickly;
- deterministic enough to become a fixture or regression test;
- scrubbed of private names and credentials;
- explicit about the Maven module, plugin, phase, and command involved;
- reproducible without live provider credentials or external services.

Example:

```text
module: service-core
plugin: maven-surefire-plugin
phase: test
command: mvn -pl service-core test
failure: ExampleServiceTest.shouldRejectInvalidInput
message: expected status <400> but was <200>
workspace: <PROJECT_ROOT>
```

## Fixture And Issue Expectations

When adding a fixture, follow [Fixture Notes](fixtures.md). Sanitized fixtures
should remain stable, inspectable, and safe to publish.

When opening an issue or pull request:

- describe the failure in plain language;
- attach the smallest sanitized artifact that demonstrates it;
- say whether the excerpt came from local Maven output, a CI log, or a report
  artifact;
- avoid asking maintainers to inspect private logs or credentials.
