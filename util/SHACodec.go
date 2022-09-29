package util

import(
	"encoding/base64"
	"crypto/sha256"
)

func SHAEncoding(target string) (output string) {
	buf := []byte(target)
	encr := sha256.New()
	encr.Write(buf)

	output = base64.StdEncoding.EncodeToString(encr.Sum(nil))
	return
}