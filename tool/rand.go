package tool

import (
	"math/rand"
	"time"
)

/**
 * @Author: Lee
 * @Date: 2021/6/1 14:39
 * @Desc:
 */
//生成随机字符串
//strType:1.仅小写字母，2.仅大写字母，3.字母，4.字母+数字
func RandStr(length,strType int) string {
	str := ""
	switch strType {
	case 1:
		str = "0123456789"
	case 2:
		str = "abcdefghijklmnopqrstuvwxyz"
	case 3:
		str = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	case 4:
		str = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	default:
		str = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	}
	bytes := []byte(str)
	result := []byte{}
	lens := len(bytes)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(lens)])
	}
	return string(result)
}
//范围随机数
func RandInt64(min, max int64) int64 {
	if min >= max {
		return max
	}
	rand.Seed(time.Now().UnixNano())
	return rand.Int63n(max-min) + min
}