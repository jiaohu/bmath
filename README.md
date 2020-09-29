# bmath
math method about big numbers, such as add, sub, multiple, divide, module, fast pow.
Now, only integer number into account. Also, there are some methods about RSA.

# Install
```
go get github.com/jiaohu/bmath
```


#Example
```
import (
	"fmt"
	"github.com/jiaohu/bmath/operator"
)

func main() {
	s1, err := operator.NewBigData("1234325646546546")
	if err != nil {
		return
	}
	s2, err := operator.NewBigData("98372487423874")
	if err != nil {
		return
	}
	// BaseOperator
	r1 := operator.Add(s1, s2)
	fmt.Println(r1.ConvertString())

	r2 := operator.Subtract(s1, s2)
	fmt.Println(r2.ConvertString())

	r3 := operator.Multiple(s1, s2)
	fmt.Println(r3.ConvertString())

	r4 := operator.Divide(s1, s2)
	fmt.Println(r4.ConvertString())

	r5 := operator.Module(s1, s2)
	fmt.Println(r5.ConvertString())

	// RSA p,q,e should be a big prime number
        //here is a example, 
        //the numbers may not be correct 
	p, _ := operator.NewBigData("24324121")
	q, _ := operator.NewBigData("37826476721")
	e, _ := operator.NewBigData("65355")
	pub, self := operator.GenRSAKey(p, q, e)
	pub.Key.ConvertString() // pub key
	self.Key.ConvertString()
	pub.N.ConvertString()
}
```
