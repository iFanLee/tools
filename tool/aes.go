package tool

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"errors"
)

/**
 * @Author: Lee
 * @Date: 2021/6/1 15:14
 * @Desc:
 */
// AesDecrypt 解密函数
func Aes128CBCDecrypt(ciphertext []byte, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, iv[:blockSize])
	origData := make([]byte, len(ciphertext))
	blockMode.CryptBlocks(origData, ciphertext)
	origData = PKCS7UnPadding(origData)
	if len(origData)==0{
		return origData,errors.New("dec fail")
	}
	return origData, nil
}
//AesEncrypt 加密函数
func Aes128CBCEncrypt(plaintext []byte, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	plaintext = PKCS7Padding(plaintext, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, iv)
	crypted := make([]byte, len(plaintext))
	blockMode.CryptBlocks(crypted, plaintext)
	return crypted, nil
}

func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	if length>0{
		unpadding := int(origData[length-1])
		if d:= length-unpadding;d>=0{
			return origData[:d]
		}
	}
	return []byte{}
}