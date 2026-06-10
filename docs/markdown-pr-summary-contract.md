# Markdown PR Summary Contract

This document defines the Stage 3 Markdown PR summary output contract.

The contract exists before `-format markdown` or PR-comment output is implemented so contributors can agree on content, ordering, deterministic behavior, and provider-context boundaries first.

Use this document with [Deterministic Output](deterministic-output.md), [JSON Contract](json-contract.md), [PR Context JSON Extension Contract](pr-context-json-extension.md), [Confidence](confidence.md), and [Provider Errors And Offline Fallbacks](provider-errors-and-fallbacks.md).

## Goals

- Produce a compact Markdown summary suitable for pull request comments.
- Preserve deterministic section order and finding order.
- Show Maven findings, reproduction commands, confidence reasons, and optional provider context.
- Avoid relying on remote provider data to create Maven findings.
- Keep no-finding output clear and quiet.

## Required Sections

Markdown PR summaries should use these sections in this order:

1. Title.
2. Summary.
3. Findings.
4. Reproduction.
5. Confidence.
6. Provider Context.

Sections may be omitted only when they are not applicable and the omission is deterministic. For example, `Provider Context` may be omitted when provider context was not requested.

## Deterministic Rules

For the same report data, Markdown output must keep:

- the same section order;
- the same finding order as JSON output;
- the same stable finding IDs;
- the same module paths and report paths;
- the same reproduction commands;
- the same confidence values and reasons;
- slash-separated paths across platforms;
- bounded wording for provider diagnostics.

Markdown output may improve wording over time, but wording changes should be treated as user-facing output changes and covered by tests or golden fixtures once implementation begins.

## PR Comment Constraints

The output should be suitable for GitHub pull request comments and similar review surfaces:

- avoid excessively long raw logs;
- avoid raw stack traces unless a future issue scopes truncation rules;
- avoid secrets, tokens, signed URLs, actor emails, or private account IDs;
- prefer fenced code blocks for shell commands;
- keep one finding block concise enough to scan;
- include provider context as supporting metadata, not as the main evidence.

## Finding Block

Each finding should include:

- stable finding ID;
- module path;
- report kind;
- Maven plugin and phase;
- test class, test name, source file, rule, bug type, or threshold when available;
- one-line message when available;
- report path;
- reproduction command;
- confidence value;
- confidence reasons.

Recommended finding shape:

```md
### Finding 1: payment-core

- ID: `payment-core-paymentroundingtest-shouldrejectinvalidscale-failure`
- Source: `surefire`
- Plugin: `maven-surefire-plugin`
- Phase: `test`
- Test: `dev.prmaven.demo.PaymentRoundingTest#shouldRejectInvalidScale`
- Report: `payment-core/target/surefire-reports/TEST-dev.prmaven.demo.PaymentRoundingTest.xml`
- Message: expected:<2.50> but was:<2.49>
```

## Reproduction Section

Reproduction commands should be copyable and deterministic.

Use one fenced block per unique command:

````md
## Reproduction

```bash
mvn -pl payment-core -am -Dtest=PaymentRoundingTest test
```
````

When multiple findings share the same command, print it once and reference the related finding IDs.

## Confidence Section

Confidence should explain why the finding is trusted.

Recommended shape:

```md
## Confidence

- `payment-core-paymentroundingtest-shouldrejectinvalidscale-failure`: `high`
  - failure was found in a Maven Surefire JUnit XML report
  - report path maps to Maven module payment-core
```

Provider relevance, changed files, or check-run metadata may appear as provider-context reasons, but they must not replace local Maven confidence reasons.

## Provider Context Section

Provider context is optional and additive.

When available, include concise provider context:

- provider name;
- repository visibility when known;
- pull request number when known;
- changed-file relevance;
- related check-run names and states;
- provider diagnostics for partial or unavailable context.

Recommended shape:

```md
## Provider Context

