package sms

import (
	"encoding/json"
	"errors"
	"fmt"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dysmsapi "github.com/alibabacloud-go/dysmsapi-20170525/v3/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AliSender struct {
	RegionId     string
	AccessKey    string
	AccessSecret string
	SignName     string
	TemplateCode string
}

func (s *AliSender) Send(mobile string, content map[string]string) (err error) {
	config := &openapi.Config{
		Endpoint:        tea.String("dysmsapi.aliyuncs.com"), // Endpoint 请参考 https://api.aliyun.com/product/Dysmsapi
		AccessKeyId:     tea.String(s.AccessKey),
		AccessKeySecret: tea.String(s.AccessSecret),
	}

	client, err := dysmsapi.NewClient(config)
	if err != nil {
		return err
	}

	param, _ := json.Marshal(content)

	sendSmsRequest := &dysmsapi.SendSmsRequest{
		PhoneNumbers:  tea.String(mobile),
		SignName:      tea.String(s.SignName),
		TemplateCode:  tea.String(s.TemplateCode),
		TemplateParam: tea.String(string(param)),
	}

	runtime := &util.RuntimeOptions{}
	response := &dysmsapi.SendSmsResponse{}
	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		response, err = client.SendSmsWithOptions(sendSmsRequest, runtime)
		if err != nil {
			return err
		}
		return nil
	}()

	if tryErr != nil {
		var error = &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			error = _t
		} else {
			error.Message = tea.String(tryErr.Error())
		}
		// 如有需要，请打印 error
		_, err = util.AssertAsString(error.Message)
		if err != nil {
			return err
		}
	}
	fmt.Printf("%+v \r\n", response.Body)
	if tea.StringValue(response.Body.Code) != "OK" {
		switch tea.StringValue(response.Body.Code) {
		case "isv.SMS_SIGNATURE_SCENE_ILLEGAL":
			return errors.New("签名和模板类型不一致")
		case "isv.BUSINESS_LIMIT_CONTROL":
			return errors.New("获取短信过于频繁，请稍后再试")
		case "isv.AMOUNT_NOT_ENOUGH":
			fmt.Println("账户余额不足")
			// 触发告警
		default:
			return errors.New("阿里云短信发送失败" + tea.StringValue(response.Body.Message))
		}
	}
	return nil
}
