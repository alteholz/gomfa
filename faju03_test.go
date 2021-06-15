package gomfa

import (
	"testing"
)

func BenchmarkFaju03(b *testing.B) {
	for i := 0; i < 10000000; i++ {
		_ = Faju03(0.80)
	}
}

func TestFaju03(t *testing.T) {
	var faju float64

	faju = Faju03(0.80)

	if !CheckFloat(faju, 5.275711665202481138, 1e-12) {
		t.Errorf("faju !=  5.275711665202481138 -> %v\n", faju)
	}
}
