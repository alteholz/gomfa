package gomfa

import (
	"testing"
)

func BenchmarkZp(b *testing.B) {
	var p [3]float64

	p[0] = 0.3
	p[1] = 1.2
	p[2] = -2.5

	for i := 0; i < 100000000; i++ {
		Zp(&p)
	}
}

func TestZp(t *testing.T) {
	var p [3]float64

	p[0] = 0.3
	p[1] = 1.2
	p[2] = -2.5

	Zp(&p)

	if p[0] != 0.0 {
		t.Errorf("p[0] != 0.0\n")
	}
	if p[1] != 0.0 {
		t.Errorf("p[1] != 0.0\n")
	}
	if p[2] != 0.0 {
		t.Errorf("p[2] != 0.0\n")
	}
}
