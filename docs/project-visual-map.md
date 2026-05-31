# Project Visual Map

This document gives maintainers and contributors a quick visual view of what is done, what comes next, and how the backlog is organized.

Use it together with [ROADMAP.md](../ROADMAP.md), the [v0.1.0 release snapshot](release-snapshot-v0.1.0.md), and the GitHub issue labels.

## One-Screen Status

```mermaid
flowchart LR
    S1["Stage 1\nv0.1.0 shipped\nLocal Maven report MVP"] --> S2["Stage 2 remaining\nDocs and ergonomics\n#17 #18 #19 #20"]
    S2 --> S21["Stage 2.1\nMVP hardening\nValidation and readiness"]
    S21 --> GATE["Stage 3 readiness gate\n#79 #80 #102"]
    GATE --> S3PLAN["Stage 3 planning\nRelease acceptance and dependency map\n#100, #101 shipped"]
    S3PLAN --> S3DESIGN["Stage 3 design contracts\nProvider boundaries, mocks, fixtures, JSON, UX"]
    S3DESIGN --> S3BUILD["Stage 3 implementation\nOptional PR and CI context layer"]
    S3BUILD --> V03["v0.3.0 target\nPR and CI context release"]
```

Current interpretation:

- Stage 1 is shipped as `v0.1.0`.
- Stage 2 has a small set of ready documentation and CLI ergonomics issues left.
- Stage 2.1 is the hardening lane before broad Stage 3 work.
- Stage 3 is well mapped, but implementation-heavy issues should wait for readiness gates and design contracts.

## What Is Already Done

```mermaid
flowchart TD
    MVP["v0.1.0 release"] --> CLI["CLI\nfails, why, version"]
    MVP --> LIB["Go library\npkg/prmaven"]
    MVP --> PARSERS["Maven evidence\nSurefire, Failsafe, Checkstyle, SpotBugs, Enforcer, JaCoCo"]
    MVP --> OUTPUT["Output\ntext, JSON, file output, module filter"]
    MVP --> DEMO["Demos and fixtures\nfailure and no-failure workspaces"]
    MVP --> OSS["OSS infrastructure\nREADME, contributing, governance, security, CI, release"]
    MVP --> ARTIFACTS["Release artifacts\nLinux, macOS, Windows, SHA-256"]
```

Important shipped properties:

- Local-first core analyzer.
- No required GitHub, GitLab, CI, AI, or telemetry dependency.
- Provider-agnostic runtime.
- Maven 3.9.x production baseline.
- Maven 4 stays future investigation until the Maven 4 line is production-ready and validated.

## Backlog Lanes

```mermaid
flowchart TD
    L1["Stage 2 ready lane\nFinish remaining docs and ergonomics\n#17 #18 #19 #20"]
    L2["Stage 2.1 docs lane\nDeterministic terminology and context\n#61 #62 #65"]
    L3["Stage 2.1 validation fixtures\nGolden JSON and fixture integrity\n#69 #70 #72"]
    L4["Stage 2.1 CI and E2E\nSchema validation, smoke tests, output/module matrices\n#67 #68 #71 #73 #74 #76"]
    L5["Stage 2.1 gates\nPR checklist, Stage 3 readiness, no-network guard, label review\n#77 #79 #80 #102"]
    L6["Stage 3 planning\nv0.3.0 acceptance and visual map\n#100, #101 shipped"]
    L7["Stage 3 design contracts\nProvider boundaries, fixtures, permissions, JSON, UX\n#82-#94"]
    L8["Stage 3 implementation\nGitHub adapters, scoring, commands, summaries, evidence bundles\n#21-#29 #34 #36"]
    L9["Stage 3 research and distribution\nSARIF, annotations, Maven 4, GitLab, package managers, privacy\n#30-#33 #35 #37-#40"]

    L1 --> L2 --> L3 --> L4 --> L5 --> L6 --> L7 --> L8
    L7 --> L9
```

Recommended maintainer order:

1. Finish Stage 2 ready issues.
2. Finish Stage 2.1 documentation and deterministic-context issues.
3. Add fixture and JSON validation.
4. Move stable validation into CI.
5. Complete the Stage 3 readiness gate.
6. Unblock the first Stage 3 design issues.
7. Start implementation only after contracts and mocks exist.

