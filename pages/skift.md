---
type: project
status: active
repository: steading-ai/skift
summary: Tracks the work, as an agent-first backlog shaped like a graph.
---
# skift

skift is Steading's agent-first project management system. It tracks the work.[^skift][^thicket]

Project management systems like Trello, Pivotal Tracker, GitHub Projects, Linear, Asana, and Jira are made for people. skift is made for agents: its tools let an agent work effectively and build its own context.[^skift]

## Status

skift is built and in use.[^steading]

## A backlog that is a graph

skift's backlog is a directed acyclic graph. A goal decomposes into prerequisites the way a package manager such as Gentoo's resolves what installing ffmpeg needs, except that software has ambiguity and plans change. Agents are effective at identifying prerequisites and bad at keeping them all in context; skift keeps them.[^skift]

- An agent decomposes a goal into right-sized nodes with dependencies, and separates what an agent can act on from what needs a person.[^skift]
- A node sees the whole graph that feeds into it: the delivered artifacts, course corrections, research, findings, assessments, and decisions upstream of it, not only its own task body.[^skift]
- The graph leaves room for ambiguity and resteering.[^skift] See [plans change](principles.md#agents-and-people).

One branch of the graph for a goal of learning Japanese through full immersion: research SRS and other study methods, review the findings and decide, research libraries and algorithms to build or buy, ask a person if there's no clear winner, then build it.[^skift]

## Nodes

A node is as simple as possible. Its data model covers the specification of the work, delivered artifacts, priority, tags, groups, and accounting.[^skift] A node does not specify a model; [mill](mill.md) decides how work gets done.[^skift] Agents attach the artifacts of their deliverables as a URI or a blob: a commit, a document, an asset, or anything else.[^skift]

## A work queue

Agents tell skift what kinds of work they can take; they don't ask for specific tasks. An agent holds a lease on a node and sends heartbeats.[^skift]

## People

You can take work yourself, but mostly you talk to an agent acting on your behalf, in a separate role from the agents building. While mill works through nodes, you keep your own agent, such as a Claude Code session, talking to skift: you review what interests you, hash out ideas and screenshots with it, and have it change skift when you're ready. You do the work yourself when it needs keys or approval an agent can't have.[^skift] See [one task per agent](principles.md#agents-and-people).

The [web UI](web-ui.md)'s inbox is primarily for skift. The inbox and notifications bring eventful, actionable things to people's attention where it's needed, and nothing else.[^steading]

## Accounting

skift records estimates and actual costs; only actuals are recorded so far, and mill reports them.[^skift][^mill] The ledger opens the door to evals: run a new model through the exact same checkouts, and compare its artifacts and costs against what was delivered before. That gives feedback on cost against effectiveness for your own work.[^skift] Holds and priorities in skift stand in for budgets.[^principles]

[^skift]: [Projects interview: skift](../sources/202611150814-projects-in-ivys-words.md#steading-aiskift)
[^thicket]: [Projects interview: thicket](../sources/202611150814-projects-in-ivys-words.md#steading-aithicket)
[^steading]: [Projects interview: steading](../sources/202611150814-projects-in-ivys-words.md#steading-aisteading)
[^mill]: [Projects interview: mill](../sources/202611150814-projects-in-ivys-words.md#steading-aimill)
[^principles]: [Projects interview: principles](../sources/202611150814-projects-in-ivys-words.md#principles)
