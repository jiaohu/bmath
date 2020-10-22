package bigfloat

import (
	"fmt"
	"testing"
)

func TestNewFloat(t *testing.T) {
	f, err := NewFloat("-111000.0000003")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(f)
	fmt.Println(f.accuracy)
	fmt.Println(f.neg)
}

func TestAdd(t *testing.T) {
	f, err := NewFloat("111000.0000003")
	if err != nil {
		fmt.Println(err)
		return
	}
	f2, _ := NewFloat("1111.01")
	f3 := f.add(f2)
	fmt.Println(f3)
	fmt.Println(f3.accuracy)
}

func TestSub(t *testing.T) {
	f, err := NewFloat("1111.01")
	if err != nil {
		fmt.Println(err)
		return
	}
	f2, _ := NewFloat("111000.0000003")
	fmt.Println(f.subtract(f2).String())
}

func TestSub2(t *testing.T) {
	f, err := NewFloat("111000.0000003")
	if err != nil {
		fmt.Println(err)
		return
	}
	f2, _ := NewFloat("1111.01")
	fmt.Println(f.subtract(f2).String())
}

func TestCompare(t *testing.T) {
	f, err := NewFloat("111000.0000003")
	if err != nil {
		fmt.Println(err)
		return
	}
	f2, _ := NewFloat("111000.00000004")
	ans := f.compare(f2)
	fmt.Println(ans)
	fmt.Println(f)
	fmt.Println(f2)
}

func TestMultiple(t *testing.T) {
	f, err := NewFloat("111000.0000003")
	if err != nil {
		fmt.Println(err)
		return
	}
	f2, _ := NewFloat("111000.00000004")
	fmt.Println(f.multiple(f2).String())
}

func TestDivide(t *testing.T) {
	f, err := NewFloat("111000.0000003")
	if err != nil {
		fmt.Println(err)
		return
	}
	f2, _ := NewFloat("111000.00000004")
	fmt.Println(f.divide(f2).String())
}
