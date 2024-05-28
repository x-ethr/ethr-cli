package marshalers

import (
	"bytes"
	"encoding/json"
	"gopkg.in/yaml.v3"
)

// JSON is a UTF-8 friendly marshaler.  Go's json.Marshal is not UTF-8
// friendly because it replaces the valid UTF-8 and JSON characters "&". "<",
// ">" with the "slash u" unicode escaped forms (e.g. \u0026).  It preemptively
// escapes for HTML friendliness.  Where text may include any of these
// characters, json.Marshal should not be used.
func JSON(i interface{}) ([]byte, error) {
	var buffer bytes.Buffer

	encoder := json.NewEncoder(&buffer)
	encoder.SetEscapeHTML(false)
	encoder.SetIndent("", "    ")

	e := encoder.Encode(i)

	return bytes.TrimRight(buffer.Bytes(), "\n"), e
}

// YAML is a custom yaml marshaller that sets the encoder's indent
// to 4 spaces. Additionally, newlines are trimmed from the end
// of the bytes buffer.
func YAML(i interface{}) ([]byte, error) {
	var buffer bytes.Buffer

	encoder := yaml.NewEncoder(&buffer)
	encoder.SetIndent(4)

	e := encoder.Encode(i)

	return bytes.TrimRight(buffer.Bytes(), "\n"), e
}
