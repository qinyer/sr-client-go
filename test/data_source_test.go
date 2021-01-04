package test

import (
	"fmt"
	sr "github.com/qinyer/sr-client-go"
	"github.com/qinyer/sr-client-go/config"
	"github.com/qinyer/sr-client-go/models"
	"testing"
)

var cfg = &config.YouShuConfig{
	AppId:     "bi4fde6c11c8c44589",
	AppSecret: "1a61f2b77a2c472b94beacdec5cbd018",
}

var merchantId = "10001995"

func TestDataSourceAdd(t *testing.T) {
	res, err := sr.NewYouShuClient(cfg).DataSource().Add(merchantId, models.DsAnalysisVisitDistribution)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(res)
}

func TestDataSourceGet(t *testing.T) {
	res, err := sr.NewYouShuClient(cfg).DataSource().Get(merchantId, models.DsAnalysisVisitDistribution)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(res)
}
