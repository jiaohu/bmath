package bigfloat

import (
	"errors"
	"github.com/jiaohu/bmath/base"
	"github.com/jiaohu/bmath/bigint"
	"regexp"
)

type BigFloat struct {
	val      []byte
	accuracy int
	neg      bool
}

func NewFloat(str string) (*BigFloat, error) {
	ok, err := regexp.MatchString("^(\\-)?(([0]{1})|([1-9][0-9]*))(\\.){1}[0-9]*[1-9]$", str)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, errors.New("invalid number")
	}
	origin := []byte(str)
	var accuracy int
	var flag bool = false
	if origin[0] == base.NegativeFlag {
		flag = true
	}

	for i, v := range origin {
		if v == base.NumberPoint {
			accuracy = len(origin) - i - 1
			break
		}
	}
	return &BigFloat{val: origin, accuracy: accuracy, neg: flag}, nil
}

func (b *BigFloat) String() string {
	if b == nil {
		return "0.0"
	}
	return string(b.val)
}

func (b *BigFloat) IsNegative() bool {
	return b.neg
}

func (b *BigFloat) check() bool {
	if b.val[0] == '-' {
		return false
	}
	return true
}

func (b *BigFloat) add(str *BigFloat) *BigFloat {
	originLen := len(b.val)
	addLen := len(str.val)
	var resultMaxLen int
	if originLen > addLen {
		resultMaxLen = originLen + 1
	} else {
		resultMaxLen = addLen + 1
	}
	var (
		ptr1           = originLen - 1
		ptr2           = addLen - 1
		index      int = resultMaxLen - 1
		firstIndex int = resultMaxLen - 1
		accuracy   int
	)
	result := make([]byte, resultMaxLen)
	if b.accuracy > str.accuracy {
		for i := 1; i <= (b.accuracy - str.accuracy); i++ {
			result[index] = b.val[ptr1]
			ptr1--
			index--
		}
	} else if b.accuracy < str.accuracy {
		for i := 1; i <= (str.accuracy - b.accuracy); i++ {
			result[index] = str.val[ptr2]
			ptr2--
			index--
		}
	}

	var carry int
	for {
		if ptr1 < 0 || ptr2 < 0 {
			break
		}
		if b.val[ptr1] == base.NumberPoint && str.val[ptr2] == base.NumberPoint {
			result[index] = b.val[ptr1]
			ptr1--
			ptr2--
			index--
			continue
		}

		temp := int(b.val[ptr1]-48) + int(str.val[ptr2]-48) + carry

		if temp >= 10 {
			carry = temp / 10
		} else {
			carry = 0
		}
		result[index] = byte(temp%10 + 48)
		ptr1--
		ptr2--
		index--
	}
	if ptr1 < 0 && ptr2 >= 0 {
		for i := ptr2; i >= 0; i-- {
			temp := int(str.val[ptr2]-48) + carry
			if temp >= 10 {
				carry = temp / 10
			} else {
				carry = 0
			}
			result[index] = byte(temp%10 + 48)
			index--
			ptr2--
		}
	} else if ptr1 >= 0 && ptr2 < 0 {
		for i := ptr1; i >= 0; i-- {
			temp := int(b.val[ptr1]-48) + carry
			if temp >= 10 {
				carry = temp / 10
			} else {
				carry = 0
			}
			result[index] = byte(temp%10 + 48)
			index--
			ptr1--
		}
	}
	if carry != 0 {
		result[index] = byte(48 + carry)
	}
	for i, v := range result {
		if v == base.NumberPoint {
			accuracy = len(result) - i - 1
			break
		}
	}
	for i, c := range result {
		if c != 0 {
			firstIndex = i
			break
		}
	}
	return &BigFloat{val: result[firstIndex:], accuracy: accuracy, neg: false}
}

