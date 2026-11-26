---
type: topic
summary: What the wiki is for, and how it stays trustworthy.
---
# The wiki

The wiki is the knowledge base for Steading as an organization: its governance, strategy, roadmap, playbooks, and direction.[^1] It keeps Steading's organizational knowledge in one place, including what each project is for and what is still uncertain about it, and people and agents use it when changing a project's documentation and steering the work.[^2] Each project documents itself in its own repository and on its own site.[^1] The wiki is Steading's source of truth, for people and agents alike; [stead.ing](https://stead.ing/) and the project sites are the front doors that lead to it.[^3] See [hosting](hosting.md#where-things-live).[^4] It follows the spirit of Andrej Karpathy's [LLM Wiki](https://gist.github.com/karpathy/442a6bf555914893e9891c11519de94f): people and agents talk through where Steading is going, and the wiki records what that settles.[^1]

It is published at `wiki.stead.ing`, as a public site.[^1] It is meant to be read on the web and in Obsidian, and indexed with qmd.[^1]

## Why it exists

The documentation in Steading's projects has rotted. It encodes bad assumptions, and agent reasoning that turns passing remarks into unchanging gospel. The wiki is a shared space for the current truth and the uncertainty instead.[^2]

## How it stays trustworthy

- **Every fact cites where it came from.** Sources are dated records of what was said or written. Pages state what is true now and cite their sources in footnotes.[^1]
- **Pages state facts, not history.** A page that narrated how something changed would mislead readers and agents looking back, and its details would rot. How a fact changed lives in the source that recorded it.[^1]
- **Structure is rigid and checked by machine.** Names, layout, frontmatter, and links follow fixed rules, and hooks and CI reject a commit that breaks one.[^1] [AGENTS.md](../AGENTS.md) states the rules.
- **Search instead of curation.** Two flat folders and no hand-kept index. Each page has a type, which sets the section it is listed in, so the wiki reads as sections, not hundreds of pages; a new type arrives only when a kind of page needs fields of its own.[^1][^2] At scale, search and agents serve readers better than grooming does, and grooming becomes "curating content for an audience of no one."[^1]

[^1]: [Wiki foundations](../sources/202611130854-wiki-foundations.md)
[^3]: [Projects interview: governance](../sources/202611150814-projects-in-ivys-words.md#governance)
[^2]: [Projects interview: wiki](../sources/202611150814-projects-in-ivys-words.md#steading-aiwiki)
[^4]: [Hosting the stack](../sources/202611260037-hosting-the-stack.md)
