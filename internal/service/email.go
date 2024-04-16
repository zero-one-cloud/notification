package service

import (
	"context"
	v1 "github.com/zero-one-cloud/notification/api/notification/v1"
)

func (s *NotificationService) SendEmail(ctx context.Context, req *v1.SendEmailReq) (*v1.SendEmailReply, error) {
	err := s.emailUC.SendEmail(ctx, req.Address, req.Subject, req.Content)
	if err != nil {
		return nil, err
	}
	return &v1.SendEmailReply{Success: true}, nil
}
