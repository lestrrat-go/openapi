package openapi

import "github.com/pkg/errors"

func (v *parameter) ConvertToSchema() (Schema, error) {
	if v.in == InBody {
		return v.schema.Clone(), nil
	}

	b := NewSchema().
		Type(v.Type()).
		Format(v.Format()).
		Default(v.Default()).
		Pattern(v.Pattern()).
		Reference(v.Reference()).
		UniqueItems(v.UniqueItems())

	if items := v.Items(); items != nil {
		s, err := items.ConvertToSchema()
		if err != nil {
			return nil, errors.Wrap(err, `failed to convert items to schema`)
		}
		b.Items(s)
	}

	if v.HasMaximum() {
		b.Maximum(v.Maximum())
	}

	if v.HasExclusiveMaximum() {
		b.ExclusiveMaximum(v.ExclusiveMaximum())
	}

	if v.HasMinimum() {
		b.Minimum(v.Minimum())
	}

	if v.HasExclusiveMinimum() {
		b.ExclusiveMinimum(v.ExclusiveMinimum())
	}

	if v.HasMaxLength() {
		b.MaxLength(v.MaxLength())
	}

	if v.HasMinLength() {
		b.MinLength(v.MinLength())
	}

	if v.HasMaxItems() {
		b.MaxItems(v.MaxItems())
	}

	if v.HasMinLength() {
		b.MinLength(v.MinLength())
	}

	if v.HasMinItems() {
		b.MinItems(v.MinItems())
	}

	if v.HasMultipleOf() {
		b.MultipleOf(v.MultipleOf())
	}

	for iter := v.Enum(); iter.Next(); {
		b.Enum(iter.Item())
	}

	for iter := v.Extensions(); iter.Next(); {
		b.Extension(iter.Item())
	}
	return b.Do()
}
