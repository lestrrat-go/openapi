package entity_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/ghodss/yaml"
	"github.com/lestrrat-go/openapi/entity"
	"github.com/stretchr/testify/assert"
)

const (
	jsonKey = "json"
	yamlKey = "yaml"
)

type marshalFunc func(interface{}) ([]byte, error)
type unmarshalFunc func([]byte, interface{}) error
type postSerializer interface {
	Execute(*testing.T, []byte) bool
}
type postDeserializer interface {
	Execute(*testing.T, interface{}) bool
}

type EntityTestCase struct {
	target          interface{}
	postSerialize   map[string]postSerializer
	postDeserialize postDeserializer
}

func NewEntityTestCase(v interface{}) *EntityTestCase {
	return &EntityTestCase{
		target:        v,
		postSerialize: make(map[string]postSerializer),
	}
}

func (tc *EntityTestCase) PostSerialize(key string, check postSerializer) {
	tc.postSerialize[key] = check
}

type PostSerializeMatchString string

func (s PostSerializeMatchString) Execute(t *testing.T, buf []byte) bool {
	return assert.Equal(t, string(s), string(buf), `Serialized form should match`)
}

type PostDeserializeMatchValue struct {
	Target interface{}
}

func (p PostDeserializeMatchValue) Execute(t *testing.T, in interface{}) bool {
	return assert.Equal(t, p.Target, in)
}

func (tc *EntityTestCase) Run(t *testing.T) {
	t.Helper()

	rv := reflect.ValueOf(tc.target)
	if !assert.True(t, rv.Kind() == reflect.Ptr, `second argument to roundtrip() must be a pointer`) {
		return
	}

	dst := reflect.New(rv.Elem().Type()).Interface()

	serializers := []struct {
		Name      string
		Marshal   marshalFunc
		Unmarshal unmarshalFunc
	}{
		{Name: yamlKey, Marshal: yaml.Marshal, Unmarshal: yaml.Unmarshal},
		{Name: jsonKey, Marshal: func(v interface{}) ([]byte, error) { return json.MarshalIndent(v, "", "  ") }, Unmarshal: json.Unmarshal},
	}

	for _, s := range serializers {
		mfn := s.Marshal
		ufn := s.Unmarshal
		sname := s.Name
		t.Run(s.Name, func(t *testing.T) {
			serialized, err := mfn(tc.target)
			t.Run("Serialize data", func(t *testing.T) {
				if !assert.NoError(t, err, `Marshal should succeed`) {
					return
				}
			})

			t.Logf("serialized format:\n%s", serialized) // newline is added so we get better indentation
			if check, ok := tc.postSerialize[sname]; ok {
				t.Run("Post serialize", func(t *testing.T) {
					if !check.Execute(t, serialized) {
						return
					}
				})
			}

			t.Run("Deserialize data", func(t *testing.T) {
				if !assert.NoError(t, ufn(serialized, dst), `Unmarshal should succeed`) {
					return
				}
			})

			check := tc.postDeserialize
			if check == nil {
				check = PostDeserializeMatchValue{Target: tc.target}
			}
			t.Run("Post Deserialize", func(t *testing.T) {
				if !check.Execute(t, dst) {
					return
				}
			})
		})
	}
}

func TestContact(t *testing.T) {
	var src entity.Contact
	src.Name = "API Support"
	src.URL = "http://www.example.com/support"
	src.Email = "support@example.com"

	tc := NewEntityTestCase(&src)
	tc.Run(t)
}

func TestLicense(t *testing.T) {
	var src entity.License
	src.Name = "Apache 2.0"
	src.URL = "https://www.apache.org/licenses/LICENSE-2.0.html"

	tc := NewEntityTestCase(&src)
	tc.PostSerialize(jsonKey, PostSerializeMatchString("{\n  \"name\": \"Apache 2.0\",\n  \"url\": \"https://www.apache.org/licenses/LICENSE-2.0.html\"\n}"))
	tc.PostSerialize(yamlKey, PostSerializeMatchString("name: Apache 2.0\nurl: https://www.apache.org/licenses/LICENSE-2.0.html\n"))
	tc.Run(t)
}

func TestInfo(t *testing.T) {
	var src entity.Info
	src.Title = "Sample Pet Store App"
	src.Description = "This is a sample server for a pet store."
	src.TermsOfService = "http://example.com/terms/"
	src.Version = "1.0.1"
	src.Contact = &entity.Contact{
		Name:  "API Support",
		URL:   "http://www.example.com/support",
		Email: "support@example.com",
	}
	src.License = &entity.License{
		Name: "Apache 2.0",
		URL:  "https://www.apache.org/licenses/LICENSE-2.0.html",
	}

	tc := NewEntityTestCase(&src)
	tc.Run(t)
}
