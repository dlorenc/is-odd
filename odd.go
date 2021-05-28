package odd

import (
	"fmt"

	even "github.com/dlorenc/is-even"
)

func IsOdd(n int) bool {
	fmt.Println("IsOdd")
	return !even.IsEven(n)
}
