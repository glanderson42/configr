package configr

import (
	"encoding/json"
	"reflect"
	"strings"
)

func ParseConfig(input interface{}) error {
	t := reflect.TypeOf(input)

	if t.Kind() != reflect.Ptr {
		return InvalidTypeError(t.Kind().String())
	}

	for i := 0; i < t.Elem().NumField(); i++ {
		field := t.Elem().Field(i)
		tag := strings.ReplaceAll(field.Tag.Get("configr"), "'", "\"")

		var parser ConfigParser

		if err := json.Unmarshal([]byte(tag), &parser); err != nil {
			return err
		}

		parser.Type = field.Type.Name()
		parser.FieldName = field.Name

		if tag == "" {
			continue
		}

		v := reflect.ValueOf(input).Elem().FieldByName(field.Name)
		if err := parser.SetValue(&v); err != nil {
			return err
		}
	}

	return nil
}

func LoadConfig() bool {
	return true
}
