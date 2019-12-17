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
	sclient.request.TemplateParam ="{\"code\":\"" + code  + "\"}"  //短信模板中的验证码内容 自己生成   之前试过直接返回，但是失败，加上code成功。

	response, err := sclient.client.SendSms(sclient.request)
	if err != nil {
		fmt.Printf(err.Error(), "response is %#v\n", response)
	}
    return err
}

// //电话号码实名验证
// func 
// safrv_cert_phone_chec
// verifyKey  IVopb27dPlH7T2
// MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQCz1DichNBoOKMYnE4ijd1siWYmHMdtP1/rebZXPhdeY0dFATjCwJePgOHoX2Hzu/yXq3z8oX8nQHXcGgV5+c22p093PylQTMd/X1cQ/nGJhsUXIYk/Wn5oB0yuWtM1HQ4dkyxR+5DV8w8meMlAPp6BQefYIRMBgly19WhCGqQbacAWRKTeZmqGKdVWVbHCYOabKZFREvUzbtbDSF0YnhIBTuZ+Ct3Q3oVyj6RkJHkvsifIxxz7zFN0cbIPgm1yNa1WxTVImkRu2ee0A1a0HkJFrR3TBIZBiFmDfMX5zrw5Uzy+IQ+qGgIRtUtc72INAeUeJWLAycyDUAvZ31+qEVH7AgMBAAECggEAC7dyENCee4vlJH9an+m6WY4CN818OOP+Let6M+j5wM6bICXaOuYtec4d6fP8/9JpvMARnqwvCrGqfDK91JqwCoyQIubcnDzEasoddB//6cP8J2up85TD06dWGa4VQMqp1sd2BclsQGRbwlfwVN0/mpBeuLg+8sQ5oilhJ51wv9BhxRULmQpcDux9dn++K3mHH6n2wz3CqjATZF+MxWNFsJ2TXvqdbJZ2vu6bBjCpbpgsZNytbYFQTHo2LTN/NYiYEWsCaMDnLCZDx1ZkgRDFbpRkUdFVcQAI+JuNUVWxpTpvcwVfZTKDB3rWELYL8VeMYwQikBy4SsZhVUlrSMMIAQKBgQD6smcWvhcQshELkS+tdPKxfclz25blNW8rZ4Fo/I+h2dpeI1BE7G2zFiRqHrNkS9TsBeuQGonBtK34LAzCEI6NDp/SnqeIQ7X9J49WMXYNMiKrG67nLWGr2xVGK7DMlkVM+mJoSG4d7PoijUQ+qIo+LvBrIygrR46LmrlNfXtmAQKBgQC3ogxHIvNRHkk1BuegG5FN3SkNcGrNNRx42dy1OPODtKEEyK8P4yDiWELSYnpWSJvx9OoKppTWOmXrkJ+7Yxp/92JtluyfzrnsXv+VG1dcUxoKvqtmGmroIjKgAbxkuuk0YjYs7EnqBPbHttdYMBlNVJ3KZXA0RNQUeInrMZpP+wKBgQDFZzdfpB0ba2ualR9a0WzNMEQx86ZVJAc+wUkYd09OXfHiEKKDMzyKORRBb51Ii6HEnEDJ/uucWP5oHZ+KLGljTzGwMnevB8pE5iKq91WCvIip/alypbx0C5Yr4/laj7VQERTuEqKK8/BexqgHQsVgg6uSYWOe+MCY0yO4jedeAQKBgDjBvUm/Y7UOx1G8cJrr6l1GngmpVVlXsRUpQS5CDAIJMjtoS7N9YAyVexHLnkRM7OY77JqeIiKUqivE55njZ61lqOZ2X2yFWLHWKujLKpU+mi34AMHuKFzXNQ5/etsnaizmrhf0cGQEZMIVoTmApFFideK4t5VwMJyjJP2styVBAoGBAJSn0fxBXgWqULhsxnmM5oP6K4tNQ7XkjDVMUXJm9MlaE3qy8LVXDcq5hT+1p7T5jQJzVtVivZWhB05di5bRham1o8i4kkJQYpMcDBfcu5ocbZf5R6W7YYqdak2jYsk1xqU+RKymqCQx3PT+w7MGHRnvS6vK13EuJWZSpaefkJvL

// {"__userId":"1906140863465222","customerID":"","extData":"{\"callBack\":\" http://www.aliyun.com\",\"validator_attr_user_name\":\"张三\",\"validator_attr_card_num\":\"18位的客户身份证号\"}","verifyKey":"IVopb27dPlH7T2"}


// {"__userId":"1906140863465222","customerID":"","extData":"{\"callBack\":\" http://www.aliyun.com\",\"validator_attr_user_name\":\"黄乔加\",\"validator_attr_input_phone\":\"18627783299\",\"validator_attr_card_num\":\"422428197204075196\"}","verifyKey":"IVopb27dPlH7T2"}


// https://www.aliyun.com/?ivSign=fh3bCC8EfC452ZpjHGyAIw3HApSFgAW1wywUekj6NjH%2F6jKdFgsoUo3gCIrFd44DtS6tRO2O9qzf6amGf6G6ydAD5BIALK5KpwAnWbugBjPQlkdTpBdwUDDhnu6xM42I6IM3JE6St6QNUJSx2QqdbxjKHHFJ%2FwvXv9IrXzfDdCVi3AMnjMCpgPm%2F9KSPJ2LyNH2OCYKGTtk8sGN5bXOVzM8Ftwxg170v9db1TOvy5Oia6NE02Apnsb%2B7LRBVy1P03RBETgAR%2BtCtNVP5cVxiI5e3JnLS8NaL1QDR62%2FjIzJu3xzUmGVGRUpSCh7j%2FK%2FxfJc1uGvV%2F214YvfsxBKdqQ%3D%3D&appName=