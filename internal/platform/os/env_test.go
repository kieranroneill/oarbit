package os

import "testing"

type TestParams struct {
	name         string
	envValue     string
	defaultValue string
	expected     string
}

func TestGetEnvWithDefault(t *testing.T) {
	tests := []TestParams{
		{
			name:         "returns default when environment variable is unset",
			envValue:     "",
			defaultValue: "default-value",
			expected:     "default-value",
		},
		{
			name:         "returns default when environment variable is empty",
			envValue:     "",
			defaultValue: "fallback",
			expected:     "fallback",
		},
		{
			name:         "returns environment variable value when set",
			envValue:     "configured-value",
			defaultValue: "fallback",
			expected:     "configured-value",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Setenv("TEST_GET_ENV_WITH_DEFAULT", test.envValue)

			expected := GetEnvWithDefault(
				"TEST_GET_ENV_WITH_DEFAULT",
				test.defaultValue,
			)

			if expected != test.expected {
				t.Errorf(
					"expected \"%q\", actual \"%q\"",
					expected,
					test.expected,
				)
			}
		})
	}
}
