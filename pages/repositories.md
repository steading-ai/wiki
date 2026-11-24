---
type: topic
summary: Which repository holds each project, and which are copies.
---
# Repositories

Steading lives in the [steading-ai](https://github.com/steading-ai) GitHub organization. Concepts often start under a personal account and move to steading-ai to be shared; new work starts in the open in steading-ai.[^almanac][^mill] Where a project exists in both, the steading-ai repository is canonical.[^mill][^thicket][^almanac]

## steading-ai

| Repository | What it holds | Status |
|------------|---------------|--------|
| `skift` | [skift](skift.md)[^skift] | |
| `mill` | [mill](mill.md)[^mill] | |
| `almanac` | [almanac](almanac.md)[^almanac] | |
| `thicket` | [thicket](thicket.md)[^thicket] | A reference implementation[^thicket] |
| `croft` | [croft](croft.md)[^croft] | A concept; its pieces live in mill[^croft] |
| `steading` | The monolithic [web UI](web-ui.md)[^steading] | Private. Poorly named, since it implies the whole system. Renaming and archiving it is undecided.[^steading] |
| `cli` | A one-off for running services built on Linux in a macOS VM[^steading] | Poorly named; to be archived[^steading] |
| `llmrouter` | A sketch of [the gateway](gateway.md)'s routing[^mill] | |
| `mill-evals` | Scratch research on models[^mill] | |
| `scratch-forge` | A test repository for mill's end-to-end tests[^forge] | |
| `agent-plugins` | Nothing; the idea behind it is lost[^plugins] | Empty |
| `stead.ing` | The source of the [stead.ing](https://stead.ing/) website[^steading] | |
| `wiki` | [This wiki](wiki.md) | |

## Personal accounts

These repositories under the `ivy` account bear on Steading.

| Repository | What it holds | Status |
|------------|---------------|--------|
| `mill` | A stale copy of `steading-ai/mill`[^mill] | To be deleted[^mill] |
| `thicket` | A copy of `steading-ai/thicket`[^thicket] | |
| `almanac` | A copy of `steading-ai/almanac`[^almanac] | |
| `kb` | [kb](kb.md)[^kb] | |

[^almanac]: [Projects interview: almanac](../sources/202611150814-projects-in-ivys-words.md#steading-aialmanac)
[^mill]: [Projects interview: mill](../sources/202611150814-projects-in-ivys-words.md#steading-aimill)
[^thicket]: [Projects interview: thicket](../sources/202611150814-projects-in-ivys-words.md#steading-aithicket)
[^skift]: [Projects interview: skift](../sources/202611150814-projects-in-ivys-words.md#steading-aiskift)
[^croft]: [Projects interview: croft](../sources/202611150814-projects-in-ivys-words.md#steading-aicroft)
[^steading]: [Projects interview: steading](../sources/202611150814-projects-in-ivys-words.md#steading-aisteading)
[^forge]: [Projects interview: scratch-forge](../sources/202611150814-projects-in-ivys-words.md#steading-aiscratch-forge)
[^plugins]: [Projects interview: agent-plugins](../sources/202611150814-projects-in-ivys-words.md#steading-aiagent-plugins)
[^kb]: [Projects interview: kb](../sources/202611150814-projects-in-ivys-words.md#ivykb)
