package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CustomerPublicseaList
// @id 1718
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询客户公海列表)
func (client *Client) CustomerPublicseaList(request *CustomerPublicseaListRequest) (response *CustomerPublicseaListResponse, err error) {
	rpcResponse := CreateCustomerPublicseaListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CustomerPublicseaListRequest struct {
	*api.BaseRequest

	Wid int `json:"wid,omitempty"`
}

type CustomerPublicseaListResponse struct {
}

func CreateCustomerPublicseaListRequest() (request *CustomerPublicseaListRequest) {
	request = &CustomerPublicseaListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "customer/publicseaList", "api")
	request.Method = api.POST
	return
}

func CreateCustomerPublicseaListResponse() (response *api.BaseResponse[CustomerPublicseaListResponse]) {
	response = api.CreateResponse[CustomerPublicseaListResponse](&CustomerPublicseaListResponse{})
	return
}
