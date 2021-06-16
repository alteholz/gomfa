package gomfa

import (
	"testing"
)

func BenchmarkTtut1(b *testing.B) {
	var u1, u2 float64

	for i := 0; i < 10000000; i++ {
		_ = Ttut1(2453750.5, 0.892855139, 64.8499, &u1, &u2)
	}
}

func TestTtut1(t *testing.T) {
	var u1, u2 float64
	var j int

	j = Ttut1(2453750.5, 0.892855139, 64.8499, &u1, &u2)

	if j != 0 {
		t.Errorf("j != 0\n")
	}
	if !CheckFloat(u1, 2453750.5, 1e-6) {
		t.Errorf("u1 !=  2453750.5 -> %v\n", u1)
	}
	if !CheckFloat(u2, 0.8921045614537037037, 1e-12) {
		t.Errorf("u2 !=  0.8921045614537037037 -> %v\n", u2)
	}
}
