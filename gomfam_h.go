package gomfa

import (
	"math"
)

/*
**  - - - - - - -
**    gomfam . h
**  - - - - - - -
**
**  Macros used by GoMFA library.
 */

/* Pi */
const DPI = (3.141592653589793238462643)

/* 2Pi */
const D2PI = (6.283185307179586476925287)

/* Radians to degrees */
const DR2D = (57.29577951308232087679815)

/* Degrees to radians */
const DD2R = (1.745329251994329576923691e-2)

/* Radians to arcseconds */
const DR2AS = (206264.8062470963551564734)

/* Arcseconds to radians */
const DAS2R = (4.848136811095359935899141e-6)

/* Seconds of time to radians */
const DS2R = (7.272205216643039903848712e-5)

/* Arcseconds in a full circle */
const TURNAS = (1296000.0)

/* Milliarcseconds to radians */
const DMAS2R = (DAS2R / 1e3)

/* Length of tropical year B1900 (days) */
const DTY = (365.242198781)

/* Seconds per day. */
const DAYSEC = (86400.0)

/* Days per Julian year */
const DJY = (365.25)

/* Days per Julian century */
const DJC = (36525.0)

/* Days per Julian millennium */
const DJM = (365250.0)

/* Reference epoch (J2000.0), Julian Date */
const DJ00 = (2451545.0)

/* Julian Date of Modified Julian Date zero */
const DJM0 = (2400000.5)

/* Reference epoch (J2000.0), Modified Julian Date */
const DJM00 = (51544.5)

/* 1977 Jan 1.0 as MJD */
const DJM77 = (43144.0)

/* TT minus TAI (s) */
const TTMTAI = (32.184)

/* Astronomical unit (m, IAU 2012) */
const DAU = (149597870.7e3)

/* Speed of light (m/s) */
const CMPS = 299792458.0

/* Light time for 1 au (s) */
const AULT = (DAU / CMPS)

/* Speed of light (au per day) */
const DC = (DAYSEC / AULT)

/* L_G = 1 - d(TT)/d(TCG) */
const ELG = (6.969290134e-10)

/* L_B = 1 - d(TDB)/d(TCB), and TDB (s) at TAI 1977/1/1.0 */
const ELB = (1.550519768e-8)
const TDB0 = (-6.55e-5)

/* Schwarzschild radius of the Sun (au) */
/* = 2 * 1.32712440041e20 / (2.99792458e8)^2 / 1.49597870700e11 */
const SRS = 1.97412574336e-8

//XXX /* DINT(A) - truncate to nearest whole number towards zero (double) */
//XXX const DINT(A) = ((A)<0.0?ceil(A):floor(A))
//XXX
//XXX /* DNINT(A) - round to nearest whole number (double) */
//XXX const DNINT(A) = (fabs(A)<0.5?0.0\
//XXX                                 :((A)<0.0?ceil((A)-0.5):floor((A)+0.5)))
//XXX
/* DSIGN(A,B) - magnitude of A with sign of B (double) */
func DSIGN(a float64, b float64) float64 {
	if b < 0 {
		return -math.Abs(a)
	} else {
		return math.Abs(a)
	}
}

//XXX
//XXX /* max(A,B) - larger (most +ve) of two numbers (generic) */
//XXX const GMAX(A,B) = (((A)>(B))?(A):(B))
//XXX
//XXX /* min(A,B) - smaller (least +ve) of two numbers (generic) */
//XXX const GMIN(A,B) = (((A)<(B))?(A):(B))

/* Reference ellipsoids */
const WGS84 = 1
const GRS80 = 2
const WGS72 = 3

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
