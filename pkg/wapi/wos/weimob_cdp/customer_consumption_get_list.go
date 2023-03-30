package weimob_cdp

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CustomerConsumptionGetList
// @id 1544
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1544?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询客户消费统计信息列表)
func (client *Client) CustomerConsumptionGetList(request *CustomerConsumptionGetListRequest) (response *CustomerConsumptionGetListResponse, err error) {
	rpcResponse := CreateCustomerConsumptionGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CustomerConsumptionGetListRequest struct {
	*api.BaseRequest

	TargetProductInstanceId int64   `json:"targetProductInstanceId,omitempty"`
	UkeyList                []int64 `json:"ukeyList,omitempty"`
	BosId                   int64   `json:"bosId,omitempty"`
	EndDay                  string  `json:"endDay,omitempty"`
	StartDay                string  `json:"startDay,omitempty"`
}

type CustomerConsumptionGetListResponse struct {
	ListVo []CustomerConsumptionGetListResponseListVo `json:"listVo,omitempty"`
}

func CreateCustomerConsumptionGetListRequest() (request *CustomerConsumptionGetListRequest) {
	request = &CustomerConsumptionGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_cdp", "v2.0", "customer/consumption/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateCustomerConsumptionGetListResponse() (response *api.BaseResponse[CustomerConsumptionGetListResponse]) {
	response = api.CreateResponse[CustomerConsumptionGetListResponse](&CustomerConsumptionGetListResponse{})
	return
}

type CustomerConsumptionGetListResponseListVo struct {
	TotalConsumptionCountD1    int64  `json:"totalConsumptionCountD1,omitempty"`
	TotalConsumptionCountD2    int64  `json:"totalConsumptionCountD2,omitempty"`
	TotalConsumptionCountD3    int64  `json:"totalConsumptionCountD3,omitempty"`
	TotalConsumptionCountD7    int64  `json:"totalConsumptionCountD7,omitempty"`
	TotalConsumptionCountD30   int64  `json:"totalConsumptionCountD30,omitempty"`
	TotalConsumptionCountD60   int64  `json:"totalConsumptionCountD60,omitempty"`
	TotalConsumptionCountD90   int64  `json:"totalConsumptionCountD90,omitempty"`
	TotalConsumptionCountD180  int64  `json:"totalConsumptionCountD180,omitempty"`
	TotalConsumptionCountD365  int64  `json:"totalConsumptionCountD365,omitempty"`
	TotalConsumptionCountM1    int64  `json:"totalConsumptionCountM1,omitempty"`
	TotalConsumptionCountM2    int64  `json:"totalConsumptionCountM2,omitempty"`
	TotalConsumptionCountM3    int64  `json:"totalConsumptionCountM3,omitempty"`
	TotalConsumptionCountM6    int64  `json:"totalConsumptionCountM6,omitempty"`
	TotalConsumptionCountY1    int64  `json:"totalConsumptionCountY1,omitempty"`
	TotalConsumptionCountY2    int64  `json:"totalConsumptionCountY2,omitempty"`
	TotalConsumptionCountY3    int64  `json:"totalConsumptionCountY3,omitempty"`
	GoodsCountD1               int64  `json:"goodsCountD1,omitempty"`
	GoodsCountD2               int64  `json:"goodsCountD2,omitempty"`
	GoodsCountD3               int64  `json:"goodsCountD3,omitempty"`
	GoodsCountD7               int64  `json:"goodsCountD7,omitempty"`
	GoodsCountD30              int64  `json:"goodsCountD30,omitempty"`
	GoodsCountD60              int64  `json:"goodsCountD60,omitempty"`
	GoodsCountD90              int64  `json:"goodsCountD90,omitempty"`
	GoodsCountD180             int64  `json:"goodsCountD180,omitempty"`
	GoodsCountD365             int64  `json:"goodsCountD365,omitempty"`
	GoodsCountM1               int64  `json:"goodsCountM1,omitempty"`
	GoodsCountM2               int64  `json:"goodsCountM2,omitempty"`
	GoodsCountM3               int64  `json:"goodsCountM3,omitempty"`
	GoodsCountM6               int64  `json:"goodsCountM6,omitempty"`
	GoodsCountY1               int64  `json:"goodsCountY1,omitempty"`
	GoodsCountY2               int64  `json:"goodsCountY2,omitempty"`
	GoodsCountY3               int64  `json:"goodsCountY3,omitempty"`
	TotalConsumptionAmountD1   int64  `json:"totalConsumptionAmountD1,omitempty"`
	TotalConsumptionAmountD2   int64  `json:"totalConsumptionAmountD2,omitempty"`
	TotalConsumptionAmountD3   int64  `json:"totalConsumptionAmountD3,omitempty"`
	TotalConsumptionAmountD7   int64  `json:"totalConsumptionAmountD7,omitempty"`
	TotalConsumptionAmountD30  int64  `json:"totalConsumptionAmountD30,omitempty"`
	TotalConsumptionAmountD60  int64  `json:"totalConsumptionAmountD60,omitempty"`
	TotalConsumptionAmountD90  int64  `json:"totalConsumptionAmountD90,omitempty"`
	TotalConsumptionAmountD180 int64  `json:"totalConsumptionAmountD180,omitempty"`
	TotalConsumptionAmountD365 int64  `json:"totalConsumptionAmountD365,omitempty"`
	TotalConsumptionAmountM1   int64  `json:"totalConsumptionAmountM1,omitempty"`
	TotalConsumptionAmountM2   int64  `json:"totalConsumptionAmountM2,omitempty"`
	TotalConsumptionAmountM3   int64  `json:"totalConsumptionAmountM3,omitempty"`
	TotalConsumptionAmountM6   int64  `json:"totalConsumptionAmountM6,omitempty"`
	TotalConsumptionAmountY1   int64  `json:"totalConsumptionAmountY1,omitempty"`
	TotalConsumptionAmountY2   int64  `json:"totalConsumptionAmountY2,omitempty"`
	TotalConsumptionAmountY3   int64  `json:"totalConsumptionAmountY3,omitempty"`
	BosId                      string `json:"bosId,omitempty"`
	ChannelId                  string `json:"channelId,omitempty"`
	ChannelType                string `json:"channelType,omitempty"`
	ConsumptionAmount          int64  `json:"consumptionAmount,omitempty"`
	ConsumptionCount           int64  `json:"consumptionCount,omitempty"`
	ExternalConsumptionAmount  int64  `json:"externalConsumptionAmount,omitempty"`
	ExternalConsumptionCount   int64  `json:"externalConsumptionCount,omitempty"`
	ExternalRightsAmount       int64  `json:"externalRightsAmount,omitempty"`
	ExternalRightsCount        int64  `json:"externalRightsCount,omitempty"`
	FirstOrderTime             int64  `json:"firstOrderTime,omitempty"`
	FirstPayTime               int64  `json:"firstPayTime,omitempty"`
	GoodsCount                 int64  `json:"goodsCount,omitempty"`
	LastOrderTime              int64  `json:"lastOrderTime,omitempty"`
	LastPayAmount              int64  `json:"lastPayAmount,omitempty"`
	LastPayTime                int64  `json:"lastPayTime,omitempty"`
	OriginalAmount             int64  `json:"originalAmount,omitempty"`
	PaymentAmount              int64  `json:"paymentAmount,omitempty"`
	SchannelId                 int64  `json:"schannelId,omitempty"`
	SchannelType               int64  `json:"schannelType,omitempty"`
	RightsAmount               int64  `json:"rightsAmount,omitempty"`
	RightsCount                int64  `json:"rightsCount,omitempty"`
	TotalConsumptionAmount     int64  `json:"totalConsumptionAmount,omitempty"`
	TotalConsumptionCount      int64  `json:"totalConsumptionCount,omitempty"`
	TotalRightsAmount          int64  `json:"totalRightsAmount,omitempty"`
	TotalRightsCount           int64  `json:"totalRightsCount,omitempty"`
	Ukey                       string `json:"ukey,omitempty"`
	UkeyType                   string `json:"ukeyType,omitempty"`
	FirstGoodsCount            int64  `json:"firstGoodsCount,omitempty"`
	FirstPaymentAmount         int64  `json:"firstPaymentAmount,omitempty"`
	FirstStoreId               int64  `json:"firstStoreId,omitempty"`
	LastPaymentAmount          int64  `json:"lastPaymentAmount,omitempty"`
	LastGoodsCount             int64  `json:"lastGoodsCount,omitempty"`
	LastConsumptionStore       int64  `json:"lastConsumptionStore,omitempty"`
	SingleMaxGoodsCount        int64  `json:"singleMaxGoodsCount,omitempty"`
	SingleMaxPaymentAmount     int64  `json:"singleMaxPaymentAmount,omitempty"`
	ExternalGoodsCount         int64  `json:"externalGoodsCount,omitempty"`
	InsideGoodsCount           int64  `json:"insideGoodsCount,omitempty"`
	BuybackPeriod              int64  `json:"buybackPeriod,omitempty"`
	AvgGoodsPerOrder           int64  `json:"avgGoodsPerOrder,omitempty"`
	AverageDiscount            int64  `json:"averageDiscount,omitempty"`
}
