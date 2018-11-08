package lsss

import (
	"crypto/rand"
	"math/big"
)

// Share is D(xi) = (xi, f(xi))
type Share struct {
	x    *big.Int
	v    *big.Int
	base *big.Int
}

// Lsss is the scheme
type Lsss struct {
	base   *big.Int
	secret *big.Int
	share  []*Share
	coef   []*big.Int
	t      int
	n      int
}

// Zero is 0 of big.Int
var Zero = big.NewInt(0)

// SetBase .
func (l *Lsss) SetBase(b *big.Int) {
	l.base.Set(b)
}

// SetSecret .
func (l *Lsss) SetSecret(s *big.Int) {
	l.secret.Set(s)
}

// RandSecret .
func (l *Lsss) RandSecret() {
	s, _ := rand.Int(rand.Reader, l.base)
	l.secret.Set(s)
}

// GenShares .
func (l *Lsss) GenShares(t int, n int) []*Share {
	l.t = t
	l.n = n
	l.coef = make([]*big.Int, t, t)
	l.share = make([]*Share, n, n)

	l.coef[0] = l.secret

	for i := 1; i < t; i++ {
		for {
			a, _ := rand.Int(rand.Reader, l.base)
			if a.Cmp(Zero) != 0 {
				l.coef[i] = new(big.Int)
				l.coef[i].Set(a)
				break
			}
		}
	}

	for i := 0; i < n; i++ {
		x := big.NewInt(int64(i + 1))
		v := l.calShare(x)

		l.share[i] = &Share{x, v, l.base}
	}

	return l.share

}

// RecoverShare .
func (l *Lsss) RecoverShare(t int, n int) {

}

func (l *Lsss) calShare(x *big.Int) *big.Int {
	tmpX := new(big.Int)
	tmpA := new(big.Int)
	sum := new(big.Int)

	tmpX.Set(x)
	tmpA.SetInt64(0)
	sum.Set(l.secret)

	for i := 1; i < l.t; i++ {
		tmpA.Set(l.coef[i])
		sum.Add(sum, tmpA.Mul(tmpA, tmpX))

		tmpX.Mul(tmpX, x)
	}

	sum.Mod(sum, l.base)

	return sum
}

// NewLsss .
func NewLsss(base *big.Int) *Lsss {
	ss := new(Lsss)
	ss.base = new(big.Int)
	ss.secret = new(big.Int)
	ss.SetBase(base)
	ss.RandSecret()
	return ss
}

// CalPartOfSecret .
// for party[i]. use it's share to calculate fi(0)
func CalPartOfSecret(share *Share, Co []*big.Int) *big.Int {
	up := new(big.Int)
	down := new(big.Int)
	tmp := new(big.Int)

	up.SetInt64(1)
	down.SetInt64(1)

	for _, x := range Co {
		if x.Cmp(share.x) != 0 {
			tmp.Set(x)
			tmp.Sub(tmp, share.x)

			up.Mul(up, x)
			down.Mul(down, tmp)
		}
	}
	down.ModInverse(down, share.base)
	up.Mul(up, share.v)
	up.Mul(up, down)
	up.Mod(up, share.base)
	return up
}
