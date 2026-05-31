# Project Context and Requirements

PR Maven CLI exists to turn local Maven failure artifacts into deterministic pull request and CI context.

## Problem

Maven CI failures usually leave developers with scattered evidence:

- terminal logs;
- Surefire and Failsafe XML reports;
- quality plugin reports;
- module paths;
- plugin phases;
- CI artifacts that may not be easy to inspect locally.

Generic CI systems can show that a job failed, but they rarely answer the next engineering question in a Maven-aware way:

```text
What failed, in which module, through which Maven plugin, and how do I reproduce it locally?
```

PR Maven CLI focuses on that question.

## Product Goal

The target workflow is:

```text
local Maven workspace or CI artifact directory
  -> report discovery
  -> deterministic parsing
  -> Maven module mapping
  -> text or JSON failure context
  -> local reproduction command
```

The CLI should be useful to developers, maintainers, CI jobs, engineering platform teams, and coding agents that need reliable evidence before acting on a failing pull request.

## Execution Model

PR Maven CLI is:

- local-first;
- dependency-light;
- usable as both a CLI and a Go library;
- deterministic by default;
- provider-agnostic in the core analyzer;
- safe for private source trees and private CI logs.

Core analysis reads local files and does not require GitHub, GitLab, CI provider APIs, AI providers, telemetry, or external services.

## Current Workflow Boundary

External workflow:

```text
developer or CI runs Maven
  -> Maven writes reports and logs under target/
```

This project:

```text
local Maven reports/logs exist
  -> PR Maven CLI analyzes supported artifacts
  -> report is emitted as text, JSON, or a file
```

PR Maven CLI does not run Maven on behalf of the user in the core flow. It analyzes artifacts that already exist.

## Out of Scope

The current production baseline intentionally does not include:

- native GitHub API adapters;
- native GitLab API adapters;
- hosted dashboards;
- automatic code fixes;
- telemetry;
- uploading reports or logs to external services;
- claiming root cause without report-backed evidence;
- Maven 4 production support before the Maven 4 line is considered ready.

Stage 3 may add optional provider adapters, but local Maven report parsing must remain usable without network access or provider tokens.

## Product Principles

- Deterministic before agentic.
- Local-first before hosted.
- Maven-aware before generic.
- Report-backed evidence before broad claims.
- Human-readable and machine-readable output from day one.
- Stable contracts for CI systems and coding agents.
- Low setup cost for contributors.
- Network access stays optional.

## Domain Model

The public contract is centered on:

- `Report`: the top-level analysis result.
- `Summary`: aggregate counts for modules, artifacts, and findings.
- `Module`: a discovered Maven module.
- `Finding`: a test, quality, or build failure mapped to Maven context.

Important finding concepts include:

- module name and path;
- source report or log path;
- report kind;
- Maven plugin;
- Maven phase;
- source class, test, file, rule, or metric when available;
- failure kind and message;
- confidence and confidence reasons;
- minimal reproduction command.

## Source Requirements

Supported inputs are local Maven workspaces or fixture directories that contain:

- `pom.xml` files;
- `target/surefire-reports/*.xml`;
- `target/failsafe-reports/*.xml`;
- `target/checkstyle-result.xml`;
- `target/spotbugsXml.xml` or `target/spotbugs.xml`;
- selected Maven log artifacts such as Enforcer or JaCoCo threshold logs.

Input parsing should be deterministic, fixture-backed, and tolerant of unsupported artifacts by ignoring what it does not understand.

## Output Requirements

PR Maven CLI emits:

- text output for humans;
- JSON output for CI, bots, and coding agents;
- optional file output through `-output`;
- slash-separated relative paths in JSON, including on Windows;
- stable field names documented in the JSON contract.

Text output should stay concise and actionable. JSON output should remain compatible with consumers that ignore unknown future fields.

For the terminology behind deterministic text and deterministic JSON, read [Deterministic Output](deterministic-output.md).

## Validation Requirements

Before a change is considered production-safe, the project should validate:

- parser behavior with sanitized fixtures;
- CLI behavior through end-to-end tests;
- JSON schema alignment;
- text output through golden snapshots;
- path normalization across operating systems;
- no-failure behavior;
- module filtering behavior;
- output file behavior;
- CI behavior across supported platforms.

## Automation Requirements

Repository automation should support OSS-style maintenance:

- pull request CI;
- release builds from version tags;
- security scanning;
- dependency maintenance;
- contributor-friendly issue templates;
- standard contributor thank-you automation.

Product runtime automation should remain explicit and local-first. Any future publishing, provider API, or PR-commenting behavior must be opt-in and documented.

## Objective Final State

The long-term target is a stable Maven-aware triage tool that can:

- explain local Maven failure artifacts deterministically;
- provide stable JSON for internal platforms and agents;
- produce readable summaries for pull request workflows;
- optionally enrich findings with PR or CI provider context;
- remain safe for private repositories and private CI logs;
- keep Maven report analysis useful without network access.
