package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CustomerDetail
// @id 1717
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询客户详情)
func (client *Client) CustomerDetail(request *CustomerDetailRequest) (response *CustomerDetailResponse, err error) {
	rpcResponse := CreateCustomerDetailResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CustomerDetailRequest struct {
	*api.BaseRequest

	Wid int    `json:"wid,omitempty"`
	Key string `json:"key,omitempty"`
}

type CustomerDetailResponse struct {
}

func CreateCustomerDetailRequest() (request *CustomerDetailRequest) {
	request = &CustomerDetailRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "customer/detail", "api")
	request.Method = api.POST
	return
}

func CreateCustomerDetailResponse() (response *api.BaseResponse[CustomerDetailResponse]) {
	response = api.CreateResponse[CustomerDetailResponse](&CustomerDetailResponse{})
	return
}
