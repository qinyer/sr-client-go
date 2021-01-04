package test

import (
	"fmt"
	sr "github.com/qinyer/sr-client-go"
	"github.com/qinyer/sr-client-go/models"
	"testing"
)

// TestAnalysisVisitPage
func TestAnalysisVisitPage(t *testing.T) {
	req := &models.AnalysisVisitPageReq{
		DataSourceId: "11268",
		RawMsg: &[]models.VisitPageRawMsg{
			{
				RefDate: "20201228",
				List: &[]models.VisitPageRawMsgList{
					{
						PagePath:       "pages/magazine/index",
						PageVisitPv:    20,
						PageVisitUv:    16,
						PageStaytimePv: 8.273738,
						EntrypagePv:    20,
						ExitpagePv:     16,
						PageSharePv:    14,
						PageShareUv:    11,
					},
				},
			},
		},
	}
	res, err := sr.NewYouShuClient(cfg).Analysis().AddVisitPage(req)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(res)
}

// TestAnalysisVisitTrend
func TestAnalysisVisitTrend(t *testing.T) {
	req := &models.AnalysisVisitTrendReq{
		DataSourceId: "11257",
		RawMsg: &[]models.VisitTrendRawMsg{
			{
				List: &[]models.VisitTrendRawMsgList{
					{
						RefDate:         "20201228",
						SessionCnt:      22828,
						VisitPv:         20,
						VisitUv:         16,
						VisitUvNew:      1,
						StayTimeSession: 0,
						VisitDepth:      1.92727,
					},
				},
			},
		},
	}
	res, err := sr.NewYouShuClient(cfg).Analysis().AddVisitTrend(req)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(res)
}

// TestAnalysisVisitDistribution
func TestAnalysisVisitDistribution(t *testing.T) {
	req := &models.AnalysisVisitDistributionReq{
		DataSourceId: "11257",
		RawMsg: &[]models.VisitDistributionRawMsg{
			{
				RefDate: "20201228",
				List: &[]models.VisitDistributionRawMsgList{
					{
						Index: "access_source_session_cnt",
						ItemList: &[]models.VisitDistributionRawMsgItem{
							{
								Key:   10,
								Value: 777,
							},
						},
					},
				},
			},
		},
	}
	res, err := sr.NewYouShuClient(cfg).Analysis().AddVisitDistribution(req)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(res)
}
