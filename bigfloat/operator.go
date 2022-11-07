package bigfloat

// Add operate +
func Add(str *BigFloat, str2 *BigFloat) *BigFloat {
	isStrBigThanZero := str.Check()
	isStr2BigThanZero := str2.Check()
	if isStrBigThanZero && isStr2BigThanZero {
		return str.Add(str2)
	} else if !isStrBigThanZero && isStr2BigThanZero {
		temp := &BigFloat{val: str.val[1:]}
		return str2.Subtract(temp)
	} else if isStrBigThanZero && !isStr2BigThanZero {
		temp := &BigFloat{val: str2.val[1:]}
		return str.Subtract(temp)
	} else {
		res := &BigFloat{val: []byte("-")}
		num1 := &BigFloat{val: str.val[1:]}
		num2 := &BigFloat{val: str.val[1:]}
		temp := num1.Add(num2)
		res.val = append(res.val, temp.val...)
		return res
	}
}

// Subtract operator -
func Subtract(str *BigFloat, str2 *BigFloat) *BigFloat {
	isStrBigThanZero := str.Check()
	isStr2BigThanZero := str2.Check()
	if isStrBigThanZero && isStr2BigThanZero {
		return str.Subtract(str2)
	} else if isStrBigThanZero && !isStr2BigThanZero {
		return str.Add(&BigFloat{val: str2.val[1:]})
	} else if !isStrBigThanZero && isStr2BigThanZero {
		res := &BigFloat{val: []byte("-")}
		num1 := &BigFloat{val: str.val[1:]}
		temp := num1.Add(str2)
		res.val = append(res.val, temp.val...)
		return res
	} else if !isStrBigThanZero && !isStr2BigThanZero {
		temp := &BigFloat{val: str2.val[1:]}
		return temp.Subtract(&BigFloat{val: str.val[1:]})
	}
	return nil
}

// Multiple operator *
func Multiple(str *BigFloat, str2 *BigFloat) *BigFloat {
	isStrBigThanZero := str.Check()
	isStr2BigThanZero := str2.Check()
	if isStrBigThanZero && isStr2BigThanZero {
		return str.Multiple(str2)
	} else if !isStrBigThanZero && isStr2BigThanZero {
		res := &BigFloat{val: []byte{'-'}}
		num1 := &BigFloat{val: str.val[1:]}
		num2 := num1.Multiple(str2)
		res.val = append(res.val, num2.val...)
		return res
	} else if isStrBigThanZero && !isStr2BigThanZero {
		res := &BigFloat{val: []byte{'-'}}
		num1 := &BigFloat{val: str2.val[1:]}
		num2 := str.Multiple(num1)
		res.val = append(res.val, num2.val...)
		return res
	} else if !isStrBigThanZero && !isStr2BigThanZero {
		num1 := &BigFloat{val: str.val[1:]}
		num2 := &BigFloat{val: str2.val[1:]}
		return num1.Multiple(num2)
	}
	return nil
}
