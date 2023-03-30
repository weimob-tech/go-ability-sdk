package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// OrderDelete
// @id 2679
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=订单信息删除)
func (client *Client) OrderDelete(request *OrderDeleteRequest) (response *OrderDeleteResponse, err error) {
	rpcResponse := CreateOrderDeleteResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type OrderDeleteRequest struct {
	*api.BaseRequest

	UserWid     int64    `json:"userWid,omitempty"`
	OrderId     string   `json:"orderId,omitempty"`
	OrderIdList []string `json:"orderIdList,omitempty"`
}

type OrderDeleteResponse struct {
}

func CreateOrderDeleteRequest() (request *OrderDeleteRequest) {
	request = &OrderDeleteRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "order/delete", "api")
	request.Method = api.POST
	return
}

func CreateOrderDeleteResponse() (response *api.BaseResponse[OrderDeleteResponse]) {
	response = api.CreateResponse[OrderDeleteResponse](&OrderDeleteResponse{})
	return
}
