package goenv

import (
	"reflect"
	"testing"
)

type parserTestCase struct {
	name          string
	stringValue   string
	expectedValue interface{}
	kind          reflect.Kind
	isValid       bool
}

func testBuiltinParsers(t *testing.T, testCases []parserTestCase) {
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			parser := BuiltInParsers[testCase.kind]

			v, err := parser(testCase.stringValue)

			if testCase.isValid && err != nil {
				t.Errorf("parsing error: %v\n", err)
			}

			if testCase.isValid {
				if v != testCase.expectedValue {
					t.Errorf("expected value: %v, returned value: %v\n", testCase.expectedValue, v)
				}
			} else {
				if v == testCase.expectedValue {
					t.Errorf("parsed value equals in invalid case\n")
				}
			}
		})
	}
}

func TestBoolBuiltinParsers(t *testing.T) {
	testCases := []parserTestCase{
		{
			name:          "test parse bool",
			stringValue:   "true",
			expectedValue: true,
			kind:          reflect.Bool,
			isValid:       true,
		},
		{
			name:          "test parse bool with unequal values",
			stringValue:   "false",
			expectedValue: true,
			kind:          reflect.Bool,
			isValid:       false,
		},
	}

	testBuiltinParsers(t, testCases)
}

func TestIntBuiltinParsers(t *testing.T) {
	testCases := []parserTestCase{
		{
			name:          "test parse int",
			stringValue:   "12",
			expectedValue: int(12),
			kind:          reflect.Int,
			isValid:       true,
		},
		{
			name:          "test parse int with unequal values",
			stringValue:   "13",
			expectedValue: int(12),
			kind:          reflect.Int,
			isValid:       false,
		},
		{
			name:          "test parse int with invalid string value",
			stringValue:   "12.3",
			expectedValue: int(12),
			kind:          reflect.Int,
			isValid:       false,
		},
		{
			name:          "test parse int8",
			stringValue:   "12",
			expectedValue: int8(12),
			kind:          reflect.Int8,
			isValid:       true,
		},
		{
			name:          "test parse int8 with unequal values",
			stringValue:   "13",
			expectedValue: int8(12),
			kind:          reflect.Int8,
			isValid:       false,
		},
		{
			name:          "test parse int8 with invalid string value",
			stringValue:   "12.9",
			expectedValue: int8(12),
			kind:          reflect.Int8,
			isValid:       false,
		},
		{
			name:          "test parse int16",
			stringValue:   "12",
			expectedValue: int16(12),
			kind:          reflect.Int16,
			isValid:       true,
		},
		{
			name:          "test parse int16 with unequal values",
			stringValue:   "13",
			expectedValue: int16(12),
			kind:          reflect.Int16,
			isValid:       false,
		},
		{
			name:          "test parse int16 with invalid string value",
			stringValue:   "12.1",
			expectedValue: int16(12),
			kind:          reflect.Int16,
			isValid:       false,
		},
		{
			name:          "test parse int32",
			stringValue:   "12",
			expectedValue: int32(12),
			kind:          reflect.Int32,
			isValid:       true,
		},
		{
			name:          "test parse int32 with unequal values",
			stringValue:   "13",
			expectedValue: int32(12),
			kind:          reflect.Int32,
			isValid:       false,
		},
		{
			name:          "test parse int32 with invalid string value",
			stringValue:   "abc123",
			expectedValue: int32(12),
			kind:          reflect.Int32,
			isValid:       false,
		},
		{
			name:          "test parse int64",
			stringValue:   "12",
			expectedValue: int64(12),
			kind:          reflect.Int64,
			isValid:       true,
		},
		{
			name:          "test parse int64 with unequal values",
			stringValue:   "13",
			expectedValue: int64(12),
			kind:          reflect.Int64,
			isValid:       false,
		},
		{
			name:          "test parse int64 with invalid string value",
			stringValue:   "abc123",
			expectedValue: int64(12),
			kind:          reflect.Int64,
			isValid:       false,
		},
	}

	testBuiltinParsers(t, testCases)
}

