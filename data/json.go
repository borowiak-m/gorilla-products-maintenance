package data

import (
	"encoding/json"
	"io"
)

// Serializes contents of the collection to JSON
// NewEncoder has better performance than json.Unmarshall
// due to no usage of interim buffer in memory
func ToJSON(obj interface{}, wrt io.Writer) error {
	encoder := json.NewEncoder(wrt)
	return encoder.Encode(obj)
}

// Decoding struct from JSON
func FromJSON(obj interface{}, re io.Reader) error {
	decoder := json.NewDecoder(re)
	return decoder.Decode(obj)
}
