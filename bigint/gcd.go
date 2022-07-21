package bigint

func Gcd(a *BigInt, b *BigInt) *BigInt {
	if b.String() == "0" {
		return a
	} else {
		return Gcd(b, Module(a, b))
	}
}

// ExtGcd
//扩展欧几里的算法
//计算 ax + by = 1中的x与y的整数解（a与b互质）
func ExtGcd(a *BigInt, b *BigInt) (*BigInt, *BigInt) {
	var (
		x *BigInt = &BigInt{}
		y *BigInt = &BigInt{}
	)
	if b.String() == "0" {
		return &BigInt{val: []byte{'1'}}, &BigInt{val: []byte{'0'}}
	} else {
		tempMod := Module(a, b)
		x, y = ExtGcd(b, tempMod)
		x, y = y, Subtract(x, Multiple(Divide(a, b), y))
		return x, y
	}
}

// ModuloInverse
/// 求模逆
/// ed mod n = 1, 其中e与n互质, 且已知
func ModuloInverse(e *BigInt, d *BigInt) *BigInt {
	return d
}
