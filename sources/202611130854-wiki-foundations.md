---
participants: [ivy]
---

# Wiki foundations

A working session in which ivy set out what the Steading wiki is and the rules it keeps, with an agent researching how others build LLM wikis and proposing a design. What follows is what ivy decided, what was passed over, and what stayed open.

## What the wiki is

The wiki is the central knowledge base for the steading-ai organization, in the spirit of Andrej Karpathy's [LLM Wiki](https://gist.github.com/karpathy/442a6bf555914893e9891c11519de94f) idea. Its subjects are organization governance, playbooks, strategy, roadmap, and project direction. It is "not meant to be an index of all the repos (they have their own sites)." It is the place where ivy discusses direction with agents and has the outcome reflected.

It is meant to be hosted as a website, read in Obsidian, and indexed with qmd. qmd indexing is the end goal and is left out of the first build.

## Where it lives

- It is published at `wiki.stead.ing`.
- It is built separately from `stead.ing` for now, content included. Moving hub content into the wiki, or making the wiki the source of truth for anything, waits until the foundation is in place.
- It will be public. Until it is, the repository stays private.
- No pull requests at this stage; commits go to `main`.

## Shape

- **Two flat directories:** one for the wiki's content and one for sources, with no subdirectories in either.
- **Strict names:** lowercase, hyphenated filenames, held strictly. "The more rigid we can be with regards to frontmatter, naming conventions, and directory hierarchy, the less of a mess we'll deal with later."
- **Search over grooming:** at scale, search and agents fit better than hand-grooming documents. Grooming is how you end up "curating content for an audience of no one. It's procrastination in its finest form."
- **Source names:** sources take Obsidian Zettelkasten-style names, a timestamp followed by a slug. The timestamp became the UTC minute, `YYYYMMDDHHmm`, matching Obsidian's unique-note format.
- **No page types:** "no types until we need them. We'll grow this from scratch and remain disciplined." `type:` is the key that will select a schema when a kind of page needs one.

## Integrity

- **Validation:** YAML frontmatter has a schema, and hooks and CI validate conformance.
- **Provenance:** provenance is required, and citations are Markdown footnotes that link to sources.
- **Corrections:** sources are immutable in principle, though mistakes happen. "Amendments are for Git history."
- **No narration on pages:** "No narrating history in fact docs. Period." "It was X, then we switched to Y" belongs only in sources, and even there a source says "It became X." Narration misleads readers and agents looking back, and its details rot.
- **No raw transcripts:** "No raw transcripts, EVER." Sources are distilled records.
- **Links:** hk and lychee run so that no link stays broken. Links are Markdown links; Obsidian wikilinks are blocked outright.
- **Markdown style:** GitHub-flavored Markdown, with no line-length or table-alignment rules. The checks guard semantic data integrity, not formatting.
- **People:** people are identified by GitHub login.
- **Tooling:** written in Go.
- **Workflow linting:** Steading lints GitHub Actions workflows with jactionlint, which took the place of actionlint.

## Governance

"Steading is a solo project that's trying to grow a community. No DRIs." Work carries no assigned owners.

## Passed over

- **Karpathy's page types:** entity, concept, comparison, and the agent-proposed decision, question, bet, playbook, and topic were all set aside in favor of no types until one is needed.
- **`log.md` and a hand-kept `index.md`:** git history is the log. An index waits until search proves insufficient.
- **An `amendments` list in source frontmatter:** dropped for git history.
- **A `sources:` list in page frontmatter:** footnote citations carry provenance instead.
- **A DRI field on roadmap bets:** there are no DRIs.
- **Style linting:** markdownlint, line length, and table alignment.
- **Building on the stead.ing hub's content folder:** deferred.

## Open

- Whether the hub's governance pages (`how-we-work`, `open-for-good`) move into the wiki once it is the source of truth.
- Which notes from ivy's Obsidian vault (`Projects/Steading/`) migrate into the wiki.
- How `stead.ing` builds a site that is neither the hub nor a project.
