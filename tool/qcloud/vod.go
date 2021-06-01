package qcloud

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
	tool2 "tools/tool"
)

/**
 * @Author: Lee
 * @Date: 2021/6/1 15:36
 * @Desc:
 */
const (
	vodKey = ""
	vodDomain = ""
	vodAppID = int64(0)
	vodRlimit = 3
	vodPlayerPcfg = ""
)
//生成签名url
func PriCreateVodUrl(url string,expireTime int64,exper,rlimit int) string {
	split := strings.Split(url,"vod2.myqcloud.com")
	if len(split)!=2{
		return ""
	}
	split1 := strings.Split(split[1],"/")
	length := len(split1)-1
	dir := "/"
	for i:=1;i<length;i++{
		dir += split1[i]+"/"
	}
	txTime := strings.ToLower(fmt.Sprintf("%X",expireTime))
	us := tool2.RandStr(15,4)
	str := fmt.Sprintf("%v%v%v%v%v%v",vodKey,dir,txTime,exper,rlimit,us)
	sign := tool2.MD5(str)
	domain := vodDomain
	urlSign := fmt.Sprintf("%v%v?t=%v&exper=%v&rlimit=%v&us=%v&sign=%v",domain,split[1],txTime,exper,rlimit,us,sign)
	return  urlSign
}


type QVUrlSign struct {
	AppId int64 `json:"appId"`
	FileId string `json:"fileId"`
	CurrentTimeStamp int64 `json:"currentTimeStamp"`
	ExpireTimeStamp int64 `json:"expireTimeStamp"`
	UrlAccessInfo QVUrlSignAccessInfo `json:"urlAccessInfo"`
	Pcfg string `json:"pcfg"`
}
type QVUrlSignAccessInfo struct {
	T string `json:"t"`
	Rlimit int `json:"rlimit"`
	Us string `json:"us"`
}
//点播SDK播放token
func PriQVodUrlToken(fileId string) string {
	ts := time.Now().Unix()
	ex := ts+int64(urlLifetime)
	payload := QVUrlSign{
		AppId:            vodAppID,
		FileId:           fileId,
		CurrentTimeStamp: ts,
		ExpireTimeStamp:  ex,
		UrlAccessInfo:    QVUrlSignAccessInfo{
			T:      fmt.Sprintf("%X",ex),
			Rlimit: vodRlimit,
			Us:     tool2.RandStr(10,4),
		},
		Pcfg:vodPlayerPcfg,
	}
	payloadStr,_ := json.Marshal(&payload)
	header := `{"alg":"HS256","typ":"JWT"}`
	s1 := tool2.Base64UrlEncode(header)
	s2 := tool2.Base64UrlEncode(string(payloadStr))
	s3 := s1+"."+s2
	s4 := tool2.HmacSha256(s3,vodKey)
	token := s3+"."+tool2.Base64UrlEncode(string(s4))
	return token
}
//点播链接签名
func PriQVodUrlSign(url string) string {
	dir := tool2.GetUrlDir(url)
	t := fmt.Sprintf("%X",time.Now().Unix()+int64(urlLifetime))
	us := tool2.RandStr(10,4)
	sign := tool2.MD5(fmt.Sprintf("%v%v%v%v%v",vodKey,dir,t,vodRlimit,us))
	return fmt.Sprintf("%v?t=%v&rlimit=%v&us=%v&sign=%v",url,t,vodRlimit,us,sign)
}