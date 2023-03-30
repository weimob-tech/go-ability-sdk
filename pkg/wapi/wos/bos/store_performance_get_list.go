package bos

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// StorePerformanceGetList
// @id 1943
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1943?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=获取门店业绩明细)
func (client *Client) StorePerformanceGetList(request *StorePerformanceGetListRequest) (response *StorePerformanceGetListResponse, err error) {
	rpcResponse := CreateStorePerformanceGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type StorePerformanceGetListRequest struct {
	*api.BaseRequest

	Vid             int64  `json:"vid,omitempty"`
	PerformanceType int64  `json:"performanceType,omitempty"`
	LastId          int64  `json:"lastId,omitempty"`
	PageSize        int64  `json:"pageSize,omitempty"`
	StartTime       string `json:"startTime,omitempty"`
	EndTime         string `json:"endTime,omitempty"`
	TimeType        int64  `json:"timeType,omitempty"`
}

type StorePerformanceGetListResponse struct {
	PageList   []StorePerformanceGetListResponsePageList `json:"pageList,omitempty"`
	PageNum    int64                                     `json:"pageNum,omitempty"`
	PageSize   int64                                     `json:"pageSize,omitempty"`
	TotalCount int64                                     `json:"totalCount,omitempty"`
}

func CreateStorePerformanceGetListRequest() (request *StorePerformanceGetListRequest) {
	request = &StorePerformanceGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("bos", "v2.0", "store/performance/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateStorePerformanceGetListResponse() (response *api.BaseResponse[StorePerformanceGetListResponse]) {
	response = api.CreateResponse[StorePerformanceGetListResponse](&StorePerformanceGetListResponse{})
	return
}

type StorePerformanceGetListResponsePageList struct {
	Id                         int64  `json:"id,omitempty"`
	Vid                        int64  `json:"vid,omitempty"`
	StoreName                  string `json:"storeName,omitempty"`
	TradeNo                    int64  `json:"tradeNo,omitempty"`
	Wid                        int64  `json:"wid,omitempty"`
	TradeType                  int64  `json:"tradeType,omitempty"`
	TradeCreateTime            string `json:"tradeCreateTime,omitempty"`
	TradeUpdateTime            string `json:"tradeUpdateTime,omitempty"`
	PaymentAmount              int64  `json:"paymentAmount,omitempty"`
	BalanceDiscountAmount      int64  `json:"balanceDiscountAmount,omitempty"`
	PointsAmount               int64  `json:"pointsAmount,omitempty"`
	CardDiscountAmount         int64  `json:"cardDiscountAmount,omitempty"`
	PerformanceAmount          int64  `json:"performanceAmount,omitempty"`
	StorePayAmount             int64  `json:"storePayAmount,omitempty"`
	StoreBalanceDiscountAmount int64  `json:"storeBalanceDiscountAmount,omitempty"`
	StoredPointsAmount         int64  `json:"storedPointsAmount,omitempty"`
	StoreCardDiscountAmount    int64  `json:"storeCardDiscountAmount,omitempty"`
	StoreProportion            int64  `json:"storeProportion,omitempty"`
	StorePerformance           int64  `json:"storePerformance,omitempty"`
	PerformanceType            int64  `json:"performanceType,omitempty"`
	OrderSource                int64  `json:"orderSource,omitempty"`
}
