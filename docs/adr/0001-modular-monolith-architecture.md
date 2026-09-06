# ADR 0001: Modular Monolith Architecture

## Status

Accepted

## Context

Vitalynq v1.0 establishes a stable local-first foundation for organizing personal health data.

Vitalynq remains a single-user application exposed through a CLI for v1.0. The CLI is the first user interface, but it must not define the product architecture. Future interfaces should be able to reuse the same core without requiring a rewrite.

The system must keep diagnosis, treatment recommendations, predictions, and medical advice outside its scope.

The system must not require any network service, cloud service, external API, telemetry service, or user account for its core operation.

## Decision

Vitalynq v1.0 will use a modular monolith architecture.

The codebase will be organized around clear boundaries between:

- domain logic;
- application workflows;
- persistence;
- interfaces.

The domain layer contains the core health-data concepts and validation rules. It must stay independent from SQLite, CLI behavior, external standards, network services, cloud services, telemetry, and future user interfaces.

The application layer coordinates use cases around the domain model. It can depend on domain concepts and abstract persistence contracts.

The persistence layer is an external adapter. It implements storage details such as SQLite access, schema versioning, and migrations behind contracts required by the internal layers.

The interface layer adapts user-facing entry points to application workflows. For v1.0, the CLI is the only supported interface.

## Dependency Direction

Dependencies must point inward.

External layers may depend on internal layers, but internal layers must not depend on external layers.

Allowed direction:

```text
interfaces -> application -> domain
external adapters -> internal contracts -> domain
```

Persistence code is an external adapter. It may implement contracts required by internal layers, but this ADR does not decide yet where those contracts are defined in the Go package structure.

The reverse direction is not allowed. Domain logic must not import or depend on CLI code, SQLite code, database schemas, filesystem paths, network services, or future interface code.

## Local-First Operation

Vitalynq core functionality must work locally and offline.

No network service or cloud service is required for basic operation. Local SQLite storage is the persistence mechanism for v1.0.

Exports are produced locally. Vitalynq does not move or share exported data unless the user explicitly does so outside the application.

## Not Decided

This ADR does not define the exact Go package structure.

This ADR does not define repository interfaces yet.

Those decisions will be made in later ADRs or implementation changes when the codebase is refactored around these boundaries.

## Consequences

Benefits:

- domain rules can be tested without CLI or SQLite setup;
- future interfaces can reuse the same application and domain logic;
- SQLite migrations and schema changes stay outside the domain model;
- local-first and single-user constraints are explicit;
- long-term maintenance becomes easier because responsibilities are separated.

Tradeoffs:

- small features may require touching more than one layer;
- the codebase needs boundary discipline even while it remains a monolith;
- package and repository design must be handled carefully in later steps;
- early refactoring may add structure before the project has many features.
