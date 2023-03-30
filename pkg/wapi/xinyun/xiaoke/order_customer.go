package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// OrderCustomer
// @id 2662
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=获取客户下订单列表)
func (client *Client) OrderCustomer(request *OrderCustomerRequest) (response *OrderCustomerResponse, err error) {
	rpcResponse := CreateOrderCustomerResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type OrderCustomerRequest struct {
	*api.BaseRequest

	UserWid  int64  `json:"userWid,omitempty"`
	Key      string `json:"key,omitempty"`
	PageSize int64  `json:"pageSize,omitempty"`
	PageNum  int64  `json:"pageNum,omitempty"`
}

type OrderCustomerResponse struct {
}

func CreateOrderCustomerRequest() (request *OrderCustomerRequest) {
	request = &OrderCustomerRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "order/customer", "api")
	request.Method = api.POST
	return
}

func CreateOrderCustomerResponse() (response *api.BaseResponse[OrderCustomerResponse]) {
	response = api.CreateResponse[OrderCustomerResponse](&OrderCustomerResponse{})
	return
}
