package entity

import (
	"bytes"
	"encoding/json"
	"sort"
	"strconv"

	"github.com/pkg/errors"
)

func (p *Paths) MarshalJSON() ([]byte, error) {
	if p == nil {
		return []byte("null"), nil
	}

	var keys []string
	for key := range p.Paths {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, key := range keys {
		buf.WriteString(strconv.Quote(key))
		buf.WriteByte(':')
		if err := json.NewEncoder(&buf).Encode(p.Paths[key]); err != nil {
			return nil, errors.Wrapf(err, `failed to encode paths.%s`, key)
		}

		if i < len(keys)-1 {
			buf.WriteByte(',')
		}
	}
	buf.WriteByte('}')
	return buf.Bytes(), nil
}
