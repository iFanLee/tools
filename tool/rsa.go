package tool

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/binary"
	"encoding/pem"
	"errors"
	"math/big"
)

/**
 * @Author: Lee
 * @Date: 2021/6/1 15:05
 * @Desc:
 */
func GenRsaKey(bits int) (priKey,pubKey string,err error) {
	// 生成私钥文件
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return "","",err
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	priBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}
	// 生成公钥文件
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return "","",err
	}
	publicBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	return string(pem.EncodeToMemory(priBlock)),string(pem.EncodeToMemory(publicBlock)),nil
}
// 解密
func RsaDecrypt(privateKey,ciphertext []byte) ([]byte, error) {
	//解密
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error!")
	}
	//解析PKCS1格式的私钥
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 解密
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}
// 加密
func RsaEncrypt(publicKey,origData []byte) ([]byte, error) {
	//解密pem格式的公钥
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)
	//加密
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

func RsaPubKeyNE(nStr,eStr string) (string,error) {
	decN, err := base64.RawURLEncoding.DecodeString(nStr)
	if err != nil {
		return "",err
	}
	n := big.NewInt(0)
	n.SetBytes(decN)
	decE, err := base64.RawURLEncoding.DecodeString(eStr)
	if err != nil {
		return "",err
	}
	var eBytes []byte
	if len(decE) < 8 {
		eBytes = make([]byte, 8-len(decE), 8)
		eBytes = append(eBytes, decE...)
	} else {
		eBytes = decE
	}
	eReader := bytes.NewReader(eBytes)
	var e uint64
	err = binary.Read(eReader, binary.BigEndian, &e)
	if err != nil {
		return "",err
	}

	// 生成公钥文件
	derPkix, err := x509.MarshalPKIXPublicKey(&rsa.PublicKey{N: n, E: int(e)})
	if err != nil {
		return "",err
	}
	publicBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	return string(pem.EncodeToMemory(publicBlock)),nil
}
