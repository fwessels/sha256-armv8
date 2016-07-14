//+build !arm64 noasm appengine

package sha256

func sha256() string {
	return "hallo from _noasm"
}
