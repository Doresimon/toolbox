package euclidean

// upper = O(lgn)
func F_gcd(a int, b int) int{
	r := make([]int,64,64)
	q := make([]int,64,64)

	if a>b {
		r[0] = a
		r[1] = b
	}else{
		r[0] = b
		r[1] = a
	}

	m:=1

	for i := 0; r[m]!=0; i++ {
		// if len(r) <= (m+1) {
		// 	r = append(r, make([]int,8,8)...)
		// 	q = append(q, make([]int,8,8)...)
		// }
		q[m] = r[m-1]/r[m]
		r[m+1] = r[m-1]%r[m]
		m++
	}
	m--
	return r[m]
}