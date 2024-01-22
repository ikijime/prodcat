package utils

import (
	"errors"
	"unicode"
	"unicode/utf8"

	"github.com/go-playground/validator/v10"
)

func Capitalize(text string) string {
	r, size := utf8.DecodeRuneInString(text)
	if r == utf8.RuneError {
		return text
	}
	return string(unicode.ToUpper(r)) + text[size:]
}

func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func MsgForTag(fe validator.FieldError) error {
	switch fe.Tag() {
	case "required":
		return errors.New(fe.Field() + " is required")
	case "min":
		return errors.New(fe.Field() + " minimum is " + fe.Param())
	case "email":
		return errors.New("Email required")
	case "type":
		return errors.New("Type required")
	}
	return errors.New("not valid") // default error
}
