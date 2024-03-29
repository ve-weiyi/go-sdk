package crypto

import (
	"crypto/md5"
	"encoding/hex"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: MD5V
//@description: md5加密
//@param: str []byte
//@return: string

func MD5V(str []byte, b ...byte) string {
	h := md5.New()
	h.Write(str)
	return hex.EncodeToString(h.Sum(b))
}

// Md5v md5 加盐加密
func Md5v(str string, salt string) string {
	h := md5.New()
	h.Write([]byte(salt + str))
	var res []byte
	res = h.Sum(nil)
	return hex.EncodeToString(res)
}

// Md5vMulti  iteration:加密次数
func Md5vMulti(str string, salt string, iteration int) string {
	h := md5.New()
	h.Write([]byte(salt + str))
	var res []byte
	res = h.Sum(nil)
	for i := 0; i < iteration-1; i++ {
		h.Reset()
		h.Write(res)
		res = h.Sum(nil)
	}
	return hex.EncodeToString(res)
}
