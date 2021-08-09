package goenv

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	EnvTag           string
	DefaultValueTag  string
	RequiredValueTag string
}

type ParserFunc func(value string) (interface{}, error)

var (
	// ErrNoStruct error when given parametetr is not struct
	ErrNoStruct = errors.New("given parameter is not a struct")
	// ErrNoPtr error when given parameter is not pointer
	ErrNoPtr = errors.New("given parameter is not a pointer")
	// DefaultTag tag for loading env key names from struct
	DefaultTag = "env"

	// Default value of env variable
	DefaultValueTag = "default"
	// Required env variable
	DefaultRequiredTag = "required"

	// Built-in parser functions
	BuiltInParsers = map[reflect.Kind]ParserFunc{
		reflect.Bool: func(value string) (interface{}, error) {
			return strconv.ParseBool(value)
		},
		reflect.Int: func(value string) (interface{}, error) {
			return strconv.Atoi(value)
		},
		reflect.Int8: func(value string) (interface{}, error) {
			val, err := strconv.ParseInt(value, 10, 8)
			if err != nil {
				return nil, err
			}

			return int8(val), nil
		},
		reflect.Int16: func(value string) (interface{}, error) {
			val, err := strconv.ParseInt(value, 10, 16)
			if err != nil {
				return nil, err
			}

			return int16(val), nil
		},
		reflect.Int32: func(value string) (interface{}, error) {
			val, err := strconv.ParseInt(value, 10, 32)
			if err != nil {
				return nil, err
			}

			return int32(val), nil
		},
		reflect.Int64: func(value string) (interface{}, error) {
			return strconv.ParseInt(value, 10, 64)
		},
		reflect.Uint: func(value string) (interface{}, error) {
			val, err := strconv.ParseUint(value, 10, 0)
			if err != nil {
				return nil, err
			}

			return uint(val), nil
		},
		reflect.Uint8: func(value string) (interface{}, error) {
			val, err := strconv.ParseUint(value, 10, 8)
			if err != nil {
				return nil, err
			}

			return uint8(val), nil
		},
		reflect.Uint16: func(value string) (interface{}, error) {
			val, err := strconv.ParseUint(value, 10, 16)
			if err != nil {
				return nil, err
			}

			return uint16(val), nil
		},
		reflect.Uint32: func(value string) (interface{}, error) {
			val, err := strconv.ParseUint(value, 10, 32)
			if err != nil {
				return nil, err
			}

			return uint32(val), nil
		},
		reflect.Uint64: func(value string) (interface{}, error) {
			return strconv.ParseUint(value, 10, 64)
		},
		reflect.Float32: func(value string) (interface{}, error) {
			val, err := strconv.ParseFloat(value, 32)
			if err != nil {
				return nil, err
			}

			return float32(val), nil
		},
		reflect.Float64: func(value string) (interface{}, error) {
			return strconv.ParseFloat(value, 64)
		},
		reflect.String: func(value string) (interface{}, error) {
			return value, nil
		},
	}
)

func New(envFiles ...string) (Config, error) {
	if err := godotenv.Load(envFiles...); err != nil {
		return Config{}, err
	}

	return Config{
		EnvTag:           DefaultTag,
		DefaultValueTag:  DefaultValueTag,
		RequiredValueTag: DefaultRequiredTag,
	}, nil
}

func (e Config) Parse(s interface{}) error {
	reflectValue := reflect.ValueOf(s)
	if reflectValue.Kind() != reflect.Ptr || reflectValue.IsNil() {
		return ErrNoPtr
	}

	reflectValue = reflectValue.Elem()

	if reflectValue.Kind() != reflect.Struct {
		return ErrNoStruct
	}

	t := reflectValue.Type()

	for i := 0; i < reflectValue.NumField(); i++ {
		valueField := reflectValue.Field(i)
		if valueField.Kind() == reflect.Struct {
			if !valueField.Addr().CanInterface() {
				continue
			}

			iFace := valueField.Addr().Interface()
			if err := e.Parse(iFace); err != nil {
				return err
			}
		}

		typeField := t.Field(i)
		key, defaultValue := typeField.Tag.Get(e.EnvTag), typeField.Tag.Get(e.DefaultValueTag)
		isRequired := typeField.Tag.Get(e.RequiredValueTag) == "true"

		if key != "" {
			value := e.getOrDefault(key, defaultValue)

			parser := BuiltInParsers[typeField.Type.Kind()]

			parsedValue, err := parser(value)
			if err != nil {
				return err
			}

			if parsedValue == "" && isRequired {
				return fmt.Errorf("%s required", key)
			}

			reflectValue.Field(i).Set(reflect.ValueOf(parsedValue))
		}
	}

	return nil
}

func (e Config) getOrDefault(key string, defaultValue string) string {
	if v, ok := os.LookupEnv(key); ok {
		return v
	}

	return defaultValue
}
