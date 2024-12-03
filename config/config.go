package config

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
)

type Config struct {
	Port int `env:"port"`
}

// NewConfig reads configuration from environment variables and uses defaults
func NewConfig() (*Config, error) {
	cfg := &Config{}

	// Read environment variables with support for defaults
	t := reflect.TypeOf(cfg).Elem()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		envTag := field.Tag.Get("env")

		value, ok := os.LookupEnv(envTag)
		if !ok {
			t := field.Tag.Get("default")
			if t == "" {
				return nil, fmt.Errorf("missing required env variable '%s'", envTag)
			}
			value = t
		}
		if err := setField(cfg, field.Name, value); err != nil {
			return nil, fmt.Errorf("error parsing environment variable '%s': %w", field.Name, err)
		}
	}

	return cfg, nil
}

// setField uses reflection to set the value of a struct field
func setField(item any, name string, val string) error {
	s := reflect.ValueOf(item).Elem()
	f := s.FieldByName(name)
	if !f.IsValid() {
		return fmt.Errorf("field '%s' not found in config struct", name)
	}
	if !f.CanSet() {
		return fmt.Errorf("field '%s' is not settable", name)
	}

	switch f.Kind() {
	case reflect.String:
		f.SetString(val)
	case reflect.Bool:
		b, err := strconv.ParseBool(val)
		if err != nil {
			return fmt.Errorf("error parsing value for bool field '%s': %w", name, err)
		}
		f.SetBool(b)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		i, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return fmt.Errorf("error parsing value for int field '%s': %w", name, err)
		}
		f.SetInt(i)
	// Add more cases for other data types as needed
	default:
		return fmt.Errorf("unsupported type for field '%s': %v", name, f.Kind())
	}
	return nil
}
