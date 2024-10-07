package configr

import (
	"errors"
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
		return errors.New(c.FieldName + " is required")
	}

	return nil
}

func (c *ConfigParser) LoadValue(value *reflect.Value) error {
	switch c.Type {
	case "int":
		if os.Getenv(c.Env) != "" {
			val, err := strconv.Atoi(os.Getenv(c.Env))
			if err != nil {
				return err
			}
			value.SetInt(int64(val))
			break
		}

		value.SetInt(int64(c.Default.(float64)))
	case "string":
		if os.Getenv(c.Env) != "" {
			value.SetString(os.Getenv(c.Env))
			break
		}

		value.SetString(c.Default.(string))
	case "bool":
		if os.Getenv(c.Env) != "" {
			val, err := strconv.ParseBool(os.Getenv(c.Env))
			if err != nil {
				return err
			}
			value.SetBool(val)
			break
		}

		value.SetBool(c.Default.(bool))
	case "float64":
		if os.Getenv(c.Env) != "" {
			val, err := strconv.ParseFloat(os.Getenv(c.Env), 64)
			if err != nil {
				return err
			}
			value.SetFloat(val)
			break
		}

		value.SetFloat(c.Default.(float64))
	case "float32":
		if os.Getenv(c.Env) != "" {
			val, err := strconv.ParseFloat(os.Getenv(c.Env), 32)
			if err != nil {
				return err
			}
			value.SetFloat(val)
			break
		}

		value.SetFloat(float64(c.Default.(float32)))
	case "uint":
		if os.Getenv(c.Env) != "" {
			val, err := strconv.ParseUint(os.Getenv(c.Env), 10, 64)
			if err != nil {
				return err
			}
			value.SetUint(val)
			break
		}

		value.SetUint(uint64(c.Default.(float64)))
	case "uint8":
		if os.Getenv(c.Env) != "" {
			val, err := strconv.ParseUint(os.Getenv(c.Env), 10, 8)
			if err != nil {
				return err
			}
			value.SetUint(val)
			break
		}

		value.SetUint(uint64(c.Default.(float64)))
	case "uint16":
		if os.Getenv(c.Env) != "" {
			val, err := strconv.ParseUint(os.Getenv(c.Env), 10, 16)
			if err != nil {
				return err
			}
			value.SetUint(val)
			break
		}

		value.SetUint(uint64(c.Default.(float64)))
	case "uint32":
		if os.Getenv(c.Env) != "" {
			val, err := strconv.ParseUint(os.Getenv(c.Env), 10, 32)
			if err != nil {
				return err
			}
			value.SetUint(val)
			break
		}

		value.SetUint(uint64(c.Default.(float64)))
	case "uint64":
		if os.Getenv(c.Env) != "" {
			val, err := strconv.ParseUint(os.Getenv(c.Env), 10, 64)
			if err != nil {
				return err
			}
			value.SetUint(val)
			break
		}

		value.SetUint(uint64(c.Default.(float64)))
	case "int8":
		if os.Getenv(c.Env) != "" {
			val, err := strconv.ParseInt(os.Getenv(c.Env), 10, 8)
			if err != nil {
				return err
			}
			value.SetInt(val)
			break
		}

		value.SetInt(int64(c.Default.(float64)))
	case "int16":
		if os.Getenv(c.Env) != "" {
			val, err := strconv.ParseInt(os.Getenv(c.Env), 10, 16)
			if err != nil {
				return err
			}
			value.SetInt(val)
			break
		}

		value.SetInt(int64(c.Default.(float64)))
	case "int32":
		if os.Getenv(c.Env) != "" {
			val, err := strconv.ParseInt(os.Getenv(c.Env), 10, 32)
			if err != nil {
				return err
			}
			value.SetInt(val)
			break
		}

		value.SetInt(int64(c.Default.(float64)))
	case "int64":
		if os.Getenv(c.Env) != "" {
			val, err := strconv.ParseInt(os.Getenv(c.Env), 10, 64)
			if err != nil {
				return err
			}
			value.SetInt(val)
			break
		}

		value.SetInt(int64(c.Default.(float64)))
	default:
		return InvalidTypeError(c.Type)
	}

	return nil
}

func (c *ConfigParser) SetValue(value *reflect.Value) error {
	if c.Validate() != nil {
		return c.Validate()
	}
	return c.LoadValue(value)
}
