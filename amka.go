package amka

import (
	"strconv"
	"time"

	"github.com/theplant/luhn"
)

// Validate the given string for AMKA
func Validate(amka string) (bool, string) {
	if _, err := strconv.Atoi(amka); err != nil {
		return false, "Provided AMKA is not a numeric value"
	}

	if len(amka) != 11 {
		return false, "Provided number is not 11 digits long"
	}

	if _, err := time.Parse("020106", amka[:6]); err != nil {
		return false, "First 6 digits of the provided AMKA do not represent a valid date"
	}

	if n, _ := strconv.Atoi(amka); !luhn.Valid(n) {
		return false, "Provided AMKA does not have a valid checkdigit"
	}

	return true, ""
}
