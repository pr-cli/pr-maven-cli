# Deterministic Context

PR Maven CLI uses deterministic context to mean Maven failure context produced from local artifacts, parsers, and explicit rules.

The goal is to explain what failed, where it failed, and how to reproduce it without relying on random ordering, subjective interpretation, network access, or provider-specific state.

## Product Principle

The core principle is:

```text
parsing plus rules
  -> one package manager
  -> one failure model
  -> one deterministic contract
```

For this project, that means:

- one package manager: Maven;
- one primary source boundary: local Maven report and log artifacts;
- one failure model: a report-backed `Finding` mapped to module, plugin, phase, source, confidence, and reproduction command;
- one deterministic contract: text and JSON emitted from the same analysis result.

This focus keeps the MVP small enough to validate thoroughly and stable enough for contributors to extend through fixtures, parsers, tests, and documentation.

## What It Depends On

Deterministic context depends on:

- local `pom.xml` discovery;
- supported Maven report and log artifacts already present on disk;
- parser behavior covered by fixtures;
- explicit module mapping rules;
- stable path normalization;
- stable text and JSON output expectations;
- tests that lock down behavior before new integrations are added.

When the same supported artifacts are analyzed in the same project layout, the same finding should be mapped to the same Maven context.

## What It Does Not Depend On

Deterministic context does not depend on:

- LLM output;
- loose heuristics;
- random ordering;
- network access;
- GitHub or GitLab API calls;
- provider tokens;
- telemetry;
- hosted services;
- subjective interpretation of unstructured logs.

Future integrations may add provider context, annotations, or PR metadata, but those signals should be optional and additive. Local Maven evidence remains the base layer.

## Relationship To Output

Deterministic context is the product principle. Deterministic output is the text and JSON expression of that principle.

For output terminology, read [Deterministic Output](deterministic-output.md).
