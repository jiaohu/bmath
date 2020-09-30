package bigfloat

import (
	"errors"
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
	if origin[0] == 45 {
		flag = true
	}

	for i, v := range origin {
		if v == 46 {
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

func (b *BigFloat) add(str *BigFloat) *BigFloat {
	return nil
}

func (b *BigFloat) subtract(str *BigFloat) *BigFloat {
	return nil
}

func (b *BigFloat) multiple(str *BigFloat) *BigFloat {
	return nil
}

func (b *BigFloat) divide(str *BigFloat) *BigFloat {
	return nil
}

func (b *BigFloat) compare(str *BigFloat) bool {
	return false
}
