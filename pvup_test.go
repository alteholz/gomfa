package gomfa

import (
	"testing"
)

func BenchmarkPvup(b *testing.B) {
	var pv [2][3]float64
	var p [3]float64

	pv[0][0] = 126668.5912743160734
	pv[0][1] = 2136.792716839935565
	pv[0][2] = -245251.2339876830229

	pv[1][0] = -0.4051854035740713039e-2
	pv[1][1] = -0.6253919754866175788e-2
	pv[1][2] = 0.1189353719774107615e-1

	for i := 0; i < 100000000; i++ {
		Pvup(2920.0, &pv, &p)
	}
}

func TestPvup(t *testing.T) {
	var pv [2][3]float64
	var p [3]float64

	pv[0][0] = 126668.5912743160734
	pv[0][1] = 2136.792716839935565
	pv[0][2] = -245251.2339876830229

	pv[1][0] = -0.4051854035740713039e-2
	pv[1][1] = -0.6253919754866175788e-2
	pv[1][2] = 0.1189353719774107615e-1

	Pvup(2920.0, &pv, &p)

	if !CheckFloat(p[0], 126656.7598605317105, 1e-6) {
		t.Errorf("p[0] != 126656.7598605317105\n")
	}
	if !CheckFloat(p[1], 2118.531271155726332, 1e-8) {
		t.Errorf("p[1] != 2118.531271155726332\n")
	}
	if !CheckFloat(p[2], -245216.5048590656190, 1e-6) {
		t.Errorf("p[2] != -245216.5048590656190\n")
	}

}
