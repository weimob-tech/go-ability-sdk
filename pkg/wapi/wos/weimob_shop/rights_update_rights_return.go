package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// RightsUpdateRightsReturn
// @id 1150
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1150?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=根据售后单号填写退货物流)
func (client *Client) RightsUpdateRightsReturn(request *RightsUpdateRightsReturnRequest) (response *RightsUpdateRightsReturnResponse, err error) {
	rpcResponse := CreateRightsUpdateRightsReturnResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type RightsUpdateRightsReturnRequest struct {
	*api.BaseRequest

	Obj RightsUpdateRightsReturnRequestObj `json:"obj,omitempty"`
}

type RightsUpdateRightsReturnResponse struct {
}

func CreateRightsUpdateRightsReturnRequest() (request *RightsUpdateRightsReturnRequest) {
	request = &RightsUpdateRightsReturnRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "rights/updateRightsReturn", "apigw")
	request.Method = api.POST
	return
}

func CreateRightsUpdateRightsReturnResponse() (response *api.BaseResponse[RightsUpdateRightsReturnResponse]) {
	response = api.CreateResponse[RightsUpdateRightsReturnResponse](&RightsUpdateRightsReturnResponse{})
	return
}

type RightsUpdateRightsReturnRequestObj struct {
}
