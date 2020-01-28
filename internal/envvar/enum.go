package envvar

import (
	"strings"
)

// EnumDefault returns the string value of the environment variable named n,
// which must be in the given set of values. It returns the default d if it is
// not set.
//
// It returns an error if the value is not within the set of given values.
//
// The value comparison is CASE-INSENSITIVE. When successful, the returned value
// is always exactly as it appears in the value list.
func EnumDefault(n, d string, values ...string) (string, error) {
	v, ok, err := getEnum(n, values...)
	if ok || err != nil {
		return v, err
	}

	return d, nil
}

// EnumRequired returns the string value of the environment variable named n,
// which must be in the given set of values.
//
// It returns an error if the variable empty or unset, or the value is not
// within the set of given values.
//
// The value comparison is CASE-INSENSITIVE. When successful, the returned value
// is always exactly as it appears in the value list.
func EnumRequired(n string, values ...string) (string, error) {
	v, ok, err := getEnum(n, values...)
	if err != nil {
		return "", err
	} else if !ok {
		return "", &UndefinedError{n, values}
	}

	return v, nil
}

// getEnum returns the boolean value of the environment variable named n.
func getEnum(n string, values ...string) (value string, exists bool, err error) {
	if len(values) == 0 {
		panic("enums must have at least 1 possible value")
	}

	v := String(n)
	if v == "" {
		return "", false, nil
	}

	for _, x := range values {
		if x == "" {
			panic("enum values can not be empty")
		}

		if strings.EqualFold(v, x) {
			return x, true, nil
		}
	}

	return "", false, &UnexpectedValueError{n, v, values}
}
