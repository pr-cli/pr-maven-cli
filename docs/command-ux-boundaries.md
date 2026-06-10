# Command UX Boundaries

This document defines how current and planned commands should relate to each other as PR Maven CLI grows through Stage 3.

The goal is to preserve Stage 1 and Stage 2 behavior while leaving clear room for future `explain` and `ci` commands.

Use this document with [Usage](usage.md), [Deterministic Output](deterministic-output.md), [JSON Contract](json-contract.md), [Markdown PR Summary Contract](markdown-pr-summary-contract.md), and [Provider Context Plan](provider-context-plan.md).

## Current Commands

Stage 1 and Stage 2 currently implement:

- `fails`;
- `why`;
- `help`;
- `version`.

Current analysis commands:

```bash
prmaven fails -project .
prmaven why -project .
prmaven why -project . -format json
```

Current behavior must remain compatible:

- `fails` and `why` both analyze local Maven report artifacts.
- Both commands support `-project`, `-format`, `-module`, and `-output`.
- Both commands run without provider tokens, network access, or live CI APIs.
- Text remains the default output format.
- JSON remains opt-in through `-format json`.
- Existing exit-code behavior remains unchanged.

## Command Intent

| Command | Intended audience | Primary question | Default output | Provider context |
| --- | --- | --- | --- | --- |
| `fails` | Developer or maintainer scanning a local failure | What Maven failures were found? | Text | Not required. |
| `why` | Developer, maintainer, CI, or tool that needs structured failure context | Why did this Maven validation fail and how can it be reproduced? | Text | Optional in Stage 3. |
| `explain` | Future deeper triage flow | What evidence, context, confidence, and relationships explain the failure? | Text or Markdown, to be decided by the implementation issue. | Optional and additive. |
| `ci` | Future CI-oriented wrapper | How should CI preserve Maven status while emitting reports and summaries? | Text summary plus files, to be decided by the implementation issue. | Optional and additive. |

## `fails`

`fails` should stay the simplest human-facing analysis command.

Expected intent:

- list supported Maven findings from local report artifacts;
- keep output concise;
- keep local report parsing as the source of truth;
- avoid provider-specific behavior by default;
- preserve current flags and exit codes.

Example:

```bash
prmaven fails -project .
```

`fails` may support future output formats only when they do not change the command's simple failure-listing intent.

## `why`

`why` should remain the default explanation command.

Expected intent:

- explain what failed;
- identify the Maven module, plugin, phase, report path, and source detail;
- include reproduction commands;
- emit text or JSON;
- support output files;
- support module filtering;
- allow optional Stage 3 provider context without requiring it.

Examples:

```bash
prmaven why -project .
prmaven why -project . -format json
prmaven why -project . -module payment-core
prmaven why -project . -format json -output prmaven-report.json
```

Future provider-aware `why` behavior must be additive. A missing token, unavailable provider, or partial provider response must not suppress local Maven findings.

## `explain`

`explain` is reserved for future deeper explanation workflows.

Expected intent:

- expand on why a finding matters;
- connect local Maven evidence with optional provider context;
- show confidence reasons more prominently;
- include changed-file relevance and related check-run context when available;
- stay deterministic and fixture-driven by default.

`explain` should not become a free-form model prompt or remote interpretation layer. It should explain evidence that PR Maven CLI can represent deterministically.

Possible future examples:

```bash
prmaven explain -project .
prmaven explain -project . -format markdown
prmaven explain -project . -module payment-core
```

The implementation issue for `explain` should decide final supported formats and flags. It must not remove or change existing `fails` and `why` behavior.

## `ci`

`ci` is reserved for future CI-oriented workflows.

Expected intent:

- make CI usage easier without hiding the original Maven exit status;
- generate machine-readable reports and optional Markdown summaries;
- support artifact-directory layouts when implemented;
- preserve local-first analysis;
- avoid live provider calls by default.

Possible future examples:

```bash
prmaven ci -project . -format json -output prmaven-report.json
prmaven ci -project . -summary-output prmaven-summary.md
prmaven ci -project . -artifacts ./.ci-artifacts
```

The `ci` command should help CI jobs collect context. It should not become a replacement for running Maven itself.

## Exit-Code Expectations

Current exit codes remain the base contract:

| Exit code | Meaning |
| --- | --- |
| `0` | Analysis completed and no findings were found. |
| `1` | Analysis completed with Maven failure findings, or analysis failed. |
| `2` | Invalid CLI usage. |

Future commands should follow the same default expectations unless a command-specific issue documents a narrower behavior.

CI-specific behavior should preserve the Maven command's original status outside PR Maven CLI. The recommended pattern remains:

```bash
set +e
mvn -B verify
maven_status=$?

prmaven why -project . || true
prmaven why -project . -format json -output prmaven-report.json || true

exit "$maven_status"
```

## Output Format Expectations

Current supported formats:

- `text`;
- `json`.

Future candidate formats:

- `markdown`, after the [Markdown PR Summary Contract](markdown-pr-summary-contract.md) is implemented;
- provider-context-enriched JSON, after the [PR Context JSON Extension Contract](pr-context-json-extension.md) is implemented.

New formats must be explicit. Do not change the default output format of current commands without a breaking-change note.

## Local Developer Examples

Scan the current workspace:

```bash
prmaven fails -project .
```

Get explanation with reproduction commands:

```bash
prmaven why -project .
```

Focus on one module:

```bash
prmaven why -project . -module payment-core
```

Write JSON for local inspection:

```bash
prmaven why -project . -format json -output prmaven-report.json
```

## CI Examples

Preserve Maven status and publish PR Maven CLI context:

```bash
set +e
mvn -B verify
maven_status=$?

prmaven why -project . -format text -output prmaven-summary.txt || true
prmaven why -project . -format json -output prmaven-report.json || true

exit "$maven_status"
```

Future CI summary generation should be additive:

```bash
prmaven ci -project . -format json -output prmaven-report.json
prmaven ci -project . -summary-output prmaven-summary.md
```

## Non-Goals

This contract does not implement:

- the `explain` command;
- the `ci` command;
- Markdown output;
- live provider adapters;
- provider-token resolution;
- artifact upload;
- pull request commenting.

Those behaviors should land through focused implementation issues after their contracts are accepted.
