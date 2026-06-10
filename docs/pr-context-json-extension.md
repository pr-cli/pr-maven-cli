# PR Context JSON Extension Contract

This document defines how optional Stage 3 PR and CI context may extend PR Maven CLI JSON output without breaking existing consumers.

The current report contract is documented in [JSON Contract](json-contract.md). That contract remains the stable base: `Report`, `Summary`, `Module`, and `Finding` fields must stay compatible unless a future release explicitly documents a breaking change.

Use this document with the [Provider Context Plan](provider-context-plan.md), [Changed-Files Fixture Contract](changed-files-fixture-contract.md), [Check-Runs Fixture Contract](check-runs-fixture-contract.md), [Provider Errors And Offline Fallbacks](provider-errors-and-fallbacks.md), and [Confidence](confidence.md).

## Compatibility Rules

PR context must be additive.

Required rules:

- Keep existing `Report`, `Summary`, `Module`, and `Finding` field names and meanings.
- Do not move existing local Maven evidence under provider-specific objects.
- Do not require provider context for report-only JSON output.
- Add new provider-aware fields as optional objects or arrays.
- Omit optional provider-aware fields when provider context was not requested.
- Preserve slash-separated paths across platforms.
- Keep local Maven findings valid even when provider context is missing, partial, or failed.
- Represent provider failures as diagnostics or metadata, not Maven findings.

Consumers should ignore unknown fields. This lets future versions add provider metadata, confidence details, relevance scoring, or output formats without breaking older automation.

## Proposed Top-Level Extension

The preferred Stage 3 shape is a top-level optional `providerContext` object.

Report-only output may omit `providerContext` entirely. Provider-aware output may include it.

Recommended shape:

```json
{
  "projectRoot": "<PROJECT_ROOT>",
  "summary": {},
  "modules": [],
  "findings": [],
  "providerContext": {
    "schemaVersion": 1,
    "state": "available",
    "provider": "github",
    "repository": {
      "owner": "example-org",
      "name": "example-maven-project",
      "visibility": "public"
    },
    "pullRequest": {
      "number": 123,
      "baseRef": "main",
      "headRef": "feature/module-change",
      "headSha": "0000000000000000000000000000000000000000"
    },
    "changedFiles": {
      "state": "available",
      "files": []
    },
    "checkRuns": {
      "state": "available",
      "runs": []
    },
    "diagnostics": []
  }
}
```

## Provider Context Fields

| Field | Type | Description |
| --- | --- | --- |
| `schemaVersion` | integer | Version of the provider-context extension contract. |
| `state` | string | Overall provider context state: `available`, `partial`, `unavailable`, or `not_requested`. |
| `provider` | string | Provider name, such as `github` or `gitlab`. |
| `repository` | object | Sanitized repository identity and visibility details. |
| `pullRequest` | object | Pull request or merge request identity and refs. |
| `changedFiles` | object | Optional changed-file context following the changed-files fixture contract. |
| `checkRuns` | object | Optional check-run context following the check-runs fixture contract. |
| `diagnostics` | array | Provider diagnostics from missing token, permissions, rate limits, not found, network failure, or unsupported state. |

## Finding-Level Extension

Provider context may enrich individual Maven findings with an optional `prContext` object.

Recommended shape:

```json
{
  "id": "payment-core-paymentroundingtest-shouldrejectinvalidscale-failure",
  "module": "payment-core",
  "modulePath": "payment-core",
  "reportPath": "payment-core/target/surefire-reports/TEST-dev.prmaven.demo.PaymentRoundingTest.xml",
  "reportKind": "surefire",
  "mavenPlugin": "maven-surefire-plugin",
  "mavenPhase": "test",
  "testClass": "dev.prmaven.demo.PaymentRoundingTest",
  "testName": "shouldRejectInvalidScale",
  "failureKind": "failure",
  "message": "expected:<2.50> but was:<2.49>",
  "reproduceCommand": "mvn -pl payment-core -am -Dtest=PaymentRoundingTest test",
  "confidence": "high",
  "confidenceReasons": [
    "failure was found in a Maven Surefire JUnit XML report",
    "report path maps to Maven module payment-core"
  ],
  "sourceReportFormat": "junit-xml",
  "prContext": {
    "changedFileRelevance": "module_changed",
    "changedFiles": [
      "payment-core/src/main/java/example/PaymentService.java"
    ],
    "relatedCheckRuns": [
      "Maven test"
    ],
    "providerConfidenceReasons": [
      "pull request changed files include module payment-core",
      "check run Maven test completed with failure"
    ]
  }
}
```

Finding-level `prContext` must not replace existing report-backed fields. It can explain relevance, not evidence ownership.

## Report-Only Example

Report-only output remains valid without provider fields.

