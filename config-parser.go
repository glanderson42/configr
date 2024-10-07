package configr

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
)

type ConfigParser struct {
	Default   interface{}
	Required  bool
	Type      string
	Env       string
	FieldName string
}

func (c *ConfigParser) String() string {
	return fmt.Sprintf("Default: %v, Required: %v, Type: %s, Env: %s", c.Default, c.Required, c.Type, c.Env)
}

func (c *ConfigParser) Validate() error {
	if c.Required && os.Getenv(c.Env) == "" && c.Default == nil {
		return RequiredFieldError(c.FieldName)
	}
	return nil
}

func (c *ConfigParser) convertValue(envValue string) (interface{}, error) {
	switch c.Type {
	case "int":
		return strconv.Atoi(envValue)
	case "string":
		return envValue, nil
	case "bool":
		return strconv.ParseBool(envValue)
	case "float64":
		return strconv.ParseFloat(envValue, 64)
	case "float32":
		f, err := strconv.ParseFloat(envValue, 32)
		return float32(f), err
	case "uint", "uint8", "uint16", "uint32", "uint64":
		return strconv.ParseUint(envValue, 10, 64)
	case "int8", "int16", "int32", "int64":
		return strconv.ParseInt(envValue, 10, 64)
	default:
		return nil, InvalidTypeError(c.Type)
	}
}

func (c *ConfigParser) LoadValue(value *reflect.Value) error {
	envValue := os.Getenv(c.Env)
	if envValue != "" {
		convertedValue, err := c.convertValue(envValue)
		if err != nil {
			return err
		}
		value.Set(reflect.ValueOf(convertedValue).Convert(value.Type()))
	} else {
		value.Set(reflect.ValueOf(c.Default).Convert(value.Type()))
	}
	return nil
}

func (c *ConfigParser) SetValue(value *reflect.Value) error {
	if err := c.Validate(); err != nil {
		return err
	}
	return c.LoadValue(value)
}
