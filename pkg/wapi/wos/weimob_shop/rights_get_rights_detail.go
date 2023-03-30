package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// RightsGetRightsDetail
// @id 1146
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1146?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=根据售后单号查询售后单)
func (client *Client) RightsGetRightsDetail(request *RightsGetRightsDetailRequest) (response *RightsGetRightsDetailResponse, err error) {
	rpcResponse := CreateRightsGetRightsDetailResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type RightsGetRightsDetailRequest struct {
	*api.BaseRequest

	Obj RightsGetRightsDetailRequestObj `json:"obj,omitempty"`
}

type RightsGetRightsDetailResponse struct {
}

func CreateRightsGetRightsDetailRequest() (request *RightsGetRightsDetailRequest) {
	request = &RightsGetRightsDetailRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "rights/getRightsDetail", "apigw")
	request.Method = api.POST
	return
}

func CreateRightsGetRightsDetailResponse() (response *api.BaseResponse[RightsGetRightsDetailResponse]) {
	response = api.CreateResponse[RightsGetRightsDetailResponse](&RightsGetRightsDetailResponse{})
	return
}

type RightsGetRightsDetailRequestObj struct {
}
