package gomfa

import (
	"testing"
)

func BenchmarkP2s(b *testing.B) {
	var p [3]float64
	var theta, phi, r float64

	p[0] = 100.0
	p[1] = -50.0
	p[2] = 25.0

	for i := 0; i < 1000000; i++ {
		P2s(&p, &theta, &phi, &r)
	}
}

func TestP2s(t *testing.T) {
	var p [3]float64
	var theta, phi, r float64

	p[0] = 100.0
	p[1] = -50.0
	p[2] = 25.0

	P2s(&p, &theta, &phi, &r)

	if !CheckFloat(theta, -0.4636476090008061162, 1e-14) {
		t.Errorf("theta != -0.4636476090008061162\n")
	}
	if !CheckFloat(phi, 0.2199879773954594463, 1e-14) {
		t.Errorf("phi !=  0.2199879773954594463\n")
	}
	if !CheckFloat(r, 114.5643923738960002, 1e-9) {
		t.Errorf("phi !=  114.5643923738960002\n")
	}
}
