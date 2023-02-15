package age

import (
	"time"
)

// Calculate returns an integer-value age based on the duration
// between the given time to the present time.
func Calculate(givenTime time.Time) int {
	presentTime := time.Now()

	switch givenLocation := givenTime.Location(); givenLocation {
	case time.UTC, nil:
		presentTime = presentTime.UTC()
	default:
		presentTime = presentTime.In(givenLocation)
	}

	givenYear := givenTime.Year()
	presentYear := presentTime.Year()

	age := presentYear - givenYear

	givenYearIsLeapYear := IsLeapYear(givenYear)
	presentYearIsLeapYear := IsLeapYear(presentYear)

	givenYearDay := givenTime.YearDay()
	presentYearDay := presentTime.YearDay()

	if givenYearIsLeapYear && !presentYearIsLeapYear && givenYearDay >= leapYearDay {
		givenYearDay--
	} else if presentYearIsLeapYear && !givenYearIsLeapYear && presentYearDay >= leapYearDay {
		givenYearDay++
	}

	if presentYearDay < givenYearDay {
		age--
	}

	return age
}
