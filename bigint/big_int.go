package bigint

import (
	"errors"
	"regexp"
)

type BigInt struct {
	val []byte
}

func NewInt(str string) (*BigInt, error) {
	ok, err := regexp.MatchString("(\\-)?[1-9][0-9]*", str)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, errors.New("invalid number")
	}
	return &BigInt{val: []byte(str)}, nil
}

func (b *BigInt) String() string {
	if b == nil {
		return "0"
	}
	return string(b.val)
}

func (b *BigInt) check() bool {
	if b.val[0] == '-' {
		return false
	}
	return true
}

func (b *BigInt) add(str *BigInt) *BigInt {
	originLen := len(b.val)
	addLen := len(str.val)
	var resultMaxLen int
	if originLen > addLen {
		resultMaxLen = originLen + 1
	} else {
		resultMaxLen = addLen + 1
	}
	result := make([]byte, resultMaxLen)
	var (
		ptr1       int = originLen - 1
		ptr2       int = addLen - 1
		carry      int
		index      int = resultMaxLen - 1
		firstIndex int = resultMaxLen - 1
	)
	for {
		if ptr1 < 0 || ptr2 < 0 {
			break
		}
		temp := int(b.val[ptr1]-48) + int(str.val[ptr2]-48) + carry
		if temp >= 10 {
			carry = temp / 10
		} else {
			carry = 0
		}
		result[index] = byte(temp%10 + 48)
		index--
		ptr1--
		ptr2--
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

	for i, c := range result {
		if c != 0 {
			firstIndex = i
			break
		}
	}

	return &BigInt{val: result[firstIndex:]}
}

func (b *BigInt) subtract(str *BigInt) *BigInt {
	var isBig bool
	isBig = b.compare(str)
	result := &BigInt{}
	if isBig {
		result = b.sub(b, str)
	} else {
		result.val = []byte{'-'}
		temp := b.sub(str, b)
		result.val = append(result.val, temp.val...)
	}

	return result
}

func (b *BigInt) sub(str *BigInt, str2 *BigInt) *BigInt {
	strLen := len(str.val)
	str2Len := len(str2.val)
	var (
		ptr1      int = strLen - 1
		ptr2      int = str2Len - 1
		borrow    int
		result    []byte = make([]byte, strLen)
		index     int    = strLen - 1
		firstZero int    = strLen - 1
	)
	for {
		if ptr1 < 0 || ptr2 < 0 {
			break
		}
		temp := int(str.val[ptr1]) - borrow - int(str2.val[ptr2])
		if temp < 0 && borrow == 0 {
			borrow = 1
			temp = int(str.val[ptr1]) + borrow*10 - int(str2.val[ptr2])
		} else if temp < 0 && borrow == 1 {
			temp = int(str.val[ptr1]) + borrow*10 - 1 - int(str2.val[ptr2])
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
			temp := int(str.val[i]-48) - borrow
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

	for i, v := range result {
		if v != 48 {
			firstZero = i
			break
		}
	}

	return &BigInt{val: result[firstZero:]}
}

func (b *BigInt) multiple(str *BigInt) *BigInt {
	strLen := len(b.val)
	str2Len := len(str.val)
	temp := make([]int, strLen+str2Len)
	for i, v := range b.val {
		for j, v2 := range str.val {
			temp[strLen-i+str2Len-j-2] += int(v-48) * int(v2-48)
		}
	}

	var carry int
	for i, v := range temp {
		t := v + carry
		temp[i] = t % 10
		carry = t / 10
	}

	var res []byte
	var flag bool
	for i := len(temp) - 1; i >= 0; i-- {
		if temp[i] != 0 {
			flag = true
		}
		if flag {
			res = append(res, byte(temp[i]+48))
		}
	}
	if len(res) == 0 {
		res = []byte{48}
	}
	return &BigInt{val: res}
}

func (b *BigInt) divide(str *BigInt) *BigInt {
	var first byte
	var d *BigInt = &BigInt{}
	if !b.compare(str) {
		return &BigInt{val: []byte{48}}
	}
	for {
		lenDiff := len(b.val) - len(str.val)
		var temp = &BigInt{}
		var temp2 = &BigInt{}
		temp.val = str.val
		temp2.val = str.val
		var realLen int
		for i := 1; i <= lenDiff; i++ {
			temp.val = append(temp.val, byte(48))
			if b.compare(temp) {
				temp2.val = temp.val
				realLen++
			} else {
				break
			}
		}

		b = b.subtract(temp2)
		first = b.val[0]
		if first == '-' {
			break
		} else {
			if realLen != 0 {
				var one = &BigInt{val: []byte{49}}
				var add = one.powBase10(realLen)
				d = d.add(add)
			} else {
				d = d.add(&BigInt{val: []byte{byte(49)}})
			}
		}
	}
	if len(d.val) == 0 {
		return &BigInt{val: []byte{48}}
	}
	return d
}

func (b *BigInt) module(str *BigInt) *BigInt {
	var first byte
	var d int = 0
	var mod *BigInt = b
	if !b.compare(str) {
		return mod
	}
	for {
		lenDiff := len(b.val) - len(str.val)
		var temp = &BigInt{}
		var temp2 = &BigInt{}
		temp.val = str.val
		temp2.val = str.val
		for i := 1; i <= lenDiff; i++ {
			temp.val = append(temp.val, byte(48))
			if b.compare(temp) {
				temp2.val = temp.val
			} else {
				break
			}
		}

		b = b.subtract(temp2)
		first = b.val[0]
		if first == '-' {
			break
		} else {
			d++
			mod = b
		}
	}
	if mod != nil {
		var firstNotZero int
		for i, v := range mod.val {
			if v == byte(48) {
				firstNotZero = i
				continue
			} else {
				break
			}
		}
		return &BigInt{val: mod.val[firstNotZero:]}
	} else {
		return &BigInt{val: []byte{byte(48)}}
	}
}

func changeIntToByte(param int) []byte {
	if param == 0 {
		return []byte{48}
	}
	var (
		result  []byte
		results []byte
	)
	for param != 0 {
		mod := param % 10
		param = param / 10
		result = append(result, byte(mod+48))
	}
	for i := len(result) - 1; i >= 0; i-- {
		results = append(results, result[i])
	}
	return results
}

func (b *BigInt) compare(str2 *BigInt) bool {
	strLen := len(b.val)
	str2Len := len(str2.val)
	if strLen < str2Len {
		return false
	} else if strLen == str2Len {
		var flag = true
		for i := 0; i < strLen; i++ {
			if b.val[i] > str2.val[i] {
				break
			} else if b.val[i] == str2.val[i] {
				continue
			} else if b.val[i] < str2.val[i] {
				flag = false
				break
			}
		}
		return flag
	} else {
		return true
	}
}

func (b *BigInt) powBase10(len int) *BigInt {
	for i := 1; i <= len; i++ {
		b.val = append(b.val, byte(48))
	}
	return b
}

func (b *BigInt) GetValue() []byte {
	return b.val
}
