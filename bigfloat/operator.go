package bigfloat

func Add(str *BigFloat, str2 *BigFloat) *BigFloat {
	isStrBigThanZero := str.check()
	isStr2BigThanZero := str2.check()
	if isStrBigThanZero && isStr2BigThanZero {
		return str.add(str2)
	} else if !isStrBigThanZero && isStr2BigThanZero {
		temp := &BigFloat{val: str.val[1:]}
		return str2.subtract(temp)
	} else if isStrBigThanZero && !isStr2BigThanZero {
		temp := &BigFloat{val: str2.val[1:]}
		return str.subtract(temp)
	} else {
		res := &BigFloat{val: []byte("-")}
		num1 := &BigFloat{val: str.val[1:]}
		num2 := &BigFloat{val: str.val[1:]}
		temp := num1.add(num2)
		res.val = append(res.val, temp.val...)
		return res
	}
}


// operator -
func Subtract(str *BigFloat, str2 *BigFloat) *BigFloat {
	isStrBigThanZero := str.check()
	isStr2BigThanZero := str2.check()
	if isStrBigThanZero && isStr2BigThanZero {
		return str.subtract(str2)
	} else if isStrBigThanZero && !isStr2BigThanZero {
		return str.add(&BigFloat{val: str2.val[1:]})
	} else if !isStrBigThanZero && isStr2BigThanZero {
		res := &BigFloat{val: []byte("-")}
		num1 := &BigFloat{val: str.val[1:]}
		temp := num1.add(str2)
		res.val = append(res.val, temp.val...)
		return res
	} else if !isStrBigThanZero && !isStr2BigThanZero {
		temp := &BigFloat{val: str2.val[1:]}
		return temp.subtract(&BigFloat{val: str.val[1:]})
	}
	return nil
}

// operator *
func Multiple(str *BigFloat, str2 *BigFloat) *BigFloat {
	isStrBigThanZero := str.check()
	isStr2BigThanZero := str2.check()
	if isStrBigThanZero && isStr2BigThanZero {
		return str.multiple(str2)
	} else if !isStrBigThanZero && isStr2BigThanZero {
		res := &BigFloat{val: []byte{'-'}}
		num1 := &BigFloat{val: str.val[1:]}
		num2 := num1.multiple(str2)
		res.val = append(res.val, num2.val...)
		return res
	} else if isStrBigThanZero && !isStr2BigThanZero {
		res := &BigFloat{val: []byte{'-'}}
		num1 := &BigFloat{val: str2.val[1:]}
		num2 := str.multiple(num1)
		res.val = append(res.val, num2.val...)
		return res
	} else if !isStrBigThanZero && !isStr2BigThanZero {
		num1 := &BigFloat{val: str.val[1:]}
		num2 := &BigFloat{val: str2.val[1:]}
		return num1.multiple(num2)
	}
	return nil
}
