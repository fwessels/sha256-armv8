package sha256

import (
	"io/ioutil"
	"testing"
)

func TestSum(t *testing.T) {

	dat, _ := ioutil.ReadFile("/home/linaro/gopath/src/github.com/minio/blake2b-simd/out.bin")

	Sha256(dat)
}
