package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// RightsFlagUpdate
// @id 1841
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1841?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=商家标记售后单)
func (client *Client) RightsFlagUpdate(request *RightsFlagUpdateRequest) (response *RightsFlagUpdateResponse, err error) {
	rpcResponse := CreateRightsFlagUpdateResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type RightsFlagUpdateRequest struct {
	*api.BaseRequest

	FlagContent string  `json:"flagContent,omitempty"`
	FlagRank    int64   `json:"flagRank,omitempty"`
	RightsIds   []int64 `json:"rightsIds,omitempty"`
}

type RightsFlagUpdateResponse struct {
	Success bool `json:"success,omitempty"`
}

func CreateRightsFlagUpdateRequest() (request *RightsFlagUpdateRequest) {
	request = &RightsFlagUpdateRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "rights/flag/update", "apigw")
	request.Method = api.POST
	return
}

func CreateRightsFlagUpdateResponse() (response *api.BaseResponse[RightsFlagUpdateResponse]) {
	response = api.CreateResponse[RightsFlagUpdateResponse](&RightsFlagUpdateResponse{})
	return
}
