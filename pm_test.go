package gomfa

import (
	"testing"
)

func BenchmarkPm(b *testing.B) {
	var p [3]float64

	p[0] = 0.3
	p[1] = 1.2
	p[2] = -2.5

	for i := 0; i < 100000000; i++ {
		Pm(&p)
	}
}

func TestPm(t *testing.T) {
	var p [3]float64
	var r float64

	p[0] = 0.3
	p[1] = 1.2
	p[2] = -2.5

	r = Pm(&p)

	if r != 2.789265136196270604 {
		t.Errorf("r != 2.789265136196270604\n")
	}
}
