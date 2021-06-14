package gomfa

import (
	"testing"
)

func BenchmarkFalp03(b *testing.B) {
	for i := 0; i < 10000000; i++ {
		_ = Falp03(0.80)
	}
}

func TestFalp03(t *testing.T) {
	var falp float64

	falp = Falp03(0.80)

	if !CheckFloat(falp, 6.226797973505507345, 1e-12) {
		t.Errorf("falp !=  6.226797973505507345 -> %v\n", falp)
	}
}
