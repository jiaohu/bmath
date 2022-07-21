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

// BenchmarkName-8   	  232155	      5137 ns/op
func BenchmarkName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewFloat("-111000.0000003")
	}
}

// BenchmarkName2-8   	28090694	        42.65 ns/op
func BenchmarkName2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewFloat2("-111000.0000003")
	}
}

func TestAdd(t *testing.T) {
	f, err := NewFloat("111000.0000003")
	if err != nil {
		fmt.Println(err)
		return
	}
	f2, _ := NewFloat("1111.01")
	f3 := f.Add(f2)
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
	fmt.Println(f.Subtract(f2).String())
}

func TestSub2(t *testing.T) {
	f, err := NewFloat("111000.0000003")
	if err != nil {
		fmt.Println(err)
		return
	}
	f2, _ := NewFloat("1111.01")
	fmt.Println(f.Subtract(f2).String())
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
	fmt.Println(f.Multiple(f2).String())
}

func TestDivide(t *testing.T) {
	f, err := NewFloat("111000.0000003")
	if err != nil {
		fmt.Println(err)
		return
	}
	f2, _ := NewFloat("111000.00000004")
	fmt.Println(f.Divide(f2).String())
}
