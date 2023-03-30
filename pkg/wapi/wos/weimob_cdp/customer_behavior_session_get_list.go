package weimob_cdp

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CustomerBehaviorSessionGetList
// @id 1909
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1909?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询客户会话列表)
func (client *Client) CustomerBehaviorSessionGetList(request *CustomerBehaviorSessionGetListRequest) (response *CustomerBehaviorSessionGetListResponse, err error) {
	rpcResponse := CreateCustomerBehaviorSessionGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CustomerBehaviorSessionGetListRequest struct {
	*api.BaseRequest

	BosId     int64  `json:"bosId,omitempty"`
	ChannelId string `json:"channelId,omitempty"`
	PageNum   int64  `json:"pageNum,omitempty"`
	PageSize  int64  `json:"pageSize,omitempty"`
	Ukey      string `json:"ukey,omitempty"`
}

type CustomerBehaviorSessionGetListResponse struct {
	SessionList CustomerBehaviorSessionGetListResponseSessionList `json:"sessionList,omitempty"`
}

func CreateCustomerBehaviorSessionGetListRequest() (request *CustomerBehaviorSessionGetListRequest) {
	request = &CustomerBehaviorSessionGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_cdp", "v2.0", "customer/behavior/session/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateCustomerBehaviorSessionGetListResponse() (response *api.BaseResponse[CustomerBehaviorSessionGetListResponse]) {
	response = api.CreateResponse[CustomerBehaviorSessionGetListResponse](&CustomerBehaviorSessionGetListResponse{})
	return
}

type CustomerBehaviorSessionGetListResponseSessionList struct {
	List     []CustomerBehaviorSessionGetListResponselist `json:"list,omitempty"`
	PageSize int64                                        `json:"pageSize,omitempty"`
	PageNum  int64                                        `json:"pageNum,omitempty"`
	Total    int64                                        `json:"total,omitempty"`
}

type CustomerBehaviorSessionGetListResponselist struct {
	BosId         int64  `json:"bosId,omitempty"`
	Ukey          string `json:"ukey,omitempty"`
	UkeyType      string `json:"ukeyType,omitempty"`
	ChannelType   string `json:"channelType,omitempty"`
	ChannelId     string `json:"channelId,omitempty"`
	InPageName    string `json:"inPageName,omitempty"`
	OutPageName   string `json:"outPageName,omitempty"`
	SessionId     string `json:"sessionId,omitempty"`
	GuiderName    string `json:"guiderName,omitempty"`
	TuikeName     string `json:"tuikeName,omitempty"`
	Vid           string `json:"vid,omitempty"`
	VidType       string `json:"vidType,omitempty"`
	VidName       string `json:"vidName,omitempty"`
	StartTime     int64  `json:"startTime,omitempty"`
	EndTime       int64  `json:"endTime,omitempty"`
	StatType      string `json:"statType,omitempty"`
	ChannelSource string `json:"channelSource,omitempty"`
	Scene         string `json:"scene,omitempty"`
	SceneH5       string `json:"sceneH5,omitempty"`
	SourceEnd     string `json:"sourceEnd,omitempty"`
}
