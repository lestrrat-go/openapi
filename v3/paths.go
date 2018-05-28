package openapi

import (
	"bytes"
	"encoding/json"
	"sort"
	"strconv"

	"github.com/pkg/errors"
)

func (p *paths) MarshalJSON() ([]byte, error) {
	if p == nil {
		return []byte("null"), nil
	}

	var keys []string
	for key := range p.paths {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, key := range keys {
		buf.WriteString(strconv.Quote(key))
		buf.WriteByte(':')
		if err := json.NewEncoder(&buf).Encode(p.paths[key]); err != nil {
			return nil, errors.Wrapf(err, `failed to encode paths.%s`, key)
		}

		if i < len(keys)-1 {
			buf.WriteByte(',')
		}
	}
	buf.WriteByte('}')
	return buf.Bytes(), nil
}

func (p *paths) UnmarshalJSON(data []byte) error {
	var m map[string]json.RawMessage
	if err := json.Unmarshal(data, &m); err != nil {
		return errors.Wrap(err, `failed to unmarshal JSON`)
	}

	tmp := paths{
		paths: make(map[string]PathItem),
	}

	mutator := MutatePaths(&tmp)

	for path, data := range m {
		var pi pathItem
		if err := json.Unmarshal(data, &pi); err != nil {
			return errors.Wrap(err, `failed to unmarshal JSON`)
		}

		mutator.Path(path, &pi)
	}

	if err := mutator.Do(); err != nil {
		return errors.Wrap(err, `failed to mutate path`)
	}

	*p = tmp
	return nil
}
