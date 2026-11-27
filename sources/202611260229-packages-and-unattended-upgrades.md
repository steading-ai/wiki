---
participants: [ivy]
---

# Packages, channels, and unattended upgrades

A working session in skift where ivy settled how Steading ships packages: a signed package repository for the whole suite, release channels with automatic promotion, and upgrades that are safe to run unattended. Fedora rpms for skift come first. The decisions hold for every Steading component from here on. skift's [ADR-011](https://github.com/steading-ai/skift/blob/main/docs/adrs/011-releases-are-immutable-and-attested.md) and [ADR-015](https://github.com/steading-ai/skift/blob/main/docs/adrs/015-upgrades-and-migrations.md) will be revised to match; ivy said ADR-011 "just didn't cover" a package repository. What follows is what ivy decided, what was passed over and why, and what stayed open.

## Why

ivy wants every release, pre-releases included, on her own server as soon as it is cut, "so I could shorten the feedback loop", and wants onboarding and upgrades "as convenient as possible for end users". The project is in "this awkward teenage phase" where "it's just inaccessible to newcomers".

On upgrades: "Given the security landscape with agents, it's important that upgrades are done unattended, automatically, and are completely uneventful. I want to ship updates several times a week, if not daily. The faster, the better, and without toil on users." Upgrades are autonomous by default, with an escape hatch for an operator who later wants more care around rollouts.

## Scope

- **Fedora only.** Most newcomers likely run Debian, Ubuntu, Arch, or macOS, but ivy agrees with Fedora's approach to Linux and wants the repository for herself and "an easy path for anyone else who shares my love for this distro". CentOS and RHEL are YAGNI for now. skift's binary is static, so an EL host could likely use the same repository later with little work.
- **Not in scope:** an install script (`curl https://get.stead.ing | sh`) and self-updating binaries served from release channels. ivy holds them as an idea, "not really my style", and kept the work focused on Fedora packaging.

## One package repository for the suite

- `pkgs.stead.ing` serves every Steading component's packages, for every distribution that gets packages. Fedora is the first.
- A GitHub repository named `pkgs.stead.ing`, after the subdomain it is responsible for, holds the list of components and channels, the publishing workflow, the signing environment, and the site. A component joins by adding a line to the list, not by changing its own release.
- It is done "right, not the weird shenanigans that third-party repos tend to do":
  - You install the repository with a small `steading-release` rpm, or with a `.repo` file and a fingerprint published on the site.
  - `gpgcheck` and `repo_gpgcheck` are both on.
  - A package never edits repository configuration, never enables a service, and touches nothing outside its own files.
- One repository serves every Fedora version. The binary is static, so there is no split by `$releasever`.
- **Hosting:** ivy leans toward Cloudflare R2. GitHub Pages was passed over: a repository that keeps every release outgrows its 1 GB site limit, and ivy called serving packages from it "abuse". Another agent provisions the infrastructure.
- **Retention:** channels keep every release forever. Pruning isn't needed for correctness, since dnf downloads only the metadata, and every version stays available to downgrade to or audit. A skift release is about 60 MB of rpms across two architectures, so daily releases add about 22 GB a year.

## Packages carry the attested binaries

- A component's release workflow builds its rpms for x86_64 and aarch64 from the same binaries its release archives hold, attests them, and adds them to the release's assets. The rpm is not built separately.
- GitHub releases are immutable, so a release can carry rpms but not a repository's metadata, which changes with every release. The repository lives at `pkgs.stead.ing`.

## Signing

- The `pkgs.stead.ing` workflow checks each rpm's attestation against the component's release workflow, then signs the rpm and the repository metadata, and publishes them.
- Component releases hold no signing key. The key exists in one place.
- One OpenPGP signing key with no expiry. The private key is a secret in a protected environment of the `pkgs.stead.ing` repository, usable only by the publishing job on `main`, with a second copy in a password manager.
- It is rotated only if it is compromised. The `steading-release` rpm carries the public key, so a rotation is an ordinary package update signed with the old key, and users do nothing.
- ivy accepted the long-lived key in CI. dnf checks only OpenPGP signatures, which a GitHub attestation is not.
- **Passed over:** an expiring key with a yearly renewal. Forgetting to renew it would break every user's upgrades, and ivy doesn't want "key signing ceremonies" or to rotate keys every year. A key held in a cloud key service was passed over too: no tooling for signing rpms that way turned up.

## Channels and versions

- **Versions:** every release gets a plain SemVer version with no pre-release suffix. The version says what changed in the public API; the channel says how proven the build is. This holds for every component from here on, starting with skift. It replaces the tilde pre-release versions (`0.4.0~rc.2`).
- **Channels:** two, `edge` and `stable`. `edge` gets every release as soon as it is tagged. `stable` gets the same signed rpm when it is promoted, which changes only the repository's metadata, and nothing is rebuilt. The model follows Fedora CoreOS's streams and Claude Code's release channels.
- **Images** get `edge` and `stable` tags that move the same way, alongside `X.Y.Z`.
- **ivy's own server** subscribes to `edge` with `dnf-automatic`. It is the first canary, and the script it uses today to build and install skift from source goes away.

