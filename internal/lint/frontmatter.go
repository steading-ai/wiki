package lint

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"path/filepath"

	"github.com/santhosh-tekuri/jsonschema/v6"
	"go.yaml.in/yaml/v3"
)

// splitFrontmatter separates a YAML frontmatter block from the Markdown after
// it. offset is where the body starts in raw. A file without frontmatter has
// an empty mapping, so its schema still decides whether that is allowed.
func splitFrontmatter(raw []byte) (front map[string]any, body []byte, offset int, err error) {
	const fence = "---\n"
	if !bytes.HasPrefix(raw, []byte(fence)) {
		return map[string]any{}, raw, 0, nil
	}
	rest := raw[len(fence):]
	var block []byte
	switch {
	case bytes.HasPrefix(rest, []byte(fence)):
		block, offset = nil, 2*len(fence)
	default:
		end := bytes.Index(rest, []byte("\n"+fence))
		if end < 0 {
			if !bytes.HasSuffix(rest, []byte("\n---")) {
				return nil, nil, 0, errors.New("frontmatter opens with --- but never closes")
			}
			end = len(rest) - len("\n---")
		}
		block = rest[:end+1]
		offset = min(len(raw), len(fence)+end+1+len(fence))
	}

	var value any
	if err := yaml.Unmarshal(block, &value); err != nil {
		return nil, nil, 0, fmt.Errorf("frontmatter is not valid YAML: %w", err)
	}
	if value == nil {
		return map[string]any{}, raw[offset:], offset, nil
	}
	m, ok := value.(map[string]any)
	if !ok {
		return nil, nil, 0, errors.New("frontmatter is a mapping of keys to values")
	}
	return m, raw[offset:], offset, nil
}

// A schema is the JSON Schema in schemas/ that one kind of file's frontmatter must satisfy.
type schema struct{ compiled *jsonschema.Schema }

func loadSchemas(root string) (map[string]*schema, error) {
	compiler := jsonschema.NewCompiler()
	out := map[string]*schema{}
	for _, kind := range []string{Pages, Sources} {
		path, err := filepath.Abs(filepath.Join(root, "schemas", kind+".schema.json"))
		if err != nil {
			return nil, err
		}
		compiled, err := compiler.Compile(path)
		if err != nil {
			return nil, fmt.Errorf("schema for %s: %w", kind, err)
		}
		out[kind] = &schema{compiled: compiled}
	}
	return out, nil
}

// validate returns one line per way the frontmatter breaks the schema.
func (s *schema) validate(front map[string]any) []string {
	// The validator works on JSON values, so the YAML goes through JSON first.
	encoded, err := json.Marshal(front)
	if err != nil {
		return []string{fmt.Sprintf("holds a value JSON cannot express: %v", err)}
	}
	value, err := jsonschema.UnmarshalJSON(bytes.NewReader(encoded))
	if err != nil {
		return []string{err.Error()}
	}
	err = s.compiled.Validate(value)
	var invalid *jsonschema.ValidationError
	if !errors.As(err, &invalid) {
		if err != nil {
			return []string{err.Error()}
		}
		return nil
	}
	var out []string
	for _, unit := range invalid.BasicOutput().Errors {
		if unit.Error == nil {
			continue
		}
		at := unit.InstanceLocation
		if at == "" {
			at = "/"
		}
		out = append(out, fmt.Sprintf("%s: %s", at, unit.Error))
	}
	return out
}
