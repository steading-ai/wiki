package lint

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

const goodSource = `---
participants: [ivy]
---

# Wiki foundations

It became flat.
`

const goodPage = `# The wiki

The wiki is flat.[^1]

[^1]: [Wiki foundations](../sources/202611131420-wiki-foundations.md)
`

// topic is the frontmatter of a page that is not a project.
const topic = "---\ntype: topic\nsummary: What the wiki is.\n---\n"

// project is the frontmatter of a project page.
const project = "---\ntype: project\nstatus: concept\nrepository: steading-ai/skift\nsummary: Tracks the work.\n---\n"

const goodProject = project + `# skift

skift tracks the work.[^1]

## Status

skift is a concept.[^1]

[^1]: [Wiki foundations](../sources/202611131420-wiki-foundations.md)
`

// wiki builds a wiki in a temporary folder from a map of paths to contents,
// using the repository's own schemas, and returns where it is.
func wiki(t *testing.T, files map[string]string) string {
	t.Helper()
	root := t.TempDir()
	for _, kind := range []string{Pages, Sources} {
		schema, err := os.ReadFile(filepath.Join("..", "..", "schemas", kind+".schema.json"))
		if err != nil {
			t.Fatal(err)
		}
		write(t, root, "schemas/"+kind+".schema.json", string(schema))
	}
	for path, content := range files {
		write(t, root, path, content)
	}
	return root
}

