package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// RightsGetRightsDetailList
// @id 1881
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=根据售后父单号查询售后列表)
func (client *Client) RightsGetRightsDetailList(request *RightsGetRightsDetailListRequest) (response *RightsGetRightsDetailListResponse, err error) {
	rpcResponse := CreateRightsGetRightsDetailListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type RightsGetRightsDetailListRequest struct {
	*api.BaseRequest
}

type RightsGetRightsDetailListResponse struct {
}

func CreateRightsGetRightsDetailListRequest() (request *RightsGetRightsDetailListRequest) {
	request = &RightsGetRightsDetailListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "rights/getRightsDetailList", "api")
	request.Method = api.POST
	return
}

func CreateRightsGetRightsDetailListResponse() (response *api.BaseResponse[RightsGetRightsDetailListResponse]) {
	response = api.CreateResponse[RightsGetRightsDetailListResponse](&RightsGetRightsDetailListResponse{})
	return
}
