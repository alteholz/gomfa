package gomfa

import (
	"testing"
)

func BenchmarkNumat(b *testing.B) {
	var epsa, dpsi, deps float64
	var rmatn [3][3]float64

	epsa = 0.4090789763356509900
	dpsi = -0.9630909107115582393e-5
	deps = 0.4063239174001678826e-4

	for i := 0; i < 100000000; i++ {
		Numat(epsa, dpsi, deps, &rmatn)
	}
}

func TestNumat(t *testing.T) {
	var epsa, dpsi, deps float64
	var rmatn [3][3]float64

	epsa = 0.4090789763356509900
	dpsi = -0.9630909107115582393e-5
	deps = 0.4063239174001678826e-4

	Numat(epsa, dpsi, deps, &rmatn)

	if !CheckFloat(rmatn[0][0], 0.9999999999536227949, 1e-12) {
		t.Errorf("rmatn[0][0] != 0.9999999999536227949\n")
	}
	if !CheckFloat(rmatn[0][1], 0.8836239320236250577e-5, 1e-12) {
		t.Errorf("rmatn[0][1] != 0.8836239320236250577e-5\n")
	}
	if !CheckFloat(rmatn[0][2], 0.3830833447458251908e-5, 1e-12) {
		t.Errorf("rmatn[0][2] != 0.3830833447458251908e-5\n")
	}
	if !CheckFloat(rmatn[1][0], -0.8836083657016688588e-5, 1e-12) {
		t.Errorf("rmatn[1][0] != -0.8836083657016688588e-5\n")
	}
	if !CheckFloat(rmatn[1][1], 0.9999999991354654959, 1e-12) {
		t.Errorf("rmatn[1][1] != 0.9999999991354654959\n")
	}
	if !CheckFloat(rmatn[1][2], -0.4063240865361857698e-4, 1e-12) {
		t.Errorf("rmatn[1][2] != -0.4063240865361857698e-4\n")
	}
	if !CheckFloat(rmatn[2][0], -0.3831192481833385226e-5, 1e-12) {
		t.Errorf("rmatn[2][0] != -0.3831192481833385226e-5\n")
	}
	if !CheckFloat(rmatn[2][1], 0.4063237480216934159e-4, 1e-12) {
		t.Errorf("rmatn[2][1] != 0.4063237480216934159e-4\n")
	}
	if !CheckFloat(rmatn[2][2], 0.9999999991671660407, 1e-12) {
		t.Errorf("rmatn[2][2] != 0.9999999991671660407\n")
	}
}

/*----------------------------------------------------------------------
**
**
**  Copyright (C) 2021,      Thorsten Alteholz
**  All rights reserved.
**
**  TERMS AND CONDITIONS
**
**  Redistribution and use in source and binary forms, with or without
**  modification, are permitted provided that the following conditions
**  are met:
**
**  1 Redistributions of source code must retain the above copyright
**    notice, this list of conditions and the following disclaimer.
**
**  2 Redistributions in binary form must reproduce the above copyright
**    notice, this list of conditions and the following disclaimer in
**    the documentation and/or other materials provided with the
**    distribution.
**
**  3 Neither the name of the Standards Of Fundamental Astronomy Board,
**    the International Astronomical Union nor the names of its
**    contributors may be used to endorse or promote products derived
**    from this software without specific prior written permission.
**
**  THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
**  "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
**  LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS
**  FOR A PARTICULAR PURPOSE ARE DISCLAIMED.  IN NO EVENT SHALL THE
**  COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT,
**  INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING,
**  BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
**  LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
**  CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT
**  LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN
**  ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
**  POSSIBILITY OF SUCH DAMAGE.
**
 */