## Stage 3 Dependency Map

```mermaid
flowchart TD
    R0["Stage 2.1 readiness\n#79 #80 #102"]
    R1["v0.3.0 release acceptance\n#100"]
    R2["Visual dependency map shipped\n#101"]
    P0["Provider context planning\n#82"]
    P1["Error and offline fallback taxonomy\n#83"]
    P2["Provider adapter package boundaries\n#84"]
    P3["Provider fixture and mock contract\n#85"]
    P4["GitHub token and permission matrix\n#86"]
    F1["Changed-files fixture contract\n#87"]
    F2["Check-runs fixture contract\n#89"]
    J1["PR context JSON extension contract\n#90"]
    UX["why, explain, and ci UX boundaries\n#91"]
    CI["CI artifact directory fixture layout\n#93"]
    AG["Agent evidence bundle schema\n#94"]
    GL["GitLab parity boundary\n#92"]

    GH1["GitHub changed files adapter\n#21"]
    GH2["GitHub check runs adapter\n#23"]
    BASE["Baseline comparison model\n#24"]
    SCORE["PR-to-module relevance scoring\n#25"]
    CONF["Confidence model v2\n#26"]
    EXPLAIN["prmaven explain\n#27"]
    CICMD["prmaven ci\n#28"]
    MD["Markdown PR summary\n#29 #88"]
    BUNDLE["Agent evidence bundle output\n#34"]
    ART["CI artifact directory option\n#36"]

    R0 --> R1
    R0 --> R2
    R1 --> P0
    R2 --> P0
    P0 --> P1
    P0 --> P2
    P0 --> P3
    P0 --> P4
    P1 --> J1
    P2 --> GH1
    P2 --> GH2
    P3 --> F1
    P3 --> F2
    P3 --> CI
    P4 --> GH1
    P4 --> GH2
    F1 --> GH1
    F2 --> GH2
    J1 --> MD
    UX --> EXPLAIN
    UX --> CICMD
    CI --> ART
    AG --> BUNDLE
    GL --> P2
    GH1 --> SCORE
    GH2 --> SCORE
    SCORE --> BASE
    BASE --> CONF
    CONF --> EXPLAIN
    CONF --> MD
```

Stage 3 rule of thumb:

- Design issues should land before implementation issues.
- Fixture and mock contracts should land before provider adapters.
- Permission and offline fallback docs should land before token-aware behavior.
- JSON and Markdown contracts should land before new public output becomes expected by users.

## Issue Status Guide

```mermaid
flowchart LR
    READY["status: ready\nContributor can start"] --> PR["Focused PR\nTests/docs included"]
    BLOCKED["status: blocked\nNeeds dependency or design decision"] --> GATE["Wait for gate issue or dependency"]
    HELP["help wanted / need help\nMaintainer wants contributions"] --> READY
    FIRST["good first contribution / oss first friendly\nLow context contribution"] --> READY
```

When in doubt:

- Prefer `status: ready` issues for immediate work.
- Prefer `good first contribution` or `oss first friendly` issues for new contributors.
- Avoid implementation-heavy Stage 3 issues until the related design contract issues are done.
- Keep every issue tied to a stage label, area label, and status label.

## Fast Path For Maintainers

The shortest practical path from now to healthy Stage 3 execution is:

```mermaid
flowchart LR
    A["Close #17-#20"] --> B["Close #61 #62 #65"]
    B --> C["Close #69 #70 #72"]
    C --> D["Close #71 #73"]
    D --> E["Close #67 #68 #74 #76"]
    E --> F["Close #77 #79 #80 #102"]
    F --> G["Close #100 and keep #101 current"]
    G --> H["Unblock early Stage 3 design issues"]
```

After that, Stage 3 implementation can begin with much lower ambiguity.

## Source Of Truth

This document is a navigation aid. The source of truth remains:

- [Roadmap](../ROADMAP.md)
- [v0.1.0 release snapshot](release-snapshot-v0.1.0.md)
- [MVP acceptance checklist](mvp-acceptance.md)
- [Implementation status](implementation-status.md)
- [Label guide](labels.md)
- [OSS guardrails](oss-guardrails.md)
- GitHub issue labels and issue bodies

