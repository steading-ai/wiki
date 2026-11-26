---
participants: [ivy]
---

# Hosting the stack

A working session in skift on 2026-11-25, after skift v0.4.0 shipped. It covered three things: running skift as steading-ai's own issue tracker, how a hosted Steading service would be shaped, and where Steading's sites and services live on the web.

## The deliverable

ivy set the next goal: get skift running as a public service for steading-ai's projects, so that it becomes the issue tracker instead of GitHub issues.

The scope is narrow on purpose:

- **Maintainers and their agents only.** "No non-maintainer filing yet. Just a way for me and my team to collaborate."
- **No web UI.** "I'd leave out any form of web UI."
- **No MCP over HTTP.** It came up, through ideas for a stateful MCP over HTTP, and was set aside: "We have something that works and I feel like working on that transport is just going to balloon the scope." Agents keep the CLI and the stdio MCP server. Both reach a remote service with `--address`.

## The hosted service

A hosted service would offer basically the whole stack to anyone who signs up. ivy on its purpose: "The hosted service would be for ease of onboarding and a way to fund the project, not a commercial venture." And: "No sales calls. Self-host it, do whatever you want."

## One organization per instance

ivy decided that tenancy is the deployment, not a column: "Skift uses like 24 MB RSS last I checked and the DB is tiny. I'd just run one process/DB per org and shard it that way." And later: "single-org is where I think we should go." ivy called skift's organization id "a mistake".

What follows from it, as discussed:

- **Isolation is enforced by the deployment.** A query cannot reach another organization's data, because it is in another database. Backup, restore, upgrade, and deletion are per organization.
- **A hosted service needs a control plane**, which skift itself is not. It provisions an instance at sign-up, routes each organization's host to its instance, and rolls upgrades across the fleet.
- **Identity that spans organizations lives outside skift:** for example, a person who belongs to two.
- **skift's `org_id` columns stay for now, constant within an instance.** Removing them would be a migration on every table, and nothing needs that yet.
- **skift's own documents still describe tenancy by column:** principle 8 in its core principles, and its ADR-006. Recording the change there is skift's work, not yet done.

## Where things live

Decided in the session:

- **Landing pages:** `stead.ing/<product>/`, one per component. `<product>.stead.ing` redirects there.
- **Docs:** `docs.stead.ing/<product>/`, built as one site from every component's docs.
  - **Built from releases.** Each component's docs come from its latest release tag, and a component's release rebuilds the site. ivy: "it could publish on tags with the latest *release* rather than publishing unreleased docs."
  - **One file for agents.** The site's root carries `llms.txt` and `llms-full.txt` for the whole suite. ivy wanted the convenience of "docs.stead.ing/llms.txt as a way to discuss the entire suite with an agent" without the landing pages.
  - **Main stops publishing docs.** A component's own site workflow keeps only its landing page. The trade-off accepted: a docs-only fix ships with a release.
- **Accounts:** sign-up and sign-in, when they exist, are first-party, on `stead.ing`.
- **Tenant endpoints:** each organization's stack is reached at `<org>.steading.run`, with Connect procedures routed by their package path to each component's instance.
  - **A domain of its own.** It is kept separate from `stead.ing`, so that a tenant host is never same-site with Steading's own pages and their cookies.
  - **Registered.** ivy registered `steading.run` and `steading.cloud` on 2026-11-25. `steading.cloud` redirects to `stead.ing`.
  - **HTTPS.** `.run` is not on browsers' HSTS preload list, so the service sends HSTS itself.
  - **The Public Suffix List waits.** Its guidelines require more than two years of registration remaining and a permanent `_psl` TXT record, and they say projects "not serving more than thousands of users are quite likely to be declined". Until then, tenant hosts get their cookie isolation from being on a separate domain and from host-only cookies.
  - **Other `steading.*` names are taken.** A check of registries the same day found `steading.app`, `steading.org`, and `steading.work` registered in 2026 on Cloudflare nameservers other than `stead.ing`'s, and `steading.dev` at Amazon. They are probably held by others, which makes near-miss names poor choices for anything people sign in to.

## Pointing a repository at a host

ivy's requirements:

- **A file in each repository names its host,** such as a `.skift.json`.
- **The CLI and the MCP server have a sign-in flow,** and say clearly when sign-in to a service fails.
- **skift stores and retrieves credentials,** in the manner of Git's and Docker's credential helpers.
- **Configuration has a precise specification** for how layers merge and which wins, with `/etc` among them. ivy: "I hate when a tool has multiple ways to configure and documents absolutely nothing or worse, inconsistently applies the winner/loser logic."

Proposed in the session, not yet decided:

- **Sign-in borrows `gh`'s model:**
  - a GitHub device flow that proves identity, after which the service issues its own token;
  - tokens kept in the OS keychain, with a file fallback, keyed by host;
  - an environment variable that overrides the store, for CI;
  - commands for status and logout, and a `doctor` check.

  The stdio MCP server is the same binary, so it needs no sign-in of its own.
- **Configuration layers, lowest to highest:** built-in defaults, system (`/etc`), user, repository, environment, flags. This is Git's order.
  - **Merging:** each field is decided by the highest layer that sets it, and a layer can replace a value but not clear one.
  - **One setting for the service's location:** the socket and the network address together.
  - **Strictness:** unknown keys are errors.
  - **Repository limits:** the repository file may set only the host and the project.
  - **Credentials are never configuration.**
  - **Visibility:** a command shows each effective value with where it came from.
- **Agents use their maintainer's credential at first,** as they use their maintainer's operating-system identity today. Credentials per agent, delegated by a maintainer, come later.

## Set aside

- **Ideas for MCP over HTTP.** None were decided:
  - an "actor" (harness, version, and session) recorded with every write under the authenticated principal;
  - leases renewed by any call the holder makes, rather than by dedicated heartbeat calls. ivy: "I don't really like the idea of agents having to keep their lease alive with tool calls. It sounds wasteful from a turns/tokens/toolcall perspective.";
  - deferring work as an eligibility fact, like a hold with an end time, rather than as a new status.
- **Whether skift's events table should stay partitioned.** Shelved by ivy.

## What stayed open

- **The team's own instance:** its host, and when it moves off the tailnet.
- **The sign-in and configuration design:** which of the proposals above it takes up.
- **Recording single-organization instances in skift:** its ADR, and principle 8's wording.
