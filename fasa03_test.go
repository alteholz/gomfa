package gomfa

import (
	"testing"
)

func BenchmarkFasa03(b *testing.B) {
	for i := 0; i < 10000000; i++ {
		_ = Fasa03(0.80)
	}
}

func TestFasa03(t *testing.T) {
	var fasa float64

	fasa = Fasa03(0.80)

	if !CheckFloat(fasa, 5.371574539440827046, 1e-12) {
		t.Errorf("fasa !=  5.371574539440827046 -> %v\n", fasa)
	}
}
