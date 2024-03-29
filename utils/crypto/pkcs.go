package crypto

import "bytes"

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
func pkcs7UnPadding(data []byte) []byte {
	length := len(data)
	//获取填充的个数
	unPadding := int(data[length-1])
	return data[:(length - unPadding)]
}

func pkcs5Padding(data []byte, blockSize int) []byte {
	//判断缺少几位长度。最少1，最多 blockSize
	padding := blockSize - len(data)%blockSize
	//补足位数。把切片[]byte{byte(padding)}复制padding个
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

func pkcs5UnPadding(data []byte) []byte {
	length := len(data)
	unPadding := int(data[length-1])
	return data[:(length - unPadding)]
}
