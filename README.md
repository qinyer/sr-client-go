# 腾讯有数

## 官方文档
>[官方文档](https://mp.zhls.qq.com/youshu-docs/develop/dev_account/dev_account_access.html)
## 数据仓库类型表
| 数据仓库类型 | 类型编码| 描述 |
| :------| :------ | :------ |
| 主订单	| 0	| 订单数据仓库, 对应上报接口 /order/add_order|
| 子订单	| 1	| 子订单数据仓库，对应上报接口 /order/add_sub_order|
| 销售信息	| 2 | 	销售信息数据仓库，对应上报接口 /salesinfo/add|
| 商品 sku	| 3	| 商品 sku 数据仓库，对应上报接口 /sku/add|
| 门店信息	| 4	| 门店信息数据仓库，对应上报接口 /store/add|
| 商品 spu	| 5	| 商品 spu 数据仓库，对应上报接口 /spus/add|
| 商品类目	| 6	| 商品类目数据仓库，对应上报接口 /product_categories/add|
| 微信小程序页面访问	| 7 | 	微信小程序页面访问信息数据仓库，对应微信接口 analysis.getVisitPage，对应上报接口 /analysis/add_wxapp_visit_page|
| 微信小程序访问分布数据	| 8	| 微信小程序页面访问信息数据仓库，对应微信接口 analysis.getVisitDistribution，对应上报接口 /analysis/add_wxapp_visit_distribution|
| 用户信息	| 11 | 	用户信息数据仓库，对应上报接口 /user/add_user|
| 活动信息	| 12| 	活动信息数据仓库，对应上报接口 /order/add_promotion|
| 卡券信息	| 13| 	卡券信息数据仓库，对于上报接口 /order/add_coupon|
| 微信小程序日访问趋势	| 16 | 	微信小程序日访问趋势数据仓库，对应微信接口 analysis.getDailyVisitTrend，对应上报接口 /analysis/add_wxapp_daily_visit_trend|
| 汇总订单	| 51| 	汇总订单数据仓库, 对应上报接口 /order/add_order_sum|

----

### 配置
```go
// 有数基础配置
type YouShuConfig struct {
	AppId     string // appId 数据中心分配的有数应用ID
	AppSecret string // appSecret 数据中心分配的有数签名密钥，需要绝对保密
	Prod      bool   // prod 是否为生产环境 默认走测试环境
}
```
### 调用示例
```go
var cfg = &config.YouShuConfig{
	AppId:     "bi39c1fde6f4b3489b",
	AppSecret: "335163180ea24290ba822120858439fe",
	Prod: false,
}

func TestDataSourceAdd(t *testing.T) {
	res, err := sr.NewYouShuClient(cfg).DataSource().Add("9999999", models.DsOrderCoupon)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(res)
}
```