package gomfa

import (
	"testing"
)

func BenchmarkLdsun(b *testing.B) {
	var p [3]float64
	var e [3]float64
	var p1 [3]float64
	var em float64

	p[0] = -0.763276255
	p[1] = -0.608633767
	p[2] = -0.216735543
	e[0] = -0.973644023
	e[1] = -0.20925523
	e[2] = -0.0907169552
	em = 0.999809214

	for i := 0; i < 100000000; i++ {
		Ldsun(&p, &e, em, &p1)
	}
}

func TestLdsun(t *testing.T) {
	var p [3]float64
	var e [3]float64
	var p1 [3]float64
	var em float64

	p[0] = -0.763276255
	p[1] = -0.608633767
	p[2] = -0.216735543
	e[0] = -0.973644023
	e[1] = -0.20925523
	e[2] = -0.0907169552
	em = 0.999809214

	Ldsun(&p, &e, em, &p1)

	if !CheckFloat(p1[0], -0.7632762580731413169, 1e-12) {
		t.Errorf("p1[0] != -0.7632762580731413169\n")
	}
	if !CheckFloat(p1[1], -0.6086337635262647900, 1e-12) {
		t.Errorf("p1[1] != -0.6086337635262647900\n")
	}
	if !CheckFloat(p1[2], -0.2167355419322321302, 1e-12) {
		t.Errorf("p1[2] != -0.2167355419322321302\n")
	}
}
