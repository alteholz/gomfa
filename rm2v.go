package gomfa

import (
	"math"
)

func Rm2v(r *[3][3]float64, w *[3]float64) {
	/*
	 **  - - - - -
	 **   R m 2 v
	 **  - - - - -
	 **
	 **  Express an r-matrix as an r-vector.
	 **
	 **  Given:
	 **     r        *[3][3] float64    rotation matrix
	 **
	 **  Returned:
	 **     w        *[3] float64       rotation vector (Note 1)
	 **
	 **  Notes:
	 **
	 **  1) A rotation matrix describes a rotation through some angle about
	 **     some arbitrary axis called the Euler axis.  The "rotation vector"
	 **     returned by this function has the same direction as the Euler axis,
	 **     and its magnitude is the angle in radians.  (The magnitude and
	 **     direction can be separated by means of the function Pn.)
	 **
	 **  2) If r is null, so is the result.  If r is not a rotation matrix
	 **     the result is undefined;  r must be proper (i.e. have a positive
	 **     determinant) and real orthogonal (inverse = transpose).
	 **
	 **  3) The reference frame rotates clockwise as seen looking along
	 **     the rotation vector from the origin.
	 **
	 **  This revision:  2021 May 11
	 */
	var x, y, z, s2, c2, phi, f float64

	x = r[1][2] - r[2][1]
	y = r[2][0] - r[0][2]
	z = r[0][1] - r[1][0]
	s2 = math.Sqrt(x*x + y*y + z*z)
	if s2 > 0 {
		c2 = r[0][0] + r[1][1] + r[2][2] - 1.0
		phi = math.Atan2(s2, c2)
		f = phi / s2
		w[0] = x * f
		w[1] = y * f
		w[2] = z * f
	} else {
		w[0] = 0.0
		w[1] = 0.0
		w[2] = 0.0
	}

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
