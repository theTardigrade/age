package age

import (
	"testing"
	"time"
)

var (
	currentTime = time.Now().UTC()
	currentYear = currentTime.Year()
)

func BenchmarkCalculate(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Calculate(currentTime)
	}
}

func BenchmarkIsLeapYear(b *testing.B) {
	for n := 0; n < b.N; n++ {
		isLeapYear(currentYear)
	}
}
