package qcloud

import (
	"fmt"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20190711"

)

/**
 * @Author: Lee
 * @Date: 2021/6/1 15:34
 * @Desc:
 */

func SmsSendBySingleByTemp(appid,sign,tel string,tempID int64,params []string) int {
	credential := common.NewCredential(
		secID,
		secKey,
	)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "sms.tencentcloudapi.com"
	client, _ := sms.NewClient(credential, "", cpf)
	request := sms.NewSendSmsRequest()
	request.PhoneNumberSet = common.StringPtrs([]string{ tel })
	request.TemplateID = common.StringPtr(fmt.Sprintf("%v",tempID))
	request.SmsSdkAppid = common.StringPtr(appid)
	request.Sign = common.StringPtr(sign)
	request.TemplateParamSet = common.StringPtrs(params)
	response, err := client.SendSms(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		//fmt.Printf("An API error has returned: %s", err)
		return 3000
	}
	if err != nil {
		//fmt.Println("err--->",err)
		return 3000
	}

	/*
		失败:
		{
		  "Response": {
		    "SendStatusSet": [
		      {
		        "SerialNo": "",
		        "PhoneNumber": "+11111",
		        "Fee": 0,
		        "SessionContext": "",
		        "Code": "FailedOperation.TemplateIncorrectOrUnapproved",
		        "Message": "template is not approved or request content does not match the approved template content",
		        "IsoCode": ""
		      }
		    ],
		    "RequestId": "734100bf-f"
		  }
		}


		成功：
		{
		  "Response": {
		    "SendStatusSet": [
		      {
		        "SerialNo": "8:Z1V111111",
		        "PhoneNumber": "+861111111",
		        "Fee": 1,
		        "SessionContext": "",
		        "Code": "Ok",
		        "Message": "send success",
		        "IsoCode": "CN"
		      }
		    ],
		    "RequestId": "971eb80b-7633333"
		  }
		}
	*/

	//错误码 https://cloud.tencent.com/document/product/382/38778
	if len(response.Response.SendStatusSet)>0{
		switch *response.Response.SendStatusSet[0].Code {
		case "ok","Ok","OK":
			return 0
		case "FailedOperation.InsufficientBalanceInSmsPackage"://套餐包余量不足
			return 2000
		case "LimitExceeded.AppDailyLimit"://业务短信日下发条数超过设定的上限
			return 2001
		case "LimitExceeded.DailyLimit"://短信日下发条数超过设定的上限
			return 2002
		case "LimitExceeded.DeliveryFrequencyLimit"://下发短信命中了频率限制策略
			return 2003
		case "LimitExceeded.PhoneNumberCountLimit"://调用短信发送 API 接口单次提交的手机号个数超过200个
			return 2004
		case "LimitExceeded.PhoneNumberDailyLimit"://单个手机号日下发短信条数超过设定的上限
			return 2005
		case "LimitExceeded.PhoneNumberOneHourLimit"://单个手机号1小时内下发短信条数超过设定的上限
			return 2006
		case "LimitExceeded.PhoneNumberSameContentDailyLimit"://单个手机号下发相同内容超过设定的上限
			return 2007
		case "LimitExceeded.PhoneNumberThirtySecondLimit"://单个手机号30秒内下发短信条数超过设定的上限
			return 2008
		default:
			return 3000
		}
	}
	//fmt.Printf("%s", response.ToJsonString())
	return 3000
}