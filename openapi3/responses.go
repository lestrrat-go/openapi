package openapi3

import (
	"bytes"
	"encoding/json"
	"sort"
	"strconv"

	"github.com/pkg/errors"
)

func (r *responses) MarshalJSON() ([]byte, error) {
	if r == nil {
		return []byte("null"), nil
	}

	var buf bytes.Buffer

	buf.WriteByte('{')
	if v := r.defaultValue; v != nil {
		buf.WriteString(strconv.Quote("default"))
		buf.WriteByte(':')
		if err := json.NewEncoder(&buf).Encode(v); err != nil {
			return nil, errors.Wrap(err, `failed to encode responses.default`)
		}
	}

	if buf.Len() > 1 && len(r.responses) > 0 {
		buf.WriteByte(',')
	}

	var keys []string
	for key := range r.responses {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for i, key := range keys {
		v := r.responses[key]

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

func (r *responses) UnmarshalJSON(data []byte) error {
	var m map[string]json.RawMessage
	if err := json.Unmarshal(data, &m); err != nil {
		return errors.Wrap(err, `failed to unmarshal responses object`)
	}

	if raw, ok := m["default"]; ok {
		var res response
		if err := json.Unmarshal(raw, &res); err != nil {
			return errors.Wrap(err, `failed to unmarshal responses.default`)
		}
		r.defaultValue = &res
		delete(m, "default")
	}

	if len(m) > 0 {
		r.responses = make(ResponseMap)
		for code, raw := range m {
			// only work with status codes
			codeNum, err := strconv.Atoi(code)
			if err != nil {
				continue
			}
			if codeNum < 100 || codeNum >= 600 {
				continue
			}

			var res response
			if err := json.Unmarshal(raw, &res); err != nil {
				return errors.Wrapf(err, `failed to unmarshal responses.%s`, code)
			}

			r.responses[code] = &res
			delete(m, code)
		}
	}
	return nil
}
