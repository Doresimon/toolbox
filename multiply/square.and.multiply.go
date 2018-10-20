package multiply

/*
return x^c
*/
func Exponent(x int, c int) int {
	b := 0
	e := x
	s := 1

	for true {
		b = c & 0x1

		if b == 1 {
			s *= e
		}

		e *= e
		c = c >> 1

		if c == 0 {
			break
		}
	}

	return s
}

/*
return x^c % n
*/
func ExponentMod(x int, c int, n int) int {
	b := 0
	e := x
	s := 1

	for true {
		b = c & 0x1

		if b == 1 {
			s *= e
			s = s % n
		}

		e *= e
		e = e % n
		c = c >> 1

		if c == 0 {
			break
		}
	}

	return s
}
