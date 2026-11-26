---
participants: [ivy]
---

# The package signing key and pkgs.stead.ing's hosting

A working session in which ivy and an infrastructure agent hosted `pkgs.stead.ing` and generated the key that signs everything it serves, following the [packages session](202611260229-packages-and-unattended-upgrades.md). The procedure is [`bin/generate-signing-key`](https://github.com/steading-ai/pkgs.stead.ing/blob/main/bin/generate-signing-key) in `steading-ai/pkgs.stead.ing`, and [`docs/signing-key.md`](https://github.com/steading-ai/pkgs.stead.ing/blob/main/docs/signing-key.md) beside it holds the reasoning.

## Hosting

- `pkgs.stead.ing` is served from a Cloudflare R2 bucket, as ivy confirmed. Its custom domain is the bucket's only public address. Terraform in the private `steading-ai/infra` repository manages the bucket, the domain and the rest of the `stead.ing` zone.
- `steading-ai/pkgs.stead.ing` is public, so anyone can audit how their packages are verified, signed and published.
- The publishing workflow writes to the bucket with an R2 token scoped to that bucket alone. ivy creates it in Cloudflare's dashboard and keeps it in the publishing environment, so it never sits in Terraform's state.

## The key

- One OpenPGP key: RSA-4096, a primary key that signs, with no subkeys, no expiry and no passphrase. Its user ID is `Steading Packages (https://pkgs.stead.ing)`.
- Fingerprint: `9C08 AF55 2753 34E2 D862  5917 4F5C 40EB 8078 0443`.
- RSA-4096 because every rpm and librepo that might read the repository can verify it.
- No passphrase because both stores are encrypted, and a passphrase the publishing job needs would have to sit beside the key in the same environment.

## How it was generated

ivy asked for "a reasonably secure approach to key generation", one that would "ensure entropy and safe handling of key material", and to "keep it in RAM securely so it never hits the disk in the first place".

- **Where.** On ivy's Mac, with everything closed but the terminal and 1Password. It was not generated on a server where agent sessions run as the same user, since any process with that user's uid could read a keyring while it exists.
- **Not through the agent.** ivy offered to have the agent hand over the key for 1Password. The agent wrote a script for ivy to run instead, because anything an agent prints or reads becomes part of a session transcript that is stored and sent elsewhere.
- **In RAM.** The keyring was made on a RAM disk, mounted where Finder and Spotlight don't look, and detached on exit. macOS encrypts swap, and the script turned off crash dumps before anything else.
- **Through pipes.** The secret key moved only through pipes, into 1Password and into the environment secret. The script printed nothing but the fingerprint and the public key.
- **Checked.** It read the 1Password copy back to confirm it holds the secret key. It also checked that the key is what was asked for, and that a test signature verifies.
- **Entropy.** It comes from the operating system's kernel generator, which gpg reads. Nothing was installed to add more, since that doesn't strengthen a seeded kernel generator.

The first run stopped partway, after storing a key in 1Password, because `op` reads `-` only from a pipe. That key was discarded, and its 1Password item names a different fingerprint. The second run made the key above.

## Where it lives and who can use it

- **The `SIGNING_KEY` secret of the `publish` environment** in `steading-ai/pkgs.stead.ing`. Only `main` can deploy to the environment, and force-pushes to and deletion of `main` are blocked. One job references it: the publishing job, which signs an rpm only after its attestation verifies against the component's own release workflow at that tag. The environment's `SIGNING_KEY_FINGERPRINT` variable holds the fingerprint, and the job refuses a key that doesn't match it.
- **An offline copy in ivy's 1Password,** with the revocation certificate gpg made with the key.

## Rotation

Only if the key is compromised. A new key is generated the same way and shipped in a new `steading-release` rpm, signed with the old key. Users take it with their next upgrade and do nothing themselves.

## Passed over

- **An expiring key with a yearly renewal,** as ivy decided in the packages session. A missed renewal breaks every user's upgrades at once.
- **A key in a cloud key service,** as ivy decided in the packages session. No tooling turned up that signs rpms and `repomd.xml` with a key it cannot export.
- **COPR,** as ivy decided in the packages session. COPR builds packages itself and signs them with a key the project does not hold, so it cannot serve the attested binaries a release carries.
- **An offline primary key with a signing subkey in CI.** It was considered while writing the procedure. dnf accepts anything the subkey signs until users take an update, which is also the remedy for a compromised primary key, so the split would cost a second key's handling and buy little.

## Open

- The bucket's R2 token isn't in the publishing environment yet.
- `steading-ai/pkgs.stead.ing` needs write access to `ghcr.io/steading-ai/skift` to move the image's `stable`, `latest` and `X.Y` tags. It is granted by hand in the package's settings, and it hadn't been confirmed.
- Nothing reaches `edge` until skift's releases carry attested rpms.
