package envvar

import (
	"os"
)

// StringDefault returns the string value of the environment variable named n,
// or the default value d if the variable is empty or undefined.
func StringDefault(n, d string) string {
	if v := String(n); v != "" {
		return v
	}

	return d
}

// StringRequired returns the string value of the environment variable named n.
// It returns an error if the variable is empty or undefined.
func StringRequired(n string) (string, error) {
	if v := String(n); v != "" {
		return v, nil
	}

	return "", &UndefinedError{n, nil}
}

// String returns the string value of the environment variable named n.
func String(n string) string {
	validateName(n)
	return os.Getenv(n)
}
