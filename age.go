package age

import (
	"time"
)

const (
	leapYearDay = 60
)

// Calculate provides an age from the given time to the present time.
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

	givenYearIsLeapYear := isLeapYear(givenYear)
	presentYearIsLeapYear := isLeapYear(presentYear)

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

func isLeapYear(givenYear int) bool {
	if givenYear%400 == 0 {
		return true
	} else if givenYear%100 == 0 {
		return false
	} else if givenYear%4 == 0 {
		return true
	}

	return false
}
