package age

import "time"

const (
	leapYearDay = 60
)

// Calculate provides an age from the present time to a given time.
func Calculate(givenTime time.Time) int64 {
	presentTime := time.Now()

	switch givenLocation := givenTime.Location(); givenLocation {
	case time.UTC, nil:
		presentTime = presentTime.UTC()
	default:
		presentTime = presentTime.In(givenLocation)
	}

	age := int64(presentTime.Year()) - int64(givenTime.Year())

	givenIsLeapYear := isLeapYear(givenTime)
	presentIsLeapYear := isLeapYear(presentTime)

	givenYearDay := givenTime.YearDay()
	presentYearDay := presentTime.YearDay()

	if givenIsLeapYear && !presentIsLeapYear && givenYearDay >= leapYearDay {
		givenYearDay--
	} else if presentIsLeapYear && !givenIsLeapYear && presentYearDay >= leapYearDay {
		givenYearDay++
	}

	if presentYearDay < givenYearDay {
		age--
	}

	return age
}

func isLeapYear(givenTime time.Time) bool {
	year := givenTime.Year()

	if year%400 == 0 {
		return true
	} else if year%100 == 0 {
		return false
	} else if year%4 == 0 {
		return true
	}

	return false
}
