package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CustomerList
// @id 1716
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询客户列表)
func (client *Client) CustomerList(request *CustomerListRequest) (response *CustomerListResponse, err error) {
	rpcResponse := CreateCustomerListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CustomerListRequest struct {
	*api.BaseRequest

	Wid            int    `json:"wid,omitempty"`
	PageNum        int    `json:"pageNum,omitempty"`
	PageSize       int    `json:"pageSize,omitempty"`
	Type           int    `json:"type,omitempty"`
	ResourcesScope int    `json:"resourcesScope,omitempty"`
	OrderByRule    int    `json:"orderByRule,omitempty"`
	PublicSeaId    string `json:"publicSeaId,omitempty"`
}

type CustomerListResponse struct {
	TotalPage  int   `json:"totalPage,omitempty"`
	PageSize   int   `json:"pageSize,omitempty"`
	TotalCount int64 `json:"totalCount,omitempty"`
	PageNum    int64 `json:"pageNum,omitempty"`
}

func CreateCustomerListRequest() (request *CustomerListRequest) {
	request = &CustomerListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "customer/list", "api")
	request.Method = api.POST
	return
}

func CreateCustomerListResponse() (response *api.BaseResponse[CustomerListResponse]) {
	response = api.CreateResponse[CustomerListResponse](&CustomerListResponse{})
	return
}
