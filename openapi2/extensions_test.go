package openapi2_test

import (
	"testing"

	"github.com/lestrrat-go/openapi/openapi2"
	"github.com/stretchr/testify/assert"
)

func TestExtensions(t *testing.T) {
	contact := apiSupport.Clone()
	err := openapi2.MutateContact(contact).
		Extension(`x-foo`, `Foo`).
		Extension(`x-bar`, `Bar`).
		Apply()
	if !assert.NoError(t, err, `mutating contact should succeed`) {
		return
	}

	for iter := contact.Extensions(); iter.Next(); {
		key, value := iter.Item()
		t.Logf("%#v %#v", key, value)
	}
}
