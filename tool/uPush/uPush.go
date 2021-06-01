package uPush

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	tool2 "tools/tool"
)

const (
	urlMsg = "https://msgapi.umeng.com/api/send"
	secIOS = "sec"
	secAndroid = "sec"
	appKeyIOS = "key"
	appKeyAndroid = "key"

)
/*
ios 单播
*/
func UPushUnicastIOS(deviceToken,apsTitle,apsSubtitle,apsBody string,payloadData map[string]interface{},desc string) bool {
	body := UnicastJSONByIOS(deviceToken,apsTitle,apsSubtitle,apsBody,payloadData,desc)
	if body==""{
		return false
	}
	url := urlMsg +"?sign="+signIOS(urlMsg,body)
	_,err := post(url,body)
	if err==nil{
		return true
	}else {
		return false
	}
}
/*
android 单播
*/
func UPushUnicastAndroid(deviceToken,msgType,bodyTicker,bodyTitle,bodyText,afterOpen,goUrl,activity string,custom interface{},payloadData map[string]interface{},desc string) bool {
	body := UnicastJSONByAndroid(deviceToken,msgType,bodyTicker,bodyTitle,bodyText,afterOpen,goUrl,activity,custom,payloadData,desc)
	if body==""{
		return false
	}
	url := urlMsg +"?sign="+signAndroid(urlMsg,body)
	_,err := post(url,body)
	if err==nil{
		return true
	}else {
		return false
	}
}





/*
生成签名
   1.提取请求方法method（POST，全大写）；
   2.提取请求url信息，包括Host字段的域名(或ip:端口)和URI的path部分。
		注意不包括path的querystring。比如http://msg.umeng.com/api/send 或者 http://msg.umeng.com/api/status;
   3.提取请求的post-body；
   4.拼接请求方法、url、post-body及应用的app_master_secret；
   5.将D形成字符串计算MD5值，形成一个32位的十六进制（字母小写）字符串，
	即为本次请求sign（签名）的值；Sign=MD5($http_method$url$post-body$app_master_secret);
*/
func signIOS(url,body string) string {
	httpMethod := "POST"
	str := fmt.Sprintf("%v%v%v%v",httpMethod,url,body,secIOS)
	return tool2.MD5(str)
}
func signAndroid(url,body string) string {
	httpMethod := "POST"
	str := fmt.Sprintf("%v%v%v%v",httpMethod,url,body,secAndroid)
	return tool2.MD5(str)
}
//发起post请求  json 内容体
func post(url,jsonStr string) (string,error) {
	resp, err := http.Post(url,
		"application/json",
		strings.NewReader(jsonStr))
	if err != nil {
		return "",err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "",err
	}
	return string(body),nil
}

