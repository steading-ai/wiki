# AGENTS.md

## What this repository is

The wiki is Steading's knowledge base for the organization as a whole: how it is governed, where it is going, what it has decided, and how its work gets done. Each Steading project documents itself in its own repository and on its own site; the wiki does not repeat them. It will be published at `wiki.stead.ing`, and every commit is written to be read in public.

People and agents write it together. A conversation about Steading's direction ends with the wiki saying what it settled.

## Layout

```text
pages/      what is true now, one subject per page
sources/    dated records of what was said or written, which pages cite
schemas/    the JSON Schema each kind of file's frontmatter satisfies
cmd/wikilint, internal/lint/
            the checker that enforces this file
```

Both content folders are flat: no folders inside them, and nothing in them but Markdown. Markdown anywhere else sits at the top of the repository.

## Sources

A source records what was said or written at one moment: a conversation, a meeting, a thread, an issue, an article.

- **Name:** the UTC minute it was captured, then a slug: `sources/202611131420-wiki-foundations.md`. Lowercase words joined by hyphens.
- **Frontmatter:** `participants`, the GitHub logins of the people who took part in the conversation or wrote the record. Quoting someone, such as a pasted chat, doesn't make them a participant. An agent acting as a GitHub App is its bot login, quoted in a flow list: `participants: [ivy, "mill[bot]"]`. An agent in an interactive session has no login and is not a participant. [`schemas/sources.schema.json`](schemas/sources.schema.json) is the whole schema.
- **Body:** a title as the first and only level-one heading, then a distilled record: what was decided, what was passed over and why, and what stayed open. Quote exact wording where the wording matters.
- **Never a raw transcript.** Transcripts are noise and carry details of the machine they ran on.
- **A source speaks as of its moment.** A change it records reads "It became X", not "It was X, now it is Y".
- **Sources are not rewritten.** A mistake is corrected in place, and git keeps the history; nothing in the file narrates the correction.

## Pages

A page states what is true now about one subject.

- **Name:** lowercase words joined by hyphens: `pages/decision-records.md`. The name is the page's identity, so pick one that will stay true.
- **Frontmatter:** a `type`, which sets the section the page is listed in, and a one-sentence `summary` for listings. A `project` page, one per Steading component, adds its `status` and, if it has one, its `repository`. Every other page is a `topic`. [`schemas/pages.schema.json`](schemas/pages.schema.json) is the whole schema.
- **Title:** the first and only level-one heading.
- **A project states its status in a `## Status` section** that cites a source, since frontmatter can't hold a footnote.
- **Every page cites its sources** with footnotes that link into `sources/`:

  ```markdown
  The wiki is flat.[^1]

  [^1]: [Wiki foundations](../sources/202611131420-wiki-foundations.md)
  ```

  Every footnote links to at least one source, and every footnote is cited in the text.
- **Facts, not history.** A page never narrates how something changed: no "previously", "no longer", "switched to", "is now". When a fact changes, rewrite the page to state the new fact and cite the source that recorded the change. The checker rejects the common phrasings; rewording past a false match makes the sentence plainer anyway.
- **One page owns each fact.** Other pages link to it instead of restating it.
- **No open questions.** A page states what is known, including what is undecided. It doesn't list questions; every reader draws their own.

## Writing

- GitHub-flavored Markdown. No line length, no table alignment, no formatter.
- Links are Markdown links with relative paths, `[text](other-page.md)` and `[text](../sources/…)`. `[[wikilinks]]` are rejected. The Obsidian settings in [`.obsidian/app.json`](.obsidian/app.json) make Obsidian write the same kind.
- Pages speak to the reader, a user or contributor, as *you*, and for the project, never for one person. No one's opinions, habits, or anecdotes.
- People are named by GitHub login.
- A new page type, and its fields, arrives when a kind of page needs fields of its own, not before.

## Recording a conversation

1. Write the source.
2. Update the pages it changes, or start one when it establishes a subject no page owns. Each changed sentence cites the new source.
3. Run `hk check --all`.
4. Commit with a conventional message. Commits go straight to `main`.

## Checks

`hk check --all` runs every check, and the pre-commit hook and CI run the same ones from [`hk.pkl`](hk.pkl):

| Check | What it holds |
|-------|---------------|
| `go run ./cmd/wikilint` | names, layout, frontmatter against `schemas/`, titles, wikilinks, footnotes, citations, project status sections, and narration on pages |
| `lychee` | every link between files, and every heading a link points at, resolves ([`lychee.toml`](lychee.toml), offline) |
| `go test ./...` | the checker does what this file says |

A scheduled workflow, [`links.yml`](.github/workflows/links.yml), checks links to other sites.

Tools come from [`mise.toml`](mise.toml): `mise install`.
