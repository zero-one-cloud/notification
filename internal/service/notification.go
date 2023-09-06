package service

import (
	"github.com/go-kratos/kratos/v2/log"
	v1 "github.com/zero-one-cloud/shop-notification/api/notification/v1"
	"github.com/zero-one-cloud/shop-notification/internal/biz"
)

type NotificationService struct {
	v1.UnimplementedNotificationServiceServer

	smsUC   *biz.SmsUseCase
	emailUC *biz.EmailUseCase
	log     *log.Helper
}

func NewNotificationService(smsUC *biz.SmsUseCase, emailUc *biz.EmailUseCase, logger log.Logger) *NotificationService {
	return &NotificationService{
		log:     log.NewHelper(log.With(logger, "module", "service/NotificationService")),
		smsUC:   smsUC,
		emailUC: emailUc,
	}
}
