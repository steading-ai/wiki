// Package lint checks that the wiki keeps its shape: where files live, what
// they are named, what their frontmatter holds, how pages cite sources, and
// that pages state facts instead of narrating how they changed. AGENTS.md
// states each rule for the people and agents who write the wiki; this package
// enforces them.
package lint

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strings"
	"time"
)

// The two directories that hold the wiki's content, and the kind of file each holds.
const (
	Pages   = "pages"
	Sources = "sources"
)

// A Problem is one rule a file breaks, at the line where it breaks it.
// Line is zero when the problem is with the file as a whole.
type Problem struct {
	Path    string
	Line    int
	Message string
}

func (p Problem) String() string {
	if p.Line == 0 {
		return fmt.Sprintf("%s: %s", p.Path, p.Message)
	}
	return fmt.Sprintf("%s:%d: %s", p.Path, p.Line, p.Message)
}

var (
	slug       = `[a-z0-9]+(?:-[a-z0-9]+)*`
	pageName   = regexp.MustCompile(`^` + slug + `\.md$`)
	sourceName = regexp.MustCompile(`^(\d{12})-` + slug + `\.md$`)
)

// Source filenames open with the minute they were captured, in UTC.
const stamp = "200601021504"

// Check walks the wiki rooted at root and returns every problem it finds,
// ordered by path and line. An error means the check itself could not run.
func Check(root string) ([]Problem, error) {
	schemas, err := loadSchemas(root)
	if err != nil {
		return nil, err
	}
	c := &checker{root: root, schemas: schemas}
	if err := c.layout(); err != nil {
		return nil, err
	}
	for _, kind := range []string{Pages, Sources} {
		entries, err := os.ReadDir(filepath.Join(root, kind))
		if err != nil {
			if os.IsNotExist(err) {
				continue
			}
			return nil, err
		}
		for _, e := range entries {
			if hidden(e.Name()) {
				continue
			}
			rel := kind + "/" + e.Name()
			if !e.Type().IsRegular() {
				c.report(rel, 0, "%s/ holds only Markdown files, with no folders or links inside it", kind)
				continue
			}
			if err := c.file(kind, rel); err != nil {
				return nil, err
			}
		}
	}
	slices.SortStableFunc(c.problems, func(a, b Problem) int {
		if n := strings.Compare(a.Path, b.Path); n != 0 {
			return n
		}
		return a.Line - b.Line
	})
	return c.problems, nil
}

type checker struct {
	root     string
	schemas  map[string]*schema
	problems []Problem
}

func (c *checker) report(path string, line int, format string, args ...any) {
	c.problems = append(c.problems, Problem{Path: path, Line: line, Message: fmt.Sprintf(format, args...)})
}

// layout finds Markdown kept anywhere but pages/, sources/, and the top of the repository.
func (c *checker) layout() error {
	return filepath.WalkDir(c.root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(c.root, path)
		if err != nil {
			return err
		}
		rel = filepath.ToSlash(rel)
		if rel == "." {
			return nil
		}
		if hidden(d.Name()) {
			if d.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}
		if d.IsDir() {
			if rel == Pages || rel == Sources {
				return filepath.SkipDir
			}
			return nil
		}
		if strings.HasSuffix(rel, ".md") && strings.Contains(rel, "/") {
			c.report(rel, 0, "Markdown lives in pages/, sources/, or the top of the repository")
		}
		return nil
	})
}

func (c *checker) file(kind, rel string) error {
	name := filepath.Base(rel)
	switch kind {
	case Pages:
		if !pageName.MatchString(name) {
			c.report(rel, 0, "a page's name is lowercase words joined by hyphens, ending in .md: decision-records.md")
		}
	case Sources:
		m := sourceName.FindStringSubmatch(name)
		if m == nil {
			c.report(rel, 0, "a source's name is the UTC minute it was captured, a hyphen, and lowercase words joined by hyphens: 202611131420-wiki-foundations.md")
		} else if _, err := time.Parse(stamp, m[1]); err != nil {
			c.report(rel, 0, "%s is not a real minute; a source's name opens with YYYYMMDDHHmm in UTC", m[1])
		}
	}

	raw, err := os.ReadFile(filepath.Join(c.root, rel))
	if err != nil {
		return err
	}
	front, body, offset, err := splitFrontmatter(raw)
	if err != nil {
		c.report(rel, 1, "%v", err)
		return nil
	}
	for _, p := range c.schemas[kind].validate(front) {
		c.report(rel, 1, "frontmatter: %s", p)
	}

	doc := parse(body)
	line := func(pos int) int { return lineOf(raw, offset+pos) }
	doc.title(func(pos int, msg string) { c.report(rel, line(pos), "%s", msg) })
	doc.wikilinks(func(pos int) {
		c.report(rel, line(pos), "use a Markdown link, [text](other-page.md), not a [[wikilink]]")
	})
	doc.footnotes(func(pos int, msg string) { c.report(rel, line(pos), "%s", msg) })
	if kind == Pages {
		doc.citations(func(pos int, msg string) { c.report(rel, line(pos), "%s", msg) }, c.sourceExists)
		if front["type"] == "project" {
			doc.status(func(pos int, msg string) { c.report(rel, line(pos), "%s", msg) })
		}
		doc.narration(func(pos int, phrase string) {
			c.report(rel, line(pos), "%q narrates a change; a page states what is true now, and the source records how it changed", phrase)
		})
	}
	return nil
}

// sourceExists reports whether a link written on a page leads to a file in sources/.
func (c *checker) sourceExists(dest string) bool {
	dest, _, _ = strings.Cut(dest, "#")
	target := filepath.ToSlash(filepath.Clean(filepath.Join(Pages, dest)))
	dir, name := filepath.Split(target)
	if dir != Sources+"/" || !strings.HasSuffix(name, ".md") {
		return false
	}
	info, err := os.Stat(filepath.Join(c.root, target))
	return err == nil && info.Mode().IsRegular()
}

func hidden(name string) bool { return strings.HasPrefix(name, ".") }

// lineOf is the one-based line of the byte at pos.
func lineOf(raw []byte, pos int) int {
	pos = min(pos, len(raw))
	return strings.Count(string(raw[:pos]), "\n") + 1
}
