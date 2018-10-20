package euclidean

import (
	"testing"
	"fmt"
)

func TestEuclidean(t *testing.T) {
	a := 1000000
	b := 1

	fmt.Printf("gcd(%v,%v) = %v \n", a, b, F_gcd(a, b))
}
func TestEuclideanExtend(tt *testing.T) {
	a := 1000000
	b := 1

	r, s, t := F_gcd_extend(a, b)
	fmt.Printf("gcd(%v,%v) = %v, %v*a+%v*b = %v \n",a, b, r, s, t, r )
}