- Provider: `github`
- Pull request: `#123`
- Context state: `available`
- Changed files: module `payment-core` changed
- Related checks: `Maven test` completed with `failure`
```

If provider context is unavailable after being requested:

```md
## Provider Context

- Context state: `unavailable`
- Diagnostic: `missing_token`
- Local Maven findings were preserved.
```

## Example: One Finding

````md
# PR Maven CLI Summary

## Summary

- Findings: 1
- Modules with findings: 1
- Reports analyzed: 1

## Findings

### Finding 1: payment-core

- ID: `payment-core-paymentroundingtest-shouldrejectinvalidscale-failure`
- Source: `surefire`
- Plugin: `maven-surefire-plugin`
- Phase: `test`
- Test: `dev.prmaven.demo.PaymentRoundingTest#shouldRejectInvalidScale`
- Report: `payment-core/target/surefire-reports/TEST-dev.prmaven.demo.PaymentRoundingTest.xml`
- Message: expected:<2.50> but was:<2.49>

## Reproduction

```bash
mvn -pl payment-core -am -Dtest=PaymentRoundingTest test
```

## Confidence

- `payment-core-paymentroundingtest-shouldrejectinvalidscale-failure`: `high`
  - failure was found in a Maven Surefire JUnit XML report
  - report path maps to Maven module payment-core

## Provider Context

- Context state: `not_requested`
````

## Example: Multiple Findings

````md
# PR Maven CLI Summary

## Summary

- Findings: 2
- Modules with findings: 2
- Reports analyzed: 2

## Findings

### Finding 1: payment-api

- ID: `payment-api-paymentapiit-shouldexposepaymentsummary-error`
- Source: `failsafe`
- Plugin: `maven-failsafe-plugin`
- Phase: `verify`
- Test: `dev.prmaven.demo.PaymentApiIT#shouldExposePaymentSummary`
- Report: `payment-api/target/failsafe-reports/TEST-dev.prmaven.demo.PaymentApiIT.xml`
- Message: java.net.ConnectException: connection refused

### Finding 2: payment-core

- ID: `payment-core-paymentroundingtest-shouldrejectinvalidscale-failure`
- Source: `surefire`
- Plugin: `maven-surefire-plugin`
- Phase: `test`
- Test: `dev.prmaven.demo.PaymentRoundingTest#shouldRejectInvalidScale`
- Report: `payment-core/target/surefire-reports/TEST-dev.prmaven.demo.PaymentRoundingTest.xml`
- Message: expected:<2.50> but was:<2.49>

## Reproduction

```bash
mvn -pl payment-api -am -Dit.test=PaymentApiIT verify
```

```bash
mvn -pl payment-core -am -Dtest=PaymentRoundingTest test
```

## Confidence

- `payment-api-paymentapiit-shouldexposepaymentsummary-error`: `high`
  - failure was found in a Maven Failsafe JUnit XML report
  - report path maps to Maven module payment-api
- `payment-core-paymentroundingtest-shouldrejectinvalidscale-failure`: `high`
  - failure was found in a Maven Surefire JUnit XML report
  - report path maps to Maven module payment-core

## Provider Context

- Provider: `github`
- Pull request: `#123`
- Context state: `available`
- Changed modules: `payment-api`, `payment-core`
- Related checks: `Maven test` completed with `failure`
````

## Example: No Findings

```md
# PR Maven CLI Summary

## Summary

- Findings: 0
- Modules discovered: 3
- Reports analyzed: 2

## Findings

No supported Maven failures, violations, bugs, rules, or thresholds were found in the analyzed report artifacts.

## Reproduction

No reproduction command is available because there are no findings.

## Confidence

No findings were emitted.

## Provider Context

- Context state: `not_requested`
```

## Non-Goals

The Markdown summary contract does not define:

- automatic pull request commenting;
- comment update or deduplication behavior;
- GitHub annotation output;
- SARIF output;
- raw CI log rendering;
- live provider calls;
- provider-token resolution.

Those behaviors should be handled by separate issues.
