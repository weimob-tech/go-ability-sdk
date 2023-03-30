package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// OrderChangeowner
// @id 2665
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=变更所属人)
func (client *Client) OrderChangeowner(request *OrderChangeownerRequest) (response *OrderChangeownerResponse, err error) {
	rpcResponse := CreateOrderChangeownerResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type OrderChangeownerRequest struct {
	*api.BaseRequest

	Keys       []map[string]any `json:"keys,omitempty"`
	NewOwnerId int64            `json:"newOwnerId,omitempty"`
	UserWid    int64            `json:"userWid,omitempty"`
}

type OrderChangeownerResponse struct {
}

func CreateOrderChangeownerRequest() (request *OrderChangeownerRequest) {
	request = &OrderChangeownerRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "order/changeowner", "api")
	request.Method = api.POST
	return
}

func CreateOrderChangeownerResponse() (response *api.BaseResponse[OrderChangeownerResponse]) {
	response = api.CreateResponse[OrderChangeownerResponse](&OrderChangeownerResponse{})
	return
}
