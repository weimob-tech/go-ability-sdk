package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// RightsQueryRightsList
// @id 1118
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1118?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=根据时间分页查询售后列表)
func (client *Client) RightsQueryRightsList(request *RightsQueryRightsListRequest) (response *RightsQueryRightsListResponse, err error) {
	rpcResponse := CreateRightsQueryRightsListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type RightsQueryRightsListRequest struct {
	*api.BaseRequest

	Obj RightsQueryRightsListRequestObj `json:"obj,omitempty"`
}

type RightsQueryRightsListResponse struct {
}

func CreateRightsQueryRightsListRequest() (request *RightsQueryRightsListRequest) {
	request = &RightsQueryRightsListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "rights/queryRightsList", "apigw")
	request.Method = api.POST
	return
}

func CreateRightsQueryRightsListResponse() (response *api.BaseResponse[RightsQueryRightsListResponse]) {
	response = api.CreateResponse[RightsQueryRightsListResponse](&RightsQueryRightsListResponse{})
	return
}

type RightsQueryRightsListRequestObj struct {
}
