package gomfa

import (
	"testing"
)

func BenchmarkRy(b *testing.B) {
	var r [3][3]float64
	var theta float64

	theta = 0.3456789

	r[0][0] = 2.0
	r[1][0] = 3.0
	r[2][0] = 2.0

	r[0][1] = 3.0
	r[1][1] = 2.0
	r[2][1] = 3.0

	r[0][2] = 3.0
	r[1][2] = 4.0
	r[2][2] = 5.0

	for i := 0; i < 100000000; i++ {
		Ry(theta, &r)
	}
}

func TestRy(t *testing.T) {
	var r [3][3]float64
	var theta float64

	theta = 0.3456789

	r[0][0] = 2.0
	r[0][1] = 3.0
	r[0][2] = 2.0

	r[1][0] = 3.0
	r[1][1] = 2.0
	r[1][2] = 3.0

	r[2][0] = 3.0
	r[2][1] = 4.0
	r[2][2] = 5.0

	Ry(theta, &r)

	if !CheckFloat(r[0][0], 0.8651847818978159930, 1e-12) {
		t.Errorf("r[0][0] != 0.8651847818978159930\n")
	}
	if !CheckFloat(r[0][1], 1.467194920539316554, 1e-12) {
		t.Errorf("r[0][1] != 1.467194920539316554\n")
	}
	if !CheckFloat(r[0][2], 0.1875137911274457342, 1e-12) {
		t.Errorf("r[0][2] != 0.1875137911274457342\n")
	}
	if r[1][0] != 3.0 {
		t.Errorf("r[1][0] != 3.0\n")
	}
	if r[1][1] != 2.0 {
		t.Errorf("r[1][1] != 2.0\n")
	}
	if r[1][2] != 3.0 {
		t.Errorf("r[1][2] != 3.0\n")
	}
	if !CheckFloat(r[2][0], 3.500207892850427330, 1e-12) {
		t.Errorf("r[2][0] != 3.500207892850427330\n")
	}
	if !CheckFloat(r[2][1], 4.779889022262298150, 1e-12) {
		t.Errorf("r[2][1] != 4.779889022262298150\n")
	}
	if !CheckFloat(r[2][2], 5.381899160903798712, 1e-12) {
		t.Errorf("r[2][2] != 5.381899160903798712\n")
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
