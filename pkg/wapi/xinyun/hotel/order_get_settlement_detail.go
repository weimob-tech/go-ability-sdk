package hotel

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// OrderGetSettlementDetail
// @id 1597
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=获取结算明细)
func (client *Client) OrderGetSettlementDetail(request *OrderGetSettlementDetailRequest) (response *OrderGetSettlementDetailResponse, err error) {
	rpcResponse := CreateOrderGetSettlementDetailResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type OrderGetSettlementDetailRequest struct {
	*api.BaseRequest

	StoreId   int64    `json:"storeId,omitempty"`
	BookingNo string   `json:"bookingNo,omitempty"`
	RecordNos []string `json:"recordNos,omitempty"`
}

type OrderGetSettlementDetailResponse struct {
}

func CreateOrderGetSettlementDetailRequest() (request *OrderGetSettlementDetailRequest) {
	request = &OrderGetSettlementDetailRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("hotel", "1_0", "order/getSettlementDetail", "api")
	request.Method = api.POST
	return
}

func CreateOrderGetSettlementDetailResponse() (response *api.BaseResponse[OrderGetSettlementDetailResponse]) {
	response = api.CreateResponse[OrderGetSettlementDetailResponse](&OrderGetSettlementDetailResponse{})
	return
}
