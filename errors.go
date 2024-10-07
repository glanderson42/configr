package configr

import (
	"fmt"

	"github.com/pkg/errors"
)

func InvalidTypeError(t string) error {
	return errors.New(fmt.Sprintf("Invalid type: %s", t))
}

func RequiredFieldError(field string) error {
	return errors.New(fmt.Sprintf("%s is required", field))
}
