package gomfa

import (
	"testing"
)

func BenchmarkFaur03(b *testing.B) {
	for i := 0; i < 10000000; i++ {
		_ = Faur03(0.80)
	}
}

func TestFaur03(t *testing.T) {
	var faur float64

	faur = Faur03(0.80)

	if !CheckFloat(faur, 5.180636450180413523, 1e-12) {
		t.Errorf("faur !=  5.180636450180413523 -> %v\n", faur)
	}
}
