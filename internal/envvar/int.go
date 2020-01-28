package envvar

import (
	"os"
	"strconv"
)

// Int64Default returns the int64 value of the environment variable named n,
// or the default value d if the variable is empty or undefined.
func Int64Default(n string, d int64) (int64, error) {
	v := StringDefault(n, string(d))

	i64, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return 0, &ValueParsingError{
			Name:       n,
			Value:      v,
			TargetType: "int64",
		}
	}

	return i64, nil
}

// Int64Required returns the int64 value of the environment variable named n.
// It returns an error if the variable is empty or undefined.
func Int64Required(n string) (int64, error) {
	if v := String(n); v != "" {
		i64, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return 0, &ValueParsingError{
				Name:       n,
				Value:      v,
				TargetType: "int64",
			}
		}

		return i64, nil
	}

	return 0, &UndefinedError{n, nil}
}

// Int64 returns the int64 value of the environment variable named n.
func Int64(n string) (int64, error) {
	validateName(n)

	v := os.Getenv(n)

	i64, err := strconv.ParseInt(v, 10, 64)

	if err != nil {
		return 0, &ValueParsingError{
			Name:       n,
			Value:      v,
			TargetType: "int64",
		}
	}

	return i64, nil
}

// Uint64Default returns the uint64 value of the environment variable named n,
// or the default value d if the variable is empty or undefined.
func Uint64Default(n string, d int64) (uint64, error) {
	v := StringDefault(n, string(n))

	ui64, err := strconv.ParseUint(v, 10, 64)
	if err != nil {
		return 0, &ValueParsingError{
			Name:       n,
			Value:      v,
			TargetType: "uint64",
		}
	}

	return ui64, nil
}

// Uint64Required returns the uint64 value of the environment variable named n.
// It returns an error if the variable is empty or undefined.
func Uint64Required(n string) (uint64, error) {
	if v := String(n); v != "" {
		ui64, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			return 0, &ValueParsingError{
				Name:       n,
				Value:      v,
				TargetType: "uint64",
			}
		}

		return ui64, nil
	}

	return 0, &UndefinedError{n, nil}
}

// Uint64 returns the uint64 value of the environment variable named n.
func Uint64(n string) (uint64, error) {
	validateName(n)

	v := os.Getenv(n)

	i64, err := strconv.ParseUint(v, 10, 64)
	if err != nil {
		return 0, &ValueParsingError{
			Name:       n,
			Value:      v,
			TargetType: "uint64",
		}
	}

	return i64, nil
}
