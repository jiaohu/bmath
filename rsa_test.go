package bmath

import (
	"bmath/bigint"
	"fmt"
	"testing"
)

func TestGenRSAKey(t *testing.T) {
	p, _ := bigint.NewInt("7001587")
	q, _ := bigint.NewInt("4002821")
	e, _ := bigint.NewInt("65537")
	pub, self := GenRSAKey(p, q, e)
	fmt.Println(pub.Key)
	fmt.Println(self.Key)
	fmt.Println(pub.N)
}

func TestRSAEncrypt(t *testing.T) {
	m, _ := bigint.NewInt("179")
	e, _ := bigint.NewInt("2566681157393")
	n, _ := bigint.NewInt("28026099476927")
	res := RSAEncrypt(m, e, n)
	fmt.Println(res)
	pub, _ := bigint.NewInt("65537")
	ans := RSADecrypt(res, pub, n)
	fmt.Println(ans)
}
