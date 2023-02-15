package age

import (
	"time"

	leapYear "github.com/theTardigrade/golang-leapYear"
)

const (
	calculateLeapYearDay = 60
)

func calculate(startTime, endTime time.Time) int {
	startYear := startTime.Year()
	endYear := endTime.Year()

	age := endYear - startYear

	startYearIsLeapYear := leapYear.Is(startYear)
	endYearIsLeapYear := leapYear.Is(endYear)

	startYearDay := startTime.YearDay()
	endYearDay := endTime.YearDay()

	if startYearIsLeapYear && !endYearIsLeapYear && startYearDay >= calculateLeapYearDay {
		startYearDay--
	} else if endYearIsLeapYear && !startYearIsLeapYear && endYearDay >= calculateLeapYearDay {
		startYearDay++
	}

	if endYearDay < startYearDay {
		age--
	}

	return age
}

// Calculate returns an integer-value age based on the duration
// between the two times that are given as arguments.
func Calculate(startTime, endTime time.Time) int {
	switch endLocation := endTime.Location(); endLocation {
	case time.UTC, nil:
		startTime = startTime.UTC()
	default:
		startTime = startTime.In(endLocation)
	}

	return calculate(startTime, endTime)
}

// Calculate returns an integer-value age based on the duration
// between the given time and the present time.
func CalculateToNow(givenTime time.Time) int {
	presentTime := time.Now().UTC()

	switch givenLocation := givenTime.Location(); givenLocation {
	case time.UTC, nil:
		presentTime = presentTime.UTC()
	default:
		presentTime = presentTime.In(givenLocation)
	}

	return calculate(givenTime, presentTime)
}
