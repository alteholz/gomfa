package gomfa

import (
	"math"
)

func Tcgtt(tcg1 float64, tcg2 float64, tt1 *float64, tt2 *float64) int {
	/*
	 **  - - - - - -
	 **   T c g t t
	 **  - - - - - -
	 **
	 **  Time scale transformation:  Geocentric Coordinate Time, TCG, to
	 **  Terrestrial Time, TT.
	 **
	 **  Given:
	 **     tcg1,tcg2  float64    TCG as a 2-part Julian Date
	 **
	 **  Returned:
	 **     tt1,tt2    float64    TT as a 2-part Julian Date
	 **
	 **  Returned (function value):
	 **                int       status:  0 = OK
	 **
	 **  Note:
	 **
	 **     tcg1+tcg2 is Julian Date, apportioned in any convenient way
	 **     between the two arguments, for example where tcg1 is the Julian
	 **     Day Number and tcg22 is the fraction of a day.  The returned
	 **     tt1,tt2 follow suit.
	 **
	 **  References:
	 **
	 **     McCarthy, D. D., Petit, G. (eds.), IERS Conventions (2003),
	 **     IERS Technical Note No. 32, BKG (2004)
	 **
	 **     IAU 2000 Resolution B1.9
	 **
	 **  This revision:  2021 May 11
	 */

	/* 1977 Jan 1 00:00:32.184 TT, as MJD */
	const t77t float64 = DJM77 + TTMTAI/DAYSEC

	/* Result, safeguarding precision. */
	if math.Abs(tcg1) > math.Abs(tcg2) {
		*tt1 = tcg1
		*tt2 = tcg2 - ((tcg1-DJM0)+(tcg2-t77t))*ELG
	} else {
		*tt1 = tcg1 - ((tcg2-DJM0)+(tcg1-t77t))*ELG
		*tt2 = tcg2
	}

	/* OK status. */
	return 0

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
