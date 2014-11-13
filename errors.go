package properties

import (
	"errors"
)

var (
	_NON_EXISTS_ = errors.New("Not found")
	_NULL_KEY_ = errors.New("Key is null")
)
