package service

import (
	"context"
	v1 "github.com/zero-one-cloud/notification/api/notification/v1"
)

func (s *NotificationService) SendSmsVerifyCode(ctx context.Context, req *v1.SendSmsVerifyCodeReq) (*v1.SendSmsVerifyCodeReply, error) {
	smsId, status, code, err := s.smsUC.SendSmsVerifyCode(ctx, 1, req.Phone)
	if err != nil {
		return nil, err
	}
	return &v1.SendSmsVerifyCodeReply{
		SmsId:  smsId,
		Status: status,
		Code:   code,
	}, nil
}

func (s *NotificationService) SendSms(ctx context.Context, req *v1.SendSmsReq) (*v1.SendSmsReply, error) {
	return &v1.SendSmsReply{}, nil
}
