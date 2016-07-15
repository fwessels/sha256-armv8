//+build !noasm

package sha256

import (
	"fmt"
	"io/ioutil"
	"time"
)

//go:noescape
func block(h []uint32, p, x2 []uint8, constants []uint32)

func Sha256() string {

	h := []uint32{0x6a09e667,
		0xbb67ae85,
		0x3c6ef372,
		0xa54ff53a,
		0x510e527f,
		0x9b05688c,
		0x1f83d9ab,
		0x5be0cd19}

	constants := []uint32{0x428a2f98, 0x71374491, 0xb5c0fbcf, 0xe9b5dba5,
		0x3956c25b, 0x59f111f1, 0x923f82a4, 0xab1c5ed5,
		0xd807aa98, 0x12835b01, 0x243185be, 0x550c7dc3,
		0x72be5d74, 0x80deb1fe, 0x9bdc06a7, 0xc19bf174,
		0xe49b69c1, 0xefbe4786, 0x0fc19dc6, 0x240ca1cc,
		0x2de92c6f, 0x4a7484aa, 0x5cb0a9dc, 0x76f988da,
		0x983e5152, 0xa831c66d, 0xb00327c8, 0xbf597fc7,
		0xc6e00bf3, 0xd5a79147, 0x06ca6351, 0x14292967,
		0x27b70a85, 0x2e1b2138, 0x4d2c6dfc, 0x53380d13,
		0x650a7354, 0x766a0abb, 0x81c2c92e, 0x92722c85,
		0xa2bfe8a1, 0xa81a664b, 0xc24b8b70, 0xc76c51a3,
		0xd192e819, 0xd6990624, 0xf40e3585, 0x106aa070,
		0x19a4c116, 0x1e376c08, 0x2748774c, 0x34b0bcb5,
		0x391c0cb3, 0x4ed8aa4a, 0x5b9cca4f, 0x682e6ff3,
		0x748f82ee, 0x78a5636f, 0x84c87814, 0x8cc70208,
		0x90befffa, 0xa4506ceb, 0xbef9a3f7, 0xc67178f2}

	dat, _ := ioutil.ReadFile("/home/linaro/gopath/src/github.com/minio/blake2b-simd/out.bin")

	start := time.Now()

	blocks := len(dat) / 64

	if blocks > 0 {
		block(h[:], dat[:64*blocks], make([]uint8, 256, 256), constants[:])
	}
	l := len(dat) - blocks*64

	buffer := make([]uint8, 64, 64)
	copy(buffer, dat[64*blocks:])

	i := 0
	if l < 56 {
		// padd current block to 56 byte
		buffer[l] = 0x80
		i = l + 1
	} else {
		// fill up current block and update hash
		buffer[l] = 0x80
		for i = l + 1; i < 64; i++ {
			buffer[i] = 0x0
		}

		block(h[:], buffer[:], make([]uint8, 256, 256), constants[:])

		// add (almost) one block of zero bytes
		i = 0
	}
	for ; i < 56; i++ {
		buffer[i] = 0x0
	}

	// add message length in bits in big endian
	for i = 0; i < 8; i++ {
		buffer[63-i] = uint8((len(dat) * 8) >> (uint32(i) * 8))
	}

	block(h[:], buffer[:], make([]uint8, 256, 256), constants[:])

	elapsed := time.Since(start)
	fmt.Printf("SHA256 (size=%dM) took %s\n", len(dat)/1024/1024, elapsed)

	fmt.Printf("%08x%08x%08x%08x%08x%08x%08x%08x\n", h[0], h[1], h[2], h[3], h[4], h[5], h[6], h[7])
	return "hallo from arm64"
}
