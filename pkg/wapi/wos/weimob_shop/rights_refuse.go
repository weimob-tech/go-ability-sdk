package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// RightsRefuse
// @id 1849
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1849?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=商家拒绝售后)
func (client *Client) RightsRefuse(request *RightsRefuseRequest) (response *RightsRefuseResponse, err error) {
	rpcResponse := CreateRightsRefuseResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type RightsRefuseRequest struct {
	*api.BaseRequest

	RefuseReason string `json:"refuseReason,omitempty"`
	RightsId     int64  `json:"rightsId,omitempty"`
}

type RightsRefuseResponse struct {
	Success bool `json:"success,omitempty"`
}

func CreateRightsRefuseRequest() (request *RightsRefuseRequest) {
	request = &RightsRefuseRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "rights/refuse", "apigw")
	request.Method = api.POST
	return
}

func CreateRightsRefuseResponse() (response *api.BaseResponse[RightsRefuseResponse]) {
	response = api.CreateResponse[RightsRefuseResponse](&RightsRefuseResponse{})
	return
}
