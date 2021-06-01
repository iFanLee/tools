package tool

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
)

/**
 * @Author: Lee
 * @Date: 2021/6/1 15:09
 * @Desc:
 */
//sha1签名
func Sha1String(data string) string {
	s := sha1.New()
	s.Write([]byte(data))
	return hex.EncodeToString(s.Sum([]byte(nil)))
}
func HmacSHA1(secretToken, payloadBody string) []byte {
	key := []byte(secretToken)
	mac := hmac.New(sha1.New, key)
	mac.Write([]byte(payloadBody))
	return mac.Sum(nil)
}

func HmacSha256(data string, secret string) []byte {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(data))
	return h.Sum(nil)
}
//SHA256生成哈希值
func SHA256HashCode(message []byte)string{
	//方法一：
	//创建一个基于SHA256算法的hash.Hash接口的对象
	hash := sha256.New()
	//输入数据
	hash.Write(message)
	//计算哈希值
	bytes := hash.Sum(nil)
	//将字符串编码为16进制格式,返回字符串
	hashCode := hex.EncodeToString(bytes)
	//返回哈希值
	return hashCode

	//方法二：
	//bytes2:=sha256.Sum256(message)//计算哈希值，返回一个长度为32的数组
	//hashcode2:=hex.EncodeToString(bytes2[:])//将数组转换成切片，转换成16进制，返回字符串
	//return hashcode2
}