## Promotion

- Promotion is automatic. `stable` moves to the newest `edge` build that has been in `edge` for at least 48 hours and hasn't been yanked. ivy chose 48 hours over 24 to have "a smidge more time to use it myself".
- Choosing the newest build that has soaked, rather than resetting the clock on each new build, keeps `stable` moving when releases ship daily.
- Two manual controls: **promote now**, for security fixes, and **yank**, which removes a build from both channels. Yanking stops a bad build from spreading. dnf never downgrades a machine that already installed it, so the fix for those machines is a new, higher version.
- **Passed over for now:** manual promotion, which is toil and lets `stable` fall behind; promotion on a fixed weekly schedule, which delays security fixes; and promotion gated on health reports from canaries, which waits until there is more than one canary.

## Unattended upgrades

An upgrade can fail in four ways, and each has its own guard:

- **A bad build reaches everyone:** channels and the 48-hour soak.
- **A migration fails on real data:** the upgrade spec, squawk, and transactional migrations, which leave the schema untouched when they fail.
- **A release is broken after it has migrated:** safe downgrades, below.
- **A bug destroys data:** the pre-install snapshot, below.

Two more things would make an upgrade eventful, and both are ruled out:

- **No required stops.** A required stop is a release every operator must install and run before upgrading further. dnf always installs the newest version, so it cannot enforce one, and ivy called the idea "awful" and "a no-go". Anything that would have needed one becomes a migration, or a background job the new binary carries.
- **No `.rpmnew` files.** Every default moves into the binary, so the package never changes `/etc/skift/env` after installing it, and the file holds only what the operator set.

Monitoring unattended upgrades is the operator's job. skift's part is to fail clearly: in dnf-automatic's report, in systemd, and in `skift doctor`. skift doesn't page anyone.

## Snapshots before migrating

- ivy chose GitLab Omnibus's model, which backs up the database in the package's pre-install step and aborts the install when the backup fails.
- The new package's pre-install step runs the installed, older binary's `skift admin snapshot`. It skips the snapshot when the database is already at the new release's schema, so most upgrades do nothing extra.
- If the snapshot fails, the upgrade doesn't happen: the install aborts, the older release keeps running, and dnf-automatic reports the failure and tries again on its next run.
- A snapshot taken when `skift serve` starts was passed over. By then the rpm has replaced the binary, so a failed snapshot would leave the service down instead of not upgraded.
- skift writes the snapshot itself, in Go over pgx: every table copied out within one consistent transaction, to `/var/lib/skift/backups/`, labelled with its schema version, keeping the last few. `pg_dump` was passed over because its version must match the server's, and because a snapshot has to work against a remote database.
- `skift admin snapshot` and `skift admin restore` are in scope as the mechanism the packaging calls. `restore` rebuilds the snapshot's schema from the migrations the binary carries, which never change once shipped, then copies the rows back. Only an operator runs it, after data loss; nothing runs it automatically. Scheduled backups remain in skift's M4 milestone.
- The first upgrade that takes a snapshot is the one from the first release that has the command.

## Safe downgrades

- A plain `dnf downgrade` is a safe way back from here on.
- Each schema has a **compatibility floor**: the oldest release that can serve it. A release serves any schema whose floor it meets, even one newer than itself, instead of refusing every schema newer than itself. A migration that only adds leaves the floor alone. One that removes or changes something raises it, and may raise it only to a release that has been on `stable` for at least seven days. Anyone who upgrades at least weekly can always downgrade to what they ran before; anyone else restores the pre-install snapshot.
- "One release back" was passed over: with daily releases and automatic promotion, `stable` can jump several versions at once, so the previous version is not what a `stable` user last ran.
- v0.4.0 refuses any newer schema, so downgrading to it works only if the release after it adds no migration. Safe downgrades start from the first release that has the floor.
- An operator who downgrades also holds the version with `dnf versionlock`, or dnf-automatic upgrades it again on its next run.

## Open

- Whether dnf5 cleanly aborts only the one package when its pre-install step fails, leaving the rest of the transaction alone. A spike settles it before the design relies on it.
- How the floor a migration sets is checked by a machine rather than by review.
- Whether the rpm or the containers are the primary way to run a server. That turns on how much callers on the host matter, since only a host install gives them identity from the socket's uid.
- skift's admin commands need `DATABASE_URL`, which sits awkwardly with the principle that the API is the only door, and `skift doctor` doesn't compare the CLI's version with the server's. Both were noted, and neither was decided.