```json
{
  "projectRoot": "<PROJECT_ROOT>/demo/multi-module-failure",
  "summary": {
    "moduleCount": 3,
    "reportCount": 2,
    "findingCount": 2
  },
  "modules": [
    {
      "name": "payment-core",
      "path": "payment-core",
      "pom": "payment-core/pom.xml"
    }
  ],
  "findings": [
    {
      "id": "payment-core-paymentroundingtest-shouldrejectinvalidscale-failure",
      "module": "payment-core",
      "modulePath": "payment-core",
      "reportPath": "payment-core/target/surefire-reports/TEST-dev.prmaven.demo.PaymentRoundingTest.xml",
      "reportKind": "surefire",
      "mavenPlugin": "maven-surefire-plugin",
      "mavenPhase": "test",
      "testClass": "dev.prmaven.demo.PaymentRoundingTest",
      "testName": "shouldRejectInvalidScale",
      "failureKind": "failure",
      "message": "expected:<2.50> but was:<2.49>",
      "reproduceCommand": "mvn -pl payment-core -am -Dtest=PaymentRoundingTest test",
      "confidence": "high",
      "confidenceReasons": [
        "failure was found in a Maven Surefire JUnit XML report",
        "report path maps to Maven module payment-core"
      ],
      "sourceReportFormat": "junit-xml"
    }
  ]
}
```

## PR-Context-Enriched Example

Provider-aware output may add `providerContext` and finding-level `prContext`.

```json
{
  "projectRoot": "<PROJECT_ROOT>/demo/multi-module-failure",
  "summary": {
    "moduleCount": 3,
    "reportCount": 2,
    "findingCount": 1
  },
  "modules": [
    {
      "name": "payment-core",
      "path": "payment-core",
      "pom": "payment-core/pom.xml"
    }
  ],
  "findings": [
    {
      "id": "payment-core-paymentroundingtest-shouldrejectinvalidscale-failure",
      "module": "payment-core",
      "modulePath": "payment-core",
      "reportPath": "payment-core/target/surefire-reports/TEST-dev.prmaven.demo.PaymentRoundingTest.xml",
      "reportKind": "surefire",
      "mavenPlugin": "maven-surefire-plugin",
      "mavenPhase": "test",
      "testClass": "dev.prmaven.demo.PaymentRoundingTest",
      "testName": "shouldRejectInvalidScale",
      "failureKind": "failure",
      "message": "expected:<2.50> but was:<2.49>",
      "reproduceCommand": "mvn -pl payment-core -am -Dtest=PaymentRoundingTest test",
      "confidence": "high",
      "confidenceReasons": [
        "failure was found in a Maven Surefire JUnit XML report",
        "report path maps to Maven module payment-core"
      ],
      "sourceReportFormat": "junit-xml",
      "prContext": {
        "changedFileRelevance": "module_changed",
        "changedFiles": [
          "payment-core/src/main/java/example/PaymentService.java"
        ],
        "relatedCheckRuns": [
          "Maven test"
        ],
        "providerConfidenceReasons": [
          "pull request changed files include module payment-core",
          "check run Maven test completed with failure"
        ]
      }
    }
  ],
  "providerContext": {
    "schemaVersion": 1,
    "state": "available",
    "provider": "github",
    "repository": {
      "owner": "example-org",
      "name": "example-maven-project",
      "visibility": "public"
    },
    "pullRequest": {
      "number": 123,
      "baseRef": "main",
      "headRef": "feature/module-change",
      "headSha": "0000000000000000000000000000000000000000"
    },
    "changedFiles": {
      "state": "available",
      "files": [
        {
          "path": "payment-core/src/main/java/example/PaymentService.java",
          "status": "modified"
        }
      ]
    },
    "checkRuns": {
      "state": "available",
      "runs": [
        {
          "name": "Maven test",
          "status": "completed",
          "conclusion": "failure"
        }
      ]
    },
    "diagnostics": []
  }
}
```

## Partial Context Example

Provider context can be partial without invalidating the Maven report.

```json
{
  "projectRoot": "<PROJECT_ROOT>",
  "summary": {
    "moduleCount": 1,
    "reportCount": 1,
    "findingCount": 1
  },
  "modules": [],
  "findings": [],
  "providerContext": {
    "schemaVersion": 1,
    "state": "partial",
    "provider": "github",
    "repository": {
      "owner": "example-org",
      "name": "example-maven-project",
      "visibility": "public"
    },
    "pullRequest": {
      "number": 123,
      "baseRef": "main",
      "headRef": "feature/module-change",
      "headSha": "0000000000000000000000000000000000000000"
    },
    "changedFiles": {
      "state": "available",
      "files": []
    },
    "checkRuns": {
      "state": "unavailable",
      "runs": []
    },
    "diagnostics": [
      {
        "code": "insufficient_permissions",
        "message": "GitHub token cannot read check runs",
        "source": "checkRuns"
      }
    ]
  }
}
```

## Consumer Guidance

Consumers should:

- parse the known `Report`, `Summary`, `Module`, and `Finding` fields first;
- ignore unknown top-level fields;
- ignore unknown finding-level fields;
- treat `providerContext.state` as advisory context;
- avoid failing report parsing because provider context is unavailable or partial;
- prefer local Maven finding fields when provider context disagrees with local report evidence.

Provider-aware consumers may opt into the new fields, but report-only consumers should continue to work unchanged.

## Schema Guidance

When this extension is implemented, schema changes should:

- add `providerContext` as an optional top-level property;
- allow optional `prContext` on findings;
- avoid making provider fields required for report-only output;
- avoid rejecting future unknown fields unless the schema mode is intentionally strict;
- include fixtures for report-only, provider-context-enriched, and partial-context output.
