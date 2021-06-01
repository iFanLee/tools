package tool

import (
	"regexp"
	"strings"
)

/**
 * @Author: Lee
 * @Date: 2021/6/1 15:01
 * @Desc:
 */
//去掉JavaScript
func RemoveJavascript(str string) string {
	re, _ := regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	return re.ReplaceAllString(str, "")
}
//抓取HTML中图片链接地址
func GetHtmlImage(str string) []string {
	arr := make([]string,0)
	var hrefRegexp = regexp.MustCompile("<img.*?>")
	match := hrefRegexp.FindAllString(str, -1)
	var src = regexp.MustCompile("src=\".*?\"")
	if match != nil {
		for _, v := range match {
			s := src.FindString(v)
			replace,_ := regexp.Compile("src=\"")
			s1 := replace.ReplaceAllString(s,"")
			s2 := strings.TrimRight(s1,"\"")
			if s2==""{
				continue
			}
			arr = append(arr,s2)
		}
	}
	return  arr
}
//抓取HTML中图片链接地址
func GetHtmlVideo(str string) []string {
	arr := make([]string,0)
	var hrefRegexp = regexp.MustCompile("<video.*?>")
	match := hrefRegexp.FindAllString(str, -1)
	var src = regexp.MustCompile("src=\".*?\"")
	if match != nil {
		for _, v := range match {
			s := src.FindString(v)
			replace,_ := regexp.Compile("src=\"")
			s1 := replace.ReplaceAllString(s,"")
			s2 := strings.TrimRight(s1,"\"")
			if s2==""{
				continue
			}
			arr = append(arr,s2)
		}
	}
	return  arr
}