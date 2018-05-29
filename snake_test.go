package snake

import (
	"testing"
)

func TestSnakeCase(t *testing.T) {
	tests := map[string]string{
		"camelCase":    "camel_case",
		"camelCase":    "camel_case",
		"snake_case":   "snake_case",
		"snake123Case": "snake_123_case",
		"base":         "base",
		"UpperCase":    "upper_case",
	}

	for input, output := range tests {
		if got := Snake(input); got != output {
			t.Fatalf("Input %v != %v. Got %v", input, output, got)
		}
	}
}
