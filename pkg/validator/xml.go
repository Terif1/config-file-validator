package validator

import (
	"bytes"
	"encoding/xml"
	"io"
)

type XmlValidator struct{}

// Validate implements the Validator interface by attempting to
// unmarshall a byte array of xml
func (xv XmlValidator) Validate(b []byte) (bool, error) {
	var output interface{}
	d := xml.NewDecoder(bytes.NewReader(b))
	// pass an idempotent charset reader
	// avoids errors like
	// xml: encoding "utf-16" declared but Decoder.CharsetReader is nil
	d.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		return input, nil
	}
	err := d.Decode(&output)
	if err != nil {
		return false, err
	}
	return true, nil
}
