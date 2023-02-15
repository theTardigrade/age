package age

import (
	"testing"
	"time"
)

func BenchmarkCalculateToNow(b *testing.B) {
	currentTime := time.Now().UTC()

	for n := 0; n < b.N; n++ {
		CalculateToNow(currentTime)
	}
}
