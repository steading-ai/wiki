package lint

import (
	"fmt"
	"regexp"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/extension"
	extast "github.com/yuin/goldmark/extension/ast"
	"github.com/yuin/goldmark/text"
)

// markdown parses GitHub-flavored Markdown with footnotes, as GitHub, Obsidian and the site render it.
var markdown = goldmark.New(goldmark.WithExtensions(extension.GFM, extension.Footnote))

// A document is one file's Markdown body, parsed, with its prose picked out.
type document struct {
	src  []byte
	root ast.Node
	// prose is src with everything that is not running text blanked to
	// spaces: code, HTML, link destinations, and footnote markers. Newlines
	// stay, so a match's offset is its offset in src.
	prose []byte
}

func parse(src []byte) *document {
	root := markdown.Parser().Parse(text.NewReader(src))
	prose := make([]byte, len(src))
	for i, b := range src {
		if b == '\n' {
			prose[i] = '\n'
		} else {
			prose[i] = ' '
		}
	}
	_ = ast.Walk(root, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
		if !entering {
			return ast.WalkContinue, nil
		}
		switch n := n.(type) {
		case *ast.CodeSpan, *ast.RawHTML, *ast.AutoLink:
			return ast.WalkSkipChildren, nil
		case *ast.Text:
			s := n.Segment
			copy(prose[s.Start:s.Stop], src[s.Start:s.Stop])
		}
		return ast.WalkContinue, nil
	})
	return &document{src: src, root: root, prose: prose}
}

// start is where a node begins in src, found through the first text inside it.
func (d *document) start(n ast.Node) int {
	if n.Type() == ast.TypeBlock && n.Lines().Len() > 0 {
		return n.Lines().At(0).Start
	}
	pos := -1
	_ = ast.Walk(n, func(c ast.Node, entering bool) (ast.WalkStatus, error) {
		if !entering {
			return ast.WalkContinue, nil
		}
		if t, ok := c.(*ast.Text); ok {
			pos = t.Segment.Start
			return ast.WalkStop, nil
		}
		if c.Type() == ast.TypeBlock && c.Lines().Len() > 0 {
			pos = c.Lines().At(0).Start
			return ast.WalkStop, nil
		}
		return ast.WalkContinue, nil
	})
	return max(pos, 0)
}

// title requires the file to open with its one level-one heading, which is its title.
func (d *document) title(report func(pos int, msg string)) {
	first := d.root.FirstChild()
	if h, ok := first.(*ast.Heading); !ok || h.Level != 1 {
		pos := 0
		if first != nil {
			pos = d.start(first)
		}
		report(pos, "a file opens with its title, as a level-one heading: # Title")
	}
	for n := d.root.FirstChild(); n != nil; n = n.NextSibling() {
		if h, ok := n.(*ast.Heading); ok && h.Level == 1 && n != first {
			report(d.start(n), "only the title is a level-one heading; sections start at ##")
		}
	}
}

var wikilink = regexp.MustCompile(`!?\[\[`)

func (d *document) wikilinks(report func(pos int)) {
	for _, m := range wikilink.FindAllIndex(d.prose, -1) {
		report(m[0])
	}
}

// footnoteMarker is a footnote reference left in the text, which happens only
// when nothing defines its label.
var footnoteMarker = regexp.MustCompile(`\[\^[^\]\s]+\]`)

// footnoteDefinition is the label of a footnote definition at the start of a line.
var footnoteDefinition = regexp.MustCompile(`(?m)^ {0,3}\[\^([^\]\s]+)\]:`)

// footnotes requires every footnote reference to have a definition, and every
// definition to be referenced. The parser drops both kinds of mistake silently,
// so neither would show on a rendered page.
func (d *document) footnotes(report func(pos int, msg string)) {
	for _, m := range footnoteMarker.FindAllIndex(d.prose, -1) {
		report(m[0], fmt.Sprintf("footnote %s has no definition", d.src[m[0]:m[1]]))
	}
	used := map[string]bool{}
	for _, fn := range d.definitions() {
		used[string(fn.Ref)] = true
	}
	for _, m := range footnoteDefinition.FindAllSubmatchIndex(d.src, -1) {
		if inCode(d.root, m[0]) {
			continue
		}
		label := string(d.src[m[2]:m[3]])
		if !used[label] {
			report(m[0], fmt.Sprintf("footnote [^%s] is defined but never cited in the text", label))
		}
	}
}

// citations requires a page to rest on at least one source, and every one of
// its footnotes to cite a source with a link into sources/.
func (d *document) citations(report func(pos int, msg string), isSource func(dest string) bool) {
	defs := d.definitions()
	if len(defs) == 0 {
		report(0, "a page cites at least one source, with a footnote that links to it: [^1] … [^1]: [Title](../sources/202611131420-slug.md)")
		return
	}
	for _, fn := range defs {
		cites := false
		_ = ast.Walk(fn, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
			if l, ok := n.(*ast.Link); ok && entering && isSource(string(l.Destination)) {
				cites = true
				return ast.WalkStop, nil
			}
			return ast.WalkContinue, nil
		})
		if !cites {
			report(d.start(fn), fmt.Sprintf("footnote [^%s] does not link to a file in sources/", fn.Ref))
		}
	}
}

// definitions are the footnote definitions that something in the text cites.
func (d *document) definitions() []*extast.Footnote {
	var out []*extast.Footnote
	_ = ast.Walk(d.root, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
		if fn, ok := n.(*extast.Footnote); ok && entering {
			out = append(out, fn)
			return ast.WalkSkipChildren, nil
		}
		return ast.WalkContinue, nil
	})
	return out
}

// narrating catches wording that tells how something changed instead of
// stating what is true. Each phrase is one a fact can almost always be
// written without; a page that needs to say what changed cites the source
// that recorded the change.
var narrating = regexp.MustCompile(`(?i)\b(` +
	`previously|originally|formerly|initially|` +
	`no\s+longer|anymore|used\s+to\s+be|at\s+first|until\s+recently|` +
	`(?:is|are|was|were)\s+now|now\s+(?:uses?|lives?|has|is|are)|` +
	`switched\s+(?:from|to)|(?:was|were|has\s+been|have\s+been)\s+(?:changed|replaced|renamed|moved|dropped|removed)` +
	`)\b`)

func (d *document) narration(report func(pos int, phrase string)) {
	for _, m := range narrating.FindAllIndex(d.prose, -1) {
		report(m[0], string(d.src[m[0]:m[1]]))
	}
}

// inCode reports whether pos falls inside a code block.
func inCode(root ast.Node, pos int) bool {
	found := false
	_ = ast.Walk(root, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
		if !entering {
			return ast.WalkContinue, nil
		}
		switch n.(type) {
		case *ast.FencedCodeBlock, *ast.CodeBlock, *ast.HTMLBlock:
			lines := n.Lines()
			for i := range lines.Len() {
				if s := lines.At(i); pos >= s.Start && pos < s.Stop {
					found = true
					return ast.WalkStop, nil
				}
			}
			return ast.WalkSkipChildren, nil
		}
		return ast.WalkContinue, nil
	})
	return found
}
