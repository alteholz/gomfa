package gomfa

import (
	"fmt"
	"math"
)

func CheckFloat(value float64, valueok float64, delta float64) bool {
	/*
	 **  - - - -
	 **   CompareFloat
	 **  - - - -
	 **
	 **  Validate a double result.
	 **
	 **  Given:
	 **     value    float64       value computed by function under test
	 **     valueok  float64       expected value
	 **     delta	float64       maximum allowable error
	 **
	 **  Given and returned:
	 **     status   int          set to TRUE if test fails
	 **
	 */
	var a float64

	a = value - valueok
	if a != 0.0 && math.Abs(a) > math.Abs(delta) {
		/* failed */
		var f float64
		f = math.Abs(valueok / a)
		fmt.Printf("CheckFloat failed: want %v got %v (1/%.3g)\n",
			valueok, value, f)

		return false
	} else {
		return true
	}

}
