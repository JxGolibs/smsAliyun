package smsAliyun

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
)

type (
	Config struct {
		RegionId      string  `default:"cn-hangzhou"`
		AccessKeyId   string
		AccessSecret  string
		SignName      string //阿里云验证过的项目名 自己设置
		TemplateCode  string //阿里云的短信模板号 自己设置
	}
	
)

var sclient = struct {
	client  *dysmsapi.Client
	request *dysmsapi.SendSmsRequest
}{}

func New(cnf Config)   {
	sclient.client, _ = dysmsapi.NewClientWithAccessKey(cnf.RegionId, cnf.AccessKeyId, cnf.AccessSecret)
	sclient.request = dysmsapi.CreateSendSmsRequest()
	sclient.request.Scheme = "https"
	sclient.request.Domain = "dysmsapi.aliyuncs.com"
	sclient.request.SignName = cnf.SignName
	sclient.request.TemplateCode = cnf.TemplateCode
}

func  SendSms(PhoneNumbers string, code string) error {
	sclient.request.PhoneNumbers = PhoneNumbers
	sclient.request.TemplateParam ="{\"code\":" + code  + "}"  //短信模板中的验证码内容 自己生成   之前试过直接返回，但是失败，加上code成功。

	response, err := sclient.client.SendSms(sclient.request)
	if err != nil {
		fmt.Printf(err.Error(), "response is %#v\n", response)
	}
    return err
}
