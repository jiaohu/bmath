package operator

func Gcd(a *BigData, b *BigData) *BigData {
	if b.ConvertString() == "0" {
		return a
	} else {
		return Gcd(b, Module(a, b))
	}
}

//扩展欧几里的算法
//计算 ax + by = 1中的x与y的整数解（a与b互质）
func ExtGcd(a *BigData, b *BigData) (*BigData, *BigData) {
	var (
		x *BigData = &BigData{}
		y *BigData = &BigData{}
	)
	if b.ConvertString() == "0" {
		return &BigData{val: []byte{'1'}}, &BigData{val: []byte{'0'}}
	} else {
		tempMod := Module(a, b)
		x, y = ExtGcd(b, tempMod)
		x, y = y, Subtract(x, Multiple(Divide(a, b), y))
		return x, y
	}
}

//求模逆
//ed mod n = 1, 其中e与n互质, 且已知
func ModuloInverse(e *BigData, d *BigData) *BigData {
	return d
}
