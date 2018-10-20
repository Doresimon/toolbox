package main

import (
	"fmt"
	"./euclidean"
)

func main(){
	a := 1000000
	b := 1

	fmt.Printf("gcd(%v,%v) = %v \n", a, b, euclidean.F_gcd(a, b))



	r, s, t := euclidean.F_gcd_extend(a, b)
	fmt.Printf("gcd(%v,%v) = %v, %v*a+%v*b = %v \n",a, b, r, s, t, r )
}