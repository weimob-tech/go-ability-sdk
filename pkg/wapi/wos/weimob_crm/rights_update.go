package weimob_crm

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// RightsUpdate
// @id 2323
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2323?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=更新售后单)
func (client *Client) RightsUpdate(request *RightsUpdateRequest) (response *RightsUpdateResponse, err error) {
	rpcResponse := CreateRightsUpdateResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type RightsUpdateRequest struct {
	*api.BaseRequest

	RightsOrder      RightsUpdateRequestRightsOrder        `json:"rightsOrder,omitempty"`
	OriginOrder      RightsUpdateRequestOriginOrder        `json:"originOrder,omitempty"`
	AssociateBizList []RightsUpdateRequestAssociateBizList `json:"associateBizList,omitempty"`
}

type RightsUpdateResponse struct {
}

func CreateRightsUpdateRequest() (request *RightsUpdateRequest) {
	request = &RightsUpdateRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_crm", "v2.0", "rights/update", "apigw")
	request.Method = api.POST
	return
}

func CreateRightsUpdateResponse() (response *api.BaseResponse[RightsUpdateResponse]) {
	response = api.CreateResponse[RightsUpdateResponse](&RightsUpdateResponse{})
	return
}

type RightsUpdateRequestRightsOrder struct {
	OutRightsNo  string `json:"outRightsNo,omitempty"`
	ChannelType  int64  `json:"channelType,omitempty"`
	PlatformType string `json:"platformType,omitempty"`
}

type RightsUpdateRequestOriginOrder struct {
	OutOriginalOrderNo string `json:"outOriginalOrderNo,omitempty"`
}

type RightsUpdateRequestAssociateBizList struct {
	BizValue int64 `json:"bizValue,omitempty"`
}
