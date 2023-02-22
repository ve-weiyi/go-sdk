package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"errors"
)

//加密模式(英文） 		加密模式（中文） 			介绍
//Electronic Code Book(ECB)
//电子密码本模式
//最基本的加密模式，也就是通常理解的加密，相同的明文将永远加密成相同的密文，无初始向量，容易受到密码本重放攻击，一般情况下很少用

//Cipher Block Chaining(CBC)
//密码分组链接模式
//明文被加密前要与前面的密文进行异或运算后再加密，因此只要选择不同的初始向量，相同的密文加密后会形成不同的密文，这是目前应用最广泛的模式。CBC加密后的密文是上下文相关的，但明文的错误不会传递到后续分组，但如果一个分组丢失，后面的分组将全部作废(同步错误)。

//Cipher Feedback Mode(CFB)
//加密反馈模式
//类似于自同步序列密码，分组加密后，按8位分组将密文和明文进行移位异或后得到输出同时反馈回移位寄存器，优点最小可以按字节进行加解密，也可以是n位的，CFB也是上下文相关的，CFB模式下，明文的一个错误会影响后面的密文(错误扩散)。

//Output Feedback Mode(OFB)
//输出反馈模式
//将分组密码作为同步序列密码运行，和CFB相似，不过OFB用的是前一个n位密文输出分组反馈回移位寄存器，OFB没有错误扩散问题。

var AesEncryptError = errors.New("加密字符串错误！")

// AesEncrypt 加密
func AesEncryptCBC(data []byte, key []byte) []byte {
	// NewCipher creates and returns a new cipher.Block. The key argument should be the AES key, either 16, 24, or 32 bytes to select AES-128, AES-192, or AES-256.
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	//判断加密快的大小
	blockSize := block.BlockSize()
	//填充
	encryptBytes := pkcs7Padding(data, blockSize)
	//初始化加密数据接收切片
	crypted := make([]byte, len(encryptBytes))
	//加密向量
	iv := key[:blockSize]
	//使用cbc加密模式
	blockMode := cipher.NewCBCEncrypter(block, iv)
	//执行加密
	blockMode.CryptBlocks(crypted, encryptBytes)
	return crypted
}

// AesDecrypt 解密
func AesDecryptCBC(data []byte, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	//获取块的大小
	blockSize := block.BlockSize()
	//使用cbc
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	//初始化解密数据接收切片
	crypted := make([]byte, len(data))
	//执行解密
	blockMode.CryptBlocks(crypted, data)
	//去填充
	crypted, err = pkcs7UnPadding(crypted)
	if err != nil {
		panic(err)
	}
	return crypted
}

//PKCS7和PKCS5的区别是数据块的大小；
//PKCS5填充块的大小为8bytes(64位)
//PKCS7填充块的大小可以在1-255bytes之间。
//因为AES并没有64位的块, 如果采用PKCS5, 那么实质上就是采用PKCS7

// pkcs7Padding 填充
func pkcs7Padding(data []byte, blockSize int) []byte {
	//判断缺少几位长度。最少1，最多 blockSize
	padding := blockSize - len(data)%blockSize
	//补足位数。把切片[]byte{byte(padding)}复制padding个
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

// pkcs7UnPadding 移除
func pkcs7UnPadding(data []byte) ([]byte, error) {
	length := len(data)
	if length == 0 {
		return nil, AesEncryptError
	}
	//获取填充的个数
	unPadding := int(data[length-1])
	return data[:(length - unPadding)], nil
}
