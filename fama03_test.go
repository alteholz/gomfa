package gomfa

import (
	"testing"
)

func BenchmarkFama03(b *testing.B) {
	for i := 0; i < 10000000; i++ {
		_ = Fama03(0.80)
	}
}

func TestFama03(t *testing.T) {
	var fama float64

	fama = Fama03(0.80)

	if !CheckFloat(fama, 3.275506840277781492, 1e-12) {
		t.Errorf("fama !=  3.275506840277781492 -> %v\n", fama)
	}
}
