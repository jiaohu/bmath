# bmath
math method about big numbers, such as add, sub, multiple, divide, module, fast pow.
Now, only integer number into account. Also, there are some methods about RSA.

# Install
```
go get github.com/jiaohu/bmath
```


# Example
```
import (
	"fmt"
    "github.com/jiaohu/bmath"
	"github.com/jiaohu/bmath/bigint"
)

func main() {
	s1, err := bigint.NewInt("1234325646546546")
	if err != nil {
		return
	}
	s2, err := bigint.NewInt("98372487423874")
	if err != nil {
		return
	}

	r1 := bigint.Add(s1, s2)
	fmt.Println(r1)

	r2 := bigint.Subtract(s1, s2)
	fmt.Println(r2)

	r3 := bigint.Multiple(s1, s2)
	fmt.Println(r3)

	r4 := bigint.Divide(s1, s2)
	fmt.Println(r4)

	r5 := bigint.Module(s1, s2)
	fmt.Println(r5)

	// RSA p,q,e should be a big prime number. here is a example(the numbers may not be correct)
	p, _ := bigint.NewInt("7021557")
    q, _ := bigint.NewInt("4202321")
    e, _ := bigint.NewInt("65537")
	pub, self := bmath.GenRSAKey(p, q, e)
	fmt.Println(pub.Key)
	fmt.Println(self.Key)
}
```
