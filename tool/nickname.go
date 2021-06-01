package tool

import (
	"regexp"
)

/**
 * @Author: Lee
 * @Date: 2021/6/1 14:58
 * @Desc:
 */

/*
判断昵称  中英文数字+emoji
1.判读格式是否正确
2.字母数字占1字节，中文占2字节，emoji4字节，总长度不超过32字节
*/
func IsNickname(str string,maxLength int) int {
	if !isNickname(str){
		return 100
	}
	length := getNicknameLength(str)
	if length > maxLength{
		return 101
	}
	return 200
}
func getNicknameLength(str string) int {
	removeZh := RemoveZh(str)
	zhLength := CountZh(str)
	length := len(removeZh)+zhLength*2
	return length
}
/*
昵称为中英文数字+emoji
*/
func isNickname(str string) bool {
	/*
		1F601 - 1F64F
		2702 - 27B0
		1F680 - 1F6C0
		24C2 - 1F251
		1F600 - 1F636
		1F681 - 1F6C5
		1F30D - 1F567
	*/
	reg := `^[\da-zA-Z|\x{4e00}-\x{9fa5}|\x{1F600}-\x{1F9FF}|\x{2600}-\x{26FF}|\x{2702}-\x{27B0}|\x{24C2}-\x{1F251}|\x{1F30D}-\x{1F567}]+$`
	rgx := regexp.MustCompile(reg)
	if rgx.MatchString(str) {
		return  true
	}else {
		return  false
	}
}

func RemoveZh(str string) string {
	reg := `[\x{4e00}-\x{9fa5}]`
	rgx := regexp.MustCompile(reg)
	d := rgx.ReplaceAllString(str,"")
	return d
}
func CountZh(str string) int {
	reg := `[\x{4e00}-\x{9fa5}]`
	rgx := regexp.MustCompile(reg)
	d := rgx.FindAllStringSubmatch(str,-1)
	return len(d)
}
