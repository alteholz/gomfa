package gomfa

/*
**  - - - - - - -
**   gomfa_h.go
**  - - - - - - -
**
**  Prototype function declarations for GoMFA library.
 */

/* Star-independent astrometry parameters */
type ASTROM struct {
	pmt    float64       /* PM time interval (SSB, Julian years) */
	eb     [3]float64    /* SSB to observer (vector, au) */
	eh     [3]float64    /* Sun to observer (unit vector) */
	em     float64       /* distance from Sun to observer (au) */
	v      [3]float64    /* barycentric observer velocity (vector, c) */
	bm1    float64       /* sqrt(1-|v|^2): reciprocal of Lorenz factor */
	bpn    [3][3]float64 /* bias-precession-nutation matrix */
	along  float64       /* longitude + s' + dERA(DUT) (radians) */
	phi    float64       /* geodetic latitude (radians) */
	xpl    float64       /* polar motion xp wrt local meridian (radians) */
	ypl    float64       /* polar motion yp wrt local meridian (radians) */
	sphi   float64       /* sine of geodetic latitude */
	cphi   float64       /* cosine of geodetic latitude */
	diurab float64       /* magnitude of diurnal aberration vector */
	eral   float64       /* "local" Earth rotation angle (radians) */
	refa   float64       /* refraction constant A (radians) */
	refb   float64       /* refraction constant B (radians) */
}

/* (Vectors eb, eh, em and v are all with respect to BCRS axes.) */

/* Body parameters for light deflection */
type LDBODY struct {
	bm float64       /* mass of the body (solar masses) */
	dl float64       /* deflection limiter (radians^2/2) */
	pv [2][3]float64 /* barycentric PV of the body (au, au/day) */
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
