# The wiki

The wiki is the knowledge base for Steading as an organization: its governance, strategy, roadmap, playbooks, and direction.[^1] Each Steading project documents itself in its own repository and on its own site, and the wiki does not index them.[^1] It follows the spirit of Andrej Karpathy's [LLM Wiki](https://gist.github.com/karpathy/442a6bf555914893e9891c11519de94f): people and agents talk through where Steading is going, and the wiki records what that settles.[^1]

It is published at `wiki.stead.ing`, as a public site.[^1] It is meant to be read on the web and in Obsidian, and indexed with qmd.[^1]

## How it stays trustworthy

- **Every fact cites where it came from.** Sources are dated records of what was said or written. Pages state what is true now and cite their sources in footnotes.[^1]
- **Pages state facts, not history.** A page that narrated how something changed would mislead readers and agents looking back, and its details would rot. How a fact changed lives in the source that recorded it.[^1]
- **Structure is rigid and checked by machine.** Names, layout, frontmatter, and links follow fixed rules, and hooks and CI reject a commit that breaks one.[^1] [AGENTS.md](../AGENTS.md) states the rules.
- **Search instead of curation.** Two flat folders, no hand-kept index, and no page types until one is needed. At scale, search and agents serve readers better than grooming does, and grooming becomes "curating content for an audience of no one."[^1]

[^1]: [Wiki foundations](../sources/202611130854-wiki-foundations.md)
