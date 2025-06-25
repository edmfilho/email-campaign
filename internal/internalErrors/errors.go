package internalerrors

import "errors"

var InternalServerError error = errors.New("internal server error")
