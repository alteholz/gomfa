package gomfa

import (
	"testing"
)

func BenchmarkS2pv(b *testing.B) {
	var pv [2][3]float64

	for i := 0; i < 10000000; i++ {
		S2pv(-3.21, 0.123, 0.456, -7.8e-6, 9.01e-6, -1.23e-5, &pv)
	}
}

func TestS2pv(t *testing.T) {
	var pv [2][3]float64

	S2pv(-3.21, 0.123, 0.456, -7.8e-6, 9.01e-6, -1.23e-5, &pv)

	if !CheckFloat(pv[0][0], -0.4514964673880165228, 1e-12) {
		t.Errorf("pv[0][0] != -0.4514964673880165228 -> %v\n", pv[1][0])
	}
	if !CheckFloat(pv[0][1], 0.0309339427734258688, 1e-12) {
		t.Errorf("pv[0][1] != 0.0309339427734258688 -> %v\n", pv[1][1])
	}
	if !CheckFloat(pv[0][2], 0.0559466810510877933, 1e-12) {
		t.Errorf("pv[0][2] != 0.0559466810510877933 -> %v\n", pv[1][2])
	}
	if !CheckFloat(pv[1][0], 0.1292270850663260170e-4, 1e-16) {
		t.Errorf("pv[1][0] != 0.1292270850663260170e-4 -> %v\n", pv[1][0])
	}
	if !CheckFloat(pv[1][1], 0.2652814182060691422e-5, 1e-16) {
		t.Errorf("pv[1][1] != 0.2652814182060691422e-5 -> %v\n", pv[1][1])
	}
	if !CheckFloat(pv[1][2], 0.2568431853930292259e-5, 1e-16) {
		t.Errorf("pv[1][2] != 0.2568431853930292259e-5 -> %v\n", pv[1][2])
	}
}
