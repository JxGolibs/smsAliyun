package smsAliyun

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//电话号码实名验证
func PhoneCertCheck(realName, cardNo, phone, appcode string) error {
	req, err := http.NewRequest("GET", "https://auditphone.showapi.com/phoneAudit", nil)
	if err != nil {
		fmt.Println("err:", err)
		return err
	}
	q := req.URL.Query()
	q.Add("phone", phone)
	q.Add("idCard", cardNo)
	q.Add("name", realName)
	q.Add("needBelongArea", "true")
	req.URL.RawQuery = q.Encode()
	req.Header.Set("Authorization", "APPCODE "+appcode)

	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()

	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	result := struct {
		ShowapiResBody struct {
			BelongArea struct {
				AreaCode string `json:"areaCode"`
				City     string `json:"city"`
				Name     string `json:"name"`
				Num      int    `json:"num"`
				PostCode string `json:"postCode"`
				Prov     string `json:"prov"`
				ProvCode string `json:"provCode"`
				Type     int    `json:"type"`
			} `json:"belongArea"`
			Code    int    `json:"code"`
			Msg     string `json:"msg"`
			RetCode int    `json:"ret_code"`
		} `json:"showapi_res_body"`
		ShowapiResCode  int    `json:"showapi_res_code"`
		ShowapiResError string `json:"showapi_res_error"`
	}{}

	json.Unmarshal(body, &result)

	if result.ShowapiResBody.RetCode != 0 {
		fmt.Println(string(body))
		return fmt.Errorf(result.ShowapiResBody.Msg)
	}
	// fmt.Println(result)
	return nil
}
