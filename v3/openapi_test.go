package openapi_test

/*

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/lestrrat-go/openapi/v3/entity"
	"github.com/stretchr/testify/assert"
	yaml "gopkg.in/yaml.v2"
)

func TestOpenAPI(t *testing.T) {
	root := filepath.Join("..", "spec", "examples", "v3.0")
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		data, err := ioutil.ReadFile(path)
		if !assert.NoError(t, err, `reading from %s should succeed`, path) {
			return err
		}
		//		t.Logf("%s", data)
		buf := bytes.NewBuffer(data)
		rdr := bytes.NewReader(buf.Bytes())

		t.Run(path, func(t *testing.T) {
			var v entity.OpenAPI
			if !assert.NoError(t, yaml.NewDecoder(rdr).Decode(&v), `Decoding spec should succeed`) {
				rdr.Seek(0, 0)
				scanner := bufio.NewScanner(rdr)
				lineno := 1
				for scanner.Scan() {
					t.Logf("%4d: %s", lineno, scanner.Text())
					lineno++
				}
				return
			}
		})

		return nil
	})
}
*/
