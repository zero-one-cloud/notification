package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type EmailRepo interface {
	SendEmail(ctx context.Context, email string, subject, content string) error
}

type EmailUseCase struct {
	repo EmailRepo
	log  *log.Helper
}

func NewEmailUseCase(repo EmailRepo, logger log.Logger) *EmailUseCase {
	return &EmailUseCase{repo: repo, log: log.NewHelper(log.With(logger, "module", "useCase/email"))}
}

func (suc *EmailUseCase) SendEmail(ctx context.Context, email string, subject, content string) error {
	return suc.repo.SendEmail(ctx, email, subject, content)
}
