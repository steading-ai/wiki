---
type: topic
summary: The principles every Steading component holds to.
---
# Principles

These hold across every Steading component.[^principles]

## Composable

- **Unix philosophy.** Steading is a composable system in which each component can be freely tossed, rewritten, or replaced.[^steading][^principles]
- **One responsibility each.** Each component has a clear, unique responsibility and an API contract. Components work together and are interchangeable, so you adopt the parts you haven't already built, or start from nothing and get it all running. Designing a distributed architecture from day one is a risk, and it buys an architecture that is well understood and can be diagrammed.[^almanac][^principles]
- **Useful alone.** Any single component can be deployed on its own and give immediate value. Users can't do this yet, and it is a major problem for adoption and onboarding.[^steading][^principles]
- **A single static binary.** A binary with everything it needs runs on a scratch base image, is fingerprinted with a checksum, and is replaced easily.[^steading][^principles]
- **Every client uses the same APIs.** The [web UI](web-ui.md), the CLI, and the MCP server use exactly the same APIs. Every service offers its own CLI, which is both client and server. Services listen on Unix sockets locally or on the network, and the CLI and MCP support either. MCP runs through the CLI over stdio, or as HTTP(S) served by each service.[^steading][^principles]
- **Platform-agnostic.** Steading needs to run on Linux and macOS, because collaborators use both. The services don't properly support macOS yet.[^almanac][^steading][^principles]
- **Assistant-agnostic.** Any assistant can be plugged into the suite; [thicket](thicket.md) is one.[^thicket][^principles]

## Agents and people

- **Every user is an executive.** The system runs autonomously instead of acting as a project manager that bothers people with tiny details. A user is an executive, not a supervisor or an individual contributor; whether they get into the small details is theirs to decide. The system flags a person only when something absolutely needs one, such as credentials, billing, and work that isn't sufficiently automated yet.[^principles]
- **Agent-first tools.** Agents get tools built so they can work effectively and build their own context. A tool made for people with an MCP server retrofitted onto it lets agents *use* it, not use it *well*.[^skift][^principles]
- **One task per agent.** Agents work best on a single task and nothing else. Separate roles preserve focus: a person talks through feedback and new ideas with an agent of their own, and that agent changes the plan, so agents building the work aren't interrupted. The less an agent has to keep track of, the better it works: keep its *instruction density* low.[^skift][^principles]
- **Plans change.** Leave room for ambiguity and resteering: putting a feature on hold, prioritizing one thing over another, respeccing from what was learned.[^skift][^principles]
- **The repository gates what merges.** It is each repository's responsibility to decide what merges without human review. With no status checks or code ownership, [mill](mill.md) merges.[^mill][^principles]

## How it's built

- **[One tech stack](tech-stack.md)** for every component.[^thicket][^principles]
- **[Cottagecore names](steading.md#names).**[^thicket][^principles]
- **[Built in the open](steading.md#in-the-open).**[^principles]

## Benchmark against your own work

Not a principle so much as a way to reason about the system: your corpus of past [skift](skift.md) nodes, with their costs, is the benchmark for evaluating new models, alternative harnesses, and planning, in place of public benchmarks.[^principles][^skift] Budgets aren't needed yet; holds and priorities in skift are enough. Budgeting matters once the system builds efficiently enough to consistently exhaust your tokens.[^principles]

[^principles]: [Projects interview: principles](../sources/202611150814-projects-in-ivys-words.md#principles)
[^steading]: [Projects interview: steading](../sources/202611150814-projects-in-ivys-words.md#steading-aisteading)
[^almanac]: [Projects interview: almanac](../sources/202611150814-projects-in-ivys-words.md#steading-aialmanac)
[^thicket]: [Projects interview: thicket](../sources/202611150814-projects-in-ivys-words.md#steading-aithicket)
[^skift]: [Projects interview: skift](../sources/202611150814-projects-in-ivys-words.md#steading-aiskift)
[^mill]: [Projects interview: mill](../sources/202611150814-projects-in-ivys-words.md#steading-aimill)
