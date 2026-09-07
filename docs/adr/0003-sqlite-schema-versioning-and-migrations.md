# ADR 0003: SQLite Schema Versioning and Migrations

## Status

Accepted

## Context

Vitalynq v1.0 stores personal health data in a local SQLite database.

The database schema will evolve over time. Vitalynq must preserve supported user data across application upgrades and make schema evolution explicit, testable, and maintainable.

The domain model must remain independent from database schema versioning and migration details.

## Decision

SQLite remains the v1.0 persistence engine.

Vitalynq will store the database schema version explicitly inside the database.

Schema changes will be implemented as ordered migrations.

Migrations are applied exactly once, in version order. They must fail cleanly without leaving the database partially migrated, and they must be transactional when SQLite allows it.

Supported user data must be preserved across upgrades.

Migration paths between supported schema versions must be covered by tests.

Domain logic must not depend on schema versioning, migration order, SQL statements, or migration storage details.

No destructive migration is allowed without explicit justification, tests, and a documented compatibility impact.

## Consequences

Benefits:

- database evolution becomes explicit and reviewable;
- supported upgrades can be tested;
- persistence changes stay outside domain logic;
- user data preservation becomes a release requirement;
- future schema changes have a clear process.

Tradeoffs:

- schema changes require migration code and tests;
- persistence setup becomes more structured;
- failed migrations need careful error handling;
- destructive changes require stronger justification and documentation.

## Current Baseline

The existing SQLite schema created by `initializeSQLiteSchema` is considered the initial supported baseline for the current v1.0 development line.

This baseline is not yet represented as an applied migration in existing databases. Defining this baseline does not modify any existing database and does not insert any row into `schema_migrations`.

The project will define a separate baseline procedure before automatic migrations are enabled at startup.