package crypto

import (
	"encoding/base64"
	"encoding/hex"
	"golang.org/x/crypto/bcrypt"
	"log"
	"testing"
)

// 一个字节Byte由8位bit。8*64=512
var Bit256 = "012345678912345.012345678912345."
var Bit512 = "012345678912345.012345678912345." + "012345678912345.012345678912345."
var Bit1024 = Bit512 + Bit512
var plaintext = Bit512 + Bit512

// 组成32位16进制 128bit x 4
var ciphertext = ""

func TestBcrypt(t *testing.T) {

	//$2a$10$D3LbggNCcxz95XOr5CkLdeDadDc22xoSISMSADvM3p2BCmO49x1Yu
	//$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy
	//\__/\/ \____________________/\_____________________________/
	//Alg Cost      Salt 128bits               Hash 192bits
	//ciphertext = BcryptHash(plaintext)
	//log.Println(BcryptHash(Bit512))
	//log.Println(BcryptCheck(plaintext, ciphertext))
	//log.Println(plainxtext)
	//log.Println(ciphertext)

	//6.60s
	for i := 0; i < 10; i++ {
		bcrypt.GenerateFromPassword([]byte(Bit256), bcrypt.DefaultCost)
	}
}

func TestMd5(t *testing.T) {
	//1bdf247646854ad6d841ba6b0cd376fe
	//log.Println(Md5v(plaintext, "123"))

	for i := 0; i < 1000000; i++ {
		Md5v(plaintext, Bit512)
	}
}

func TestSha256(t *testing.T) {
	//log.Println(Sha256v(plaintext, "123"))

	for i := 0; i < 1000000; i++ {
		Sha256v(Bit512, Bit512)
	}
}

func BenchmarkMd5v(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// 494.4 ns/op
		Md5v(Bit1024, Bit512)
	}
}

func BenchmarkSha256v(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// 249.4 ns/op
		Sha256v(Bit1024, Bit512)
	}
}

func BenchmarkBcrypt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// cost=4:1075624 ns/op    cost=10:65876130 ns/op
		bcrypt.GenerateFromPassword([]byte(Bit256), bcrypt.DefaultCost)
	}
}

var key = []byte("1234567.1234567.") // 加密的密钥

func BenchmarkAES(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// 456.7 ns/op    2017 ns/op  15062 ns/op
		// 473.0 ns/op    2130 ns/op   16414 ns/op
		// 526.8 ns/op 	  2388 ns/op	18430 ns/op
		//encrypted := AesEncryptCBC([]byte(Bit1024), key)
		encrypted := AesEncryptCBC([]byte(Bit1024), key)
		AesDecryptCBC(encrypted, key)
	}
}

func BitsX10(src string) string {
	return src + src + src + src + src + src + src + src + src + src
}

func Test_B_1(t *testing.T) {
	plaintext := []byte("460154561234") // 待加密的数据
	key := []byte("9876787656785679")   // 加密的密钥
	log.Println("原文：", string(plaintext))

	log.Println("------------------ CBC模式 --------------------")
	encrypted := AesEncryptCBC(plaintext, key)
	log.Println("密文(hex)：", hex.EncodeToString(encrypted))
	log.Println("密文(base64)：", base64.StdEncoding.EncodeToString(encrypted))
	decrypted := AesDecryptCBC(encrypted, key)
	log.Println("解密结果：", string(decrypted))
}

func Test_B_2(t *testing.T) {
	plaintext := []byte("460154561234") // 待加密的数据
	key := []byte("9876787656785679")   // 加密的密钥
	log.Println("原文：", string(plaintext))

	log.Println("------------------ ECB模式 --------------------")
	encrypted := AesEncryptECB(plaintext, key)
	log.Println("密文(hex)：", hex.EncodeToString(encrypted))
	log.Println("密文(base64)：", base64.StdEncoding.EncodeToString(encrypted))
	decrypted := AesDecryptECB(encrypted, key)
	log.Println("解密结果：", string(decrypted))
}

func Test_B_3(t *testing.T) {
	plaintext := []byte("460154561234") // 待加密的数据
	key := []byte("9876787656785679")   // 加密的密钥
	log.Println("原文：", string(plaintext))

	log.Println("------------------ CFB模式 --------------------")
	encrypted := AesEncryptCFB(plaintext, key)
	log.Println("密文(hex)：", hex.EncodeToString(encrypted))
	log.Println("密文(base64)：", base64.StdEncoding.EncodeToString(encrypted))
	decrypted := AesDecryptCFB(encrypted, key)
	log.Println("解密结果：", string(decrypted))
}
