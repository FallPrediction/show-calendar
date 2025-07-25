package errors

import "errors"

var (
	ErrPasswordIncorrect  = errors.New("the provided password is incorrect")
	ErrInvalidToken       = errors.New("the provided token is invalid")
	ErrInvalidContentType = errors.New("the content type is invalid")
)
