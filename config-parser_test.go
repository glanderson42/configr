package configr

import (
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestConfigParser_SetValue(t *testing.T) {
	tests := []struct {
		name           string
		parser         ConfigParser
		envValue       string
		expected       interface{}
		expectError    bool
		expectedErrMsg string
	}{
		{
			name: "int type with env variable",
			parser: ConfigParser{
				Default:   100,
				Required:  true,
				Type:      "int",
				Env:       "TEST_INT",
				FieldName: "TestInt",
			},
			envValue:    "200",
			expected:    200,
			expectError: false,
		},
		{
			name: "string type with default",
			parser: ConfigParser{
				Default:   "defaultValue",
				Required:  false,
				Type:      "string",
				Env:       "TEST_STRING",
				FieldName: "TestString",
			},
			envValue:    "",
			expected:    "defaultValue",
			expectError: false,
		},
		{
			name: "bool type with env variable",
			parser: ConfigParser{
				Default:   false,
				Required:  true,
				Type:      "bool",
				Env:       "TEST_BOOL",
				FieldName: "TestBool",
			},
			envValue:    "true",
			expected:    true,
			expectError: false,
		},
		{
			name: "missing required field",
			parser: ConfigParser{
				Default:   nil,
				Required:  true,
				Type:      "string",
				Env:       "TEST_MISSING",
				FieldName: "TestMissing",
			},
			envValue:       "",
			expected:       nil,
			expectError:    true,
			expectedErrMsg: "TestMissing is required",
		},
		{
			name: "invalid float conversion",
			parser: ConfigParser{
				Default:   1.23,
				Required:  true,
				Type:      "float64",
				Env:       "TEST_INVALID_FLOAT",
				FieldName: "TestInvalidFloat",
			},
			envValue:       "not_a_float",
			expected:       nil,
			expectError:    true,
			expectedErrMsg: "invalid syntax",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set environment variable for the test case
			if tt.envValue != "" {
				os.Setenv(tt.parser.Env, tt.envValue)
				defer os.Unsetenv(tt.parser.Env) // Cleanup after test
			}

			// Handle the case where expected is nil
			var value reflect.Value
			if tt.expected != nil {
				// Create a reflect.Value of the expected type
				value = reflect.New(reflect.TypeOf(tt.expected)).Elem()
			} else {
				// Assign a dummy value for nil expected cases to avoid reflect.New(nil)
				value = reflect.New(reflect.TypeOf("")).Elem()
			}

			err := tt.parser.SetValue(&value)
			if (err != nil) != tt.expectError {
				t.Errorf("SetValue() error = %v, wantErr %v", err, tt.expectError)
				return
			}

			if err != nil && tt.expectError && !strings.Contains(err.Error(), tt.expectedErrMsg) {
				t.Errorf("SetValue() error = %v, expectedErrMsg %v", err.Error(), tt.expectedErrMsg)
			}

			// Compare expected and actual values only if expected is not nil
			if !tt.expectError && tt.expected != nil && value.Interface() != tt.expected {
				t.Errorf("SetValue() = %v, expected %v", value.Interface(), tt.expected)
			}
		})
	}
}
