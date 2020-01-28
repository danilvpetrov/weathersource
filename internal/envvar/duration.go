package envvar

import (
	"os"
	"time"
)

// DurationDefault returns the time.Duration value of the environment variable
// named n, or the default value d if the variable is empty or undefined.
func DurationDefault(n string, d time.Duration) (time.Duration, error) {
	v := StringDefault(
		n,
		d.String(),
	)

	dr, err := time.ParseDuration(v)
	if err != nil {
		return 0, &ValueParsingError{
			Name:       n,
			Value:      v,
			TargetType: "time.Duration",
		}
	}

	return dr, nil
}

// DurationRequired returns the time.Duration value of the environment variable
// named n. It returns an error if the variable is empty or undefined.
func DurationRequired(n string) (time.Duration, error) {
	if v := String(n); v != "" {
		dr, err := time.ParseDuration(v)
		if err != nil {
			return 0, &ValueParsingError{
				Name:       n,
				Value:      v,
				TargetType: "time.Duration",
			}
		}

		return dr, nil
	}

	return 0, &UndefinedError{n, nil}
}

// Duration returns the time.Duration value of the environment variable named n.
func Duration(n string) (time.Duration, error) {
	validateName(n)

	v := os.Getenv(n)

	dr, err := time.ParseDuration(v)
	if err != nil {
		return 0, &ValueParsingError{
			Name:       n,
			Value:      v,
			TargetType: "time.Duration",
		}
	}

	return dr, nil
}
