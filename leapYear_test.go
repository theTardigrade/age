package age

import (
	"testing"
	"time"
)

func BenchmarkIsLeapYear(b *testing.B) {
	currentYear := time.Now().UTC().Year()

	for n := 0; n < b.N; n++ {
		IsLeapYear(currentYear)
	}
}
