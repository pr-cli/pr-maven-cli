# Fixture Notes

PR Maven CLI uses committed, sanitized Maven report artifacts as deterministic test fixtures.

Fixtures are part of the project contract. They should be easy to inspect, safe to publish, and stable enough for contributors, maintainers, CI systems, and coding agents to use as evidence.

## Maven Compatibility Baseline

The production fixture baseline is Maven 3.9.x.

The project is currently documented against Maven 3.9.16 because Apache Maven lists Maven 3.9.16 as the recommended release for all users on the [Apache Maven download page](https://maven.apache.org/download.cgi). New report fixtures should be compatible with Maven 3.9.x behavior unless an issue explicitly documents a different compatibility target.

Maven 4 is tracked separately. The Maven 4 line may change behavior around build output, plugin defaults, model handling, and report generation. PR Maven CLI should not declare Maven 4 production support, or add Maven 4-specific fixtures as required compatibility evidence, until Maven 4 behavior is stable enough for a dedicated compatibility investigation.

## Fixture Scope

Fixtures may live under:

- `demo/` for user-facing CLI demonstrations;
- `pkg/prmaven/testdata/` for focused parser and analyzer tests;
- `pkg/prmaven/testdata/golden/` for human-readable output snapshots.

Committed fixtures may include:

- `pom.xml` files;
- `target/surefire-reports/*.xml`;
- `target/failsafe-reports/*.xml`;
- `target/checkstyle-result.xml`;
- `target/spotbugsXml.xml`;
- selected sanitized `target/*.log` files for Maven Enforcer, JaCoCo, or similar plugin output.

Generated artifacts outside intentional fixture paths should not be committed.

## Compatibility Expectations

When adding or updating fixtures:

- keep the fixture local-first and self-contained;
- avoid private source names, private dependency coordinates, secrets, tokens, hostnames, or proprietary logs;
- prefer Maven 3.9.x-compatible report shapes;
- document any dependency on a specific Maven plugin behavior;
- keep Maven 4 behavior out of the required baseline until a dedicated Maven 4 compatibility issue accepts it;
- update golden files in the same pull request when text output intentionally changes;
- keep JSON output additive and compatible unless the issue explicitly allows a breaking contract change.

## Why Maven 4 Is Separate

Maven 4 support is important, but it belongs in a future compatibility lane rather than the Stage 2 fixture baseline.

Keeping Maven 4 separate lets the project:

- preserve a stable production baseline for current users;
- avoid mixing preview-line behavior into normal parser expectations;
- add Maven 4 fixtures intentionally after Maven 4 report behavior is validated;
- document any differences between Maven 3.9.x and Maven 4 instead of treating them as accidental parser regressions.

Until then, Maven 3.9.x fixtures remain the required compatibility evidence for production-ready behavior.
