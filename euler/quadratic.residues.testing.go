package euler

import (
	"../multiply"
)

/*
CheckQuadrcticResidues ...
Complexity = Order((log(p))^3)
*/
func CheckQuadrcticResidues(a int, p int) bool {
	// check p is prime and p is odd

	// calculate a^((p-1)/2) mod p
	r := multiply.ExponentMod(a, (p-1)/2, p)

	if r == 1 {
		return true
	} else {
		return false
	}
}
