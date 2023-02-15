package age

const (
	leapYearDay = 60
)

// IsLeapYear returns true if a given year contains a leap day.
func IsLeapYear(givenYear int) bool {
	if givenYear%400 == 0 {
		return true
	} else if givenYear%100 == 0 {
		return false
	} else if givenYear%4 == 0 {
		return true
	}

	return false
}
