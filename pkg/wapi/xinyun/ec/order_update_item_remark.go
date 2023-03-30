package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// OrderUpdateItemRemark
// @id 1176
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=更新订单中商品附加属性)
func (client *Client) OrderUpdateItemRemark(request *OrderUpdateItemRemarkRequest) (response *OrderUpdateItemRemarkResponse, err error) {
	rpcResponse := CreateOrderUpdateItemRemarkResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type OrderUpdateItemRemarkRequest struct {
	*api.BaseRequest

	ItemRemarkList []OrderUpdateItemRemarkRequestItemRemarkList `json:"itemRemarkList,omitempty"`
	OrderNo        int64                                        `json:"orderNo,omitempty"`
}

type OrderUpdateItemRemarkResponse struct {
}

func CreateOrderUpdateItemRemarkRequest() (request *OrderUpdateItemRemarkRequest) {
	request = &OrderUpdateItemRemarkRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "order/updateItemRemark", "api")
	request.Method = api.POST
	return
}

func CreateOrderUpdateItemRemarkResponse() (response *api.BaseResponse[OrderUpdateItemRemarkResponse]) {
	response = api.CreateResponse[OrderUpdateItemRemarkResponse](&OrderUpdateItemRemarkResponse{})
	return
}

type OrderUpdateItemRemarkRequestItemRemarkList struct {
	CustomFieldGroup []OrderUpdateItemRemarkRequestCustomFieldGroup `json:"customFieldGroup,omitempty"`
	ItemId           int64                                          `json:"itemId,omitempty"`
}

type OrderUpdateItemRemarkRequestCustomFieldGroup struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}
