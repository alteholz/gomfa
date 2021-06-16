package gomfa

import (
	"testing"
)

func BenchmarkTttdb(b *testing.B) {
	var b1, b2 float64

	for i := 0; i < 10000000; i++ {
		_ = Tttdb(2453750.5, 0.892855139, -0.000201, &b1, &b2)
	}
}

func TestTttdb(t *testing.T) {
	var b1, b2 float64
	var j int

	j = Tttdb(2453750.5, 0.892855139, -0.000201, &b1, &b2)

	if j != 0 {
		t.Errorf("j != 0\n")
	}
	if !CheckFloat(b1, 2453750.5, 1e-6) {
		t.Errorf("b1 !=  2453750.5 -> %v\n", b1)
	}
	if !CheckFloat(b2, 0.8928551366736111111, 1e-12) {
		t.Errorf("b2 !=  0.8928551366736111111 -> %v\n", b2)
	}
}
