package tool

import (
	"net/url"
	"regexp"
	"strings"
)

/**
 * @Author: Lee
 * @Date: 2021/6/1 15:03
 * @Desc:
 */
func TrimSpace(str string) string {
	return strings.Replace(str, " ", "", -1)
}

func TrimNum(str string) string {
	re, _ := regexp.Compile(`[\d]`);
	return re.ReplaceAllString(str, "");
}
func InArrayString(data string,arr []string) bool {
	for _,n := range arr{
		if data==n{
			return true
		}
	}
	return false
}
func InArrayInt(data int,arr []int) bool {
	for _,n := range arr{
		if data==n{
			return true
		}
	}
	return false
}
/*
去掉重复内容
*/
func RemoveRepeatedElement(arr []string) (newArr []string) {
	newArr = make([]string, 0)
	length := len(arr)
	for i := 0; i < length; i++ {
		repeat := false
		for j := i + 1; j < length; j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat && arr[i] != ""{
			newArr = append(newArr, arr[i])
		}
	}
	return
}

func GetUrlDir(urlStr string) string {
	u,_ := url.Parse(urlStr)
	s1 := strings.LastIndex(u.Path,"/")
	if s1==-1{
		return ""
	}
	return u.Path[:s1+1]
}