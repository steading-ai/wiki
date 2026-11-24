---
type: project
status: active
repository: steading-ai/almanac
summary: Watches the world and gets eventful things into skift.
---
# almanac

almanac watches the world.[^almanac] Named after a farmer's almanac, its role is to:
- Receive events such as webhooks, RSS feeds, and email
- Run routines on a schedule
- Get things into [skift](skift.md) so they can be worked[^almanac]
- Filter and de-duplicate events so only noteworthy/actionable work is started[^almanac]

almanac is deterministic and event-driven. It doesn't run agents itself: agent work goes through skift, and other components run agents and open and merge pull requests.[^almanac]

Example workflows:[^almanac]

- Fix Renovate pull requests whose status checks fail, so Renovate can merge them automatically.
- Watch Alpaca's websocket ticker and act on activity.
- Watch RSS feeds for security vulnerabilities, investigate the blast radius, and make preemptive fixes.
- Any case where an agent belongs in the loop when something eventful happens.

In the vulnerability example, almanac files a node to investigate, and an [assistant](thicket.md) checks your systems for impact and decides what to do: file more nodes, or flag your attention.[^thicket]

## Status

almanac's first iteration, in Rust, is built and in use. It duplicates parts of the suite instead of integrating with them, and will be rewritten in Go, Steading's [tech stack](tech-stack.md), against the other components.[^steading][^almanac]

## Workflows

almanac's core supplies an SDK, in Go to start, and an API contract for workflows. The core is deterministic; workflows needn't be. They are the glue that defines the flow.[^almanac]

## Direction

almanac's direction is a single responsibility, watching the world, with everything else left to the other components.[^almanac]

The intended architecture, fairly settled though no work has started on it, borrows Terraform's provider model:[^almanac]

- Each workflow is its own binary that talks to the core over stdio, using Connect-RPC and protocol buffers with a versioned schema, so workflows can be written in any language.
- The core handles the essentials and primitives. Workflows are free to crash and restart under process supervision.
- A core with a small enough surface allows zero-downtime restarts: handing open file descriptors and serialized state to the new process through systemd's [file descriptor store](https://systemd.io/FILE_DESCRIPTOR_STORE/).
- Each workflow is [socket-activated](https://0pointer.de/blog/projects/socket-activation.html), so it starts lazily and holds no memory while idle.

[^almanac]: [Projects interview: almanac](../sources/202611150814-projects-in-ivys-words.md#steading-aialmanac)
[^steading]: [Projects interview: steading](../sources/202611150814-projects-in-ivys-words.md#steading-aisteading)
[^thicket]: [Projects interview: thicket](../sources/202611150814-projects-in-ivys-words.md#steading-aithicket)
