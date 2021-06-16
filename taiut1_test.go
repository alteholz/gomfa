package gomfa

import (
	"testing"
)

func BenchmarkTaiut1(b *testing.B) {
	var u1, u2 float64

	for i := 0; i < 10000000; i++ {
		_ = Taiut1(2453750.5, 0.892482639, -32.6659, &u1, &u2)
	}
}

func TestTaiut1(t *testing.T) {
	var u1, u2 float64
	var j int

	j = Taiut1(2453750.5, 0.892482639, -32.6659, &u1, &u2)

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
