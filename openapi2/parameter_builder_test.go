package openapi2_test

import (
	"testing"

	"github.com/lestrrat-go/openapi/openapi2"
	"github.com/stretchr/testify/assert"
)

func TestParameterValidate(t *testing.T) {
	t.Run("invalid name", func(t *testing.T) {
		_, err := openapi2.NewParameter("", openapi2.InQuery).Build()
		if !assert.Error(t, err, "empty name should result in error") {
			return
		}
	})
	t.Run("invalid location", func(t *testing.T) {
		_, err := openapi2.NewParameter("foo", openapi2.Location("")).Build()
		if !assert.Error(t, err, "empty location should result in error") {
			return
		}
	})
	t.Run("location = InBody + nil schema", func(t *testing.T) {
		_, err := openapi2.NewParameter("foo", openapi2.InBody).Build()
		if !assert.Error(t, err, "empty schema when location = body should result in error") {
			return
		}
	})
	t.Run("allowEmptyValue with InQuery or InForm", func(t *testing.T) {
		for _, loc := range []openapi2.Location{openapi2.InQuery, openapi2.InForm} {
			_, err := openapi2.NewParameter("foo", loc).
				AllowEmptyValue(true).
				Build()
			if !assert.Error(t, err, "allowEmptyValue = true when location = %s should result in error", loc) {
				return
			}
		}
	})
	t.Run("type = array and items = nil", func(t *testing.T) {
		_, err := openapi2.NewParameter("foo", openapi2.InQuery).
			Type(openapi2.Array).
			Build()
		if !assert.Error(t, err, "empty item when type = array should result in error") {
			return
		}
	})
	t.Run("type = file and in != formData", func(t *testing.T) {
		for _, loc := range []openapi2.Location{openapi2.InPath, openapi2.InQuery, openapi2.InHeader, openapi2.InBody} {
			_, err := openapi2.NewParameter("foo", loc).
				Type(openapi2.File).
				Build()
			if !assert.Error(t, err, "loc != formData (%s) when type = file should result in error", loc) {
				return
			}
		}
	})
	t.Run("type must be in string, number, integer, boolean, array, file", func(t *testing.T) {
		for _, typ := range []openapi2.PrimitiveType{openapi2.Object, openapi2.Null} {
			_, err := openapi2.NewParameter("foo", openapi2.InQuery).
				Type(typ).
				Build()
			if !assert.Error(t, err, "type = %s should result in error", typ) {
				return
			}
		}
	})
}
