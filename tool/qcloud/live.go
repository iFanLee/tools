package qcloud

import (
	"fmt"
	"tools/tool"
)

/**
 * @Author: Lee
 * @Date: 2021/6/1 15:31
 * @Desc:
 */
const (
	pullKey = ""
	pullDomain = ""
)
//获取腾讯云直播拉流地址
type pu struct {
	Rtmp,Flv,Hls string
}
func GenLivePullUrl(streamName string,expireTime int64)  *pu{
	secret := tool.MD5(fmt.Sprintf("%v%v%v",pullKey,streamName,expireTime))
	domain := fmt.Sprintf("%v%v",pullDomain,streamName)
	sec := fmt.Sprintf("?txSecret=%v&txTime=%v",secret,expireTime)
	return &pu{
		Rtmp: "rtmp://"+domain + sec,
		Flv:  "https://"+domain + ".flv"+ sec,
		Hls:  "https://"+domain + ".m3u8"+ sec,
	}
}