func write(t *testing.T, root, path, content string) {
	t.Helper()
	full := filepath.Join(root, path)
	if err := os.MkdirAll(filepath.Dir(full), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(full, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
}

func check(t *testing.T, files map[string]string) []string {
	t.Helper()
	problems, err := Check(wiki(t, files))
	if err != nil {
		t.Fatal(err)
	}
	var out []string
	for _, p := range problems {
		out = append(out, p.String())
	}
	return out
}

func TestAWellFormedWikiPasses(t *testing.T) {
	got := check(t, map[string]string{
		"README.md": "# Wiki\n",
		"sources/202611131420-wiki-foundations.md":     goodSource,
		"pages/the-wiki.md":                            topic + goodPage,
		"pages/skift.md":                               goodProject,
		"pages/second-page.md":                         topic + "# Second\n\nSee [the wiki](the-wiki.md).[^a]\n\n[^a]: Recorded in [foundations](../sources/202611131420-wiki-foundations.md#wiki-foundations).\n",
		"pages/code-is-not-prose.md":                   topic + "# Code\n\nFacts.[^1] `[[not a link]]` and `previously`.\n\n```text\n[[still not a link]] used to be\n[^2]: not a definition\n```\n\n[^1]: [Wiki foundations](../sources/202611131420-wiki-foundations.md)\n",
		"sources/202611131421-history-is-fine-here.md": "---\nparticipants: [ivy, \"mill[bot]\"]\n---\n\n# History\n\nThe wiki was previously nested. It became flat.\n",
	})
	if len(got) != 0 {
		t.Fatalf("want no problems, got:\n%s", strings.Join(got, "\n"))
	}
}

func TestProblems(t *testing.T) {
	cases := []struct {
		name  string
		files map[string]string
		want  string
	}{
		{
			name:  "page name with capitals",
			files: map[string]string{"pages/The-Wiki.md": goodPage},
			want:  "pages/The-Wiki.md: a page's name is lowercase",
		},
		{
			name:  "page name with underscore",
			files: map[string]string{"pages/the_wiki.md": goodPage},
			want:  "pages/the_wiki.md: a page's name is lowercase",
		},
		{
			name:  "source without a timestamp",
			files: map[string]string{"sources/wiki-foundations.md": goodSource},
			want:  "sources/wiki-foundations.md: a source's name is the UTC minute",
		},
		{
			name:  "source with seconds in its timestamp",
			files: map[string]string{"sources/20261113142000-wiki.md": goodSource},
			want:  "sources/20261113142000-wiki.md: a source's name is the UTC minute",
		},
		{
			name:  "source with an impossible minute",
			files: map[string]string{"sources/202611131475-wiki.md": goodSource},
			want:  "sources/202611131475-wiki.md: 202611131475 is not a real minute",
		},
		{
			name:  "folder inside pages",
			files: map[string]string{"pages/nested/the-wiki.md": goodPage},
			want:  "pages/nested: pages/ holds only Markdown files",
		},
		{
			name:  "Markdown outside the content folders",
			files: map[string]string{"notes/idea.md": "# Idea\n"},
			want:  "notes/idea.md: Markdown lives in pages/, sources/, or the top",
		},
		{
			name:  "page without a type",
			files: map[string]string{"pages/the-wiki.md": goodPage},
			want:  "pages/the-wiki.md:1: frontmatter: /: missing properties 'type', 'summary'",
		},
		{
			name:  "page with an unknown type",
			files: map[string]string{"pages/the-wiki.md": "---\ntype: decision\nsummary: What the wiki is.\n---\n" + goodPage},
			want:  "pages/the-wiki.md:1: frontmatter: /type:",
		},
		{
			name:  "project without a status",
			files: map[string]string{"pages/skift.md": strings.Replace(goodProject, "status: concept\n", "", 1)},
			want:  "pages/skift.md:1: frontmatter: /: missing property 'status'",
		},
		{
			name:  "project with an unknown status",
			files: map[string]string{"pages/skift.md": strings.Replace(goodProject, "status: concept", "status: shipped", 1)},
			want:  "pages/skift.md:1: frontmatter: /status:",
		},
		{
			name:  "topic with a status",
			files: map[string]string{"pages/the-wiki.md": "---\ntype: topic\nstatus: active\nsummary: What the wiki is.\n---\n" + goodPage},
			want:  "pages/the-wiki.md:1: frontmatter: /status:",
		},
		{
			name:  "project without a Status section",
			files: map[string]string{"pages/skift.md": project + goodPage},
			want:  "pages/skift.md:7: a project page states its status in a ## Status section",
		},
		{
			name:  "project whose Status section cites nothing",
			files: map[string]string{"pages/skift.md": strings.Replace(goodProject, "skift is a concept.[^1]", "skift is a concept.", 1)},
			want:  "pages/skift.md:11: a project's ## Status section cites a source",
		},
		{
			name:  "source without participants",
			files: map[string]string{"sources/202611131420-wiki.md": "# Wiki\n"},
			want:  "sources/202611131420-wiki.md:1: frontmatter: /: missing property 'participants'",
		},
		{
			name:  "participant that is not a GitHub login",
			files: map[string]string{"sources/202611131420-wiki.md": "---\nparticipants: [Ivy Evans]\n---\n\n# Wiki\n"},
			want:  "sources/202611131420-wiki.md:1: frontmatter: /participants/0:",
		},
		{
			name:  "frontmatter that is not YAML",
			files: map[string]string{"sources/202611131420-wiki.md": "---\nparticipants: [ivy\n---\n\n# Wiki\n"},
			want:  "sources/202611131420-wiki.md:1: frontmatter is not valid YAML",
		},
		{
			name:  "no title",
			files: map[string]string{"sources/202611131420-wiki.md": "---\nparticipants: [ivy]\n---\n\nIt became flat.\n"},
			want:  "sources/202611131420-wiki.md:5: a file opens with its title",
		},
		{
			name:  "second level-one heading",
			files: map[string]string{"sources/202611131420-wiki.md": goodSource + "\n# Another\n"},
			want:  "sources/202611131420-wiki.md:9: only the title is a level-one heading",
		},
		{
			name:  "wikilink",
			files: map[string]string{"sources/202611131420-wiki.md": goodSource + "\nSee [[the-wiki]].\n"},
			want:  "sources/202611131420-wiki.md:9: use a Markdown link",
		},
		{
			name:  "page without citations",
			files: map[string]string{"pages/the-wiki.md": "# The wiki\n\nThe wiki is flat.\n"},
			want:  "pages/the-wiki.md:1: a page cites at least one source",
		},
		{
			name: "footnote that cites no source",
			files: map[string]string{
				"pages/the-wiki.md": "# The wiki\n\nThe wiki is flat.[^1]\n\n[^1]: Because we said so.\n",
			},
			want: "pages/the-wiki.md:5: footnote [^1] does not link to a file in sources/",
		},
		{
			name: "footnote that links to a missing source",
			files: map[string]string{
				"pages/the-wiki.md": "# The wiki\n\nThe wiki is flat.[^1]\n\n[^1]: [Gone](../sources/202611131420-gone.md)\n",
			},
			want: "pages/the-wiki.md:5: footnote [^1] does not link to a file in sources/",
		},
		{
			name: "footnote that links to a page",
			files: map[string]string{
				"pages/other.md":    goodPage,
				"pages/the-wiki.md": "# The wiki\n\nThe wiki is flat.[^1]\n\n[^1]: [Other](other.md)\n",
			},
			want: "pages/the-wiki.md:5: footnote [^1] does not link to a file in sources/",
		},
		{
			name:  "footnote reference without a definition",
			files: map[string]string{"pages/the-wiki.md": goodPage + "\nAlso true.[^2]\n"},
			want:  "pages/the-wiki.md:7: footnote [^2] has no definition",
		},
		{
			name:  "footnote definition never cited",
			files: map[string]string{"pages/the-wiki.md": goodPage + "[^2]: [Wiki foundations](../sources/202611131420-wiki-foundations.md)\n"},
			want:  "pages/the-wiki.md:6: footnote [^2] is defined but never cited",
		},
		{
			name:  "narration on a page",
			files: map[string]string{"pages/the-wiki.md": "# The wiki\n\nThe wiki was previously nested.[^1]\n\n[^1]: [Wiki foundations](../sources/202611131420-wiki-foundations.md)\n"},
			want:  `pages/the-wiki.md:3: "previously" narrates a change`,
		},
		{
			name:  "narration across a line break",
			files: map[string]string{"pages/the-wiki.md": "# The wiki\n\nThe wiki is no\nlonger nested.[^1]\n\n[^1]: [Wiki foundations](../sources/202611131420-wiki-foundations.md)\n"},
			want:  `pages/the-wiki.md:3: "no\nlonger" narrates a change`,
		},
		{
			name:  "narration in a footnote",
			files: map[string]string{"pages/the-wiki.md": "# The wiki\n\nThe wiki is flat.[^1]\n\n[^1]: It switched to flat in [Wiki foundations](../sources/202611131420-wiki-foundations.md)\n"},
			want:  `pages/the-wiki.md:5: "switched to" narrates a change`,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			files := map[string]string{"sources/202611131420-wiki-foundations.md": goodSource}
			for k, v := range c.files {
				files[k] = v
			}
			got := check(t, files)
			for _, p := range got {
				if strings.HasPrefix(p, c.want) {
					return
				}
			}
			t.Fatalf("want a problem starting %q, got:\n%s", c.want, strings.Join(got, "\n"))
		})
	}
}
