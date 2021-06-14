package gomfa

import (
	"testing"
)

func BenchmarkFae03(b *testing.B) {
	for i := 0; i < 10000000; i++ {
		_ = Fae03(0.80)
	}
}

func TestFae03(t *testing.T) {
	var fae float64

	fae = Fae03(0.80)

	if !CheckFloat(fae, 1.744713738913081846, 1e-12) {
		t.Errorf("fae !=  1.744713738913081846 -> %v\n", fae)
	}
}
