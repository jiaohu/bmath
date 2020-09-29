package operator

import (
	"fmt"
	"testing"
)

func TestNewBigData(t *testing.T) {
	_, err := NewBigData("0")
	fmt.Println(err)
}

func TestBigData_Add(t *testing.T) {
	origin, err := NewBigData("1300000000000")
	if err != nil {
		fmt.Println(err)
		return
	}
	add, err := NewBigData("20000000000")
	if err != nil {
		fmt.Println(err)
		return
	}
	res := Add(origin, add)
	fmt.Println(res.ConvertString())
}

func TestBigData_sub(t *testing.T) {
	first, err := NewBigData("-1")
	if err != nil {
		fmt.Println(err)
		return
	}

	second, err := NewBigData("24")
	if err != nil {
		fmt.Println(err)
		return
	}
	res := Subtract(first, second)
	fmt.Println(res.ConvertString())
}

func TestBigData_multiple(t *testing.T) {
	first, err := NewBigData("73891748364368348390248933948")
	if err != nil {
		fmt.Println(err)
		return
	}

	second, err := NewBigData("73891748364368348390248933948")
	if err != nil {
		fmt.Println(err)
		return
	}
	res := Multiple(first, second)
	fmt.Println(res.ConvertString())
}

func TestMultiple(t *testing.T) {
	a := &BigData{val: []byte{48}}
	b := &BigData{val: []byte{49}}
	res := a.multiple(b)
	fmt.Println(res.ConvertString())
}

func TestDivide(t *testing.T) {
	first, err := NewBigData("4780")
	if err != nil {
		fmt.Println(err)
		return
	}

	second, err := NewBigData("2")
	if err != nil {
		fmt.Println(err)
		return
	}
	res := Divide(first, second)
	fmt.Println(res.ConvertString())
}

func TestModule(t *testing.T) {
	first, err := NewBigData("28026088472520")
	if err != nil {
		fmt.Println(err)
		return
	}

	second, err := NewBigData("65537")
	if err != nil {
		fmt.Println(err)
		return
	}
	res := Module(first, second)
	fmt.Println(res.ConvertString())
}

func TestMod(t *testing.T) {
	fmt.Println((-7) % (-2))
}

func TestFastPow(t *testing.T) {
	first, err := NewBigData("33")
	if err != nil {
		fmt.Println(err)
		return
	}

	second, err := NewBigData("13")
	if err != nil {
		fmt.Println(err)
		return
	}

	third, err := NewBigData("13")
	if err != nil {
		fmt.Println(err)
		return
	}
	res := FastPow(first, second, third)
	fmt.Println("res.ConvertString()", res.ConvertString())
}

func TestCompare(t *testing.T) {
	temp, _ := NewBigData("28026088472520")
	temp2, _ := NewBigData("65537000000000")
	fmt.Println(temp.compare(temp2))
}
