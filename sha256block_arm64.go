//+build !noasm

package sha256

//go:noescape
func block(h []uint32, message []uint8)

func Sha256(dat []byte) {

	h := []uint32{0x6a09e667,
		0xbb67ae85,
		0x3c6ef372,
		0xa54ff53a,
		0x510e527f,
		0x9b05688c,
		0x1f83d9ab,
		0x5be0cd19}

	blocks := len(dat) / 64

	if blocks > 0 {
		block(h[:], dat[:64*blocks])
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

		block(h[:], buffer[:])

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

	block(h[:], buffer[:])
}
