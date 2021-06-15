package gomfa

import (
	"testing"
)

func BenchmarkFaom03(b *testing.B) {
	for i := 0; i < 10000000; i++ {
		_ = Faom03(0.80)
	}
}

func TestFaom03(t *testing.T) {
	var faom float64

	faom = Faom03(0.80)

	if !CheckFloat(faom, -5.973618440951302183, 1e-12) {
		t.Errorf("faom !=  -5.973618440951302183 -> %v\n", faom)
	}
}
