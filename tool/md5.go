package tool

import (
	"crypto/md5"
	"encoding/hex"
)

/**
 * @Author: Lee
 * @Date: 2021/6/1 14:55
 * @Desc:
 */
//md5加密
func MD5(str string) string  {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}