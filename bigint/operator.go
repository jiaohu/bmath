package bigint

// operator +
func Add(str *BigInt, str2 *BigInt) *BigInt {
	isStrBigThanZero := str.Check()
	isStr2BigThanZero := str2.Check()
	if isStrBigThanZero && isStr2BigThanZero {
		return str.Add(str2)
	} else if !isStrBigThanZero && isStr2BigThanZero {
		temp := &BigInt{val: str.val[1:]}
		return str2.Subtract(temp)
	} else if isStrBigThanZero && !isStr2BigThanZero {
		temp := &BigInt{val: str2.val[1:]}
		return str.Subtract(temp)
	} else {
		res := &BigInt{val: []byte("-")}
		num1 := &BigInt{val: str.val[1:]}
		num2 := &BigInt{val: str.val[1:]}
		temp := num1.Add(num2)
		res.val = append(res.val, temp.val...)
		return res
	}
}

// operator -
func Subtract(str *BigInt, str2 *BigInt) *BigInt {
	isStrBigThanZero := str.Check()
	isStr2BigThanZero := str2.Check()
	if isStrBigThanZero && isStr2BigThanZero {
		return str.Subtract(str2)
	} else if isStrBigThanZero && !isStr2BigThanZero {
		return str.Add(&BigInt{val: str2.val[1:]})
	} else if !isStrBigThanZero && isStr2BigThanZero {
		res := &BigInt{val: []byte("-")}
		num1 := &BigInt{val: str.val[1:]}
		temp := num1.Add(str2)
		res.val = append(res.val, temp.val...)
		return res
	} else if !isStrBigThanZero && !isStr2BigThanZero {
		temp := &BigInt{val: str2.val[1:]}
		return temp.Subtract(&BigInt{val: str.val[1:]})
	}
	return nil
}

// operator *
func Multiple(str *BigInt, str2 *BigInt) *BigInt {
	isStrBigThanZero := str.Check()
	isStr2BigThanZero := str2.Check()
	if isStrBigThanZero && isStr2BigThanZero {
		return str.Multiple(str2)
	} else if !isStrBigThanZero && isStr2BigThanZero {
		res := &BigInt{val: []byte{'-'}}
		num1 := &BigInt{val: str.val[1:]}
		num2 := num1.Multiple(str2)
		res.val = append(res.val, num2.val...)
		return res
	} else if isStrBigThanZero && !isStr2BigThanZero {
		res := &BigInt{val: []byte{'-'}}
		num1 := &BigInt{val: str2.val[1:]}
		num2 := str.Multiple(num1)
		res.val = append(res.val, num2.val...)
		return res
	} else if !isStrBigThanZero && !isStr2BigThanZero {
		num1 := &BigInt{val: str.val[1:]}
		num2 := &BigInt{val: str2.val[1:]}
		return num1.Multiple(num2)
	}
	return nil
}

// operator /
func Divide(str *BigInt, str2 *BigInt) *BigInt {
	isStrBigThanZero := str.Check()
	isStr2BigThanZero := str2.Check()
	if isStrBigThanZero && isStr2BigThanZero {
		return str.Divide(str2)
	} else if isStrBigThanZero && !isStr2BigThanZero {
		res := &BigInt{val: []byte{'-'}}
		num1 := &BigInt{val: str2.val[1:]}
		num2 := str.Divide(num1)
		res.val = append(res.val, num2.val...)
		return res
	} else if !isStrBigThanZero && isStr2BigThanZero {
		res := &BigInt{val: []byte{'-'}}
		num1 := &BigInt{val: str.val[1:]}
		num2 := num1.Divide(str2)
		res.val = append(res.val, num2.val...)
		return res
	} else if !isStrBigThanZero && !isStr2BigThanZero {
		num1 := &BigInt{val: str.val[1:]}
		num2 := &BigInt{val: str2.val[1:]}
		return num1.Divide(num2)
	}
	return nil
}

// operator %
func Module(str *BigInt, str2 *BigInt) *BigInt {
	isStrBigThanZero := str.Check()
	isStr2BigThanZero := str2.Check()
	if isStrBigThanZero && isStr2BigThanZero {
		return str.Mod(str2)
	} else if isStrBigThanZero && !isStr2BigThanZero {
		res := &BigInt{val: []byte{}}
		num1 := &BigInt{val: str2.val[1:]}
		num2 := str.Mod(num1)
		res.val = append(res.val, num2.val...)
		return res
	} else if !isStrBigThanZero && isStr2BigThanZero {
		res := &BigInt{val: []byte{'-'}}
		num1 := &BigInt{val: str.val[1:]}
		num2 := num1.Mod(str2)
		res.val = append(res.val, num2.val...)
		return res
	} else if !isStrBigThanZero && !isStr2BigThanZero {
		res := &BigInt{val: []byte{'-'}}
		num1 := &BigInt{val: str.val[1:]}
		num2 := &BigInt{val: str2.val[1:]}
		num3 := num1.Mod(num2)
		res.val = append(res.val, num3.val...)
		return res
	}
	return nil
}

// fast Mod (a ^ b) mod n
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
