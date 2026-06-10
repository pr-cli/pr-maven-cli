# Deterministic Output

PR Maven CLI uses deterministic output to mean that the same supported Maven artifacts, in the same project layout, produce the same explanation and the same structured report.

Deterministic output is important because developers, maintainers, CI jobs, and automated tooling need stable evidence before deciding what failed and how to reproduce it.

## Deterministic Text

Deterministic text is the human-readable output printed by commands such as:

```bash
prmaven fails -project .
```

For the same issue in the same project, deterministic text should keep:

- the same section order;
- the same field order;
- the same module, plugin, phase, report, test, message, confidence, and reproduction content when the source artifacts do not change;
- the same no-failure wording when no supported findings are present.

The text format is meant for humans, so wording may improve over time. When wording changes intentionally, golden tests should be updated in the same pull request and the change should be documented as user-facing output.

Future Markdown PR summary output has a separate Stage 3 contract in [Markdown PR Summary Contract](markdown-pr-summary-contract.md).

## Deterministic JSON

Deterministic JSON is the structured output printed by commands such as:

```bash
prmaven why -project . -format json
```

For the same issue in the same project, deterministic JSON should keep:

- the same top-level structure;
- the same field names;
- the same stable field ordering produced by the CLI;
- the same values when the source artifacts do not change;
- slash-separated relative paths, including on Windows;
- additive compatibility for consumers that ignore unknown future fields.

The JSON format is meant for CI systems, bots, internal tools, and automated coding workflows. Its field-level contract is documented in [JSON Contract](json-contract.md).

## What Deterministic Does Not Mean

Deterministic output does not mean the project can never improve wording, add JSON fields, support new report formats, or change behavior through an explicit compatibility note.

It means changes should be intentional, reviewable, covered by tests or fixtures, and explainable from local Maven evidence rather than random ordering or subjective interpretation.
