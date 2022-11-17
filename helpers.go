package goscale

import (
	"reflect"
	"testing"
)

func assertEqual(t *testing.T, result interface{}, expectation interface{}) {
	if reflect.DeepEqual(result, expectation) {
		switch result := result.(type) {
		case []byte:
			// t.Logf("OK( %08b )\n", result)
			t.Logf("\n\n\b\bOK( %#x )\n\n", result)
		default:
			t.Logf("\n\n\b\bOK( %v )\n\n", result)
		}
		t.Log("\n")
		return
	}
	// t.Errorf("WRONG(\nReceived: %08b (type %v)\n\n", result, reflect.TypeOf(result))
	t.Errorf("\n\n\b\bWRONG(\nReceived: %#x (Type: %v)\nExpected: %#x (Type: %v)\n\b\b)\n\n",
		result, reflect.TypeOf(result),
		expectation, reflect.TypeOf(expectation),
	)
}
