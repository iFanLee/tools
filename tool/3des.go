package tool

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
)

/**
 * @Author: Lee
 * @Date: 2021/6/1 15:14
 * @Desc:
 */
var tripleSecKey  = []byte("1234567890qwertyuiopasdf")//长度一定是24
//解密
func TripleDESDeCrypt(crypted []byte)[]byte{
	//获取block块
	block,_ :=des.NewTripleDESCipher(tripleSecKey)
	//创建切片
	context := make([]byte,len(crypted))
	//设置解密方式
	blockMode := cipher.NewCBCDecrypter(block,tripleSecKey[:8])
	//解密密文到数组
	blockMode.CryptBlocks(context,crypted)
	//去补码
	context = PKCSUnPadding(context)
	return context
}
//3des加密
func TripleDESEnCrypt(origData []byte) []byte {
	//获取block块
	block,_ :=des.NewTripleDESCipher(tripleSecKey)
	//补码
	origData = PKSCPadding(origData, block.BlockSize())
	//设置加密方式为 3DES  使用3条56位的密钥对数据进行三次加密
	blockMode := cipher.NewCBCEncrypter(block,tripleSecKey[:8])
	//创建明文长度的数组
	crypted := make([]byte,len(origData))
	//加密明文
	blockMode.CryptBlocks(crypted,origData)
	return crypted
}
//去补码
func PKCSUnPadding(origData []byte)[]byte{
	length := len(origData)
	if length<1{
		return []byte{}
	}
	unpadding := int(origData[length-1])
	if length<unpadding{
		return []byte{}
	}
	return origData[:length-unpadding]
}
//补码
func PKSCPadding(origData []byte,blockSize int) []byte  {
	//计算需要补的位数
	padding := blockSize-len(origData)%blockSize
	//在切片后面追加char数量的byte(char)
	padtext := bytes.Repeat([]byte{byte(padding)},padding)
	return append(origData,padtext...)
}