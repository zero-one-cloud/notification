package service

import (
	"github.com/go-kratos/kratos/v2/log"
	v1 "github.com/zero-one-cloud/shop-notification/api/notification/v1"
	"github.com/zero-one-cloud/shop-notification/internal/biz"
)

type NotificationService struct {
	v1.UnimplementedNotificationServiceServer

	smsUC *biz.SmsUseCase
	log   *log.Helper
}

func NewNotificationService(smsUC *biz.SmsUseCase, logger log.Logger) *NotificationService {
	return &NotificationService{
		log:   log.NewHelper(log.With(logger, "module", "service/NotificationService")),
		smsUC: smsUC,
	}
}
