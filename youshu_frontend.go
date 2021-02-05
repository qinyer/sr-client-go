package sr_client_go

import (
	"github.com/qinyer/sr-client-go/api/frontend"
	"github.com/qinyer/sr-client-go/config"
	"github.com/qinyer/sr-client-go/context"
	"github.com/qinyer/sr-client-go/utils"
	"github.com/sirupsen/logrus"
	"os"
)

// YouShuFrontend
type YouShuFrontend struct {
	ctx *context.Context
}

// NewYouShuFrontendClient 获取有数前端上报实例
func NewYouShuFrontendClient(cfg *config.YouShuConfig) *YouShuFrontend {
	return &YouShuFrontend{&context.Context{
		YouShuConfig: cfg,
		Logger: &logrus.Logger{
			Out:          os.Stdout,
			Formatter:    &utils.CustomerFormatter{Prefix: "youshu-frontend"},
			Level:        logrus.DebugLevel,
			ExitFunc:     os.Exit,
			ReportCaller: true,
		},
	}}
}

// Order 订单相关
func (ys *YouShuFrontend) Order() *frontend.Order {
	return frontend.NewOrder(ys.ctx)
}

// Cart 购物车相关
func (ys *YouShuFrontend) Cart() *frontend.Cart {
	return frontend.NewCart(ys.ctx)
}
