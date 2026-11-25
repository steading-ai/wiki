---
participants: [ivy]
---

# Tests that run in parallel

A working session in skift where ivy made one testing rule hold across Steading: "it's important that parallel agents sharing a machine are able to run tests in parallel without interfering with each other."

## What prompted it

Two of skift's compose specs used fixed podman names: the compose project `skift-spec` and the image `localhost/skift-contrib-spec:9.9.9`. Three agents running `make test` at once broke each other's runs every time. Fixing that is skift#65. skift's database specs already kept the rule, through testcontainers.

## The rule

Agents work in parallel git worktrees on one machine and run test suites at the same time, so test runs must not interfere with each other. Everything a test run creates gets a name of its own for that run, and is cleaned up afterwards, even when the run fails:

- databases;
- ports: bind to port 0 or to a free port, never a fixed one;
- socket paths;
- compose projects;
- containers, images, volumes, and networks.

Nothing a test uses sits at a shared, well-known name.

## Where it comes from

steading-ai/steading's ADR-006, "The browser suite drives real services with a scripted harness" (`docs/adrs/006-real-services-under-every-test.md`, accepted 2026-09-13), held the rule for one suite. Its decision driver: "Agents run this suite in parallel worktrees on one machine and may not share a database or a port". Its consequence: "Parallel worktrees share nothing but podman."

This makes the rule hold for every component and carries it inside podman: test runs may share the podman daemon, but not names within it, such as compose projects and image tags.
