# ADR 0002: Local-First Offline Operation

## Status

Accepted

## Context

Vitalynq v1.0 is a local-first, single-user application for organizing personal health data.

Personal health data is sensitive. Core application behavior must not depend on accounts, cloud services, telemetry, external APIs, or silent network access.

Vitalynq may expose different interfaces in the future, but v1.0 must preserve an offline-capable user experience.

## Decision

Vitalynq core functionality must work locally and offline.

The v1.0 storage model is a local SQLite database controlled by the user.

Vitalynq v1.0 must not require:

- a user account;
- a cloud service;
- a telemetry service;
- an external API;
- silent external network access.

Exports are produced locally and remain local, regardless of their format or local destination.

Backups, sharing, syncing, copying, or moving exported data must remain explicit user-controlled actions.

Any future feature that requires network access, cloud storage, remote sync, accounts, or external services must be introduced through a separate ADR and must require explicit user action.

## Consequences

Benefits:

- the application remains usable offline;
- users can inspect where their data is stored;
- privacy expectations are easier to explain and test;
- local-first behavior does not depend on third-party service availability;
- future networked features cannot be added silently.

Tradeoffs:

- v1.0 does not provide built-in cloud sync;
- v1.0 does not provide account-based recovery;
- users remain responsible for protecting local databases, exports, and backups;
- future collaboration or multi-device workflows will require explicit architectural decisions.
