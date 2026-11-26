---
type: topic
summary: How Steading ships its components as signed packages from pkgs.stead.ing.
---
# Packages

`pkgs.stead.ing` is one package repository for every Steading component, on every distribution that gets packages. Fedora rpms come first; other distributions aren't packaged yet.[^packages] The channels are empty until skift's releases carry attested rpms.[^key] How a release reaches the repository, and how it upgrades your machine, is on [releases](releases.md#channels).

## Installing the repository

Install the repository with the small `steading-release` rpm, or with the `.repo` file and the key fingerprint published on the site. `gpgcheck` and `repo_gpgcheck` are both on.[^packages] The key is `Steading Packages (https://pkgs.stead.ing)`, with the fingerprint:[^key]

```text
9C08 AF55 2753 34E2 D862  5917 4F5C 40EB 8078 0443
```

A Steading package never edits repository configuration, never enables a service, and touches nothing outside its own files.[^packages] Every default lives in the binary, so a package never changes its configuration file after installing it, the file holds only what you set, and an upgrade leaves no `.rpmnew` behind.[^packages]

One repository serves every Fedora version: the binaries are static, so nothing is split by `$releasever`. Channels keep every release, so every version stays available to downgrade to or audit; dnf downloads only the metadata, so the size doesn't slow you down.[^packages]

## How a package gets there

A component's release workflow builds its rpms for x86_64 and aarch64 from the same binaries its release archives hold, attests them, and adds them to the release's assets. The rpm isn't built separately. A GitHub release is immutable, so it carries the rpms but not the repository's metadata, which changes with every release.[^packages]

The [`steading-ai/pkgs.stead.ing`](https://github.com/steading-ai/pkgs.stead.ing) repository holds the list of components and channels, the publishing workflow, the signing environment, and the site. A component joins by adding a line to the list, not by changing its own release.[^packages] The repository is public, so anyone can audit how their packages are verified, signed, and published. The packages are served from a Cloudflare R2 bucket whose only public address is `pkgs.stead.ing`, and the publishing workflow writes to it with a token scoped to that bucket alone.[^key]

## Signing

The `pkgs.stead.ing` workflow checks each rpm's attestation against the component's release workflow at that tag, then signs the rpm and the repository metadata, and publishes them. Component releases hold no signing key; the key exists in one place.[^packages][^key]

It's one OpenPGP key: RSA-4096, a primary key that only signs, with no subkeys, no expiry, and no passphrase.[^packages][^key] RSA-4096 is one every rpm and librepo that might read the repository can verify. A passphrase would have to sit beside the key in the same environment, so it would add nothing.[^key]

The private key is the `SIGNING_KEY` secret of the `publish` environment, which only `main` can deploy to; force-pushes to `main` and its deletion are blocked. Only the publishing job uses it, and the job refuses a key whose fingerprint doesn't match. A second copy, with the key's revocation certificate, sits offline in a password manager. The key was generated in RAM, away from the servers where agents run, and moved only through pipes into those two places, so no agent has seen it. [`docs/signing-key.md`](https://github.com/steading-ai/pkgs.stead.ing/blob/main/docs/signing-key.md) has the procedure and its reasoning.[^key]

The key is rotated only if it's compromised: `steading-release` carries the public key, so a rotation arrives as an ordinary package update signed with the old key, and you do nothing.[^packages][^key]

The key lives in CI because dnf checks only OpenPGP signatures, which a GitHub attestation isn't. It doesn't expire, because a missed renewal would break every user's upgrades.[^packages]

[^packages]: [Packages, channels, and unattended upgrades](../sources/202611260229-packages-and-unattended-upgrades.md)
[^key]: [The package signing key and pkgs.stead.ing's hosting](../sources/202611260350-package-signing-key.md)