func (b *BigFloat) subtract(str *BigFloat) *BigFloat {
	var acc int = b.accuracy
	if b.accuracy > str.accuracy {
		for i := 1; i <= (b.accuracy - str.accuracy); i++ {
			str.val = append(str.val, 48)
		}
		acc = b.accuracy
		str.accuracy = b.accuracy
	} else if str.accuracy > b.accuracy {
		for i := 1; i <= (str.accuracy - b.accuracy); i++ {
			b.val = append(b.val, 48)
		}
		acc = str.accuracy
		b.accuracy = str.accuracy
	}
	result := &BigFloat{
		val:      nil,
		accuracy: acc,
		neg:      false,
	}
	isBig := b.compare(str)
	if isBig {
		temp := b.sub(str)
		result.val = temp.val
	} else {
		result.val = []byte{'-'}
		temp := str.sub(b)
		result.neg = true
		result.val = append(result.val, temp.val...)
	}
	return result
}

func (b *BigFloat) sub(str *BigFloat) *BigFloat {
	originLen := len(b.val)
	strLen := len(str.val)
	var (
		ptr1   int = originLen - 1
		ptr2   int = strLen - 1
		borrow int
		index  int    = originLen - 1
		result []byte = make([]byte, originLen)
	)

	for {
		if ptr1 < 0 || ptr2 < 0 {
			break
		}
		if b.val[ptr1] == base.NumberPoint {
			result[index] = b.val[ptr1]
			index--
			ptr1--
			ptr2--
			continue
		}
		temp := int(b.val[ptr1]) - borrow - int(str.val[ptr2])
		if temp < 0 && borrow == 0 {
			borrow = 1
			temp = int(b.val[ptr1]) + borrow*10 - int(str.val[ptr2])
		} else if temp < 0 && borrow == 1 {
			temp = int(b.val[ptr1]) + borrow*10 - 1 - int(str.val[ptr2])
			borrow = 1
		} else {
			borrow = 0
		}
		result[index] = byte(temp + 48)
		index--
		ptr1--
		ptr2--
	}

	if ptr1 >= 0 {
		for i := ptr1; i >= 0; i-- {
			temp := int(b.val[i]-48) - borrow
			if temp < 0 && borrow == 0 {
				borrow = 1
				temp = borrow * 10
			} else if temp < 0 && borrow != 0 {
				temp = borrow*10 - 1
				borrow = 1
			} else {
				borrow = 0
			}
			result[index] = byte(temp + 48)
			index--
		}
	}

	return &BigFloat{
		val: result,
		neg: false,
	}
}

func (b *BigFloat) multiple(str *BigFloat) *BigFloat {
	acc := b.accuracy + str.accuracy
	var (
		first  []byte
		second []byte
		res    []byte
	)
	for _, v := range b.val {
		if v == base.NumberPoint {
			continue
		}
		first = append(first, v)
	}
	for _, v := range str.val {
		if v == base.NumberPoint {
			continue
		}
		second = append(second, v)
	}
	param1, _ := bigint.NewInt(string(first))
	param2, _ := bigint.NewInt(string(second))
	result := bigint.Multiple(param1, param2)
	for i, v := range result.GetValue() {
		if i == len(result.GetValue())-acc {
			res = append(res, base.NumberPoint)
		}
		res = append(res, v)
	}
	return &BigFloat{
		val:      res,
		accuracy: acc,
		neg:      false,
	}
}

func (b *BigFloat) divide(str *BigFloat) *BigFloat {
	return nil
}

func (b *BigFloat) compare(str *BigFloat) bool {
	if b.accuracy > str.accuracy {
		for i := 1; i <= (b.accuracy - str.accuracy); i++ {
			str.val = append(str.val, 48)
		}
		str.accuracy = b.accuracy
	} else if str.accuracy > b.accuracy {
		for i := 1; i <= (str.accuracy - b.accuracy); i++ {
			b.val = append(b.val, 48)
		}
		b.accuracy = str.accuracy
	}
	originLen := len(b.val)
	strLen := len(str.val)
	if originLen-b.accuracy > strLen-str.accuracy {
		return true
	} else if originLen-b.accuracy < strLen-str.accuracy {
		return false
	} else {
		var flag bool = true
		for i := 0; i < originLen; i++ {
			if b.val[i] > str.val[i] {
				break
			} else if b.val[i] == str.val[i] {
				continue
			} else if b.val[i] < str.val[i] {
				flag = false
				break
			}
		}
		return flag
	}
}
