package customerrors

import (
	"errors"
)

var (
	ErrToLoadConfiguration = errors.New("Error to load config.json")
	ErrToReadConfigurayion = errors.New("Error to read config.json")
)
