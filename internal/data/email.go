package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "github.com/zero-one-cloud/notification/api/notification/v1"
	"github.com/zero-one-cloud/notification/internal/biz"
	emailHelper "github.com/zero-one-cloud/notification/pkg/email"
)

const SyncSendEmail = 1  // 同步
const AsyncSendEmail = 2 // 异步

const (
	SendEmailWait = iota
	SendEmailSuccess
	SendEmailFail
)

type EmailEntity struct {
	BaseFields
	UserID  uint   `gorm:"not null;type:int(10) unsigned;comment:'用户id'"`
	Address string `gorm:"not null;type:varchar(255);comment:'邮箱地址'"`
	Subject string `gorm:"not null;type:varchar(255);comment:'主题'"`
	Content string `gorm:"not null;type:text;comment:'内容'"`
	Type    uint8  `gorm:"not null;default:1;type:tinyint(1) unsigned;comment:'1 同步 2 异步'"`
	Status  uint8  `gorm:"not null;default:0;type:tinyint(1) unsigned;comment:'状态0 未发送 1 成功 2 失败'"`
}

func (EmailEntity) TableName() string {
	return "email"
}

type EmailRepo struct {
	data *Data
	log  *log.Helper
}

func NewEmailRepo(data *Data, logger log.Logger) biz.EmailRepo {
	return &EmailRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/email")),
	}
}

func (s EmailRepo) SendEmail(ctx context.Context, email string, subject, content string) error {
	ent := &EmailEntity{}
	ent.UserID = 0
	ent.Address = email
	ent.Subject = subject
	ent.Content = content
	ent.Type = 1
	ent.Status = SendEmailWait
	// 保存记录
	err := s.data.db.Model(&EmailEntity{}).Create(&ent).Error
	if err != nil {
		return v1.ErrorSystemError("保存失败").WithCause(err)
	}
	status := SendEmailSuccess
	err = emailHelper.SendEmail(
		s.data.cfg.Email.Host,
		int(s.data.cfg.Email.Port),
		s.data.cfg.Email.User,
		s.data.cfg.Email.Password,
		email,
		subject,
		content,
	)
	if err != nil {
		status = SendEmailFail
		return v1.ErrorSystemError("发送失败").WithCause(err)
	}
	// todo: 异步kafka发送
	// 更新发送状态
	err = s.data.db.Model(&EmailEntity{}).Where("id = ?", ent.Id).UpdateColumns(map[string]interface{}{
		"status": status,
	}).Error
	return nil
}
