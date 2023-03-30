package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CustomerNicheList
// @id 1711
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=客户下关联商机列表)
func (client *Client) CustomerNicheList(request *CustomerNicheListRequest) (response *CustomerNicheListResponse, err error) {
	rpcResponse := CreateCustomerNicheListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CustomerNicheListRequest struct {
	*api.BaseRequest

	Wid      int    `json:"wid,omitempty"`
	PageNum  int    `json:"pageNum,omitempty"`
	PageSize int    `json:"pageSize,omitempty"`
	Key      string `json:"key,omitempty"`
}

type CustomerNicheListResponse struct {
}

func CreateCustomerNicheListRequest() (request *CustomerNicheListRequest) {
	request = &CustomerNicheListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "customer/nicheList", "api")
	request.Method = api.POST
	return
}

func CreateCustomerNicheListResponse() (response *api.BaseResponse[CustomerNicheListResponse]) {
	response = api.CreateResponse[CustomerNicheListResponse](&CustomerNicheListResponse{})
	return
}
