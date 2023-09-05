package service

import (
	"context"
	v1 "github.com/zero-one-cloud/shop-notification/api/notification/v1"
)

func (s *NotificationService) SendEmail(ctx context.Context, req *v1.SendEmailReq) (*v1.SendEmailReply, error) {
	return &v1.SendEmailReply{}, nil
}
