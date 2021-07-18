package util

import "errors"

func ErrMessage(text string) error {
	return errors.New(text)
}
