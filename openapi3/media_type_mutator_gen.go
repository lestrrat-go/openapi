package openapi3

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import "log"

var _ = log.Printf

// MediaTypeMutator is used to build an instance of MediaType. The user must
// call `Do()` after providing all the necessary information to
// the new instance of MediaType with new values
type MediaTypeMutator struct {
	proxy  *mediaType
	target *mediaType
}

// Do finalizes the matuation process for MediaType and returns the result
func (b *MediaTypeMutator) Do() error {
	*b.target = *b.proxy
	return nil
}

// MutateMediaType creates a new mutator object for MediaType
func MutateMediaType(v MediaType) *MediaTypeMutator {
	return &MediaTypeMutator{
		target: v.(*mediaType),
		proxy:  v.Clone().(*mediaType),
	}
}

// Name sets the Name field for object MediaType.
func (b *MediaTypeMutator) Name(v string) *MediaTypeMutator {
	b.proxy.name = v
	return b
}

// Mime sets the Mime field for object MediaType.
func (b *MediaTypeMutator) Mime(v string) *MediaTypeMutator {
	b.proxy.mime = v
	return b
}

// Schema sets the Schema field for object MediaType.
func (b *MediaTypeMutator) Schema(v Schema) *MediaTypeMutator {
	b.proxy.schema = v
	return b
}

func (b *MediaTypeMutator) ClearExamples() *MediaTypeMutator {
	b.proxy.examples.Clear()
	return b
}

func (b *MediaTypeMutator) Example(key ExampleMapKey, value Example) *MediaTypeMutator {
	if b.proxy.examples == nil {
		b.proxy.examples = ExampleMap{}
	}

	b.proxy.examples[key] = value
	return b
}

func (b *MediaTypeMutator) ClearEncoding() *MediaTypeMutator {
	b.proxy.encoding.Clear()
	return b
}

func (b *MediaTypeMutator) Encoding(key EncodingMapKey, value Encoding) *MediaTypeMutator {
	if b.proxy.encoding == nil {
		b.proxy.encoding = EncodingMap{}
	}

	b.proxy.encoding[key] = value
	return b
}
