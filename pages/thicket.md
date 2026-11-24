---
type: project
status: dormant
repository: steading-ai/thicket
summary: An assistant to talk to, plugged into the rest of Steading.
---
# thicket

thicket is an assistant: an easy way to talk to an agent that is plugged into the rest of Steading. A thicket agent can file [skift](skift.md) nodes, check on system status, investigate, and gather details, and it can ping you when your attention is needed.[^thicket][^steading]

thicket exists because Anthropic's policy on using Claude subscriptions with third-party harnesses is ambiguous.[^thicket] It is also a place to explore other ways of interacting with agents, like a live assistant you talk to over the phone.[^thicket]

In the [web UI](web-ui.md), a thicket agent is who you talk to about an inbox item, and the side panel for steering a running agent.[^steading]

## Status

thicket is a reference implementation for making that work, and isn't under active development. The need it serves remains.[^thicket]

## Direction

OpenClaw and Hermes are far more mature assistants. The likely direction is to use them, and to embed any assistant in the suite in an assistant-agnostic way, through A2A, AG-UI, or ACP.[^thicket] See [assistant-agnostic](principles.md#composable).

[^thicket]: [Projects interview: thicket](../sources/202611150814-projects-in-ivys-words.md#steading-aithicket)
[^steading]: [Projects interview: steading](../sources/202611150814-projects-in-ivys-words.md#steading-aisteading)
