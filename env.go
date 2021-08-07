package goenv

import (
	"errors"
	"os"
	"reflect"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	EnvTag          string
	DefaultValueTag string
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
		EnvTag:          DefaultTag,
		DefaultValueTag: DefaultValueTag,
	}, nil
}

func (e Config) Parse(s interface{}) error {
	if err := e.isStructPointer(s); err != nil {
		return err
	}

	v := reflect.ValueOf(s).Elem()
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		sf := t.Field(i)

		if err := e.isStructPointer(sf); err == nil {
			if err := e.Parse(sf); err != nil {
				return err
			}
		}

		key, defaultValue := sf.Tag.Get(e.EnvTag), sf.Tag.Get(e.DefaultValueTag)
		if key != "" {
			value := e.getOrDefault(key, defaultValue)

			parser := BuiltInParsers[sf.Type.Kind()]

			parsedValue, err := parser(value)
			if err != nil {
				return err
			}

			v.Field(i).Set(reflect.ValueOf(parsedValue))
		}
	}

	return nil
}

func (e Config) isStructPointer(s interface{}) error {
	v := reflect.ValueOf(s)
	if v.Kind() != reflect.Ptr {
		return ErrNoPtr
	}

	v = v.Elem()

	if v.Kind() != reflect.Struct {
		return ErrNoStruct
	}

	return nil
}

func (e Config) getOrDefault(key string, defaultValue string) string {
	if v, ok := os.LookupEnv(key); ok {
		return v
	}

	return defaultValue
}
