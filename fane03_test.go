package gomfa

import (
	"testing"
)

func BenchmarkFane03(b *testing.B) {
	for i := 0; i < 10000000; i++ {
		_ = Fane03(0.80)
	}
}

func TestFane03(t *testing.T) {
	var fane float64

	fane = Fane03(0.80)

	if !CheckFloat(fane, 2.079343830860413523, 1e-12) {
		t.Errorf("fane !=  2.079343830860413523 -> %v\n", fane)
	}
}
