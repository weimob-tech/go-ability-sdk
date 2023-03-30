package weimob_guide

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// ProcessShareinfoGet
// @id 1378
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1378?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询导购分享信息)
func (client *Client) ProcessShareinfoGet(request *ProcessShareinfoGetRequest) (response *ProcessShareinfoGetResponse, err error) {
	rpcResponse := CreateProcessShareinfoGetResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type ProcessShareinfoGetRequest struct {
	*api.BaseRequest

	BindSceneContent ProcessShareinfoGetRequestBindSceneContent `json:"bindSceneContent	,omitempty"`
	ExtendMap        ProcessShareinfoGetRequestExtendMap        `json:"extendMap,omitempty"`
	GuiderWid        int64                                      `json:"guiderWid,omitempty"`
	BindSceneType    int64                                      `json:"bindSceneType,omitempty"`
	BizVid           int64                                      `json:"bizVid,omitempty"`
	BizVidType       int64                                      `json:"bizVidType,omitempty"`
	Cid              int64                                      `json:"cid,omitempty"`
}

type ProcessShareinfoGetResponse struct {
	GuideShareUrl ProcessShareinfoGetResponseGuideShareUrl `json:"guideShareUrl,omitempty"`
	TicketValue   string                                   `json:"ticketValue,omitempty"`
	TicketName    string                                   `json:"ticketName,omitempty"`
}

func CreateProcessShareinfoGetRequest() (request *ProcessShareinfoGetRequest) {
	request = &ProcessShareinfoGetRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_guide", "v2.0", "process/shareinfo/get", "apigw")
	request.Method = api.POST
	return
}

func CreateProcessShareinfoGetResponse() (response *api.BaseResponse[ProcessShareinfoGetResponse]) {
	response = api.CreateResponse[ProcessShareinfoGetResponse](&ProcessShareinfoGetResponse{})
	return
}

type ProcessShareinfoGetRequestBindSceneContent struct {
}

type ProcessShareinfoGetRequestExtendMap struct {
}

type ProcessShareinfoGetResponseGuideShareUrl struct {
}
