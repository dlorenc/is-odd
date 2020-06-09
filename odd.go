package odd

import (
	even "github.com/dlorenc/is-even"
)

func IsOdd(n int) bool {
	return !even.IsEven(n)
}
