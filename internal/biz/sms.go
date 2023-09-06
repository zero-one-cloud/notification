package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type SmsRepo interface {
	SendSmsVerifyCode(ctx context.Context, mobile string, sendType int64) (int64, int64, string, error)
}

type SmsUseCase struct {
	repo SmsRepo
	log  *log.Helper
}

func NewSmsUseCase(repo SmsRepo, logger log.Logger) *SmsUseCase {
	return &SmsUseCase{repo: repo, log: log.NewHelper(log.With(logger, "module", "useCase/sms"))}
}

func (suc *SmsUseCase) SendSmsVerifyCode(ctx context.Context, sendType int64, mobile string) (int64, int64, string, error) {
	return suc.repo.SendSmsVerifyCode(ctx, mobile, sendType)
}
