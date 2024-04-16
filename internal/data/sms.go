package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "github.com/zero-one-cloud/notification/api/notification/v1"
	"github.com/zero-one-cloud/notification/internal/biz"
	"github.com/zero-one-cloud/notification/pkg/sms"
	"github.com/zero-one-cloud/notification/pkg/utils/rand"
)

const smsCodeLength = 6

const SyncSendSms = 1  // 同步
const AsyncSendSms = 2 // 异步

const (
	SendSmsWait = iota
	SendSmsSuccess
	SendSmsFail
)

type SmsEntity struct {
	BaseFields
	SignName     string
	TemplateCode string
	Mobile       string
	Content      string
	Type         int64
	Status       int64
	ErrReason    string
}

func (SmsEntity) TableName() string {
	return "sms_verify_code"
}

type SmsRepo struct {
	data *Data
	log  *log.Helper
}

func NewSmsRepo(data *Data, logger log.Logger) biz.SmsRepo {
	return &SmsRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/notification")),
	}
}

func (s SmsRepo) SendSmsVerifyCode(ctx context.Context, mobile string, sendType int64) (int64, int64, string, error) {
	code := rand.GenerateSmsCode(smsCodeLength)
	status := int64(SendSmsWait)

	ent := &SmsEntity{}
	ent.SignName = s.data.cfg.Sms.Ali.SignName
	ent.TemplateCode = s.data.cfg.Sms.Ali.TemplateCode
	ent.Mobile = mobile
	ent.Content = code
	ent.Type = sendType
	ent.Status = SendSmsWait
	// 保存记录
	err := s.data.db.Model(&SmsEntity{}).Create(&ent).Error
	if err != nil {
		return 0, 0, "", v1.ErrorSystemError("保存失败").WithCause(err)
	}
	if sendType == SyncSendSms {
		sender := &sms.AliSender{
			AccessKey:    s.data.cfg.Sms.Ali.AccessKey,
			AccessSecret: s.data.cfg.Sms.Ali.AccessSecret,
			SignName:     s.data.cfg.Sms.Ali.SignName,
			TemplateCode: s.data.cfg.Sms.Ali.TemplateCode,
		}
		err = sender.Send(mobile, map[string]string{"code": code})
		status = SendSmsSuccess
		errReason := ""
		if err != nil {
			status = SendSmsFail
			errReason = err.Error()
		}
		// 更新发送状态
		err = s.data.db.Model(&SmsEntity{}).Where("id = ?", ent.Id).UpdateColumns(map[string]interface{}{
			"status":     status,
			"err_reason": errReason,
		}).Error
		if err != nil {
			return 0, 0, "", v1.ErrorSystemError("保存失败").WithCause(err)
		}
	}
	return ent.Id, status, code, nil
}
