package format

import (
	"fmt"
	"strings"
	"time"
)

// DurationToHuman format time.Duration value to human-readable string value.
func DurationToHuman(d time.Duration) string {
	var b strings.Builder
	if days := int64(d / (24 * time.Hour)); days > 0 {
		fmt.Fprintf(&b, "%d %s ",
			days, singOrPlural(days, "day", "days"))
	}
	if hrs := int64(d.Hours()) % 24; hrs > 0 {
		fmt.Fprintf(&b, "%d %s ", hrs,
			singOrPlural(hrs, "hour", "hours"))
	}
	if min := int64(d.Minutes()) % 60; min > 0 {
		fmt.Fprintf(&b, "%d min. ", min)
	}
	if ms := (d / time.Millisecond); ms > 0 {
		s, msr := int64(d.Seconds())%60, ms%1000/100
		if s > 0 || msr > 0 {
			fmt.Fprintf(&b, "%d.%d sec.", s, msr)
		}
	}
	if b.String() == "" {
		return "0.0 sec."
	}
	return strings.TrimSpace(b.String())
}

func singOrPlural(n int64, singular, plural string) string {
	if n != 1 {
		return plural
	}
	return singular
}
