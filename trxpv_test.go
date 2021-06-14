package gomfa

import (
	"testing"
)

func BenchmarkTrxpv(b *testing.B) {
	var r [3][3]float64
	var pv [2][3]float64
	var rpv [2][3]float64

	r[0][0] = 2.0
	r[1][0] = 3.0
	r[2][0] = 2.0

	r[0][1] = 3.0
	r[1][1] = 2.0
	r[2][1] = 3.0

	r[0][2] = 3.0
	r[1][2] = 4.0
	r[2][2] = 5.0

	pv[0][0] = 0.2
	pv[0][1] = 1.5
	pv[0][2] = 0.1

	pv[1][0] = 1.5
	pv[1][1] = 0.2
	pv[1][2] = 0.1

	for i := 0; i < 10000000; i++ {
		Trxpv(&r, &pv, &rpv)
	}
}

func TestTrxpv(t *testing.T) {
	var r [3][3]float64
	var pv [2][3]float64
	var rpv [2][3]float64

	r[0][0] = 2.0
	r[0][1] = 3.0
	r[0][2] = 2.0

	r[1][0] = 3.0
	r[1][1] = 2.0
	r[1][2] = 3.0

	r[2][0] = 3.0
	r[2][1] = 4.0
	r[2][2] = 5.0

	pv[0][0] = 0.2
	pv[0][1] = 1.5
	pv[0][2] = 0.1

	pv[1][0] = 1.5
	pv[1][1] = 0.2
	pv[1][2] = 0.1

	Trxpv(&r, &pv, &rpv)

	if !CheckFloat(rpv[0][0], 5.2, 1e-12) {
		t.Errorf("rpv[0][0] != 5.2\n")
	}
	if !CheckFloat(rpv[0][1], 4.0, 1e-12) {
		t.Errorf("rpv[0][1] != 4,0\n")
	}
	if !CheckFloat(rpv[0][2], 5.4, 1e-12) {
		t.Errorf("rpv[0][2] != 5.4\n")
	}

	if !CheckFloat(rpv[1][0], 3.9, 1e-12) {
		t.Errorf("rpv[1][0] != 3.9\n")
	}
	if !CheckFloat(rpv[1][1], 5.3, 1e-12) {
		t.Errorf("rpv[1][1] != 5.3\n")
	}
	if !CheckFloat(rpv[1][2], 4.1, 1e-12) {
		t.Errorf("rpv[1][2] != 4.1\n")
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
