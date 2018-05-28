package openapi_test

import (
	"testing"

	openapi "github.com/lestrrat-go/openapi/v2"
	"github.com/stretchr/testify/assert"
)

func TestParameterValidate(t *testing.T) {
	t.Run("invalid name", func(t *testing.T) {
		_, err := openapi.NewParameter("", openapi.InQuery).Do()
		if !assert.Error(t, err, "empty name should result in error") {
			return
		}
	})
	t.Run("invalid location", func(t *testing.T) {
		_, err := openapi.NewParameter("foo", openapi.Location("")).Do()
		if !assert.Error(t, err, "empty location should result in error") {
			return
		}
	})
	t.Run("location = InBody + nil schema", func(t *testing.T) {
		_, err := openapi.NewParameter("foo", openapi.InBody).Do()
		if !assert.Error(t, err, "empty schema when location = body should result in error") {
			return
		}
	})
	t.Run("allowEmptyValue with InQuery or InForm", func(t *testing.T) {
		for _, loc := range []openapi.Location{openapi.InQuery, openapi.InForm} {
			_, err := openapi.NewParameter("foo", loc).
				AllowEmptyValue(true).
				Do()
			if !assert.Error(t, err, "allowEmptyValue = true when location = %s should result in error", loc) {
				return
			}
		}
	})
	t.Run("type = array and items = nil", func(t *testing.T) {
		_, err := openapi.NewParameter("foo", openapi.InQuery).
			Type(openapi.Array).
			Do()
		if !assert.Error(t, err, "empty item when type = array should result in error") {
			return
		}
	})
	t.Run("type = file and in != formData", func(t *testing.T) {
		for _, loc := range []openapi.Location{openapi.InPath, openapi.InQuery, openapi.InHeader, openapi.InBody} {
			_, err := openapi.NewParameter("foo", loc).
				Type(openapi.File).
				Do()
			if !assert.Error(t, err, "loc != formData (%s) when type = file should result in error", loc) {
				return
			}
		}
	})
	t.Run("type must be in string, number, integer, boolean, array, file", func(t *testing.T) {
		for _, typ := range []openapi.PrimitiveType{openapi.Object, openapi.Null} {
			_, err := openapi.NewParameter("foo", openapi.InQuery).
				Type(typ).
				Do()
			if !assert.Error(t, err, "type = %s should result in error", typ) {
				return
			}
		}
	})
}
