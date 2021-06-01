package tool

import (
	"errors"
	"fmt"
	"github.com/mozillazg/go-pinyin"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

/**
 * @Author: Lee
 * @Date: 2021/6/1 15:03
 * @Desc:
 */
func IsMob(str string) bool {
	reg := `^1[3456789]\d{9}$`
	rgx := regexp.MustCompile(reg)
	if rgx.MatchString(str) {
		return  true
	}else {
		return  false
	}
}
/*
正则判断是否为英文数字
*/
func IsEnNum(str string) bool {
	reg := `^[\da-zA-Z]+$`
	rgx := regexp.MustCompile(reg)
	if rgx.MatchString(str) {
		return  true
	}else {
		return  false
	}
}
/*
正则判断是否为数字
*/
func IsNum(str string) bool {
	reg := `^[\d]+$`
	rgx := regexp.MustCompile(reg)
	if rgx.MatchString(str) {
		return  true
	}else {
		return  false
	}
}
/*
正则判断是否为中英文数字
*/
func IsZhEnNum(str string) bool {
	reg := `^[\da-zA-Z\p{Han}]+$`
	rgx := regexp.MustCompile(reg)
	if rgx.MatchString(str) {
		return  true
	}else {
		return  false
	}
}
/*
正则判断是否为英文
*/
func IsEn(str string) bool {
	reg := `^[a-zA-Z]+$`
	rgx := regexp.MustCompile(reg)
	if rgx.MatchString(str) {
		return  true
	}else {
		return  false
	}
}

/*
正则判断是否为中英文
*/
func IsZhEn(str string) bool {
	reg := `^[a-zA-Z\p{Han}]+$`
	rgx := regexp.MustCompile(reg)
	if rgx.MatchString(str) {
		return  true
	}else {
		return  false
	}
}
//判断是否为中文
func IsHan(r rune) bool {
	return unicode.Is(unicode.Han, r)
}
/*
获取文字首字母
*/
func GetFirstLetterPinYin(r rune) string {
	result := pinyin.Pinyin(string(r), pinyin.NewArgs())
	if len(result)>0{
		if len(result[0])>0{
			if len(result[0][0])>0{
				return string(result[0][0][0])
			}
		}
	}
	return ""
}
//判断是否为金额格式
func IsMoney(str string) (money float64,err error) {
	if f1,err := strconv.ParseFloat(str,64);err==nil{
		m := strings.Split(fmt.Sprintf("%v",str),".")
		if len(m)>1{
			if len(m[1])>2{
				return 0,errors.New("not Money Format")
			}else {
				return f1,nil
			}
		}else {
			return f1,nil
		}
	}else {
		return 0,err
	}
}