package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CommunityGrouponAddLogisticsInfo
// @id 3624
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=接龙团发货)
func (client *Client) CommunityGrouponAddLogisticsInfo(request *CommunityGrouponAddLogisticsInfoRequest) (response *CommunityGrouponAddLogisticsInfoResponse, err error) {
	rpcResponse := CreateCommunityGrouponAddLogisticsInfoResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CommunityGrouponAddLogisticsInfoRequest struct {
	*api.BaseRequest

	DeliveryNo          string `json:"deliveryNo,omitempty"`
	DeliveryCompanyCode string `json:"deliveryCompanyCode,omitempty"`
	GrouponId           int64  `json:"grouponId,omitempty"`
}

type CommunityGrouponAddLogisticsInfoResponse struct {
}

func CreateCommunityGrouponAddLogisticsInfoRequest() (request *CommunityGrouponAddLogisticsInfoRequest) {
	request = &CommunityGrouponAddLogisticsInfoRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "communityGroupon/addLogisticsInfo", "api")
	request.Method = api.POST
	return
}

func CreateCommunityGrouponAddLogisticsInfoResponse() (response *api.BaseResponse[CommunityGrouponAddLogisticsInfoResponse]) {
	response = api.CreateResponse[CommunityGrouponAddLogisticsInfoResponse](&CommunityGrouponAddLogisticsInfoResponse{})
	return
}
