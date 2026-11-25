---
type: topic
summary: How a Steading component is versioned, what its version promises, and how you upgrade it.
---
# Releases and versioning

Every Steading component follows [SemVer](https://semver.org/). A release's tag is `v<semver>`, and its version is the tag without the `v`.[^releases]

## What the version covers

A component's version covers its public API:[^releases]

- **The Connect-RPC API,** including JSON field and enum names and how they behave, not just the protobuf schema.
- **The CLI:** its commands, flags, exit codes, environment variables, and `-o json` output. Text output is for people and may change.
- **Operator configuration:** config keys, server environment variables, systemd units, and compose files.
- **Policy input:** the document that Rego policies read.

These are not part of the public API, and can change in any release:[^releases]

- **MCP tools,** including their names, schemas, and descriptions. Agents read the schemas each session, their permissions come from an assistant's permission classifier rather than per-tool allow rules, and skills don't name tools.
- **The database schema.** A component's [migrations](tech-stack.md#migrations) carry you from one schema to the next.
- **Go packages.**

## Before 1.0

Before 1.0, a minor release may change the public API without a deprecation period. Its release notes say what changed and what to do instead. There are no deprecation windows or timed-removal markers.[^releases]

After 1.0, deprecations follow mise's model: a `deprecated_at!` marker warns, then fails the tests once the removal version arrives, and a `remove_by!` marker fails the build.[^releases]

## Upgrading

You can upgrade directly from any earlier release to any later one, unless a migration forces a required stop on the way, which the release notes announce.[^releases]

Back up before you upgrade. To downgrade, restore that backup: down migrations are a development tool, not a production rollback.[^releases]

## Guarding the API

`buf breaking` runs as an [hk](tech-stack.md#hk) step against the latest release tag, so a change that breaks the Connect-RPC API shows up before it's released.[^releases]

[^releases]: [Releases, migrations, and background work](../sources/202611251316-releases-and-migrations.md)
