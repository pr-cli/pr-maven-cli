# Confidence

PR Maven CLI uses confidence to describe how strong the local evidence is for a finding.

Confidence is about evidence quality. It is not a probability score, severity score, or guarantee that a developer made a mistake. A high-confidence finding means the CLI found deterministic Maven evidence on disk and can explain why it trusts that evidence.

## Current Stage 1 Behavior

Stage 1 emits `high` confidence for every finding because all current findings come from supported Maven report or log artifacts:

- Surefire JUnit XML reports;
- Failsafe JUnit XML reports;
- Checkstyle XML reports;
- SpotBugs XML reports;
- Maven Enforcer log artifacts;
- JaCoCo threshold log artifacts.

The analyzer does not infer failures from remote provider state, pull request metadata, free-form CI logs, or model output. It reads local artifacts, maps them back to Maven modules, and emits a finding only when the source artifact contains a deterministic failure, violation, bug, rule failure, or threshold failure.

## Levels

### `high`

Use `high` when the finding is backed by a supported local artifact and the CLI can explain the evidence path.

Current examples:

- a test failure found in a Surefire or Failsafe JUnit XML report;
- a Checkstyle violation found in a Checkstyle XML report;
- a SpotBugs bug found in a SpotBugs XML report;
- a Maven Enforcer rule failure found in a Maven log artifact;
- a JaCoCo threshold failure found in a Maven log artifact.

High-confidence findings should include confidence reasons such as:

- the source report or log format that contained the finding;
- the Maven module mapped from the report path;
- the test class, source file, rule, or threshold metric used to build the reproduction context.

### `medium`

`medium` is reserved for future releases.

It should be used when a finding is supported by useful evidence but one part of the chain is incomplete or inferred. Examples could include partial CI artifact layouts, provider metadata that points to a failing module without the full report payload, or log excerpts that identify the plugin but not every source detail.

Stage 1 does not emit `medium`.

### `low`

`low` is reserved for future releases.

It should be used only for weak or heuristic hints that may help triage but are not strong enough to claim a deterministic Maven failure. Examples could include changed-file relevance, broad log keywords, or provider-level failure summaries without a parseable Maven artifact.

Stage 1 does not emit `low`.

## JSON Contract

The current JSON schema allows `high` because the current implementation only emits report-backed deterministic findings. Future releases may add `medium` and `low` through an explicit compatibility note.

Consumers should:

- read `confidence` as a trust level for the evidence;
- read `confidenceReasons` before automating decisions;
- ignore unknown future fields;
- treat unknown future confidence values conservatively until their workflow explicitly supports them.

## Future Model

The Stage 3 confidence model may combine local Maven evidence with optional provider context, changed-file relevance, annotations, or PR metadata. Local report-backed evidence should remain the strongest signal, and every confidence level should keep human-readable reasons in JSON output.
