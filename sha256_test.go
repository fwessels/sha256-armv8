package sha256

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	buf := make([]byte, 128)
	for i := range buf {
		buf[i] = byte(i)
	}

	fmt.Println(len(buf))

	fmt.Println(sha256())
}
