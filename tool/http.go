package tool

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"strings"
)

/**
 * @Author: Lee
 * @Date: 2021/6/1 14:52
 * @Desc:
 */
func HttpReq(url,method string,header,cookies,form map[string]string) (result []byte,rspHeader *http.Response,err error) {
	formStr := ""
	for k,v := range form{
		if formStr==""{
			formStr = fmt.Sprintf("%v=%v",k,v)
		}else {
			formStr += fmt.Sprintf("&%v=%v",k,v)
		}
	}
	req, err := http.NewRequest(method, url,strings.NewReader(formStr))
	if err!=nil{
		return []byte{},nil,err
	}
	for k,v := range header {
		req.Header.Set(k,v)
	}
	for k,v := range cookies{
		req.AddCookie(&http.Cookie{Name: k,Value: v})
	}
	j, _ := cookiejar.New(nil)
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	var resp *http.Response
	client.Jar = j
	resp, err = client.Do(req)
	if err!=nil{
		return []byte{},nil,err
	}
	defer resp.Body.Close()
	// 正文信息
	var buf []byte
	buf, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{},resp,err
	}
	return buf,resp,nil
}
