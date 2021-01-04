package sr_client_go

import (
	"github.com/sirupsen/logrus"
	"os"
	"sr-client-go/api"
	"sr-client-go/config"
	"sr-client-go/context"
)

// init log
func init() {
	// Log as JSON instead of the default ASCII formatter.
	logrus.SetFormatter(&logrus.TextFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	logrus.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	logrus.SetLevel(logrus.DebugLevel)
}

// YouShu
type YouShu struct {
	ctx *context.Context
}

// NewYouShuClient 获取有数上报实例
func NewYouShuClient(cfg *config.YouShuConfig) *YouShu {
	ctx := &context.Context{
		YouShuConfig: cfg,
	}
	return &YouShu{ctx: ctx}
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
