package operator

type RSAKey struct {
	N   *BigData
	Key *BigData
}

func GenRSAKey(p *BigData, q *BigData, e *BigData) (*RSAKey, *RSAKey) {
	n := Multiple(p, q)
	fy := Multiple(Subtract(p, &BigData{val: []byte{'1'}}), Subtract(q, &BigData{val: []byte{'1'}}))
	a := e
	b := fy
	x, _ := ExtGcd(a, b)
	if x.val[0] == '-' {
		x = Add(x, fy)
	}
	d := x

	return &RSAKey{N: n, Key: e}, &RSAKey{N: n, Key: d}
}

func RSAEncrypt(m *BigData, e *BigData, n *BigData) *BigData {
	return FastPow(m, e, n)
}

func RSADecrypt(c *BigData, d *BigData, n *BigData) *BigData {
	return FastPow(c, d, n)
}
