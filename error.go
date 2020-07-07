package riot

import "errors"

func NewError(code int, message string) (err error) {
	return errors.New(message)
}
