package gomfa

import (
	"testing"
)

func BenchmarkBpn2xy(b *testing.B) {
	var rbpn [3][3]float64
	var x, y float64

	rbpn[0][0] = 9.999962358680738e-1
	rbpn[0][1] = -2.516417057665452e-3
	rbpn[0][2] = -1.093569785342370e-3

	rbpn[1][0] = 2.516462370370876e-3
	rbpn[1][1] = 9.999968329010883e-1
	rbpn[1][2] = 4.006159587358310e-5

	rbpn[2][0] = 1.093465510215479e-3
	rbpn[2][1] = -4.281337229063151e-5
	rbpn[2][2] = 9.999994012499173e-1

	for i := 0; i < 100000000; i++ {
		Bpn2xy(&rbpn, &x, &y)
	}
}

func TestBpn2xy(t *testing.T) {
	var rbpn [3][3]float64
	var x, y float64

	rbpn[0][0] = 9.999962358680738e-1
	rbpn[0][1] = -2.516417057665452e-3
	rbpn[0][2] = -1.093569785342370e-3

	rbpn[1][0] = 2.516462370370876e-3
	rbpn[1][1] = 9.999968329010883e-1
	rbpn[1][2] = 4.006159587358310e-5

	rbpn[2][0] = 1.093465510215479e-3
	rbpn[2][1] = -4.281337229063151e-5
	rbpn[2][2] = 9.999994012499173e-1

	Bpn2xy(&rbpn, &x, &y)

	if !CheckFloat(x, 1.093465510215479e-3, 1e-12) {
		t.Errorf("x != 1.093465510215479e-3\n")
	}
	if !CheckFloat(y, -4.281337229063151e-5, 1e-12) {
		t.Errorf("y != -4.281337229063151e-5\n")
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
