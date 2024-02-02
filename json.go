package easclient

import "github.com/go-json-experiment/json"

// unmarshalJSON is a wrapper around the json library we want to use for unmarshaling.
//
// Since the std library does not handle our edge cases, a specialized library is used.
// Have a look at the tests for details.
func unmarshalJSON(data []byte, v any, opts ...json.Options) error {
	return json.Unmarshal(data, v, opts...)
}
