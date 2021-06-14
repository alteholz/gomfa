package gomfa

import (
	"testing"
)

func BenchmarkFal03(b *testing.B) {
	for i := 0; i < 10000000; i++ {
		_ = Fal03(0.80)
	}
}

func TestFal03(t *testing.T) {
	var fal float64

	fal = Fal03(0.80)

	if !CheckFloat(fal, 5.132369751108684150, 1e-12) {
		t.Errorf("fal !=  5.132369751108684150 -> %v\n", fal)
	}
}
