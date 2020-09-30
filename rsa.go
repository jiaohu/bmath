package bmath

import (
	"github.com/jiaohu/bmath/bigint"
)

type RSAKey struct {
	N   *bigint.BigInt
	Key *bigint.BigInt
}

func GenRSAKey(p *bigint.BigInt, q *bigint.BigInt, e *bigint.BigInt) (*RSAKey, *RSAKey) {
	n := bigint.Multiple(p, q)
	one, _ := bigint.NewInt("1")
	fy := bigint.Multiple(bigint.Subtract(p, one), bigint.Subtract(q, one))
	a := e
	b := fy
	x, _ := bigint.ExtGcd(a, b)
	if x.GetValue()[0] == '-' {
		x = bigint.Add(x, fy)
	}
	d := x

	return &RSAKey{N: n, Key: e}, &RSAKey{N: n, Key: d}
}

func RSAEncrypt(m *bigint.BigInt, e *bigint.BigInt, n *bigint.BigInt) *bigint.BigInt {
	return bigint.FastPow(m, e, n)
}

func RSADecrypt(c *bigint.BigInt, d *bigint.BigInt, n *bigint.BigInt) *bigint.BigInt {
	return bigint.FastPow(c, d, n)
}
