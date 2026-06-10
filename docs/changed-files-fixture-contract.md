# Changed-Files Fixture Contract

This document defines the Stage 3 fixture contract for pull request changed-file data.

The contract exists before a live GitHub changed-files adapter is implemented so contributors can add parser, normalizer, fake-client, and mapping tests without needing provider tokens or network access.

Use this document with the [Provider Context Plan](provider-context-plan.md), [Provider Fixtures And Mocks](provider-fixtures-and-mocks.md), [Provider Adapter Package Boundaries](provider-adapter-boundaries.md), and [GitHub Provider Permissions](github-provider-permissions.md).

## Goals

- Define a sanitized JSON shape for changed-file provider fixtures.
- Cover added, modified, deleted, renamed, and nested-module paths.
- Preserve enough provider-like fields to test normalizers without storing volatile data.
- Make expected Maven module mapping explicit.
- Keep default tests provider-token-free and network-free.

## Fixture Location

When implementation begins, changed-file fixtures should live under:

```text
pkg/prmaven/provider/testdata/github/changed-files
```

Recommended first fixtures:

```text
added-modified-renamed.json
deleted-files.json
nested-modules.json
root-only-files.json
mixed-module-and-root-files.json
```

Each fixture should be small, reviewable, and safe to publish.

## Sanitized Fixture Shape

Changed-file fixtures should use JSON.

Recommended top-level shape:

```json
{
  "provider": "github",
  "resource": "pull_request_changed_files",
  "schemaVersion": 1,
  "scenario": "added modified renamed and deleted files across modules",
  "repository": {
    "owner": "example-org",
    "name": "example-maven-project",
    "visibility": "public"
  },
  "pullRequest": {
    "number": 123,
    "baseRef": "main",
    "headRef": "feature/module-change"
  },
  "files": [
    {
      "path": "payment-core/src/main/java/example/PaymentService.java",
      "previousPath": null,
      "status": "modified",
      "additions": 8,
      "deletions": 3
    },
    {
      "path": "payment-api/src/test/java/example/PaymentApiTest.java",
      "previousPath": null,
      "status": "added",
      "additions": 41,
      "deletions": 0
    },
    {
      "path": "billing-core/src/main/java/example/BillingService.java",
      "previousPath": "billing-core/src/main/java/example/InvoiceService.java",
      "status": "renamed",
      "additions": 5,
      "deletions": 5
    },
    {
      "path": "legacy-module/src/test/java/example/LegacyTest.java",
      "previousPath": null,
      "status": "deleted",
      "additions": 0,
      "deletions": 32
    }
  ],
  "expected": {
    "changedFiles": 4,
    "moduleImpacts": [
      {
        "modulePath": "payment-core",
        "paths": ["payment-core/src/main/java/example/PaymentService.java"],
        "statuses": ["modified"]
      },
      {
        "modulePath": "payment-api",
        "paths": ["payment-api/src/test/java/example/PaymentApiTest.java"],
        "statuses": ["added"]
      },
      {
        "modulePath": "billing-core",
        "paths": ["billing-core/src/main/java/example/BillingService.java"],
        "previousPaths": ["billing-core/src/main/java/example/InvoiceService.java"],
        "statuses": ["renamed"]
      },
      {
        "modulePath": "legacy-module",
        "paths": ["legacy-module/src/test/java/example/LegacyTest.java"],
        "statuses": ["deleted"]
      }
    ],
    "unmappedPaths": []
  }
}
```

## Required Fields

Each fixture must include:

- `provider`: provider name, initially `github`;
- `resource`: stable fixture resource name, initially `pull_request_changed_files`;
- `schemaVersion`: integer contract version;
- `scenario`: short human-readable scenario;
- `files`: changed file entries;
- `expected`: normalized expectations used by tests.

Each file entry must include:

- `path`: normalized slash-separated path relative to repository root;
- `status`: one of `added`, `modified`, `deleted`, `renamed`, `copied`, `unchanged`, or `unknown`;
- `previousPath`: previous slash-separated path for renames, otherwise `null`;
- `additions`: non-negative integer when available;
- `deletions`: non-negative integer when available.

Optional provider-like fields are allowed only when they are useful for parser behavior:

- `changes`;
- `patchPresent`;
- `binary`;
- `tooLarge`;
- `generated`.

Avoid storing raw patch hunks unless a future issue explicitly needs them.

## Sanitization Rules

Fixtures must not include:

- real private repository names;
- real private branch names;
- real author names, emails, tokens, or account IDs;
- signed URLs;
- artifact URLs;
- volatile API URLs;
- raw patches containing confidential code;
- timestamps unless a test explicitly needs ordering behavior.

Use stable placeholder values such as `example-org`, `example-maven-project`, `main`, and `feature/module-change`.

## Required Scenarios

The first changed-files fixture set should cover:

| Scenario | Required file statuses | Required mapping behavior |
| --- | --- | --- |
| Added and modified files | `added`, `modified` | Paths map to their nearest Maven module. |
| Deleted files | `deleted` | Deleted paths still map by their deleted path when the module path is known. |
| Renamed files | `renamed` | Current path and `previousPath` are preserved; module impact uses the current path and may record previous module path if it differs. |
| Nested modules | `added` or `modified` | Deepest matching Maven module wins. |
| Root-only files | `modified` | Root files map to `"."` or remain unmapped depending on future relevance rules. |
| Mixed root and module files | multiple | Module impacts and unmapped paths stay deterministic. |

## Maven Module Mapping Expectations

Changed-file module mapping should be deterministic and provider-neutral.

Recommended rules:

1. Normalize all changed-file paths to slash-separated repository-relative paths.
2. Match paths against discovered Maven module paths from local `pom.xml` discovery.
3. Choose the deepest matching module path when multiple module prefixes match.
4. Treat files under the root Maven project as module path `"."` only when the root project is an actual Maven module.
5. Preserve unmapped paths in `expected.unmappedPaths` instead of inventing a module.
6. For `renamed` files, preserve both `path` and `previousPath`; if the rename crosses module boundaries, expected output should record both current and previous module impact.
7. For `deleted` files, map by the deleted `path` when possible because the path is still part of the provider payload.

Mapping must not require a provider token. Tests should combine a local Maven module fixture with sanitized changed-file fixtures.

## Fake Client Expectations

Fake clients for changed files should support:

- successful changed-file response;
- empty changed-file response;
- paginated response behavior without live network calls;
- missing token;
- insufficient permissions;
- rate limit;
- not found;
- unsupported file status.

Error behavior should follow [Provider Errors And Offline Fallbacks](provider-errors-and-fallbacks.md).

## Test Expectations

Default tests should:

- load committed changed-file fixtures from disk;
- validate required fixture fields;
- normalize file paths and statuses;
- assert expected module impacts;
- assert expected unmapped paths;
- run without GitHub credentials;
- run without network access;
- avoid depending on real GitHub API responses.

Optional live-provider tests may exist later, but they must be skipped by default and must not be required by `go test ./...`.

## Compatibility With Future JSON Output

Changed-file context should enrich future provider-aware output, not replace Maven evidence.

Future JSON fields should remain additive. A local Maven finding should continue to expose the report-backed module, report path, plugin, phase, source, confidence reasons, and reproduction command even when changed-file context is absent or partial.
