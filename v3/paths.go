package openapi

import (
	"bytes"
	"encoding/json"
	"sort"
	"strconv"

	"github.com/pkg/errors"
)

// Items returns an iterator that you can use to iterate through all
// registered PathItem objects 
func (p *paths) Items() *PathItemIterator {
	var keys []string
	for key := range p.paths {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	var items []PathItem
	for _, key := range keys {
		items = append(items, p.paths[key])
	}

	return &PathItemIterator{
		items: items,
	}
}

// Next returns true if there are more elements to process in this iterator
func (iter *PathItemIterator) Next() bool {
	iter.mu.RLock()
	defer iter.mu.RUnlock()
	return iter.nextNoLock()
}

func (iter *PathItemIterator) nextNoLock() bool {
	return len(iter.items) > 0
}

// Item returns the next item in the iterator
func (iter *PathItemIterator) Item() PathItem {
	iter.mu.Lock()
	defer iter.mu.Unlock()

	if !iter.nextNoLock() {
		return nil
	}

	item := iter.items[0]
	iter.items = iter.items[1:]
	return item
}

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

	*p = paths{
		paths: make(map[string]PathItem),
	}

	for path, data := range m {
		var pi pathItem
		if err := json.Unmarshal(data, &pi); err != nil {
			return errors.Wrap(err, `failed to unmarshal JSON`)
		}

		p.paths[path] = &pi
		pi.setPath(path)
	}
	return nil
}
