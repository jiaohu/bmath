package bigint

import (
	"fmt"
	"testing"
)

func TestNewInt(t *testing.T) {
	s, err := NewInt("3473892658465897239748365838758137583658368596")
	fmt.Println(err)
	fmt.Println(s.multiple(&BigInt{val: []byte("397483274")}).String())
}

func TestBigInt_Add(t *testing.T) {
	origin, err := NewInt("1300000000000")
	if err != nil {
		fmt.Println(err)
		return
	}
	add, err := NewInt("20000000000")
	if err != nil {
		fmt.Println(err)
		return
	}
	res := Add(origin, add)
	fmt.Println(res.String())
}

func TestBigInt_sub(t *testing.T) {
	first, err := NewInt("-1")
	if err != nil {
		fmt.Println(err)
		return
	}

	second, err := NewInt("24")
	if err != nil {
		fmt.Println(err)
		return
	}
	res := Subtract(first, second)
	fmt.Println(res.String())
}

func TestBigInt_multiple(t *testing.T) {
	// 5459990476343132489621253077746345103081986255550466866704
	first, err := NewInt("73891748364368348390248933948")
	if err != nil {
		fmt.Println(err)
		return
	}

	second, err := NewInt("73891748364368348390248933948")
	if err != nil {
		fmt.Println(err)
		return
	}
	res := Multiple(first, second)
	fmt.Println(res.String())
}

func TestMultiple(t *testing.T) {
	a := &BigInt{val: []byte{48}}
	b := &BigInt{val: []byte{49}}
	res := a.multiple(b)
	fmt.Println(res.String())
}

func TestDivide(t *testing.T) {
	first, err := NewInt("4780")
	if err != nil {
		fmt.Println(err)
		return
	}

	second, err := NewInt("2")
	if err != nil {
		fmt.Println(err)
		return
	}
	res := Divide(first, second)
	fmt.Println(res.String())
}

func TestModule(t *testing.T) {
	first, err := NewInt("28026088472520")
	if err != nil {
		fmt.Println(err)
		return
	}

	second, err := NewInt("65537")
	if err != nil {
		fmt.Println(err)
		return
	}
	res := Module(first, second)
	fmt.Println(res.String())
}

func TestMod(t *testing.T) {
	fmt.Println((-7) % (-2))
}

func TestFastPow(t *testing.T) {
	first, err := NewInt("33")
	if err != nil {
		fmt.Println(err)
		return
	}

	second, err := NewInt("13")
	if err != nil {
		fmt.Println(err)
		return
	}

	third, err := NewInt("13")
	if err != nil {
		fmt.Println(err)
		return
	}
	res := FastPow(first, second, third)
	fmt.Println("res.String()", res.String())
}

func TestCompare(t *testing.T) {
	temp, _ := NewInt("121")
	temp2, _ := NewInt("211")
	fmt.Println(temp.compare(temp2))
}

