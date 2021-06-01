package tool

import "encoding/base64"

/**
 * @Author: Lee
 * @Date: 2021/6/1 15:39
 * @Desc:
 */

//base64加密
func Base64Encode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}
//base64解密
func Base64Dcode(str string) string {
	res,err := base64.StdEncoding.DecodeString(str)
	if err !=nil{
		return ""
	}
	return string(res)
}
//base64加密
func Base64UrlEncode(str string) string {
	return base64.RawURLEncoding.EncodeToString([]byte(str))
}
func Base64UrlDcode(str string) string {
	res,err := base64.RawURLEncoding.DecodeString(str)
	if err !=nil{
		return ""
	}
	return string(res)
}