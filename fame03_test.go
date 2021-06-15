package gomfa

import (
	"testing"
)

func BenchmarkFame03(b *testing.B) {
	for i := 0; i < 10000000; i++ {
		_ = Fame03(0.80)
	}
}

func TestFame03(t *testing.T) {
	var fame float64

	fame = Fame03(0.80)

	if !CheckFloat(fame, 5.417338184297289661, 1e-12) {
		t.Errorf("fame !=  5.417338184297289661 -> %v\n", fame)
	}
}
