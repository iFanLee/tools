package pay

import (
	"github.com/objcoding/wxpay"
	"time"
)

/**
 * @Author: Lee
 * @Date: 2021/6/1 15:41
 * @Desc:
 */

const (
	wxAppID = ""
	wxMchid = ""
	wxApiKey = ""
	notifyUrl = ""
)
func c() *wxpay.Client {
	acc := wxpay.NewAccount(wxAppID,wxMchid,wxApiKey,false)
	client := wxpay.NewClient(acc)
	client.SetHttpConnectTimeoutMs(2000)
	client.SetHttpReadTimeoutMs(2000)
	client.SetSignType(wxpay.HMACSHA256)
	return client
}
func s(orderNo,body string,fee int64,productID,attach,spIP string,expireAt int64) wxpay.Params {
	//trade_type   JSAPI--JSAPI支付（或小程序支付）、NATIVE--Native支付、APP--app支付，MWEB--H5支付
	params := make(wxpay.Params)
	params.SetString("body", body).//商品描述
		SetString("attach", attach).//附加数据
		SetString("out_trade_no",orderNo).//商户订单号
		SetInt64("total_fee", fee).//标价金额
		SetString("spbill_create_ip", spIP).//终端IP
		SetString("notify_url", notifyUrl).//通知地址
		SetString("time_expire",time.Unix(expireAt,0).Format("20060102150405")).
		SetString("product_id", productID)//trade_type=NATIVE时，此参数必传。此参数为二维码中包含的商品ID，商户自行定义。
	return params
}

//向微信生成Native预付订单
func OrderPreNative(orderNo,body string,fee int64,productID,attach,spIP string,expireAt int64)  (codeUrl string,code int){
	//下单准备
	client := c()
	params := s(orderNo,body,fee,productID,attach,spIP,expireAt)
	params.SetString("device_info", "WEB")//设备号
	params.SetString("trade_type", "NATIVE")//交易类型
	if ret, err := client.UnifiedOrder(params);err ==nil{
		if ret["result_code"]=="SUCCESS"&&ret["return_code"]=="SUCCESS"{
			return ret["code_url"],200
		}else {
			if ret["err_code"]=="ORDERPAID"{
				return "",201
			}else {
				return "",-100
			}
		}
	}else {
		return "",-100
	}
}
//向微信生成H5预付订单
func OrderPreH5(orderNo,body string,fee int64,productID,attach,spIP string,expireAt int64)  (codeUrl string,code int){
	//下单准备
	client := c()
	params := s(orderNo,body,fee,productID,attach,spIP,expireAt)
	params.SetString("device_info", "H5")//设备号
	params.SetString("trade_type", "MWEB")//交易类型
	if ret, err := client.UnifiedOrder(params);err ==nil{
		if ret["result_code"]=="SUCCESS"&&ret["return_code"]=="SUCCESS"{
			return ret["mweb_url"],200
		}else {
			if ret["err_code"]=="ORDERPAID"{
				return "",201
			}else {
				return "",-100
			}
		}
	}else {
		return "",-100
	}
}
//向微信生成APP预付订单
func OrderPreApp(orderNo,body string,fee int64,productID,attach,spIP string,expireAt int64)  (codeUrl string,code int){
	//下单准备
	client := c()
	params := s(orderNo,body,fee,productID,attach,spIP,expireAt)
	params.SetString("device_info", "app")//设备号
	params.SetString("trade_type", "APP")//交易类型
	if ret, err := client.UnifiedOrder(params);err ==nil{
		if ret["result_code"]=="SUCCESS"&&ret["return_code"]=="SUCCESS"{
			return ret["prepay_id"],200
		}else {
			if ret["err_code"]=="ORDERPAID"{
				return "",201
			}else {
				return "",-100
			}
		}
	}else {
		return "",-100
	}
}
//向微信生成jsapi预付订单
func OrderPreJsapi(orderNo,body string,fee int64,productID,attach,spIP string,expireAt int64)  (codeUrl string,code int){
	//下单准备
	client := c()
	params := s(orderNo,body,fee,productID,attach,spIP,expireAt)
	params.SetString("device_info", "jsapi")//设备号
	params.SetString("trade_type", "JSAPI")//交易类型
	if ret, err := client.UnifiedOrder(params);err ==nil{
		if ret["result_code"]=="SUCCESS"&&ret["return_code"]=="SUCCESS"{
			return ret["prepay_id"],200
		}else {
			if ret["err_code"]=="ORDERPAID"{
				return "",201
			}else {
				return "",-100
			}
		}
	}else {
		return "",-100
	}
}










//校验sign
func NotifySignCheck(bt []byte) (wxpay.Params,bool) {
	//将xml转换成map
	xmlMap := wxpay.XmlToMap(string(bt))
	params := make(wxpay.Params)
	for k,v := range xmlMap{
		params.SetString(k,v)
	}
	acc := wxpay.NewAccount(wxAppID,wxMchid,wxApiKey,false)
	client := wxpay.NewClient(acc)
	client.SetSignType(wxpay.HMACSHA256)
	flag := false
	if ch := client.ValidSign(params);ch{//校验通过
		flag = true
	}
	return xmlMap,flag
}