package storage_error

import (
	"errors"
)

var NotFound = errors.New("resource not found")
var IOError = errors.New("IO error on resource read/write")
