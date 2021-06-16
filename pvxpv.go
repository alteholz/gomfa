package gomfa

func Pvxpv(a *[2][3]float64, b *[2][3]float64, axb *[2][3]float64) {
	/*
	 **  - - - - - -
	 **   P v x p v
	 **  - - - - - -
	 **
	 **  Outer (=vector=cross) product of two pv-vectors.
	 **
	 **  Given:
	 **     a        *[2][3] float64      first pv-vector
	 **     b        *[2][3] float64      second pv-vector
	 **
	 **  Returned:
	 **     axb      *[2][3] float64      a x b
	 **
	 **  Notes:
	 **
	 **  1) If the position and velocity components of the two pv-vectors are
	 **     ( ap, av ) and ( bp, bv ), the result, a x b, is the pair of
	 **     vectors ( ap x bp, ap x bv + av x bp ).  The two vectors are the
	 **     cross-product of the two p-vectors and its derivative.
	 **
	 **  2) It is permissible to re-use the same array for any of the
	 **     arguments.
	 **
	 **  Called:
	 **     Cpv       copy pv-vector
	 **     Pxp       vector product of two p-vectors
	 **     Ppp       p-vector plus p-vector
	 **
	 **  This revision:  2021 May 11
	 */
	var wa [2][3]float64
	var wb [2][3]float64
	var axbd [3]float64
	var adxb [3]float64

	/* Make copies of the inputs. */
	Cpv(a, &wa)
	Cpv(b, &wb)

	/* a x b = position part of result. */
	Pxp(&wa[0], &wb[0], &axb[0])

	/* a x bdot + adot x b = velocity part of result. */
	Pxp(&wa[0], &wb[1], &axbd)
	Pxp(&wa[1], &wb[0], &adxb)
	Ppp(&axbd, &adxb, &axb[1])

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
