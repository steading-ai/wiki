---
type: project
status: active
repository: steading-ai/mill
summary: Works the stream of tasks in skift, from claiming a node to merging its pull request.
---
# mill

mill works the stream of tasks in [skift](skift.md). It tracks agent definitions, plugins, and skills, and a router, itself a skill, decides which of them run and in what order.[^mill][^thicket] It is what many would call an orchestrator, or a software factory: it supervises claiming skift nodes, executing them, and delivering their artifacts, and it opens pull requests and gets them merged.[^mill]

## Status

mill is built and in use. It also does croft's and the gateway's work for now: see [what belongs elsewhere](#what-belongs-elsewhere).[^steading][^mill]

## Who decides what

A skift node doesn't say how it gets done. mill decides the workflow, the models, and the harness.[^skift] mill picks the shape of the workflow, and administrators manage the prompts and models and what to use when.[^mill] Workflows are linear, not graphs.[^mill]

A mill deployment specifies which projects or tags it takes, so separate deployments can take different kinds of work, such as iOS and macOS development, or work in another ecosystem or with constrained resources.[^mill]

## The golden workflow

The golden example workflows separate an implementer from a reviewer:[^mill]

1. The implementer takes a skift node and plans.
2. The reviewer approves the plan or asks for changes, until the two agree.
3. The implementer builds it.
4. The reviewer checks that the node's acceptance criteria are met.

The reviewer doesn't review code; that is a later consideration. What matters at this stage is that agents don't cut corners or deviate from the vision. Each step remembers earlier rounds, so goalposts don't move and decisions don't waffle.[^mill]

## Merging

Each repository decides what merges without human review. With no status checks or code ownership, mill merges.[^mill] See [the repository gates what merges](principles.md#agents-and-people).

## Costs

mill reports actual costs back into skift.[^mill]

## What belongs elsewhere

mill does work that belongs to other components:[^mill]

- It provisions sandboxes, sets up dev environments, and operates the agent harnesses. That belongs in [croft](croft.md).
- It brokers inference credentials, with no guardrails on what agents can reach. That belongs in [the gateway](gateway.md).

[^mill]: [Projects interview: mill](../sources/202611150814-projects-in-ivys-words.md#steading-aimill)
[^thicket]: [Projects interview: thicket](../sources/202611150814-projects-in-ivys-words.md#steading-aithicket)
[^steading]: [Projects interview: steading](../sources/202611150814-projects-in-ivys-words.md#steading-aisteading)
[^skift]: [Projects interview: skift](../sources/202611150814-projects-in-ivys-words.md#steading-aiskift)
