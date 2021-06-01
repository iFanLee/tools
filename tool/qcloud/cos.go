package qcloud

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sts "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sts/v20180813"
	"github.com/tencentyun/cos-go-sdk-v5"

)

/**
 * @Author: Lee
 * @Date: 2021/6/1 15:25
 * @Desc:
 */
const (
	secID = ""
	secKey = ""
	region = ""
	uid = ""
	bucket = ""
	expireTime = 10
	urlLifetime = 20
	urlBase = "https://"
)
//获取腾讯云上传凭证
func GetUploadSts()  (string,error){
	var  err error
	m := make(map[string]interface{})
	m["Name"] = "app"
	stat := make(map[string]interface{})
	stat["effect"] = "allow"
	var action = []string{"name/cos:PutObject","name/cos:InitiateMultipartUpload","name/cos:ListParts",
		"name/cos:UploadPart","name/cos:CompleteMultipartUpload","name/cos:AbortMultipartUpload",
		"name/cos:PostObject","name/cos:HeadObject","name/cos:GetObject","name/cos:ListMultipartUploads"}
	stat["action"] = action
	var resource = []string{fmt.Sprintf("qcs::cos:%s:uid/%s:%s/pub/*",region,uid,bucket),fmt.Sprintf("qcs::cos:%s:uid/%s:%s/pri/*",region,uid,bucket)}
	stat["resource"] = resource
	statement := make(map[string]interface{})
	statement["statement"] = stat
	statement["version"] = "2.0"
	str,err := json.Marshal(statement)
	if nil != err{
		return "",err
	}
	m["Policy"] = string(str)
	m["DurationSeconds"] = expireTime
	data,err := json.Marshal(m)
	if nil != err{
		return "",err
	}
	credential := common.NewCredential(
		secID,
		secKey,
	)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "sts.tencentcloudapi.com"
	client, _ := sts.NewClient(credential,region, cpf)
	request := sts.NewGetFederationTokenRequest()
	params := string(data)
	err = request.FromJsonString(params)
	if err != nil {
		return "",err
	}
	response, err := client.GetFederationToken(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		return "",err
	}
	if err != nil {
		return "",err
	}
	return response.ToJsonString(),nil
}
//腾讯云对象存储预签名
func QcloudPreSignUrl(str string) string {
	path :=strings.Split(str,".com/")
	lifetime := urlLifetime
	u,_ := url.Parse(urlBase)
	b := &cos.BaseURL{BucketURL:u}
	c := cos.NewClient(b,&http.Client{Transport:&cos.AuthorizationTransport{SecretID:secID,SecretKey:secKey}})
	if urls,err := c.Object.GetPresignedURL(context.Background(),http.MethodGet,path[1],secID,secKey,time.Second*time.Duration(lifetime),nil);err==nil{
		return urls.String()
	}
	return ""
}