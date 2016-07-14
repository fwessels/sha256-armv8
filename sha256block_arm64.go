//+build !noasm

package sha256

import "fmt"

//go:noescape
func block(h []uint32, p []uint8)

func sha256() string {

	h := []uint32{0x6a09e667,
		0xbb67ae85,
		0x3c6ef372,
		0xa54ff53a,
		0x510e527f,
		0x9b05688c,
		0x1f83d9ab,
		0x5be0cd19}

	fmt.Printf("%08x\n", h[0])
	block(h[:], make([]uint8, 8, 8))
	fmt.Printf("%08x\n", h[0])
	return "hallo from arm64"
}
