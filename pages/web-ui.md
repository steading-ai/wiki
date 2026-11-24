---
type: project
status: dormant
repository: steading-ai/steading
summary: Shows the system to people, through the same APIs every client uses.
---
# The web UI

Steading's web UI shows the system to people. It uses exactly the same APIs as the CLI and the MCP server, and it is disposable: its purpose is to push on those APIs' ergonomics for future apps.[^steading] Asking for live streaming updates on everything happening in the system, filtered so they aren't a firehose, produced something that carries over to every other client.[^steading]

The web UI is a client, not a server or a proxy. It obeys CORS, and the browser makes requests to each service directly.[^steading] For example, each service can be a Tailscale service with tagged ACLs.[^steading]

## Status

The web UI is a monolith in the private [steading-ai/steading](https://github.com/steading-ai/steading) repository, built as one app on purpose to prove the concept of a single UI tying the suite together. It didn't reach its full vision.[^steading] Parts of its functionality were beautiful, but it has no life and no taste, and it doesn't feel like something you'd want to use. It needs a rebrand that keeps its most successful parts.[^steading]

## What it shows

- **An inbox**, the main place a person acts. It collects what needs them: an agent flags that it needs attention, a decision, or taste, and the person talks it through with a [thicket](thicket.md) agent.[^steading] The inbox is primarily for [skift](skift.md).[^steading]
- **Metrics for past [mill](mill.md) runs:** what they cost in money and tokens, which models and skills they used, and their success rate.[^steading]
- **Active and past runs,** with live output streamed from the running agent harnesses, and a side panel for steering a run through a thicket agent.[^steading]

It has no view of the backlog. The assistant to talk things through with matters more for trust.[^steading]

## Direction

One possible next phase is for every component to bundle its own web UI, in keeping with a [single static binary](principles.md#composable). That means decomposing the monolith into a reusable package of UI components and assets, so each component serves its own piece.[^steading]

[^steading]: [Projects interview: steading](../sources/202611150814-projects-in-ivys-words.md#steading-aisteading)
