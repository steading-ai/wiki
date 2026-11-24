---
type: project
status: concept
repository: ivy/kb
summary: A knowledge base that agents query and that corrects itself.
---
# kb

kb is a knowledge base with agentic retrieval built in and a feedback loop. Agents ask it for details or information, and it answers with citations. Inaccuracies, stale docs, and conflicting facts become corrections.[^kb] kb ingests knowledge from a variety of sources, starting with Git repositories, which it indexes and keeps up to date. That tracks provenance and gives one way to query everything.[^kb]

kb is an iteration from before Steading was a concept. It tries to make every correction on its own; in Steading, a correction to an incorrect doc is a [skift](skift.md) node. kb probably didn't need its own agent, since a subagent would do, and it probably didn't need an MCP server.[^kb]

## Status

The concept is still taking shape, toward something lighter weight.[^kb]

[^kb]: [Projects interview: kb](../sources/202611150814-projects-in-ivys-words.md#ivykb)
