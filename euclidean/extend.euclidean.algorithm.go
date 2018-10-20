package euclidean

func F_gcd_extend(a int, b int) (int,int,int){
	_a := 0
	_b := 0
	_t := 0
	_s := 0

	if a>b {
		_a = a
		_b = b
	}else{
		_a = b
		_b = a
	}

	_t = 0
	_s = 1
	t := 1
	s := 0

	q:=_a/_b
	r:=_a%_b

	for i := 0; r>0; i++ {
		// tmp := _t - q*t
		// _t = t
		// t = tmp
		_t, t = t, _t - q*t
		_s, s = s, _s - q*s
		_a, _b, q = _b, r, _a/_b
		r = _a - q*_b
	}
	r = _b
	return r,s,t
}