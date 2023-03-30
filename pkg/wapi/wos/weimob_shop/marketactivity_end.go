package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MarketactivityEnd
// @id 1456
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1456?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=结束活动)
func (client *Client) MarketactivityEnd(request *MarketactivityEndRequest) (response *MarketactivityEndResponse, err error) {
	rpcResponse := CreateMarketactivityEndResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MarketactivityEndRequest struct {
	*api.BaseRequest

	BasicInfo    MarketactivityEndRequestBasicInfo `json:"basicInfo,omitempty"`
	ActivityId   int64                             `json:"activityId,omitempty"`
	ActivityType int64                             `json:"activityType,omitempty"`
}

type MarketactivityEndResponse struct {
	Result bool `json:"result,omitempty"`
}

func CreateMarketactivityEndRequest() (request *MarketactivityEndRequest) {
	request = &MarketactivityEndRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "marketactivity/end", "apigw")
	request.Method = api.POST
	return
}

func CreateMarketactivityEndResponse() (response *api.BaseResponse[MarketactivityEndResponse]) {
	response = api.CreateResponse[MarketactivityEndResponse](&MarketactivityEndResponse{})
	return
}

type MarketactivityEndRequestBasicInfo struct {
	Vid     int64 `json:"vid,omitempty"`
	VidType int64 `json:"vidType,omitempty"`
}
