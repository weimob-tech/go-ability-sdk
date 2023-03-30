package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CustomerStageList
// @id 1884
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询客户阶段列表)
func (client *Client) CustomerStageList(request *CustomerStageListRequest) (response *CustomerStageListResponse, err error) {
	rpcResponse := CreateCustomerStageListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CustomerStageListRequest struct {
	*api.BaseRequest

	Wid int64 `json:"wid,omitempty"`
}

type CustomerStageListResponse struct {
}

func CreateCustomerStageListRequest() (request *CustomerStageListRequest) {
	request = &CustomerStageListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "customer/stageList", "api")
	request.Method = api.POST
	return
}

func CreateCustomerStageListResponse() (response *api.BaseResponse[CustomerStageListResponse]) {
	response = api.CreateResponse[CustomerStageListResponse](&CustomerStageListResponse{})
	return
}
