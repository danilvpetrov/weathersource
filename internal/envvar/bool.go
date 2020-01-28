package envvar

// BoolDefault returns the boolean value of the environment variable named n, or
// the default d if it is not set.
func BoolDefault(n string, d bool) (bool, error) {
	v, ok, err := getBool(n)
	if ok || err != nil {
		return v, err
	}

	return d, nil
}

// BoolT returns the boolean value of the environment variable named n, or true
// if it is not set.
func BoolT(n string) (bool, error) {
	return BoolDefault(n, true)
}

// BoolF returns the boolean value of the environment variable named n, or false
// if it is not set.
func BoolF(n string) (bool, error) {
	return BoolDefault(n, false)
}

// getBool returns the boolean value of the environment variable named n.
func getBool(n string) (value bool, exists bool, err error) {
	v, exists, err := getEnum(
		n,
		append(trueValues, falseValues...)...,
	)

	for _, x := range trueValues {
		if x == v {
			value = true
			return
		}
	}

	return
}

var (
	trueValues  = []string{"1", "t", "true", "yes"}
	falseValues = []string{"0", "f", "false", "no"}
)
