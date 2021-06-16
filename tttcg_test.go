package gomfa

import (
	"testing"
)

func BenchmarkTttcg(b *testing.B) {
	var g1, g2 float64

	for i := 0; i < 10000000; i++ {
		_ = Tttcg(2453750.5, 0.892482639, &g1, &g2)
	}
}

func TestTttcg(t *testing.T) {
	var g1, g2 float64
	var j int

	j = Tttcg(2453750.5, 0.892482639, &g1, &g2)

	if j != 0 {
		t.Errorf("j != 0\n")
	}
	if !CheckFloat(g1, 2453750.5, 1e-6) {
		t.Errorf("g1 !=  2453750.5 -> %v\n", g1)
	}
	if !CheckFloat(g2, 0.8924900312508587113, 1e-12) {
		t.Errorf("g2 !=  0.8924900312508587113 -> %v\n", g2)
	}
}
