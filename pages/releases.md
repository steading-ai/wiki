---
type: topic
summary: How a Steading component is versioned, how releases reach you, and how upgrades stay uneventful.
---
# Releases and versioning

Every Steading component follows [SemVer](https://semver.org/). A release's tag is `v<semver>`, and its version is the tag without the `v`.[^releases] Every release gets a plain version with no pre-release suffix: the version says what changed in the public API, and the [channel](#channels) says how proven the build is.[^packages]

Releases ship often, several times a week if not daily, and upgrade your machine unattended, automatically, and uneventfully. Given the security landscape with agents, an upgrade that waits on you is a risk.[^packages]

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

## Channels

A release ships to two channels, `edge` and `stable`, in [Steading's package repository](packages.md) and as image tags alongside `X.Y.Z`.[^packages]

- **`edge`** gets every release as soon as it's tagged.
- **`stable`** moves automatically to the newest `edge` build that has been in `edge` for at least 48 hours and hasn't been yanked. Taking the newest build that has soaked, rather than restarting the clock with each new build, keeps `stable` moving when releases ship daily. Promotion changes only the repository's metadata; the signed package is the same one, and nothing is rebuilt.

Two controls override that:[^packages]

- **Promote now** moves a build to `stable` early, for a security fix.
- **Yank** removes a build from both channels, so a bad build stops spreading. dnf never downgrades a machine that already installed it, so the fix for those machines is a new, higher version.

## Upgrading

You can upgrade directly from any earlier release to any later one. There are no required stops: dnf always installs the newest version, so it can't enforce one. Anything that would need one is a migration, or a background job the new binary carries.[^packages] Down migrations are a development tool, not a production rollback.[^releases]

Each way an unattended upgrade can fail has its own guard:[^packages]

- **A bad build reaching everyone:** the channels and the 48-hour soak.
- **A migration failing on real data:** the upgrade tests, [squawk](tech-stack.md#migrations), and transactional migrations, which leave the schema untouched when they fail.
- **A release broken after it has migrated:** [safe downgrades](#downgrading).
- **A bug that destroys data:** the [snapshot](#snapshots) taken before migrating.

Watching unattended upgrades is your job as the operator. A component's part is to fail clearly: in dnf-automatic's report, in systemd, and in its `doctor` command. It doesn't page anyone.[^packages]

### Snapshots

Before an upgrade migrates the database, the new package's pre-install step runs the installed release's `admin snapshot`, which copies every table out in one consistent transaction. If there's nothing to migrate, it skips the snapshot. If the snapshot fails, the install aborts, the older release keeps running, and dnf-automatic tries again on its next run.[^packages]

`admin restore` rebuilds a snapshot's schema from the migrations the binary carries, which never change once shipped, then copies the rows back. Only you run it, after data loss; nothing runs it automatically.[^packages]

### Downgrading

A plain `dnf downgrade` is a safe way back. Each schema has a compatibility floor, the oldest release that can serve it, and a release serves any schema whose floor it meets, even one newer than itself. A migration that only adds leaves the floor alone. One that removes or changes something raises it, but only to a release that has been on `stable` for at least seven days.[^packages]

So if you upgrade at least weekly, you can always downgrade to what you ran before; otherwise, restore the pre-install snapshot. After downgrading, hold the version with `dnf versionlock`, or dnf-automatic upgrades it again on its next run.[^packages]

## Guarding the API

`buf breaking` runs as an [hk](tech-stack.md#hk) step against the latest release tag, so a change that breaks the Connect-RPC API shows up before it's released.[^releases]

[^releases]: [Releases, migrations, and background work](../sources/202611251316-releases-and-migrations.md)
[^packages]: [Packages, channels, and unattended upgrades](../sources/202611260229-packages-and-unattended-upgrades.md)
