package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// FieldsOrder
// @id 2661
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询订单字段列表)
func (client *Client) FieldsOrder(request *FieldsOrderRequest) (response *FieldsOrderResponse, err error) {
	rpcResponse := CreateFieldsOrderResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type FieldsOrderRequest struct {
	*api.BaseRequest

	UserWid int64  `json:"userWid,omitempty"`
	Stage   string `json:"stage,omitempty"`
}

type FieldsOrderResponse struct {
}

func CreateFieldsOrderRequest() (request *FieldsOrderRequest) {
	request = &FieldsOrderRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "fields/order", "api")
	request.Method = api.POST
	return
}

func CreateFieldsOrderResponse() (response *api.BaseResponse[FieldsOrderResponse]) {
	response = api.CreateResponse[FieldsOrderResponse](&FieldsOrderResponse{})
	return
}
