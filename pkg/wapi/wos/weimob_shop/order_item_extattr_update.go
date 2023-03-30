package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// OrderItemExtattrUpdate
// @id 1817
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1817?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=添加商品附加信息)
func (client *Client) OrderItemExtattrUpdate(request *OrderItemExtattrUpdateRequest) (response *OrderItemExtattrUpdateResponse, err error) {
	rpcResponse := CreateOrderItemExtattrUpdateResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type OrderItemExtattrUpdateRequest struct {
	*api.BaseRequest

	ItemRemarks []OrderItemExtattrUpdateRequestItemRemarks `json:"itemRemarks,omitempty"`
	OrderNo     int64                                      `json:"orderNo,omitempty"`
}

type OrderItemExtattrUpdateResponse struct {
	Success bool `json:"success,omitempty"`
}

func CreateOrderItemExtattrUpdateRequest() (request *OrderItemExtattrUpdateRequest) {
	request = &OrderItemExtattrUpdateRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "order/item/extattr/update", "apigw")
	request.Method = api.POST
	return
}

func CreateOrderItemExtattrUpdateResponse() (response *api.BaseResponse[OrderItemExtattrUpdateResponse]) {
	response = api.CreateResponse[OrderItemExtattrUpdateResponse](&OrderItemExtattrUpdateResponse{})
	return
}

type OrderItemExtattrUpdateRequestItemRemarks struct {
	CustomFieldGroup []OrderItemExtattrUpdateRequestCustomFieldGroup `json:"customFieldGroup,omitempty"`
	ItemId           int64                                           `json:"itemId,omitempty"`
}

type OrderItemExtattrUpdateRequestCustomFieldGroup struct {
}
