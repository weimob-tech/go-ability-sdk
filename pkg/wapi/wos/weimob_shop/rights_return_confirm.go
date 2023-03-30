package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// RightsReturnConfirm
// @id 1845
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1845?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=商家确认退货)
func (client *Client) RightsReturnConfirm(request *RightsReturnConfirmRequest) (response *RightsReturnConfirmResponse, err error) {
	rpcResponse := CreateRightsReturnConfirmResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type RightsReturnConfirmRequest struct {
	*api.BaseRequest

	RightsId int64 `json:"rightsId,omitempty"`
}

type RightsReturnConfirmResponse struct {
	Success bool `json:"success,omitempty"`
}

func CreateRightsReturnConfirmRequest() (request *RightsReturnConfirmRequest) {
	request = &RightsReturnConfirmRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "rights/return/confirm", "apigw")
	request.Method = api.POST
	return
}

func CreateRightsReturnConfirmResponse() (response *api.BaseResponse[RightsReturnConfirmResponse]) {
	response = api.CreateResponse[RightsReturnConfirmResponse](&RightsReturnConfirmResponse{})
	return
}
