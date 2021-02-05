package sr_client_go

import (
	"github.com/qinyer/sr-client-go/api"
	"github.com/qinyer/sr-client-go/config"
	"github.com/qinyer/sr-client-go/context"
	"github.com/qinyer/sr-client-go/utils"
	"github.com/sirupsen/logrus"
	"os"
)

// YouShu
type YouShu struct {
	ctx *context.Context
}

// NewYouShuClient 获取有数上报实例
func NewYouShuClient(cfg *config.YouShuConfig) *YouShu {
	return &YouShu{&context.Context{
		YouShuConfig: cfg,
		Logger: &logrus.Logger{
			Out:          os.Stdout,
			Formatter:    &utils.CustomerFormatter{Prefix: "youshu-backend"},
			Level:        logrus.DebugLevel,
			ExitFunc:     os.Exit,
			ReportCaller: true,
		},
	}}
}

// DataSource 数据仓库
func (ys *YouShu) DataSource() *api.DataSource {
	return api.NewDataSource(ys.ctx)
}

// Analysis 微信数据
func (ys *YouShu) Analysis() *api.Analysis {
	return api.NewAnalysis(ys.ctx)
}

// Order 订单
func (ys *YouShu) Order() *api.Order {
	return api.NewOrder(ys.ctx)
}

// ProductCategory 商品类目
func (ys *YouShu) ProductCategory() *api.ProductCategory {
	return api.NewProductCategory(ys.ctx)
}

// SalesInfo 销售信息
func (ys *YouShu) SalesInfo() *api.SalesInfo {
	return api.NewSalesInfo(ys.ctx)
}

// Sku 商品Sku
func (ys *YouShu) Sku() *api.Sku {
	return api.NewSku(ys.ctx)
}

// Spu 商品Spu
func (ys *YouShu) Spu() *api.Spu {
	return api.NewSpu(ys.ctx)
}

// Store 门店仓库
func (ys *YouShu) Store() *api.Store {
	return api.NewStore(ys.ctx)
}

// User 会员
func (ys *YouShu) User() *api.User {
	return api.NewUser(ys.ctx)
}
