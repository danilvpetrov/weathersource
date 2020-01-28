package envvar

import (
	"os"
	"strconv"
)

// Float64Default returns the float64 value of the environment variable named n,
// or the default value d if the variable is empty or undefined.
func Float64Default(n string, d float64) (float64, error) {
	v := StringDefault(
		n,
		strconv.FormatFloat(d, 'f', -1, 64),
	)

	f64, err := strconv.ParseFloat(v, 64)
	if err != nil {
		return 0, &ValueParsingError{
			Name:       n,
			Value:      v,
			TargetType: "float64",
		}
	}

	return f64, nil
}

// Float64Required returns the float64 value of the environment variable named n.
// It returns an error if the variable is empty or undefined.
func Float64Required(n string) (float64, error) {
	if v := String(n); v != "" {
		f64, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return 0, &ValueParsingError{
				Name:       n,
				Value:      v,
				TargetType: "float64",
			}
		}

		return f64, nil
	}

	return 0, &UndefinedError{n, nil}
}

// Float64 returns the float64 value of the environment variable named n.
func Float64(n string) (float64, error) {
	validateName(n)

	v := os.Getenv(n)

	f64, err := strconv.ParseFloat(v, 64)
	if err != nil {
		return 0, &ValueParsingError{
			Name:       n,
			Value:      v,
			TargetType: "float64",
		}
	}

	return f64, nil
}
