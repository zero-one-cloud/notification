package server

import (
	"github.com/go-kratos/kratos/v2/middleware/validate"
	v1 "github.com/zero-one-cloud/notification/api/notification/v1"
	"github.com/zero-one-cloud/notification/internal/conf"
	"github.com/zero-one-cloud/notification/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, notification *service.NotificationService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.ResponseEncoder(respEncoder), // 用来将用户pb定义里的reply结构体序列化后写入Response Body中
		http.ErrorEncoder(errorEncoder),   // 用来将业务抛出的error序列化后写入Response Body中
		http.Middleware(
			recovery.Recovery(),
			validate.Validator(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	v1.RegisterNotificationServiceHTTPServer(srv, notification)
	return srv
}
