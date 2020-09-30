package bigint

// operator +
func Add(str *BigInt, str2 *BigInt) *BigInt {
	isStrBigThanZero := str.check()
	isStr2BigThanZero := str2.check()
	if isStrBigThanZero && isStr2BigThanZero {
		return str.add(str2)
	} else if !isStrBigThanZero && isStr2BigThanZero {
		temp := &BigInt{val: str.val[1:]}
		return str2.subtract(temp)
	} else if isStrBigThanZero && !isStr2BigThanZero {
		temp := &BigInt{val: str2.val[1:]}
		return str.subtract(temp)
	} else {
		res := &BigInt{val: []byte("-")}
		num1 := &BigInt{val: str.val[1:]}
		num2 := &BigInt{val: str.val[1:]}
		temp := num1.add(num2)
		res.val = append(res.val, temp.val...)
		return res
	}
}

// operator -
func Subtract(str *BigInt, str2 *BigInt) *BigInt {
	isStrBigThanZero := str.check()
	isStr2BigThanZero := str2.check()
	if isStrBigThanZero && isStr2BigThanZero {
		return str.subtract(str2)
	} else if isStrBigThanZero && !isStr2BigThanZero {
		return str.add(&BigInt{val: str2.val[1:]})
	} else if !isStrBigThanZero && isStr2BigThanZero {
		res := &BigInt{val: []byte("-")}
		num1 := &BigInt{val: str.val[1:]}
		temp := num1.add(str2)
		res.val = append(res.val, temp.val...)
		return res
	} else if !isStrBigThanZero && !isStr2BigThanZero {
		temp := &BigInt{val: str2.val[1:]}
		return temp.subtract(&BigInt{val: str.val[1:]})
	}
	return nil
}

// operator *
func Multiple(str *BigInt, str2 *BigInt) *BigInt {
	isStrBigThanZero := str.check()
	isStr2BigThanZero := str2.check()
	if isStrBigThanZero && isStr2BigThanZero {
		return str.multiple(str2)
	} else if !isStrBigThanZero && isStr2BigThanZero {
		res := &BigInt{val: []byte{'-'}}
		num1 := &BigInt{val: str.val[1:]}
		num2 := num1.multiple(str2)
		res.val = append(res.val, num2.val...)
		return res
	} else if isStrBigThanZero && !isStr2BigThanZero {
		res := &BigInt{val: []byte{'-'}}
		num1 := &BigInt{val: str2.val[1:]}
		num2 := str.multiple(num1)
		res.val = append(res.val, num2.val...)
		return res
	} else if !isStrBigThanZero && !isStr2BigThanZero {
		num1 := &BigInt{val: str.val[1:]}
		num2 := &BigInt{val: str2.val[1:]}
		return num1.multiple(num2)
	}
	return nil
}

// operator /
func Divide(str *BigInt, str2 *BigInt) *BigInt {
	isStrBigThanZero := str.check()
	isStr2BigThanZero := str2.check()
	if isStrBigThanZero && isStr2BigThanZero {
		return str.divide(str2)
	} else if isStrBigThanZero && !isStr2BigThanZero {
		res := &BigInt{val: []byte{'-'}}
		num1 := &BigInt{val: str2.val[1:]}
		num2 := str.divide(num1)
		res.val = append(res.val, num2.val...)
		return res
	} else if !isStrBigThanZero && isStr2BigThanZero {
		res := &BigInt{val: []byte{'-'}}
		num1 := &BigInt{val: str.val[1:]}
		num2 := num1.divide(str2)
		res.val = append(res.val, num2.val...)
		return res
	} else if !isStrBigThanZero && !isStr2BigThanZero {
		num1 := &BigInt{val: str.val[1:]}
		num2 := &BigInt{val: str2.val[1:]}
		return num1.divide(num2)
	}
	return nil
}

// operator %
func Module(str *BigInt, str2 *BigInt) *BigInt {
	isStrBigThanZero := str.check()
	isStr2BigThanZero := str2.check()
	if isStrBigThanZero && isStr2BigThanZero {
		return str.module(str2)
	} else if isStrBigThanZero && !isStr2BigThanZero {
		res := &BigInt{val: []byte{}}
		num1 := &BigInt{val: str2.val[1:]}
		num2 := str.module(num1)
		res.val = append(res.val, num2.val...)
		return res
	} else if !isStrBigThanZero && isStr2BigThanZero {
		res := &BigInt{val: []byte{'-'}}
		num1 := &BigInt{val: str.val[1:]}
		num2 := num1.module(str2)
		res.val = append(res.val, num2.val...)
		return res
	} else if !isStrBigThanZero && !isStr2BigThanZero {
		res := &BigInt{val: []byte{'-'}}
		num1 := &BigInt{val: str.val[1:]}
		num2 := &BigInt{val: str2.val[1:]}
		num3 := num1.module(num2)
		res.val = append(res.val, num3.val...)
		return res
	}
	return nil
}

// fast module (a ^ b) mod n
func FastPow(a *BigInt, b *BigInt, n *BigInt) *BigInt {
	var r *BigInt = &BigInt{val: []byte("0")}
	x := a
	y := &BigInt{val: []byte("1")}
	for b.String() != "0" {
		r = Module(b, &BigInt{val: []byte("2")})
		b = Divide(b, &BigInt{val: []byte("2")})
		if r.String() == "1" {
			y = Module(Multiple(y, x), n)
		}
		x = Multiple(Module(x, n), Module(x, n))
	}
	return y
}
