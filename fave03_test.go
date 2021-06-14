package gomfa

import (
	"testing"
)

func BenchmarkFave03(b *testing.B) {
	for i := 0; i < 10000000; i++ {
		_ = Fave03(0.80)
	}
}

func TestFver03(t *testing.T) {
	var fave float64

	fave = Fave03(0.80)

	if !CheckFloat(fave, 3.424900460533758000, 1e-12) {
		t.Errorf("fave !=  3.424900460533758000 -> %v\n", fave)
	}
}