func TestUintBuiltinParsers(t *testing.T) {
	testCases := []parserTestCase{
		{
			name:          "test parse uint",
			stringValue:   "12",
			expectedValue: uint(12),
			kind:          reflect.Uint,
			isValid:       true,
		},
		{
			name:          "test parse uint with unequal values",
			stringValue:   "13",
			expectedValue: uint(12),
			kind:          reflect.Uint,
			isValid:       false,
		},
		{
			name:          "test parse uint with invalid string value",
			stringValue:   "abc",
			expectedValue: uint(12),
			kind:          reflect.Uint,
			isValid:       false,
		},
		{
			name:          "test parse uint8",
			stringValue:   "12",
			expectedValue: uint8(12),
			kind:          reflect.Uint8,
			isValid:       true,
		},
		{
			name:          "test parse uint8 with unequal values",
			stringValue:   "13",
			expectedValue: uint8(12),
			kind:          reflect.Uint8,
			isValid:       false,
		},
		{
			name:          "test parse uint8 with invalid string value",
			stringValue:   "abc123",
			expectedValue: uint8(12),
			kind:          reflect.Uint8,
			isValid:       false,
		},
		{
			name:          "test parse uint16",
			stringValue:   "12",
			expectedValue: uint16(12),
			kind:          reflect.Uint16,
			isValid:       true,
		},
		{
			name:          "test parse uint16 with invalid string value",
			stringValue:   "abc",
			expectedValue: uint16(12),
			kind:          reflect.Uint16,
			isValid:       false,
		},
		{
			name:          "test parse uint32",
			stringValue:   "12",
			expectedValue: uint32(12),
			kind:          reflect.Uint32,
			isValid:       true,
		},
		{
			name:          "test parse uint32 with unequal values",
			stringValue:   "13",
			expectedValue: uint32(12),
			kind:          reflect.Uint32,
			isValid:       false,
		},
		{
			name:          "test parse uint32 with invalid string value",
			stringValue:   "abc",
			expectedValue: uint32(12),
			kind:          reflect.Uint32,
			isValid:       false,
		},
		{
			name:          "test parse uint64",
			stringValue:   "12",
			expectedValue: uint64(12),
			kind:          reflect.Uint64,
			isValid:       true,
		},
		{
			name:          "test parse uint64 with unequal values",
			stringValue:   "13",
			expectedValue: uint64(12),
			kind:          reflect.Uint64,
			isValid:       false,
		},
		{
			name:          "test parse uint64 with invalid string value",
			stringValue:   "abc",
			expectedValue: uint64(12),
			kind:          reflect.Uint64,
			isValid:       false,
		},
	}

	testBuiltinParsers(t, testCases)
}

func TestFloatBuiltinParsers(t *testing.T) {
	testCases := []parserTestCase{
		{
			name:          "test parse float32",
			stringValue:   "12.23",
			expectedValue: float32(12.23),
			kind:          reflect.Float32,
			isValid:       true,
		},
		{
			name:          "test parse float32 with unequal values",
			stringValue:   "13.12",
			expectedValue: float32(12.12),
			kind:          reflect.Float32,
			isValid:       false,
		},
		{
			name:          "test parse float32",
			stringValue:   "abc123",
			expectedValue: float32(12.23),
			kind:          reflect.Float32,
			isValid:       false,
		},
		{
			name:          "test parse float64",
			stringValue:   "12.23",
			expectedValue: float64(12.23),
			kind:          reflect.Float64,
			isValid:       true,
		},
		{
			name:          "test parse float64 with unequal values",
			stringValue:   "13.12",
			expectedValue: float64(12.12),
			kind:          reflect.Float64,
			isValid:       false,
		},
		{
			name:          "test parse float64",
			stringValue:   "abc123",
			expectedValue: float64(12.23),
			kind:          reflect.Float64,
			isValid:       false,
		},
	}

	testBuiltinParsers(t, testCases)
}

func TestStringBuiltinParsers(t *testing.T) {
	testCases := []parserTestCase{
		{
			name:          "test parse string",
			stringValue:   "test123",
			expectedValue: "test123",
			kind:          reflect.String,
			isValid:       true,
		},
		{
			name:          "test parse string with unequal values",
			stringValue:   "teststring",
			expectedValue: "test123",
			kind:          reflect.String,
			isValid:       false,
		},
	}

	testBuiltinParsers(t, testCases)
}

func TestNew(t *testing.T) {
	_, err := New()
	if err != nil {
		t.Errorf("error loading env variables: %v\n", err)
	}

	if _, err := New("dev.env"); err == nil {
		t.Errorf("error loading env variables")
	}
}

func TestParse(t *testing.T) {
	TestNew(t)

	env, _ := New()

	conf := "config"

	if err := env.Parse(conf); err == nil {
		t.Errorf("error parsing env to struct without pointer\n")
	}

	if err := env.Parse(&conf); err == nil {
		t.Errorf("error parsing env to struct\n")
	}

	// with invalid value
	// value in the .env file
	// INVALID_VALUE=true
	cfg := struct {
		InvalidValue int64 `env:"INVALID_VALUE" default:"1"`
	}{}

	if err := env.Parse(&cfg); err == nil {
		t.Errorf("error parsing env to struct: %v\n", err)
	}

	// values in the .env file
	// DEBUG=false
	// DB_NAME not set
	// PORT=8081
	// FEE_PERCENT=3.3
	// INVALID_VALUE=true
	config := struct {
		Debug      bool    `env:"DEBUG" default:"true"`
		DBName     string  `env:"DB_NAME" default:"postgres"`
		Port       int64   `env:"PORT" default:"8080"`
		FeePercent float32 `env:"FEE_PERCENT" default:"1"`
	}{}

	if err := env.Parse(&config); err != nil {
		t.Errorf("error parsing env to struct: %v\n", err)
	}

	if config.Debug != false {
		t.Errorf("error parsing bool value\n")
	}

	if config.DBName != "postgres" {
		t.Errorf("error parsing with default value\n")
	}

	if config.Port != 8081 {
		t.Errorf("error parsing int64\n")
	}

	if config.FeePercent != float32(3.3) {
		t.Errorf("error parsing float32\n")
	}
}
