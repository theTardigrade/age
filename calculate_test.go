package age

import (
	"testing"
	"time"
)

func BenchmarkCalculate(b *testing.B) {
	currentTime := time.Now().UTC()

	for n := 0; n < b.N; n++ {
		Calculate(currentTime)
	}
}
