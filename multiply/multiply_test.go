package multiply

import (
	"testing"
	"fmt"
)

func TestExponent(t *testing.T){
	x := 2
	c := 10
	r := Exponent(x, c)

	fmt.Printf("2^10 = %v\n", r)
}

func TestExponentMod(t *testing.T){
	x := 2
	c := 10
	n := 3
	r := ExponentMod(x, c, n)

	fmt.Printf("2^10 mod 3 = %v\n", r)
}
