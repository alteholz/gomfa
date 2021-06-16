package gomfa

import (
	"testing"
)

func BenchmarkPvxpv(b *testing.B) {
	var av [2][3]float64
	var bv [2][3]float64
	var axb [2][3]float64

	av[0][0] = 2.0
	av[0][1] = 2.0
	av[0][2] = 3.0

	av[1][0] = 6.0
	av[1][1] = 0.0
	av[1][2] = 4.0

	bv[0][0] = 1.0
	bv[0][1] = 3.0
	bv[0][2] = 4.0

	bv[1][0] = 0.0
	bv[1][1] = 2.0
	bv[1][2] = 8.0

	for i := 0; i < 1000000; i++ {
		Pvxpv(&av, &bv, &axb)
	}
}

func TestPvxpv(t *testing.T) {
	var av [2][3]float64
	var bv [2][3]float64
	var axb [2][3]float64

	av[0][0] = 2.0
	av[0][1] = 2.0
	av[0][2] = 3.0

	av[1][0] = 6.0
	av[1][1] = 0.0
	av[1][2] = 4.0

	bv[0][0] = 1.0
	bv[0][1] = 3.0
	bv[0][2] = 4.0

	bv[1][0] = 0.0
	bv[1][1] = 2.0
	bv[1][2] = 8.0

	Pvxpv(&av, &bv, &axb)

	if axb[0][0] != -1.0 {
		t.Errorf("axb[0][0] != -1.0\n")
	}
	if axb[0][1] != -5.0 {
		t.Errorf("axb[0][1] != -5.0\n")
	}
	if axb[0][2] != 4.0 {
		t.Errorf("axb[0][2] != 4.0\n")
	}

	if axb[1][0] != -2.0 {
		t.Errorf("axb[1][0] != -2.0\n")
	}
	if axb[1][1] != -36.0 {
		t.Errorf("axb[1][1] != -36.0\n")
	}
	if axb[1][2] != 22.0 {
		t.Errorf("axb[1][2] != 22.0\n")
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
