package gomfa

import (
	"math"
)

func Sepp(a *[3]float64, b *[3]float64) float64 {
	/*
	 **  - - - - -
	 **   S e p p
	 **  - - - - -
	 **
	 **  Angular separation between two p-vectors.
	 **
	 **  Given:
	 **     a      *[3] float64    first p-vector (not necessarily unit length)
	 **     b      *[3] float64    second p-vector (not necessarily unit length)
	 **
	 **  Returned (function value):
	 **            float64       angular separation (radians, always positive)
	 **
	 **  Notes:
	 **
	 **  1) If either vector is null, a zero result is returned.
	 **
	 **  2) The angular separation is most simply formulated in terms of
	 **     scalar product.  However, this gives poor accuracy for angles
	 **     near zero and pi.  The present algorithm uses both cross product
	 **     and dot product, to deliver full accuracy whatever the size of
	 **     the angle.
	 **
	 **  Called:
	 **     Pxp       vector product of two p-vectors
	 **     Pm        modulus of p-vector
	 **     Pdp       scalar product of two p-vectors
	 **
	 **  This revision:  2021 May 11
	 */
	var axb [3]float64
	var ss, cs, s float64

	/* Sine of angle between the vectors, multiplied by the two moduli. */
	Pxp(a, b, &axb)
	ss = Pm(&axb)

	/* Cosine of the angle, multiplied by the two moduli. */
	cs = Pdp(a, b)

	/* The angle. */
	if (ss != 0.0) || (cs != 0.0) {
		s = math.Atan2(ss, cs)
	} else {
		s = 0.0
	}

	return s

	/* Finished. */

}

/*----------------------------------------------------------------------
**
**
**  Copyright (C) 2021,      Thorsten Alteholz
**  All rights reserved.
**
**  This library is derived, with permission, from the International
**  Astronomical Union's "Standards of Fundamental Astronomy" library,
**  available from http://www.iausofa.org.
**
**  The GOMFA version is intended to retain identical functionality to
**  the SOFA library, but made distinct through different namespaces and
**  file names, as set out in the SOFA license conditions.  The SOFA
**  original has a role as a reference standard for the IAU and IERS,
**  and consequently redistribution is permitted only in its unaltered
**  state.  The GOMFA version is not subject to this restriction and
**  therefore can be included in distributions which do not support the
**  concept of "read only" software.
**
**  Although the intent is to replicate the SOFA API (other than
**  replacement of prefix names) and results (with the exception of
**  bugs;  any that are discovered will be fixed), SOFA is not
**  responsible for any errors found in this version of the library.
**
**  If you wish to acknowledge the SOFA heritage, please acknowledge
**  that you are using a library derived from SOFA, rather than SOFA
**  itself.
**
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
