package common

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

// RoundTripDecode exists to assert that `src` can actually be decoded
// to `dst`. This is necessary when `src` is an unknown thing resolved
// via a JSON reference, and we're not sure what the value is.
// This function does a roundtrip to make sure things correctly get
// unmarshaled. Make sure to pass a pointer to the type you want to
// end up with to `dst`
func RoundTripDecode(dst, src interface{}, unmarshaler func([]byte, interface{}) error) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(src); err != nil {
		return errors.Wrap(err, `failed to marshal src into JSON`)
	}

	if err := unmarshaler(buf.Bytes(), dst); err != nil {
		return errors.Wrap(err, `failed to unmarshal JSON`)
	}

	return nil
}

func DumpCode(dst io.Writer, src io.Reader) {
	scanner := bufio.NewScanner(src)
	lineno := 1
	for scanner.Scan() {
		fmt.Fprintf(dst, "%5d: %s\n", lineno, scanner.Text())
		lineno++
	}
}

func WriteToFile(fn string, data []byte) error {
	dir := filepath.Dir(fn)
	if _, err := os.Stat(dir); err != nil {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return errors.Wrapf(err, `failed to create directory %s`, dir)
		}
	}

	f, err := os.OpenFile(fn, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return errors.Wrap(err, `failed to open file for writing`)
	}
	defer f.Close()

	if _, err := f.Write(data); err != nil {
		return errors.Wrap(err, `failed to write to file`)
	}

	return nil
}
