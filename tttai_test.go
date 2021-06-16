package gomfa

import (
	"testing"
)

func BenchmarkTttai(b *testing.B) {
	var a1, a2 float64

	for i := 0; i < 10000000; i++ {
		_ = Tttai(2453750.5, 0.892482639, &a1, &a2)
	}
}

func TestTttai(t *testing.T) {
	var a1, a2 float64
	var j int

	j = Tttai(2453750.5, 0.892482639, &a1, &a2)

	if j != 0 {
		t.Errorf("j != 0\n")
	}
	if !CheckFloat(a1, 2453750.5, 1e-6) {
		t.Errorf("a1 !=  2453750.5 -> %v\n", a1)
	}
	if !CheckFloat(a2, 0.892110139, 1e-12) {
		t.Errorf("a2 !=  0.892110139 -> %v\n", a2)
	}
}
