//+build !arm64 noasm appengine

package sha256

func Sha256() string {
	return "hallo from _noasm"
}
