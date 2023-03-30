package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// RightsUpdateRightsFlag
// @id 1880
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=用于更新售后单的标记（增加、更新、删除）)
func (client *Client) RightsUpdateRightsFlag(request *RightsUpdateRightsFlagRequest) (response *RightsUpdateRightsFlagResponse, err error) {
	rpcResponse := CreateRightsUpdateRightsFlagResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type RightsUpdateRightsFlagRequest struct {
	*api.BaseRequest

	Id          int64  `json:"id,omitempty"`
	FlagRank    int    `json:"flagRank,omitempty"`
	FlagContent string `json:"flagContent,omitempty"`
}

type RightsUpdateRightsFlagResponse struct {
}

func CreateRightsUpdateRightsFlagRequest() (request *RightsUpdateRightsFlagRequest) {
	request = &RightsUpdateRightsFlagRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "rights/updateRightsFlag", "api")
	request.Method = api.POST
	return
}

func CreateRightsUpdateRightsFlagResponse() (response *api.BaseResponse[RightsUpdateRightsFlagResponse]) {
	response = api.CreateResponse[RightsUpdateRightsFlagResponse](&RightsUpdateRightsFlagResponse{})
	return
}
