package envvar

import (
	"fmt"
	"strings"
)

// UndefinedError is an error that occurs when a required environment variable
// has not been defined.
type UndefinedError struct {
	Name    string
	Allowed []string
}

func (e *UndefinedError) Error() string {
	if len(e.Allowed) == 0 {
		return fmt.Sprintf(
			"'%s' must be defined",
			e.Name,
		)
	}

	return fmt.Sprintf(
		"'%s' is not defined, expected its value to be one of '%s'",
		e.Name,
		strings.Join(e.Allowed, "', '"),
	)
}

// UnexpectedValueError is an error that occurs when an environment variable
// contains an invalid value.
type UnexpectedValueError struct {
	Name    string
	Value   string
	Allowed []string
}

func (e *UnexpectedValueError) Error() string {
	return fmt.Sprintf(
		"'%s' is set to '%s', expected its value to be one of '%s'",
		e.Name,
		e.Value,
		strings.Join(e.Allowed, "', '"),
	)
}

func validateName(n string) {
	if n == "" {
		panic("environment variable name must not be empty")
	}
}
