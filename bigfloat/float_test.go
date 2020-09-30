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
	fmt.Println(f.val)
}
