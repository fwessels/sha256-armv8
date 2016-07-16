package sha256

import (
	"io/ioutil"
	"testing"
)

var dat []byte

func benchmarkSize(b *testing.B, size int) {
	b.SetBytes(int64(size))
	for i := 0; i < b.N; i++ {
		Sha256(dat)
	}
}

// Benchmark writes of 100MB
func BenchmarkSize100MB(b *testing.B) {
	b.StopTimer()
	dat, _ = ioutil.ReadFile("/home/linaro/gopath/src/github.com/minio/blake2b-simd/out.bin")
	b.StartTimer()

	benchmarkSize(b, 100*1024*1024)
}
