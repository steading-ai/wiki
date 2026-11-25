---
participants: [ivy]
---

# Releases, migrations, and background work

A working session in skift where ivy settled how a component is versioned and released, how its migrations are kept safe, and where its background work and database code live. The decisions were made in skift but hold for Steading components generally. skift's ADRs will record them: skift#51 (ADR-014), skift#52 (ADR-015), and skift#53 (ADR-016). What follows is what ivy decided, what was passed over and why, and what stayed open.

## Versions

Versions follow SemVer. The tag is `v<semver>`, and the version is the tag without its `v`.

## What the version covers

The version covers the public API:

- the Connect and protobuf API, including JSON field and enum names and their behavior, not just the schema;
- the CLI's commands, flags, exit codes, environment variables, and `-o json` output. Text output is for people and may change;
- operator configuration: config keys, server environment variables, systemd units, and compose files;
- the policy input document that Rego reads.

These are not public:

- MCP tool names, schemas, and descriptions. Agents read the schemas each session, ivy expects agents' permissions to come from the auto-mode classifier rather than per-tool allow rules, and skills shouldn't name tools.
- The database schema.
- Go packages.

## Before 1.0

Before 1.0, a minor release may change the public API without a deprecation period, and its release notes say what changed and what to do instead. There are no deprecation windows or timed-removal markers until 1.0; ivy called the alternative "maturity theater". After 1.0, mise's model is the reference for timed obligations: `deprecated_at!` warns, then fails tests at a removal version, and `remove_by!` fails the build.

## Upgrades

Any released version upgrades directly from any earlier release, until a migration forces a required stop, which the release notes announce. Back up before upgrading. A downgrade means restoring that backup; down migrations are a development tool, not a production rollback.

`buf breaking` runs as an hk step against the latest release tag.

## Migration safety

Migration safety comes from tools, not a style guide:

- **Lint:** squawk lints each new migration for statements that take unsafe locks, as an hk step. It won a hands-on comparison. eugene was passed over because it missed a case and has no gcc-style output, and Atlas's `migrate lint` because it is Pro-only.
- **Code:** goose's session lock, `lock_timeout` on the migration connection with bounded retries, and refusing to run against a schema newer than the binary.
- **Tests:** an upgrade from every released schema with data in every table, and a check that shipped migrations never change.
- **Review** covers only what tools can't judge: what a default means for existing rows, the size of a backfill, and data that only running code wrote.

## Background work

River (riverqueue/river, PostgreSQL-backed, on pgx v5) runs a component's own background work as periodic jobs inside the server, with leader election. It takes the place of systemd timers, which compose and launchd installs never had. River's schema is embedded as goose migrations, so there is one migration system.

## Package boundaries

- `store` owns PostgreSQL, including River's client.
- `domain` owns the core rules and use cases only.
- An `admin` package owns migrations, bootstrap, opening the database, and background jobs.

`admin` owning the database means it decides when and how to open it, by calling the store's constructor; it never imports pgx. The pgx pool and River's PostgreSQL drivers stay in `store`. Lint enforces the split with depguard:

- `domain` and `admin` may import `store`.
- pgx and River's drivers are allowed only in `store`.
- River's core package is allowed only in `store` and `admin`, for workers and the periodic schedule, whose APIs carry no pgx types.

Operator commands live under the CLI's `admin` subcommand.

## What stayed open

- Deriving the CLI and MCP from the proto, held in skift#49.
- Whether the software takes the pre-upgrade backup itself, held in skift#63.
