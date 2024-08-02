package server

import "errors"

var ErrNoApplicationJsonHeader error = errors.New("There is no application/json header")
var ErrWhileReadingBody error = errors.New("There is an error while reading the body which should be the JSON to be validated")
var ErrBodyIsTooLong error = errors.New("Please provide a string that is less than 1025 character")
