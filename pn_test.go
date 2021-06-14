package gomfa

import (
	"testing"
)

func BenchmarkPn(b *testing.B) {
	var p [3]float64
	var u [3]float64
	var r float64

	p[0] = 0.3
	p[1] = 1.2
	p[2] = -2.5

	for i := 0; i < 100000000; i++ {
		Pn(&p, &r, &u)
	}
}

func TestPn(t *testing.T) {
	var p [3]float64
	var u [3]float64
	var r float64

	p[0] = 0.3
	p[1] = 1.2
	p[2] = -2.5

	Pn(&p, &r, &u)

	if !CheckFloat(r, 2.789265136196270604, 1e-12) {
		t.Errorf("r != 2.789265136196270604\n")
	}
	if !CheckFloat(u[0], 0.1075552109073112058, 1e-12) {
		t.Errorf("u[0] != 0.1075552109073112058\n")
	}
	if !CheckFloat(u[1], 0.4302208436292448232, 1e-12) {
		t.Errorf("u[1] != 0.4302208436292448232\n")
	}
	if !CheckFloat(u[2], -0.8962934242275933816, 1e-12) {
		t.Errorf("u[2] != -0.8962934242275933816\n")
	}
}
