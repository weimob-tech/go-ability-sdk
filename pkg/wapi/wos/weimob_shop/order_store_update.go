package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// OrderStoreUpdate
// @id 1812
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1812?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=订单转单)
func (client *Client) OrderStoreUpdate(request *OrderStoreUpdateRequest) (response *OrderStoreUpdateResponse, err error) {
	rpcResponse := CreateOrderStoreUpdateResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type OrderStoreUpdateRequest struct {
	*api.BaseRequest

	TransferOrder OrderStoreUpdateRequestTransferOrder `json:"transferOrder,omitempty"`
}

type OrderStoreUpdateResponse struct {
	Success bool `json:"success,omitempty"`
}

func CreateOrderStoreUpdateRequest() (request *OrderStoreUpdateRequest) {
	request = &OrderStoreUpdateRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "order/store/update", "apigw")
	request.Method = api.POST
	return
}

func CreateOrderStoreUpdateResponse() (response *api.BaseResponse[OrderStoreUpdateResponse]) {
	response = api.CreateResponse[OrderStoreUpdateResponse](&OrderStoreUpdateResponse{})
	return
}

type OrderStoreUpdateRequestTransferOrder struct {
	OrderNo      int64 `json:"orderNo,omitempty"`
	TransferType int64 `json:"transferType,omitempty"`
	ProcessVid   int64 `json:"processVid,omitempty"`
}
