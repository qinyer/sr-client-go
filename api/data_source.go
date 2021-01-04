package api

import (
	"encoding/json"
	"net/url"
	"sr-client-go/context"
	"sr-client-go/models"
	"sr-client-go/utils"
	"strconv"
)

type DataSource struct {
	*context.Context
}

// NewDataSource 获取数据仓库上报实例
func NewDataSource(ctx *context.Context) *DataSource {
	dataSource := new(DataSource)
	dataSource.Context = ctx
	return dataSource
}

// Add 添加数据仓库
func (ds *DataSource) Add(merchantId string, dataSourceType int) (dataSource models.DataSourceRes, err error) {
	params := map[string]interface{}{
		"merchantId":     merchantId,
		"dataSourceType": dataSourceType,
	}
	response, err := utils.NewHttpClient(ds.AppId, ds.AppSecret, ds.Prod).HttpPostJson("/data_source/add", params)
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(response), &dataSource)
	if err != nil {
		return
	}
	return
}

// Get 获取数据仓库
func (ds *DataSource) Get(merchantId string, dataSourceType int) (dataSource models.DataSourceGetRes, err error) {
	params := url.Values{}
	params.Add("merchantId", merchantId)
	params.Add("dataSourceType", strconv.Itoa(dataSourceType))
	response, err := utils.NewHttpClient(ds.AppId, ds.AppSecret, ds.Prod).HttpGetJson("/data_source/get", params)
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(response), &dataSource)
	if err != nil {
		return
	}
	return
}
