package ihash

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(req string) string {
	sum := md5.Sum([]byte(req))
	return hex.EncodeToString(sum[:])
}
