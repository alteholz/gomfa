package gomfa

import (
	"testing"
)

func BenchmarkFw2xy(b *testing.B) {

	var gamb, phib, psi, eps, x, y float64

	gamb = -0.2243387670997992368e-5
	phib = 0.4091014602391312982
	psi = -0.9501954178013015092e-3
	eps = 0.4091014316587367472

	for i := 0; i < 10000000; i++ {
		Fw2xy(gamb, phib, psi, eps, &x, &y)
	}
}

func TestFw2xy(t *testing.T) {
	var gamb, phib, psi, eps, x, y float64

	gamb = -0.2243387670997992368e-5
	phib = 0.4091014602391312982
	psi = -0.9501954178013015092e-3
	eps = 0.4091014316587367472

	Fw2xy(gamb, phib, psi, eps, &x, &y)

	if !CheckFloat(x, -0.3779734957034082790e-3, 1e-14) {
		t.Errorf("x !=  -0.3779734957034082790e-3 -> %v\n", x)
	}
	if !CheckFloat(y, -0.1924880848087615651e-6, 1e-14) {
		t.Errorf("y !=  -0.1924880848087615651e-6 -> %v\n", y)
	}
}
