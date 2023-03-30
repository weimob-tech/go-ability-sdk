package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GuideQueryGuideTradeListByScroll
// @id 1587
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=获取导购业绩明细（即将下线）)
func (client *Client) GuideQueryGuideTradeListByScroll(request *GuideQueryGuideTradeListByScrollRequest) (response *GuideQueryGuideTradeListByScrollResponse, err error) {
	rpcResponse := CreateGuideQueryGuideTradeListByScrollResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GuideQueryGuideTradeListByScrollRequest struct {
	*api.BaseRequest

	StartId         int64  `json:"startId,omitempty"`
	PageSize        int    `json:"pageSize,omitempty"`
	PerformanceType int    `json:"performanceType,omitempty"`
	StoreId         int64  `json:"storeId,omitempty"`
	GuiderWid       int64  `json:"guiderWid,omitempty"`
	StartTime       string `json:"startTime,omitempty"`
	EndTime         string `json:"endTime,omitempty"`
}

type GuideQueryGuideTradeListByScrollResponse struct {
	Id                          int64  `json:"id,omitempty"`
	StoreId                     int64  `json:"storeId,omitempty"`
	TradeId                     int64  `json:"tradeId,omitempty"`
	GuiderWid                   int64  `json:"guiderWid,omitempty"`
	GuiderName                  string `json:"guiderName,omitempty"`
	GuiderPhone                 string `json:"guiderPhone,omitempty"`
	TradeType                   string `json:"tradeType,omitempty"`
	OrderType                   string `json:"orderType,omitempty"`
	PaymentAmount               int64  `json:"paymentAmount,omitempty"`
	BalanceDiscountAmount       int64  `json:"balanceDiscountAmount,omitempty"`
	PointsAmount                int64  `json:"pointsAmount,omitempty"`
	SharedPayAmount             int64  `json:"sharedPayAmount,omitempty"`
	SharedBalanceDiscountAmount int64  `json:"sharedBalanceDiscountAmount,omitempty"`
	SharedPointsAmount          int64  `json:"sharedPointsAmount,omitempty"`
	Wid                         int64  `json:"wid,omitempty"`
	UserNickname                string `json:"userNickname,omitempty"`
	TradeTime                   string `json:"tradeTime,omitempty"`
	TradeUpdateTime             string `json:"tradeUpdateTime,omitempty"`
	CommissionAmount            int64  `json:"commissionAmount,omitempty"`
	PerformanceType             string `json:"performanceType,omitempty"`
}

func CreateGuideQueryGuideTradeListByScrollRequest() (request *GuideQueryGuideTradeListByScrollRequest) {
	request = &GuideQueryGuideTradeListByScrollRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "guide/queryGuideTradeListByScroll", "api")
	request.Method = api.POST
	return
}

func CreateGuideQueryGuideTradeListByScrollResponse() (response *api.BaseResponse[GuideQueryGuideTradeListByScrollResponse]) {
	response = api.CreateResponse[GuideQueryGuideTradeListByScrollResponse](&GuideQueryGuideTradeListByScrollResponse{})
	return
}
