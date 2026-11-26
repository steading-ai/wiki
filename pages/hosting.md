---
type: topic
summary: Where Steading's sites and services live on the web, and how a hosted Steading is shaped.
---
# Hosting

## Where things live

- **Landing pages:** each component has one at `stead.ing/<product>/`, and `<product>.stead.ing` redirects there.[^hosting]
- **Docs:** `docs.stead.ing/<product>/`, one site built from every component's docs.[^hosting]
  - Each component's docs come from its latest release tag, and a component's [release](releases.md) rebuilds the site, so the docs never describe unreleased work. A docs-only fix ships with a release.[^hosting]
  - The site's root carries `llms.txt` and `llms-full.txt` for the whole suite, so you can hand an agent all of Steading's docs at once.[^hosting]
- **Packages:** `pkgs.stead.ing`, one signed package repository for the whole suite. See [packages](packages.md).[^packages]
- **Accounts:** sign-up and sign-in, when they exist, are first-party, on `stead.ing`.[^hosting]
- **Tenant endpoints:** each organization's stack is reached at `<org>.steading.run`, which routes each Connect-RPC procedure by its package path to that component's instance.[^hosting]
- **This wiki:** `wiki.stead.ing`. See [the wiki](wiki.md).[^foundations]

### Domains

Steading's domains are `stead.ing`, `steading.run`, and `steading.cloud`, which redirects to `stead.ing`. Other `steading.*` names, such as `steading.app`, `steading.org`, `steading.work`, and `steading.dev`, aren't Steading's.[^hosting]

Tenant hosts live on `steading.run`, apart from `stead.ing`, so a tenant host is never same-site with Steading's own pages or their cookies. `.run` isn't on browsers' HSTS preload list, so the service sends HSTS itself. `steading.run` isn't on the Public Suffix List, whose guidelines decline small projects; tenant hosts get their cookie isolation from the separate domain and from host-only cookies.[^hosting]

## The hosted service

A hosted Steading offers the whole stack to anyone who signs up. It exists to make onboarding easy and to fund the project, not as a commercial venture: there are no sales calls, and you can self-host and do whatever you want.[^hosting]

### One organization per instance

Tenancy is the deployment, not a column: each organization gets its own process and its own database for each component. That's cheap, since skift runs in tens of megabytes of memory with a tiny database, and isolation comes from the deployment: a query can't reach another organization's data, because that data is in another database. Backup, restore, upgrade, and deletion are each per organization.[^hosting]

That puts two things outside the components:[^hosting]

- **A control plane**, which provisions an instance at sign-up, routes each organization's host to its instance, and rolls upgrades across the fleet.
- **Identity that spans organizations**, such as a person who belongs to two.

[^hosting]: [Hosting the stack](../sources/202611260037-hosting-the-stack.md)
[^foundations]: [Wiki foundations](../sources/202611130854-wiki-foundations.md)
[^packages]: [Packages, channels, and unattended upgrades](../sources/202611260229-packages-and-unattended-upgrades.md)
