package lsss

import (
	"math/big"
	"testing"
)

// $> go test -timeout 30s -run ^TestLsss$ -v
func TestLsss(t *testing.T) {
	// notice that the base must be a prime...
	base := new(big.Int)
	base.SetString("179426549", 10)

	ss := NewLsss(base)
	threshold := 15
	n := 21

	t.Logf("\n base: %v\n", ss.base)
	t.Logf("\n secret: %v\n", ss.secret)

	shares := ss.GenShares(threshold, n)

	t.Logf("\n coef: %#v\n", ss.coef)
	t.Logf("\n shares: %#v\n", shares[0].x)
	t.Logf("\n shares: %#v\n", shares[1].x)
	t.Logf("\n shares: %#v\n", shares[2].x)

	// here we set the number of parties
	// to recover the secret
	CoNum := 20
	Co := make([]*big.Int, CoNum, CoNum)

	for i := 0; i < CoNum; i++ {
		Co[i] = big.NewInt(int64(i + 1))
	}

	sum := big.NewInt(0)

	for i := range Co {
		sum.Add(sum, CalPartOfSecret(shares[i], Co))
	}

	sum.Mod(sum, base)
	t.Logf("\n sum: %v\n", sum)
	t.Logf("\n result: %v\n", sum.Cmp(ss.secret) == 0)

}
