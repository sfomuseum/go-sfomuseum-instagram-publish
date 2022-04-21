package secret

import (
	"testing"
)

func TestDeriveSecret(t *testing.T) {

	tests := map[string]string{
		"7784f12fc1c2315991c8afd2542fbc09": "7bc63e2e17",
	}

	for input, expected := range tests {

		s := DeriveSecret(input)

		if s != expected {
			t.Fatalf("Invalid secret. Expected '%s', got '%s'", expected, s)
		}
	}
}
