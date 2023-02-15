package age

import (
	"time"
)

func calculate(startTime, endTime time.Time) int {
	startYear := startTime.Year()
	endYear := endTime.Year()

	age := endYear - startYear

	startYearIsLeapYear := IsLeapYear(startYear)
	endYearIsLeapYear := IsLeapYear(endYear)

	startYearDay := startTime.YearDay()
	endYearDay := endTime.YearDay()

	if startYearIsLeapYear && !endYearIsLeapYear && startYearDay >= leapYearDay {
		startYearDay--
	} else if endYearIsLeapYear && !startYearIsLeapYear && endYearDay >= leapYearDay {
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
