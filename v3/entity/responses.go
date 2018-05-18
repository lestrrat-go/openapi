package entity

import (
	"bytes"
	"encoding/json"
	"sort"
	"strconv"

	"github.com/pkg/errors"
)

func (r *Responses) MarshalJSON() ([]byte, error) {
	if r == nil {
		return []byte("null"), nil
	}

	var buf bytes.Buffer

	buf.WriteByte('{')
	if v := r.Default; v != nil {
		buf.WriteString(strconv.Quote("default"))
		buf.WriteByte(':')
		if err := json.NewEncoder(&buf).Encode(v); err != nil {
			return nil, errors.Wrap(err, `failed to encode responses.default`)
		}
	}

	if buf.Len() > 1 && len(r.StatusCodes) > 0 {
		buf.WriteByte(',')
	}

	var keys []string
	for key := range r.StatusCodes {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for i, key := range keys {
		v := r.StatusCodes[key]

		buf.WriteString(strconv.Quote(key))
		buf.WriteByte(':')
		if err := json.NewEncoder(&buf).Encode(v); err != nil {
			return nil, errors.Wrapf(err, `failed to encode responses.%s`, key)
		}

		if i < len(keys)-1 {
			buf.WriteByte(',')
		}
	}

	buf.WriteByte('}')
	return buf.Bytes(), nil
}
